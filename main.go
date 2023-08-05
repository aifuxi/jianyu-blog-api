package main

import (
	"fmt"

	"github.com/aifuxi/jianyu-blog-api/models"
	"github.com/aifuxi/jianyu-blog-api/pkg/logging"
	"github.com/aifuxi/jianyu-blog-api/pkg/setting"
	"github.com/aifuxi/jianyu-blog-api/routers"
)

func main() {
	setting.Setup()
	logging.Setup()
	models.Setup()

	router := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()

	router.Run(fmt.Sprintf(":%d", setting.ServerSetting.HttpPort))
}
