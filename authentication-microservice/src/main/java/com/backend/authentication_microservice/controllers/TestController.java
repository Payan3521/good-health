package com.backend.authentication_microservice;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {
    @GetMapping("/")
    public String hello() {
        return "Authentication Microservice is running with Java and Springboot";
    }
}