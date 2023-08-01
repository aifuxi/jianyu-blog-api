package main

import (
	"fmt"
	"net/http"

	"github.com/aifuxi/jianyu-blog-api/pkg/setting"
	"github.com/aifuxi/jianyu-blog-api/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
