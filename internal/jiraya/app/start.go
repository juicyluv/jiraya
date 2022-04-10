package app

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/juicyluv/jiraya/internal/jiraya/core"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw"
	"github.com/juicyluv/jiraya/internal/jiraya/interfaces/grpc_gw/protobuf"
	"github.com/juicyluv/jiraya/internal/jiraya/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
)

func (a *app) Start() error {
	db, err := postgres.New(a.cfg)

	if err != nil {
		a.logger.Error(err.Error())
	}

	c := core.New(db)

	srv := grpc_gw.New(c)

	grpcListener, err := net.Listen("tcp", a.cfg.GRPC.Port)

	if err != nil {
		a.logger.Fatal("cannot create grpc listener", zap.Error(err))
	}

	httpListener, err := net.Listen("tcp", a.cfg.Http.Port)

	if err != nil {
		a.logger.Fatal("cannot create http listener", zap.Error(err))
	}

	go func() {
		s := grpc.NewServer()
		protobuf.RegisterJirayaServer(s, srv)

		a.logger.Info("grpc server started", zap.String("address", grpcListener.Addr().String()))

		if err := s.Serve(grpcListener); err != nil {
			a.logger.Fatal("cannot run grpc server", zap.Error(err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		grpcRouter = runtime.NewServeMux(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						Multiline:       false,
						Indent:          "  ",
						AllowPartial:    false,
						UseProtoNames:   true,
						UseEnumNumbers:  false,
						EmitUnpopulated: true,
						Resolver:        nil,
					},
				},
			}),
		)
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	)

	err = protobuf.RegisterJirayaHandlerFromEndpoint(
		ctx,
		grpcRouter,
		grpcListener.Addr().String(),
		opts,
	)

	if err != nil {
		a.logger.Fatal("cannot register handler", zap.Error(err))
	}

	routesWrapper := chi.NewRouter()
	routesWrapper.Mount(`/`, grpcRouter)

	a.logger.Info("http server started", zap.String("address", a.cfg.Http.Port))

	return http.Serve(httpListener, routesWrapper)
}
