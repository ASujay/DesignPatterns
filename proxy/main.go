package main

import (
    "fmt"
)

func main() {
    nginxServer := newNginxServer()
    appStatusURL := "/app/status"
    createuserURL := "/create/user"

    httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
    fmt.Printf("\nURL: %s\nHTTP CODE: %d\nBODY: %s\n", appStatusURL, httpCode, body)

    httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
    fmt.Printf("\nURL: %s\nHTTP CODE: %d\nBODY: %s\n", appStatusURL, httpCode, body)

    httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
    fmt.Printf("\nURL: %s\nHTTP CODE: %d\nBODY: %s\n", appStatusURL, httpCode, body)

    httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
    fmt.Printf("\nURL: %s\nHTTP CODE: %d\nBODY: %s\n", createuserURL, httpCode, body)

    httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
    fmt.Printf("\nURL: %s\nHTTP CODE: %d\nBODY: %s\n", createuserURL, httpCode, body)

}
