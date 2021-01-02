package http

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type GorillaWebsocket struct {
	response_writer http.ResponseWriter
	request         *http.Request
	websocket       *websocket.Conn
}

func (g *GorillaWebsocket) OpenConnection() error {
	var error_opening_connection error
	websocket_options := websocket.Upgrader{}
	g.websocket, error_opening_connection = websocket_options.Upgrade(g.response_writer, g.request, nil)
	return error_opening_connection
}

func (g *GorillaWebsocket) CloseConnection() error {
	return g.websocket.Close()
}

func (g *GorillaWebsocket) SetHttpConfiguration(response_writer http.ResponseWriter, request *http.Request) {
	g.request = request
	g.response_writer = response_writer
}

func (g *GorillaWebsocket) ReadMessage() (WebsocketMessage, error) {
	var websocket_message WebsocketMessage
	message_type, message, read_error := g.websocket.ReadMessage()
	websocket_message.MessageType = message_type
	websocket_message.Message = message
	return websocket_message, read_error
}

func (g *GorillaWebsocket) WriteMessage(message WebsocketMessage) error {
	return g.websocket.WriteMessage(message.MessageType, message.Message)
}
