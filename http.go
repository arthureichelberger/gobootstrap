package gobootstrap

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type RestHandler func(g *gin.Engine)

func InitHTTP(cfg Config, handler RestHandler) error {
	engine := gin.New()
	engine.Use(logger.SetLogger())
	handler(engine)

	return engine.Run()
}
