package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/liluhao/ginstarter/pkg/server"

	"go.uber.org/zap"

	"github.com/liluhao/ginstarter/pkg/middleware"
	"github.com/liluhao/ginstarter/pkg/validation"
	"github.com/liluhao/lib/loglib"
	"github.com/liluhao/lib/migrationlib"
	"github.com/liluhao/lib/pglib"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"

	"github.com/gin-gonic/gin"
	"github.com/liluhao/ginstarter/pkg/dao"
	"github.com/liluhao/ginstarter/pkg/service"

	"github.com/jessevdk/go-flags"
)

//以下Tag的规定都在由"github.com/jessevdk/go-flags"所规定
//long：选项的长名称;
//description：选项的描述（可选）;
//required：如果非空，则使该选项在命令中出现。如果该选项不存在，解析器将返回 ErrRequired（可选）;
//default:默认值;
//env:如果已经定义了环境变量，则该选项的默认值会从指定的环境变量中重写。(可选)
//group: 在结构字段上指定时，使结构字段成为具有给定名称的单独组（可选）;
//namesapce：在组结构字段上指定名称空间时，名称空间被置于每个选项的长名称和该组子组的名称空间的前面，由解析器的namespace分隔符分隔（可选）。
//env-namespace: 当在组结构字段上指定时，env-namespace被前缀到每个选项的env键和该组的子组的env-namespace，以解析器的env-namespace分隔符(可选)
type PostgresConfig struct {
	URL              string `long:"url" description:"database url" env:"URL" required:"true"`
	PoolSize         int    `long:"pool-size" description:"database pool size" env:"POOL_SIZE" default:"10"`
	MigrationFileDir string `long:"migration-file-dir" description:"migration file dir" env:"MIGRATION_FILE_DIR" default:"file://migrations"`
}

type GinConfig struct {
	Port string `long:"port" description:"port" env:"PORT" default:":8080"`
	Mode string `long:"mode" description:"mode" env:"MODE" default:"debug"`
}

type Environment struct {
	GinConfig      GinConfig      `group:"gin" namespace:"Gin" env-namespace:"GIN"`
	PostgresConfig PostgresConfig `group:"postgres" namespace:"postgres" env-namespace:"POSTGRES"`
}

func main() {
	var env Environment
	parser := flags.NewParser(&env, flags.Default) //创建一个新的解析器；第二个参数是为解析器指定一组选项
	if _, err := parser.Parse(); err != nil {      //解析来自os.Args 的命令行参数
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0) //Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。
		} else {
			os.Exit(1)
		}
	}

	//数据库信息
	migrationLib := migrationlib.NewMigrateLib(migrationlib.Config{
		DatabaseDriver: migrationlib.PostgresDriver,
		DatabaseURL:    env.PostgresConfig.URL,
		SourceDriver:   migrationlib.FileDriver,
		SourceURL:      env.PostgresConfig.MigrationFileDir,
		TableName:      "migrate_version",
	})
	if err := migrationLib.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("run database migration fail:%v", err)
	}
	pgClient, err := pglib.NewDefaultGOPGClient(pglib.GOPGConfig{
		URL:       env.PostgresConfig.URL,
		DebugMode: false,
		PoolSize:  env.PostgresConfig.PoolSize,
	})

	//创建日志logger
	logger, err := loglib.NewProductionLogger()
	if err != nil {
		log.Fatalf("fail to init logger:%v", err)
	}

	memberDAO := dao.NewPGMemberDAO(logger, pgClient)

	//创建校验器
	bindingValidator, _ := binding.Validator.Engine().(*validator.Validate)
	CustomValidator, err := validation.NewValidationTranslator(bindingValidator, "en")
	if err != nil {
		log.Fatalf("fail to init validation translator:%v", err)
	}

	svc := service.NewService(memberDAO)
	mwe := middleware.NewMiddleware(logger, CustomValidator)
	gin.SetMode(env.GinConfig.Mode)
	GracefulRun(logger, StartFunc(logger, server.NewHTTPServer(gin.Default(), env.GinConfig.Port, mwe, svc)))
}

func GracefulRun(logger *loglib.Logger, fn func(ctx context.Context) error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1) //设置error类型有缓冲通道

	go func() {
		done <- fn(ctx) //传入子goroutine
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-done:
		return
	case <-shutdown:
		cancel()
		timeoutCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		select {
		case <-done:
			return
		case <-timeoutCtx.Done():
			logger.Error("shutdown timeout", zap.Error(timeoutCtx.Err()))
			return
		}
	}
}

func StartFunc(logger *loglib.Logger, server *http.Server) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("http server listen error", zap.Error(err))
			}
		}()

		<-ctx.Done()

		ctx1, cancel1 := context.WithCancel(context.Background())
		go func() {
			logger.Info("shutdown http server...")
			server.Shutdown(ctx1)
			cancel1()
		}()
		<-ctx1.Done()
		logger.Info("http server existing")
		return nil
	}
}
