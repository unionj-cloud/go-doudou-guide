package main

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/framework/logger"
	"os"
	service "statsvc"
	"statsvc/config"
	"statsvc/internal/reportsvcj"
	"statsvc/transport/httpsrv"
)

func main() {
	conf := config.LoadFromEnv()

	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"localhost",
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        6060,
		ServiceName: os.Getenv("GDD_SERVICE_NAME"),
		Weight:      5,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"goVer": "go1.17"},
	})
	if err != nil {
		panic(err)
	}
	if success {
		logger.Info("joined nacos default cluster successfully")
	}
	defer func() {
		success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          "localhost",
			Port:        6060,
			ServiceName: os.Getenv("GDD_SERVICE_NAME"),
			Ephemeral:   true,
		})
		if err != nil {
			logger.Error(err)
		}
		if success {
			logger.Info("left nacos default cluster successfully")
		}
	}()

	svc := service.NewStatsvc(conf, reportsvcj.NewEcho(ddhttp.WithRootPath("/report-svc-j"), ddhttp.WithProvider(ddhttp.NewNacosWRRServiceProvider(namingClient, "report-svc-j"))))
	handler := httpsrv.NewStatsvcHandler(svc)
	srv := ddhttp.NewDefaultHttpSrv()
	srv.AddRoute(httpsrv.Routes(handler)...)
	srv.Run()
}
