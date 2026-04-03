package main

import (
    "doctorSchedule-microservice/internal/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    handler := handlers.NewTestHandler()
    router.GET("/", handler.TestHandler)
    router.Run(":8081")
}