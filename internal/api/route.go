package api

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	http "github.com/gvidow/go-chat-transport/internal/pkg/delivery/http/v1"
	"github.com/gvidow/go-chat-transport/internal/server"
)

func RegistryHandler(serv *server.Server, h *http.Handler) error {
	if serv == nil {
		return ErrServerNil
	}
	if h == nil {
		return ErrHandlerNil
	}

	v1 := serv.Group("/api/v1")
	{
		v1.POST("/message/send", h.SendMessage)
		v1.POST("/segment/put", h.PutSegment)

		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	}
	return nil
}
