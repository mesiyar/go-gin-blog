package main

import (
	"fmt"
	"gin-blog/src/gin-blog/pkg/setting"
	"gin-blog/src/gin-blog/routers"
	"net/http"
)

func main()  {
	router := routers.InitRouter()
	// 启动一个http服务
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
