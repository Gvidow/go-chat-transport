package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

// PutSegment godoc
// @Summary      Put a segment in the queue
// @Description  puts a message segment in the queue
// @Tags         segment
// @Accept       json
// @Produce      json
// @Param        segment   body      entity.Segment  true  "Segment"
// @Success      200  {object}  int
// @Failure      400  {object}  responseError
// @Failure      500  {object}  responseError
// @Router       /segment/put [post]
func (h *Handler) PutSegment(c *gin.Context) {
	segment := &entity.Segment{}

	err := json.NewDecoder(c.Request.Body).Decode(segment)
	defer c.Request.Body.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, responseError{Status: "error", Code: "parse_body", Message: "the request body could not be parsed segment"})
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(h.baseCtx, _timeoutPutSegment)
		defer cancel()

		err := h.segmentUsecase.SaveToQueue(ctx, segment)
		if err != nil {
			h.log.Error(err.Error())
		}
	}()

	c.JSON(http.StatusOK, responseOk{Status: "ok"})
}
