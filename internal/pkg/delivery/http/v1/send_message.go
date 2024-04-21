package http

import (
	"context"
	"encoding/json"
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
	mes := &entity.Message{}
	err := json.NewDecoder(c.Request.Body).Decode(mes)
	defer c.Request.Body.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Status: "error", Code: "parse_body", Message: "the request body could not be parsed message"})
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(h.baseCtx, _timeoutSendMessage)
		defer cancel()

		err := h.messageUsecase.Send(ctx, mes)
		if err != nil {
			h.log.Error(err.Error())
		}
	}()

	c.JSON(http.StatusOK, responseOk{Status: "ok"})
}
