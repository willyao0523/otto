package otto

import (
	"net/http"
	"otto/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr: ":8080",
	}
	server.ListenAndServe()
}