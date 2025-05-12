package main

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/postgress"
	"net"
	"os"
	"os/signal"

	"gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/config"
	"gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/logger"
	test "gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/pkg/api/api/test"

	"gitlab.crja72.ru/golang/2025/spring/course/students/308006-dsycev62-course-1343/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	ctx, _ = logger.New(ctx)

	cfg, err := config.New()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "init config success", zap.String("grpc-port", fmt.Sprintf("%v", cfg.GRPCPort)),
		zap.String("host", fmt.Sprintf("%v", cfg.Postgres.Host)))
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to load config", zap.Error(err))
	}

	pool, err := postgress.New(ctx, cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect to postgres", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to listen", zap.Error(err))
	}

	srv := service.New()
	server := grpc.NewServer(grpc.UnaryInterceptor(logger.Interceptor))
	test.RegisterOrderServiceServer(server, srv)

	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to serve", zap.Error(err))
		}
	}()

	logger.GetLoggerFromCtx(ctx).Info(ctx, "everything is fine!")
	//mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//grpcServerEndpoint := flag.String("grpc-server-endpoint", fmt.Sprintf("localhost:%d", cfg.GRPCPort), "gRPC server endpoint")
	//if err := test.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
	//	logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to register gateway", zap.Error(err))
	//}
	//
	//go func() {
	//	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", cfg.RestPort), mux); err != nil {
	//		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to listen and serve gateway", zap.Error(err))
	//	}
	//}()

	select {
	case <-ctx.Done():
		server.GracefulStop()
		pool.Close(ctx)
		logger.GetLoggerFromCtx(ctx).Info(ctx, "Server Stopped")
	}
	// err = test.RegisterOrderServiceServer()
	// err = gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
}
