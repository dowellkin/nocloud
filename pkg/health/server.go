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
package health

import (
	"context"

	"github.com/slntopp/nocloud/pkg/health/healthpb"
	"go.uber.org/zap"
)

type HealthServiceServer struct {
	healthpb.UnimplementedHealthServiceServer
	log *zap.Logger
}

func NewServer(log *zap.Logger) *HealthServiceServer {
	return &HealthServiceServer{log: log}
}

func (s *HealthServiceServer) Probe(ctx context.Context, request *healthpb.ProbeRequest) (*healthpb.ProbeResponse, error) {
	log := s.log.Named("Health Probe")
	log.Info("Probe received", zap.String("Type", request.ProbeType))
	if request.ProbeType == "PING" {
		return &healthpb.ProbeResponse{Response: "PONG"}, nil
	}

	return &healthpb.ProbeResponse{Response: "ok"}, nil
}
