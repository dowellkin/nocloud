/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package services_providers

import (
	"context"
	"fmt"

	"github.com/arangodb/go-driver"
	driverpb "github.com/slntopp/nocloud/pkg/drivers/instance/vanilla"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	sppb "github.com/slntopp/nocloud/pkg/services_providers/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type ServicesProviderServer struct {
	sppb.UnimplementedServicesProvidersServiceServer

	drivers map[string]driverpb.DriverServiceClient
	extention_servers map[string]sppb.ServicesProvidersExtentionsServiceClient
	db driver.Database
	ctrl graph.ServicesProvidersController
	ns_ctrl graph.NamespacesController

	log *zap.Logger
}

func NewServicesProviderServer(log *zap.Logger, db driver.Database) *ServicesProviderServer {
	return &ServicesProviderServer{
		log: log, db: db, ctrl: graph.NewServicesProvidersController(log, db),
		ns_ctrl: graph.NewNamespacesController(log, db),
		drivers: make(map[string]driverpb.DriverServiceClient),
		extention_servers: make(map[string]sppb.ServicesProvidersExtentionsServiceClient),
	}
}

func (s *ServicesProviderServer) RegisterDriver(type_key string, client driverpb.DriverServiceClient) {
	s.drivers[type_key] = client
}

func (s *ServicesProviderServer) RegisterExtentionServer(type_key string, client sppb.ServicesProvidersExtentionsServiceClient) {
	s.extention_servers[type_key] = client
}

func (s *ServicesProviderServer) ListExtentions(ctx context.Context, req *sppb.ListRequest) (res *sppb.ListExtentionsResponse, err error) {
	s.log.Debug("ListExtentions request received", zap.Any("request", req))

	keys := make([]string, 0, len(s.extention_servers))
	for k := range s.extention_servers {
		keys = append(keys, k)
	}
	
	return &sppb.ListExtentionsResponse{Types: keys}, nil
}

func (s *ServicesProviderServer) Test(ctx context.Context, req *sppb.ServicesProvider) (*sppb.TestResponse, error) {
	s.log.Debug("Test request received", zap.Any("request", req))

	title := req.GetTitle()
	if title == "" {
		return nil, status.Error(codes.InvalidArgument, "Services Provider 'title' is not given")
	}

	client, ok := s.drivers[req.GetType()]
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Driver type '%s' not registered", req.GetType()))
	}

	tfc, ok := ctx.Value(nocloud.TestFromCreate).(bool)
	tfc = ok && tfc
	if !tfc {
		for ext, data := range req.GetExtentions() {
			client, ok := s.extention_servers[ext]
			if !ok {
				return nil, status.Error(codes.NotFound, fmt.Sprintf("Extention Server type '%s' not registered", req.GetType()))
			}
			res, err := client.Test(ctx, &sppb.ServicesProvidersExtentionData{
				Data: data,
			})
			if err != nil {
				return nil, err
			}
			if !res.Result {
				err := fmt.Sprintf("Extention '%s': %s", ext, res.Error)
				return &sppb.TestResponse{
					Result: res.Result, Error: err,
					}, nil
			}
		}
	}

	test_req := &driverpb.TestServiceProviderConfigRequest{ServicesProvider: req}
	if !tfc && len(req.GetExtentions()) > 0 {
		test_req.SyntaxOnly = true
	}

	return client.TestServiceProviderConfig(ctx, test_req)
}

func (s *ServicesProviderServer) Create(ctx context.Context, req *sppb.ServicesProvider) (res *sppb.ServicesProvider, err error) {
	log := s.log.Named("Create")
	log.Debug("Create request received", zap.Any("request", req))

	testRes, err := s.Test(ctx, req)
	if err != nil {
		return req, err
	}
	if !testRes.Result {
		return req, status.Error(codes.Internal, testRes.Error)
	}

	sp := &graph.ServicesProvider{ServicesProvider: req}

	for ext, data := range req.GetExtentions() {
		client, ok := s.extention_servers[ext]
		if !ok {
			s.UnregisterExtentions(ctx, log, sp)
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Extention Server type '%s' not registered", req.GetType()))
		}
		res, err := client.Register(ctx, &sppb.ServicesProvidersExtentionData{
			Data: data,
		})
		if err != nil {
			s.UnregisterExtentions(ctx, log, sp)
			return nil, err
		}
		if !res.Result {
			s.UnregisterExtentions(ctx, log, sp)
			err := fmt.Sprintf("Extention '%s': %s", ext, res.Error)
			return req, status.Error(codes.Internal, err)
		}
	}

	ctx = context.WithValue(ctx, nocloud.TestFromCreate, true)
	testRes, err = s.Test(ctx, req)
	if err != nil {
		s.UnregisterExtentions(ctx, log, sp)
		return req, err
	}
	if !testRes.Result {
		s.UnregisterExtentions(ctx, log, sp)
		return req, status.Error(codes.Internal, testRes.Error)
	}

	err = s.ctrl.Create(ctx, sp)
	if err != nil {
		s.UnregisterExtentions(ctx, log, sp)
		s.log.Debug("Error allocating in DataBase", zap.Any("sp", sp), zap.Error(err))
		return req, status.Error(codes.Internal, "Error allocating in DataBase")
	}
	return sp.ServicesProvider, err
}

