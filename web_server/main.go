package main

func main() {
	server := NewServer(":3000")
	server.Handler("GET", "/", HandleRoot)
	server.Handler("GET", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handler("POST", "/create", PostRequest)
	server.Handler("POST", "/user", UserPostRequest)
	server.Listen()
}
