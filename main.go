package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karasuneo/go-ci-build/routers"
)

func main() {
	engine := gin.Default()
	// route.goのDefineRoutesを呼び出す
	routers.DefineRoutes(engine)
	engine.Run(":3000")
}
