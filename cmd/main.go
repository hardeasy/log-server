package main

import (
	"flag"
	"log-server/config"
	"log-server/internal/dao"
	"log-server/internal/server/http"
	"log-server/internal/service"
	"os"
	"os/signal"
	"syscall"
)

var pid = 0

//func preInit() bool {
//	pidfile := config.Conf.Application.PidPath
//	if exists, _ := utils.FileIsExists(pidfile); !exists {
//		fd, err := os.Create(pidfile)
//		defer fd.Close()
//		if err != nil {
//			fmt.Println("create pid file error.")
//			return false
//		}
//		pid = os.Getpid()
//		_, err = fd.Write([]byte(fmt.Sprintf("%d", pid)))
//		if err != nil {
//			fmt.Println("write pid failed.")
//			return false
//		}
//		return true;
//	}
//	//
//	fd, err := os.OpenFile(pidfile, os.O_RDWR, 0)
//	defer fd.Close()
//	pd, _ := ioutil.ReadAll(fd)
//	old_pid, err := strconv.Atoi(string(pd))
//	if err == nil {
//		pro,_ := os.FindProcess(old_pid);
//		err := pro.Signal(syscall.Signal(0x0))
//		if err == nil {
//			fmt.Println("server is already running.")
//			return false
//		}
//	}
//
//	return true
//}

func main() {
	flag.Parse()
	err := config.Init()
	if err != nil {
		panic(err)
	}
	//if (!preInit()) {
	//	return;
	//}
	d := dao.New(config.Conf)
	servs := service.NewService(config.Conf, d)
	httpServer := http.NewHttpServer(config.Conf, servs)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	go func() {
		servs.Stop()
	}()
	httpServer.Shutdown();
}
