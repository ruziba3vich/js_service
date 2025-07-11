package main

import (
	"context"
	"log"
	"net"

	"github.com/ruziba3vich/js_service/genprotos/genprotos/compiler_service"
	"github.com/ruziba3vich/js_service/internal/service"
	"github.com/ruziba3vich/js_service/pkg/config"
	logger "github.com/ruziba3vich/prodonik_lgger"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func NewNetListener(lc fx.Lifecycle, cfg *config.Config) (net.Listener, error) {
	lis, err := net.Listen("tcp", ":"+cfg.AppPort)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return lis.Close()
		},
	})

	return lis, nil
}

func NewLogger(cfg *config.Config) (*logger.Logger, error) {
	return logger.NewLogger(cfg.LogPath)
}

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			NewNetListener,
			NewLogger,
			service.NewJsExecutorServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, lis net.Listener, server *service.JsExecutorServer) {
			s := grpc.NewServer()
			compiler_service.RegisterCodeExecutorServer(s, server)

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Println("gRPC server starting . . .")
					go func() {
						if err := s.Serve(lis); err != nil {
							log.Printf("gRPC server failed: %v", err)
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Println("Stopping gRPC server...")
					s.GracefulStop()
					return nil
				},
			})
		}),
	).Run()
}
