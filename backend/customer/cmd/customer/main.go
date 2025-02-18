package main

import (
	"flag"
<<<<<<< HEAD
=======
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	"os"

	"customer/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
<<<<<<< HEAD

	_ "go.uber.org/automaxprocs"
=======
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
<<<<<<< HEAD
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
=======
	Name string = "Customer"
	// Version is the version of the compiled software.
	Version string = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	//id, _ = os.Hostname()
	id = Name + "-" + uuid.NewString()
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
<<<<<<< HEAD
=======
	// 一，获取consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "localhost:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 二，获取consul注册管理器
	reg := consul.New(consulClient)

	// 初始化 TP
	tracerURL := "http://localhost:14268/api/traces"
	if err := initTracer(tracerURL); err != nil {
		log.Error(err)
	}

>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
<<<<<<< HEAD
=======
		// 三，创建服务时，指定服务器注册
		kratos.Registrar(reg),
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
<<<<<<< HEAD
=======

// 初始化Tracer
// @param url string 指定 Jaeger 的API接口
// :14268/api/traces
func initTracer(url string) error {
	//一，建立jaeger客户端，称之为：exporter，导出器
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	// 创建 TracerProvider
	tracerProvider := traceSDK.NewTracerProvider(
		//采样设置
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		// 使用 exporter 作为批处理程序
		traceSDK.WithBatcher(exporter),
		// 将当前服务的信息，作为资源告知给TracerProvider
		traceSDK.WithResource(resource.NewSchemaless(
			// 必要的配置
			semconv.ServiceNameKey.String(Name),
			// 任意的其他属性配置
			attribute.String("exporter", "jaeger"),
		)),
	)

	// 三，设置全局的TP
	otel.SetTracerProvider(tracerProvider)

	return nil
}
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
