package graph

import (
	"context"

	"github.com/arangodb/go-driver"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ACCOUNTS_COL = "Accounts"
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

func NewAccountsController(log *zap.Logger, col, cred driver.Collection) AccountsController {
	return AccountsController{log: log, col: col, cred: cred}
}

func (ctrl *AccountsController) Get(ctx context.Context, id string) (Account, error) {
	var r Account
	_, err := ctrl.col.ReadDocument(nil, id, &r)
	return r, err
}

func (ctrl *AccountsController) Exists(ctx context.Context, id string) (bool, error) {
	return ctrl.col.DocumentExists(nil, id)
}

func (ctrl *AccountsController) Create(ctx context.Context, title string) (Account, error) {
	acc := Account{
		Title: title,
	}
	meta, err := ctrl.col.CreateDocument(ctx, acc)
	acc.DocumentMeta = meta
	return acc, err
}

// Grant account access to namespace
func (acc *Account) LinkNamespace(ctx context.Context, edge driver.Collection, ns Namespace, level int8) (error) {
	_, err := edge.CreateDocument(ctx, Access{
		From: acc.ID,
		To: ns.ID,
		Level: level,
		DocumentMeta: driver.DocumentMeta {
			Key: acc.Key + "-" + ns.Key,
		},
	})
	return err
}

// Set Account Credentials, ensure account has only one credentials document linked per credentials type
func (ctrl *AccountsController) SetCredentials(ctx context.Context, acc Account, edge driver.Collection, c Credentials) (error) {
	requestor, ok := ctx.Value("account").(string)
	if !ok {
		return status.Error(codes.Internal, "Account ID is not given")
	}
	if !HasAccess(ctx, ctrl.col.Database(), requestor, acc.ID.String(), 3) {
		return status.Error(codes.PermissionDenied, "NoAccess")
	}

	cred, err := ctrl.cred.CreateDocument(ctx, c)	
	_, err = edge.CreateDocument(ctx, CredentialsLink{
		From: acc.ID,
		To: cred.ID,
		Type: c.Type(),
		DocumentMeta: driver.DocumentMeta {
			Key: c.Type() + "-" + acc.Key, // Ensure only one credentials vertex per type
		},
	})
	return err
}

func (ctrl *AccountsController) Authorize(ctx context.Context, auth_type string, args ...string) (Account, bool) {
	var credentials Credentials;
	var ok bool;
	ctrl.log.Debug("Authorization request", zap.String("type", auth_type))
	switch auth_type {
	case "standard":
		credentials = &StandardCredentials{Username: args[0]}
		ok = credentials.Find(ctx, ctrl.col.Database())
	default:
		return Account{}, false
	}
	// Check if could find Credentials
	if !ok {
		ctrl.log.Info("Coudn't find credentials", zap.Skip())
		return Account{}, false
	}

	ok = credentials.Authorize(args...)
	// Check if could authorize
	if !ok {
		ctrl.log.Info("Coudn't authorize", zap.Skip())
		return Account{}, false
	}

	account, ok := credentials.Account(ctx, ctrl.col.Database())
	ctrl.log.Debug("Authorized account", zap.Bool("result", ok), zap.Any("account", account))
	return account, ok
}

