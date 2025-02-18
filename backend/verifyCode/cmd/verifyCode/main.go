package main

import (
	"flag"
<<<<<<< HEAD
	"os"
=======
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"os"
	"time"

>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	"verifyCode/internal/conf"

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
	Name string = "VerifyCode"
	// Version is the version of the compiled software.
	Version string = "1.0.0"
	// flagconf is the config flag.
	flagconf string

	//id, _ = os.Hostname()
	// 使用唯一的uuid，作为id
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

	// 设置meta属性，设置weight
	mate := map[string]string{
		"weight": "999",
	}
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
<<<<<<< HEAD
		kratos.Metadata(map[string]string{}),
=======
		// 设置元数据
		kratos.Metadata(mate),
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
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

<<<<<<< HEAD
	//启动之前重新设定全局rand函数随机种子	//h5fk5h
	//rand.Seed(time.Now().UnixNano()) //纳秒级别时间戳，一定时间重新生成种子，避免出现服务重启后更换代码
	//rand.New(rand.NewSource(time.Now().UnixNano()))
	//seed := 1
	//r := rand.New(rand.NewSource(int64(seed)))
	//fmt.Println(r.Uint64())
	//fmt.Println(r.Uint64())
	// start and wait for stop signal	开始并等待停止信号
=======
	// set rand seed
	rand.Seed(time.Now().UnixNano())

	// start and wait for stop signal
>>>>>>> f088d708a9cd6384c0a30c0669e6062431292c3e
	if err := app.Run(); err != nil {
		panic(err)
	}
}
