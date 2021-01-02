package http

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type EchoServer struct {
	port   int
	server *echo.Echo
}

func (e *EchoServer) RunServer() error {
	e.server = echo.New()
	e.server.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		}))

	return e.run_server()
}

func (e *EchoServer) Get(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	e.server.GET(path, handler, middleware...)
}

func (e *EchoServer) Post(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	e.server.POST(path, handler, middleware...)
}

func (e *EchoServer) Put(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	e.server.PUT(path, handler, middleware...)
}

func (e *EchoServer) Patch(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	e.server.PATCH(path, handler, middleware...)
}

func (e *EchoServer) Delete(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) {
	e.server.DELETE(path, handler, middleware...)
}

func (e *EchoServer) CloseServer() error {
	return e.server.Close()
}

func (e *EchoServer) SetPort(_port int) {
	e.port = _port
}

func (e *EchoServer) GetPort() int {
	return e.port
}

func (e *EchoServer) getAddressToStartService() string {
	var address strings.Builder

	address.WriteString(":")
	address.WriteString(strconv.Itoa(e.port))

	return address.String()
}

func (e *EchoServer) runServer() error {
	address := e.getAddressToStartService()
	return e.server.Start(address)
}

func BasicAuthorization(username string, password string) string {
	var authentication strings.Builder

	authentication.WriteString(username)
	authentication.WriteString(":")
	authentication.WriteString(password)
	authentication_basic := authentication.String()
	authentication_basic = base64.StdEncoding.EncodeToString([]byte(authentication_basic))
	authentication.Reset()
	authentication.WriteString("Basic ")
	authentication.WriteString(authentication_basic)
	authentication_basic = authentication.String()
	return authentication_basic
}
