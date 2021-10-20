package accounts

import (
	"context"

	"github.com/arangodb/go-driver"

	"github.com/slntopp/nocloud/pkg/accounts/accountspb"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/access"

	"go.uber.org/zap"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	jwt "github.com/dgrijalva/jwt-go"
)

type AccountsServiceServer struct {
	accountspb.UnimplementedAccountsServiceServer
	db driver.Database
	ctrl graph.AccountsController
	ns_ctrl graph.NamespacesController

	log *zap.Logger
	SIGNING_KEY []byte
}

func NewServer(log *zap.Logger, db driver.Database) *AccountsServiceServer {
	accountsCol, _ := db.Collection(nil, graph.ACCOUNTS_COL)
	credCol, _ := db.Collection(nil, graph.CREDENTIALS_COL)
	nsCol, _ := db.Collection(nil, graph.NAMESPACES_COL)

	return &AccountsServiceServer{
		log: log, db: db, 
		ctrl: graph.NewAccountsController(
			log.Named("AccountsController"), accountsCol, credCol,
		),
		ns_ctrl: graph.NewNamespacesController(
			log.Named("NamespacesController"), nsCol,
		),
	}
}

func ValidateMetadata(ctx context.Context, log *zap.Logger) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("Failed to get metadata from context")
		return ctx, status.Error(codes.Aborted, "Failed to get metadata from context")
	}

	//Check for Authentication
	requestor := md.Get(nocloud.NOCLOUD_ACCOUNT_CLAIM)
	if requestor == nil {
		log.Error("Failed to authenticate account")
		return ctx, status.Error(codes.Unauthenticated, "Not authenticated")
	}
	ctx = context.WithValue(ctx, nocloud.NoCloudAccount, requestor[0])

	return ctx, nil
}

func (s *AccountsServiceServer) Token(ctx context.Context, request *accountspb.TokenRequest) (*accountspb.TokenResponse, error) {
	log := s.log.Named("Token")

	log.Debug("Token request received", zap.Any("request", request))
	account, ok := s.ctrl.Authorize(ctx, request.Auth.Type, request.Auth.Data...)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wrong credentials given")
	}
	log.Debug("Authorized user", zap.String("ID", account.ID.String()))

	claims := jwt.MapClaims{}
	claims[nocloud.NOCLOUD_ACCOUNT_CLAIM] = account.Key
	claims["exp"] = request.Exp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(s.SIGNING_KEY)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to issue token")
	}

	return &accountspb.TokenResponse{Token: token_string}, nil
}

func (s *AccountsServiceServer) Create(ctx context.Context, request *accountspb.CreateRequest) (*accountspb.CreateResponse, error) {
	log := s.log.Named("CreateAccount")
	log.Debug("Create request received", zap.Any("request", request),  zap.Any("context", ctx))
	ctx, err := ValidateMetadata(ctx, log)
	if err != nil {
		return nil, err
	}
	requestor := ctx.Value(nocloud.NoCloudAccount).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns, err := s.ns_ctrl.Get(ctx, request.Namespace)
	if err != nil {
		s.log.Debug("Error getting namespace", zap.Error(err), zap.String("namespace", request.Namespace))
		return nil, err
	}

	ok, access_lvl := graph.AccessLevel(ctx, s.db, requestor, ns.ID.String())
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "No Access")
	} else if access_lvl < access.MGMT {
		return nil, status.Error(codes.PermissionDenied, "No Enough Rights")
	}

	account, err := s.ctrl.Create(ctx, request.Title)
	if err != nil {
		s.log.Debug("Error creating account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating account")
	}
	res := &accountspb.CreateResponse{Id: account.Key}

	if (*request.Access) < access_lvl {
		access_lvl = (*request.Access)
	}

	col, _ := s.db.Collection(ctx, graph.NS2ACC)
	err = account.JoinNamespace(ctx, col, ns, access_lvl)
	if err != nil {
		s.log.Debug("Error linking to namespace")
		return res, err
	}

	col, _ = s.db.Collection(ctx, graph.CREDENTIALS_EDGE_COL)
	cred, err := graph.MakeCredentials(*request.Auth)
	if err != nil {
		return res, status.Error(codes.Internal, err.Error())
	}

	err = s.ctrl.SetCredentials(ctx, account, col, cred)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *AccountsServiceServer) EnsureRootExists(passwd string) (error) {
	return s.ctrl.EnsureRootExists(passwd)
}
