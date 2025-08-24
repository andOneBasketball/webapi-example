package router

import (
	"webapi-example/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Setup 初始化Router
func initRouter(debug bool) *gin.Engine {
	//设置启动模式
	switch debug {
	case true:
		gin.SetMode(gin.DebugMode)
	case false:
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	return r
}

func Run(addr string, debug bool) {
	r := initRouter(debug)

	initWebRouter(r)
	err := r.Run(addr)
	if err != nil {
		logger.Log.Fatal("run web server error: ", zap.Any("err", err))
	}
}
