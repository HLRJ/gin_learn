package main

import (
	"HLRJ/gin_learn/common"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	//现在已经没有close方法了
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
