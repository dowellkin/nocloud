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
package graph

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/roles"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"

	"github.com/slntopp/nocloud/pkg/credentials"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Account struct {
	Title string `json:"title"`
	driver.DocumentMeta
}

type AccountsController struct {
	col driver.Collection // Accounts Collection
	cred driver.Collection // Credentials Collection

	log *zap.Logger
}

func NewAccountsController(log *zap.Logger, db driver.Database) AccountsController {
	col, _ := db.Collection(context.TODO(), schema.ACCOUNTS_COL)
	cred, _ := db.Collection(context.TODO(), schema.CREDENTIALS_COL)
	log = log.Named("AccountsController")
	return AccountsController{log: log, col: col, cred: cred}
}

func (ctrl *AccountsController) Get(ctx context.Context, id string) (Account, error) {
	if id == "me" {
		id = ctx.Value(nocloud.NoCloudAccount).(string)
	}
	var r Account
	_, err := ctrl.col.ReadDocument(context.TODO(), id, &r)
	return r, err
}

func (ctrl *AccountsController) List(ctx context.Context, requestor Account, req_depth *int32) ([]Account, error) {
	var depth int32
	if req_depth == nil || *req_depth < 2 {
		depth = 2
	} else {
		depth = *req_depth
	}

	query := `FOR node IN 0..@depth OUTBOUND @account GRAPH @permissions_graph OPTIONS {order: "bfs", uniqueVertices: "global"} FILTER IS_SAME_COLLECTION(@@accounts, node) RETURN node`
	bindVars := map[string]interface{}{
		"depth": depth,
		"account": requestor.ID.String(),
		"permissions_graph": schema.PERMISSIONS_GRAPH.Name,
		"@accounts": schema.ACCOUNTS_COL,
	}
	ctrl.log.Debug("Ready to build query", zap.Any("bindVars", bindVars))

	c, err := ctrl.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var r []Account
	for {
		var acc Account 
		_, err := c.ReadDocument(ctx, &acc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		ctrl.log.Debug("Got document", zap.Any("account", acc))
		r = append(r, acc)
	}

	return r, nil
}

func (ctrl *AccountsController) Exists(ctx context.Context, id string) (bool, error) {
	return ctrl.col.DocumentExists(context.TODO(), id)
}

func (ctrl *AccountsController) Create(ctx context.Context, title string) (acc Account, err error) {
	acc = Account{
		Title: title,
	}
	meta, err := ctrl.col.CreateDocument(ctx, acc)
	acc.DocumentMeta = meta
	return acc, err
}

func (ctrl *AccountsController) Update(ctx context.Context, acc Account, title string) (err error) {
	patch := map[string]interface{}{
		"title": title,
	}
	_, err = ctrl.col.UpdateDocument(ctx, acc.Key, patch)
	return err
}

// Grant account access to namespace
func (acc *Account) LinkNamespace(ctx context.Context, edge driver.Collection, ns Namespace, level int32, role string) (error) {
	_, err := edge.CreateDocument(ctx, Access{
		From: acc.ID,
		To: ns.ID,
		Level: level,
		Role: role,
		DocumentMeta: driver.DocumentMeta {
			Key: acc.Key + "-" + ns.Key,
		},
	})
	return err
}

// Grant namespace access to account
func (acc *Account) JoinNamespace(ctx context.Context, edge driver.Collection, ns Namespace, level int32, role string) (error) {
	_, err := edge.CreateDocument(ctx, Access{
		From: ns.ID,
		To: acc.ID,
		Level: level,
		Role: role,
		DocumentMeta: driver.DocumentMeta {
			Key: ns.Key + "-" + acc.Key,
		},
	})
	return err
}

func (acc *Account) Delete(ctx context.Context, db driver.Database) (error) {
	err := DeleteNodeChildren(ctx, db, acc.ID.String())
	if err != nil {
		return err
	}

	graph, _ := db.Graph(ctx, schema.PERMISSIONS_GRAPH.Name)
	col, _ := graph.VertexCollection(ctx, schema.ACCOUNTS_COL)
	_, err = col.RemoveDocument(ctx, acc.Key)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *AccountsController) Delete(ctx context.Context, id string) (error) {
	acc, err := ctrl.Get(ctx, id)
	if err != nil {
		return err
	}
	return acc.Delete(ctx, ctrl.col.Database())
}

// Set Account Credentials, ensure account has only one credentials document linked per credentials type
func (ctrl *AccountsController) SetCredentials(ctx context.Context, acc Account, edge driver.Collection, c credentials.Credentials) (error) {
	cred, err := ctrl.cred.CreateDocument(ctx, c)	
	_, err = edge.CreateDocument(ctx, credentials.Link{
		From: acc.ID,
		To: cred.ID,
		Type: c.Type(),
		DocumentMeta: driver.DocumentMeta {
			Key: c.Type() + "-" + acc.Key, // Ensure only one credentials vertex per type
		},
	})
	if err != nil {
		return status.Error(codes.Internal, "Couldn't create credentials")
	}
	return nil
}

func (ctrl *AccountsController) UpdateCredentials(ctx context.Context, cred string, c credentials.Credentials) (err error) {
	_, err = ctrl.cred.UpdateDocument(ctx, cred, c)
	return err
}

func (ctrl *AccountsController) GetCredentials(ctx context.Context, edge_col driver.Collection, acc Account, auth_type string) (key string, has_credentials bool) {
	cred_edge := auth_type + "-" + acc.Key
	ctrl.log.Debug("Looking for Credentials Edge(Link)", zap.String("key", cred_edge))
	var edge credentials.Link
	_, err := edge_col.ReadDocument(ctx, cred_edge, &edge)
	if err != nil {
		ctrl.log.Debug("Error getting Credentials Edge(Link)", zap.Error(err))
		return key, false
	}
	ctrl.log.Debug("Found Credentials Edge(Link)", zap.Any("edge", edge))

	cred, ok := credentials.Determine(auth_type)
	if !ok {
		return key, false
	}

	key = edge.To.Key()
	ctrl.log.Debug("Looking for Credentials", zap.Any("key", key))
	err = cred.FindByKey(ctx, ctrl.cred, key)
	if err != nil {
		ctrl.log.Debug("Error getting Credentials by Key", zap.Error(err))
		return key, false
	}
	return key, true
}

// Return Account authorisable by this Credentials
func Authorisable(ctx context.Context, cred *credentials.Credentials, db driver.Database) (Account, bool) {
	query := `FOR account IN 1 INBOUND @credentials GRAPH @credentials_graph RETURN account`
	c, err := db.Query(ctx, query, map[string]interface{}{
		"credentials": cred,
		"credentials_graph": schema.CREDENTIALS_GRAPH.Name,
	})
	if err != nil {
		return Account{}, false
	}
	defer c.Close()

	var r Account
	_, err = c.ReadDocument(ctx, &r)
	return r, err == nil
}

func (ctrl *AccountsController) Authorize(ctx context.Context, auth_type string, args ...string) (Account, bool) {
	ctrl.log.Debug("Authorization request", zap.String("type", auth_type))

	credentials, err := credentials.Find(ctx, ctrl.col.Database(), ctrl.log, auth_type, args...)
	// Check if could authorize
	if err != nil {
		ctrl.log.Info("Coudn't authorize", zap.Error(err))
		return Account{}, false
	}

	account, ok := Authorisable(ctx, &credentials, ctrl.col.Database())
	ctrl.log.Debug("Authorized account", zap.Bool("result", ok), zap.Any("account", account))
	return account, ok
}

func (ctrl *AccountsController) EnsureRootExists(passwd string) (err error) {
	exists, err := ctrl.col.DocumentExists(context.TODO(), "0")
	if err != nil {
		return err
	}

	var meta driver.DocumentMeta
	if !exists {
		meta, err = ctrl.col.CreateDocument(context.TODO(), Account{ 
			Title: "nocloud",
			DocumentMeta: driver.DocumentMeta { Key: "0" },
		})
		if err != nil {
			return err
		}
		ctrl.log.Debug("Created root Account", zap.Any("result", meta))
	}
	var root Account
	meta, err = ctrl.col.ReadDocument(context.TODO(), "0", &root)
	if err != nil {
		return err
	}
	root.DocumentMeta = meta

	ns_col, _ := ctrl.col.Database().Collection(context.TODO(), schema.NAMESPACES_COL)
	exists, err = ns_col.DocumentExists(context.TODO(), "0")
	if !exists {
		meta, err := ns_col.CreateDocument(context.TODO(), Namespace{ 
			Title: "platform",
			DocumentMeta: driver.DocumentMeta { Key: "0" },
		})
		if err != nil {
			return err
		}
		ctrl.log.Debug("Created root Namespace", zap.Any("result", meta))
	}

	var rootNS Namespace
	meta, err = ns_col.ReadDocument(context.TODO(), "0", &rootNS)
	if err != nil {
		return err
	}
	rootNS.DocumentMeta = meta

	edge_col, _ := ctrl.col.Database().Collection(context.TODO(), schema.ACC2NS)
	exists, err = edge_col.DocumentExists(context.TODO(), "0-0")
	if !exists {
		err = root.LinkNamespace(context.TODO(), edge_col, rootNS, 4, roles.OWNER)
		if err != nil {
			return err
		}
	}

	ctx := context.WithValue(context.Background(), nocloud.NoCloudAccount, "0")
	cred_edge_col, _ := ctrl.col.Database().Collection(context.TODO(), schema.ACC2CRED)
	cred, err := credentials.NewStandardCredentials("nocloud", passwd)
	if err != nil {
		return err
	}

	exists, err = cred_edge_col.DocumentExists(context.TODO(), "standard-0")
	if !exists {
		err = ctrl.SetCredentials(ctx, root, cred_edge_col, cred)
		if err != nil {
			return err
		}
	}
	_, r := ctrl.Authorize(ctx, "standard", "nocloud", passwd)
	if !r {
		return errors.New("Cannot authorize nocloud")
	}
	return nil
}