func (s *ServicesProviderServer) UnregisterExtentions(ctx context.Context, log *zap.Logger, sp *graph.ServicesProvider) {
	log.Debug("Unregistering ServicesProvider Extentions")
	for ext, data := range sp.GetExtentions() {
		client, ok := s.extention_servers[ext]
		if !ok {
			continue // TODO add Warnings
		}
		res, err := client.Unregister(ctx, &sppb.ServicesProvidersExtentionData{
			Data: data,
		})
		if err != nil {
			log.Error("Error unregistering extension", zap.Error(err))
			continue // TODO add Warnings
		}
		if !res.Result {
			log.Error("Error unregistering extension", zap.Any("result", res))
			continue // TODO add Warnings
		}
	}
}

func (s *ServicesProviderServer) Delete(ctx context.Context, req *sppb.DeleteRequest) (res *sppb.DeleteResponse, err error) {
	log := s.log.Named("Delete")
	log.Debug("Request received", zap.Any("request", req))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns, err := s.ns_ctrl.Get(ctx, "0")
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID.String(), 3)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights to perform Invoke")
	}
	
	sp, err := s.ctrl.Get(ctx, req.GetUuid())
	if err != nil {
		log.Error("Error getting ServicesProvider from DB", zap.Error(err))
		return nil, status.Error(codes.NotFound, "ServicesProvider not Found in DB")
	}

	services, err := s.ctrl.ListDeployments(ctx, sp)
	if err != nil {
		log.Error("Error getting provisioned Services from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "Couldn't get Provisioned Services")
	}

	if len(services) > 0 {
		res = &sppb.DeleteResponse{ Result: false, Services: make([]string, len(services)) }
		for i, service := range services {
			res.Services[i] = service.GetUuid()
		}
		return res, nil
	}

	err = s.ctrl.Delete(ctx, sp.GetUuid())
	if err != nil {
		log.Error("Error deleting ServicesProvider", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error deleting ServicesProvider")
	}

	s.UnregisterExtentions(ctx, log, sp)
	return &sppb.DeleteResponse{Result: true}, nil
}

func (s *ServicesProviderServer) Get(ctx context.Context, request *sppb.GetRequest) (res *sppb.ServicesProvider, err error) {
	log := s.log.Named("Get")
	log.Debug("Request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	r, err := s.ctrl.Get(ctx, request.GetUuid())
	if err != nil {
		log.Debug("Error getting ServicesProvider from DB", zap.Error(err))
		return nil, status.Error(codes.NotFound, "ServicesProvider not Found in DB")
	}

	return r.ServicesProvider, nil
}

func (s *ServicesProviderServer) List(ctx context.Context, req *sppb.ListRequest) (res *sppb.ListResponse, err error) {
	log := s.log.Named("List")
	log.Debug("Request received", zap.Any("request", req), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	r, err := s.ctrl.List(ctx, requestor)
	if err != nil {
		log.Debug("Error reading ServicesProviders from DB", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error reading ServicesProviders from DB")
	}

	res = &sppb.ListResponse{Pool: make([]*sppb.ServicesProvider, len(r))}
	for i, sp := range r {
		res.Pool[i] = sp.ServicesProvider
	}

	return res, nil
}

func (s *ServicesProviderServer) Invoke(ctx context.Context, req *sppb.ActionRequest) (res *structpb.Struct, err error) {
	log := s.log.Named("Invoke")
	log.Debug("Request received", zap.Any("request", req), zap.Any("context", ctx))

	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns, err := s.ns_ctrl.Get(ctx, "0")
	if err != nil {
		return nil, err
	}
	ok := graph.HasAccess(ctx, s.db, requestor, ns.ID.String(), 3)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights to perform Invoke")
	}

	sp, err := s.ctrl.Get(ctx, req.GetServicesProvider().GetUuid())
	if err != nil {
		log.Error("Error getting services provider",
			zap.String("services_provider", req.GetServicesProvider().GetUuid()),
			zap.Error(err),
		)
		return nil, status.Error(codes.NotFound, "ServicesProvider not found")
	}

	req.ServicesProvider = sp.ServicesProvider

	client, ok := s.drivers[sp.GetType()]
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Driver type '%s' not registered", sp.GetType()))
	}

	return client.Invoke(ctx, req)
}