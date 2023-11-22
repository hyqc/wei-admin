package serve

import (
	"admin/config"
	"admin/router"
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"os"
	"time"
)

func Run() {
	ser := manualStart()

	ser = micro.NewService(
		micro.Name(config.AppConfig.Server.Name),
		micro.Address(config.AppConfig.Server.Port),
		micro.Version(config.AppVersion),
		micro.Server(http.NewServer()),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(15*time.Second),
		micro.HandleSignal(true),
		micro.Context(context.Background()),
		micro.Client(grpc.NewClient(
			client.DialTimeout(30*time.Second),
			client.RequestTimeout(30*time.Second),
			client.Retries(0),
			client.PoolSize(1000),
		)),
	)

	ser.Init(

		micro.BeforeStart(func() error {
			fmt.Println("=================", ser.Options().Server.Options().Address)
			return nil
		}),
		micro.AfterStart(func() error {
			fmt.Println("---------------", ser.Options().Transport.Options().Addrs, ser.Options().Server.Options().Address)
			return nil
		}),
	)

	if err := micro.RegisterHandler(ser.Server(), router.Engine); err != nil {
		logger.Errorf("register http handler failed, error: %v", err)
		return
	}

	if err := ser.Run(); err != nil {
		logger.Errorf("start micro service failed, error: %v", err)
		return
	}
}

func manualStart() micro.Service {
	ops := []micro.Option{
		beforeStart(), afterStart(),
	}
	ops = append(ops, config.AppFlags.Init()...)

	ser := micro.NewService(
		ops...,
	)

	// 手动解析cmd
	if err := ser.Options().Cmd.App().Run(os.Args); err != nil {
		logger.Errorf("cmd run failed, error: %v", err)
		os.Exit(1)
		return nil
	}

	if err := config.AppConfig.Init(); err != nil {
		logger.Errorf("config init failed, error: %v", err)
		os.Exit(2)
		return nil
	}

	logger.Infof("config info: %v", config.AppConfig.Server)
	return ser
}

func afterStart() micro.Option {
	return micro.AfterStart(func() error {
		logger.Infof("server port:%s", config.AppConfig.Server)
		return nil
	})
}

func beforeStart() micro.Option {
	return micro.BeforeStart(func() error {
		if err := config.AppConfig.Init(); err != nil {
			logger.Errorf("config init failed, error: %v", err)
			return err
		}
		return nil
	})
}
