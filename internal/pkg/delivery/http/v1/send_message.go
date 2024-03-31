package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

// SendMessage godoc
// @Summary      Send message
// @Description  splits the message and sends segments to the channel layer
// @Tags         message
// @Accept       json
// @Produce      json
// @Param        message   body      entity.Message  true  "Message"
// @Success      200  {object}  responseOk{body=nil}
// @Failure      400  {object}  responseError
// @Failure      500  {object}  responseError
// @Router       /message/send [post]
func (h *Handler) SendMessage(c *gin.Context) {
	_ = entity.Message{}
	c.JSON(http.StatusInternalServerError, responseError{Status: "error", Code: "unimplemented", Message: "method send message is not implemented"})
}
