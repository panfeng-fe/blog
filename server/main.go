package main

import (
	"blog/conf"
	"blog/route"
	"blog/utils"
	"github.com/kataras/iris/v12/middleware/logger"
	"io"
	"os"

	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/middleware/cors"
)

func main() {
	appConf := conf.Config.App
	logFile := utils.NewLogFile()
	defer logFile.Close()

	app := iris.New()

	app.Logger().SetOutput(io.MultiWriter(logFile, os.Stdout))

	requestLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
	})

	app.Use(requestLogger)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
		Debug:            false,
		AllowCredentials: true,
	})

	app.Configure(iris.WithConfiguration(iris.Configuration{
		Tunneling:                         iris.TunnelingConfiguration{},
		IgnoreServerErrors:                nil,
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		DisablePathCorrectionRedirection:  false,
		EnablePathEscape:                  false,
		EnableOptimizations:               false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "utf-8",
		PostMaxMemory:                     0,
		LocaleContextKey:                  "",
		ViewLayoutContextKey:              "",
		ViewDataContextKey:                "",
		RemoteAddrHeaders:                 nil,
		Other:                             nil,
	}))

	router := app.Party("/", crs).AllowMethods(iris.MethodOptions)

	// 路由注册
	route.RouterInit(router)

	app.Run(
		iris.Addr(appConf.Port),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
