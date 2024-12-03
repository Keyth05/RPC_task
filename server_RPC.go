// server.go
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Define the structure type to be used in RPC
type HelloWorld struct{}

// Define the method that will be called remotely
func (h *HelloWorld) SayHello(request string, reply *string) error {
	*reply = "Hello " + request + " from RPC!"
	return nil
}

func main() {
	// Start an RPC server
	hello := new(HelloWorld)
	rpc.Register(hello) // Registrar el servicio

	// Open a TCP port for the server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Server is running on port 1234...")

	// Listen for client requests
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn) // Handle each client connection
	}
}
