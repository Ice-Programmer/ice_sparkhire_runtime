package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	kitexserver "github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"ice_sparkhire_runtime/gateway"
	init_ "ice_sparkhire_runtime/init"
	sparkruntime "ice_sparkhire_runtime/kitex_gen/sparkhire_runtime/sparkhireruntimeservice"
	"io"
	"os"
	"path"
	"sync"
	"time"

	"log"
	"net"
)

const (
	LogFileAddr = "./logs/"
)

func main() {
	ctx := context.Background()

	if err := initLog(); err != nil {
		log.Fatalf("init_ log error: %v", err)
	}
	if err := init_.Init(ctx); err != nil {
		klog.CtxErrorf(ctx, "Init log error: %v", err)
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := StartKitexServer(); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := gateway.StartHttpServer(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func StartKitexServer() error {

	svr := sparkruntime.NewServer(new(SparkhireRuntimeServiceImpl),
		kitexserver.WithServiceAddr(&net.TCPAddr{Port: 8888}),
		kitexserver.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		kitexserver.WithMetaHandler(transmeta.ServerHTTP2Handler),
		kitexserver.WithMiddleware(gateway.AuthInterceptor()),
	)
	err := svr.Run()
	if err != nil {
		log.Println("Thrift server stopped:", err.Error())
		return err
	}
	return nil
}

func initLog() error {
	// 创建日志目录
	if err := os.MkdirAll(LogFileAddr, 0o755); err != nil {
		return err
	}

	// 生成日志文件名（按天）
	logFileName := time.Now().Format("2003-10-21") + ".log"
	fileName := path.Join(LogFileAddr, logFileName)

	// 配置lumberjack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // 单个文件最大20M
		MaxBackups: 5,    // 保留5个备份
		MaxAge:     10,   // 日志保存10天
		Compress:   true, // 压缩备份
	}

	logger := kitexlogrus.NewLogger()

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(klog.LevelDebug)
	klog.SetLogger(logger)

	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logger.SetOutput(multiWriter)

	logger.SetLevel(klog.LevelDebug)
	klog.SetLogger(logger)
	return nil
}
