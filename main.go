package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/model"
	"go-blog/internal/routers"
	"go-blog/pkg/logger"
	"go-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSettting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	fmt.Printf("%s", global.ServerSetting)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}

func setupSetting() error {
	settingVal, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settingVal.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settingVal.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settingVal.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExit
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
