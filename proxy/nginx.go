package main


// This is the proxy struct
// This object shields the application object(actual object) by checking rate of queries, if too many requests are made to the url then it blocks the request
// But if the rate of queries is ok then it will run the actual application's handleRequest method.
// This struct's handleRequest method just does extra bit of checking before doing the actual request handling

type Nginx struct {
    application *Application
    maxAllowedRequest int
    rateLimiter map[string]int
}


func newNginxServer() *Nginx {
    return &Nginx{
        application: &Application{},
        maxAllowedRequest: 2,
        rateLimiter: make(map[string]int),
    }
}


func (n *Nginx) handleRequest(url, method string)(int, string){
    allowed := n.checkRateLimiting(url)
    if !allowed {
        return 403, "Not Allowed"
    }

    return n.application.handleRequest(url, method)
}


func (n *Nginx) checkRateLimiting(url string) bool {
    if n.rateLimiter[url] == 0 {
        n.rateLimiter[url] = 1
    }

    if n.rateLimiter[url] > n.maxAllowedRequest {
        return false
    }

    n.rateLimiter[url] = n.rateLimiter[url] + 1
    return true
}
