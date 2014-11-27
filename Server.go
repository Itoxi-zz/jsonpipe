package jsonpipe

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
)

type Server struct {
	ActionRegistry map[string]Action
	Reader         *bufio.Reader
	Encoder        *json.Encoder
}

type Message struct {
	Connection net.Conn
	Data       []byte
}

type Action struct {
	Handler Handler
	Pattern string
}

func NewServer() *Server {
	server := Server{
		ActionRegistry: make(map[string]Action),
	}
	return &server
}

func (s Server) Handle(action string, handler Handler) {
	if len(action) < 1 {
		log.Println("Error registering handler: pattern string is required")
	}
	s.ActionRegistry[action] = Action{Pattern: action, Handler: handler}
}

func (s Server) ListenAndServe(port string) {

	allClients := make(map[net.Conn]string) //map of all clients keyed on their connection
	newConnections := make(chan net.Conn)   //channel for incoming connections
	deadConnections := make(chan net.Conn)  //channel for dead connections
	messages := make(chan Message)          //channel for messages

	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("JSON Pipe Server listening on %s\n", port)

	go acceptConnections(server, newConnections)

	for {
		select {
		case conn := <-newConnections:
			addr := conn.RemoteAddr().String()
			log.Printf("Accepted new client, %v", addr)
			allClients[conn] = addr
			go read(conn, messages, deadConnections)
		case conn := <-deadConnections:
			log.Printf("Client %v disconnected", allClients[conn])
			delete(allClients, conn)
		case message := <-messages:
			go s.HandleRequest(message)
		}
	}

}

func acceptConnections(server net.Listener, newConnections chan net.Conn) {
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
		}
		newConnections <- conn
	}
}

func read(conn net.Conn, messages chan Message, deadConnections chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		incoming, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		messages <- Message{conn, []byte(incoming)}
	}
	deadConnections <- conn
}

func (server Server) HandleRequest(msg Message) {

	response := Response{}
	request := Request{}

	if err := json.Unmarshal(msg.Data, &request); err != nil {
		log.Println("Error decoding JSON:" + err.Error())
	}

	if action, ok := server.ActionRegistry[request.Action]; ok { //Get the handler for this action
		response = action.Handler.Run(&request)
	} else {
		response.Error = errors.New(fmt.Sprintf("No handler registered for %s", request.Action))
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling JSON:%s\n", err)
		return
	}

	msg.Connection.Write(bytes)

	return
}
