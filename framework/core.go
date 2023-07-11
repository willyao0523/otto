package framework

import "net/http"

// framework core
type Core struct {}

// init framework core
func NewCore() *Core {
	return &Core{}
}

// handlers
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}