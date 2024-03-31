package http

import (
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
	_ = entity.Segment{}
	c.JSON(http.StatusInternalServerError, responseError{Status: "error", Code: "unimplemented", Message: "method put segment is not implemented"})
}
