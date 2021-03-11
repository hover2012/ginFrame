package main

import (
	"fmt"
	"gin/pkg/setting"
	"gin/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d",setting.HTTPPort)

	server := endless.NewServer(endPoint,routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d",syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server err : %v",err)
	}

	/*
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//s.ListenAndServe()

	go func() {
		if err := s.ListenAndServe();err !=nil{
			logs.Info("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<- quit
	log.Println("shut down server .....")

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx);err != nil {
		log.Fatal("server shutdown :",err)
	}
	log.Println("Server exiting")
 */


}