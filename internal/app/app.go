package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wazwki/api-grpc/db/postgres"
	"github.com/wazwki/api-grpc/db/redis"
	"github.com/wazwki/api-grpc/internal/config"
	"github.com/wazwki/api-grpc/internal/controllers/grpc_c"
	v1 "github.com/wazwki/api-grpc/internal/controllers/grpc_c/v1"
	"github.com/wazwki/api-grpc/internal/repository"
	"github.com/wazwki/api-grpc/internal/service"
	"github.com/wazwki/api-grpc/pkg/jwtutil"
	"github.com/wazwki/api-grpc/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	migrateDSN string
	pool       *pgxpool.Pool
	grpcServer *grpc.Server
	grpcAddr   string
	httpServer *http.Server
}

func New(cfg *config.Config) (*App, error) {
	logger.LogInit(cfg.LogLevel)
	logger.Info("App started", zap.String("module", "name"))

	pool, err := postgres.ConnectPool(cfg.DBdsn)
	if err != nil {
		logger.Error("Fail connect pool", zap.Error(err), zap.String("module", "name"))
		return nil, err
	}

	redisClient, err := redis.Config(cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword, cfg.RedisDBNumber)
	if err != nil {
		logger.Error("Fail connect to redis", zap.Error(err), zap.String("module", "name"))
		return nil, err
	}

	nameRepository := repository.NewNameRepository(pool, redisClient)
	nameService := service.NewNameService(nameRepository)
	nameControllers := v1.NewNameControllers(nameService)

	jwt := jwtutil.NewJWTUtil(jwtutil.Config{
		AccessTokenSecret:  []byte(cfg.AccessTokenSecret),
		RefreshTokenSecret: []byte(cfg.RefreshTokenSecret),
		AccessTokenTTL:     time.Duration(cfg.AccessTokenTTL) * time.Second,
		RefreshTokenTTL:    time.Duration(cfg.RefreshTokenTTL) * time.Second,
	})

	grpcServer, srv, err := grpc_c.NewGRPCServer(cfg, nameControllers, jwt)
	if err != nil {
		logger.Error("Fail to create grpc server", zap.Error(err), zap.String("module", "name"))
		return nil, err
	}

	return &App{migrateDSN: cfg.DBdsn, pool: pool, grpcServer: grpcServer, httpServer: srv, grpcAddr: fmt.Sprintf("%v:%v", cfg.Host, cfg.GRPCPort)}, nil
}

func (a *App) Run() error {
	if err := postgres.RunMigrations(a.migrateDSN); err != nil {
		logger.Error("Fail migrate", zap.Error(err), zap.String("module", "name"))
		return err
	}
	logger.Info("Migrations applied", zap.String("module", "name"))

	lis, err := net.Listen("tcp", a.grpcAddr)
	if err != nil {
		logger.Error("Fail to listen", zap.Error(err), zap.String("module", "name"))
		return err
	}

	go func() {
		logger.Info("Starting grpc server", zap.String("module", "name"))
		if err := a.grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logger.Error("Fail to start grpc server", zap.Error(err), zap.String("module", "name"))
		}
	}()

	go func() {
		logger.Info("Starting http server", zap.String("module", "name"))
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Fail to start http server", zap.Error(err), zap.String("module", "name"))
		}
	}()

	return nil
}

func (a *App) Stop() error {
	logger.Info("App stopping", zap.String("module", "name"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.pool.Close()
	logger.Info("Pool closed", zap.String("module", "name"))

	a.grpcServer.GracefulStop()
	logger.Info("Grpc server stopped", zap.String("module", "name"))

	if err := a.httpServer.Shutdown(ctx); err != nil {
		logger.Error("Fail to stop http server", zap.Error(err), zap.String("module", "name"))
		return err
	}
	logger.Info("Http server stopped", zap.String("module", "name"))

	logger.Info("App stopped", zap.String("module", "name"))

	return nil
}
