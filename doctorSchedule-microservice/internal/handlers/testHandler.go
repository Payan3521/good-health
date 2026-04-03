package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type testHandler struct {

}

func NewTestHandler() *testHandler {
	return &testHandler{}
}

func (h *testHandler) TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Doctor Schedule Microservice is running with Go and GIN"})
}