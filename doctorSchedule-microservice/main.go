package main

import (
    "fmt"
    "net/http"
    "doctorSchedule-microservice/controller"
)

func main() {
    http.HandleFunc("/", controller.HelloHandler)
    fmt.Println("Server running on port 8082")
    http.ListenAndServe(":8082", nil)
}