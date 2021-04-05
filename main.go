package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/api", server.AddMiddleware(HandleRoot, CheckAuth()))
	server.Handle("POST", "/api", server.AddMiddleware(HandleRoot, CheckAuth()))
	server.Handle("POST", "/api/req", PostRequest)
	server.Listen()

}
