package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}

	// Request a greeting from the server
	var reply string
	err = client.Call("HelloWorld.SayHello", "Keyth", &reply)
	if err != nil {
		fmt.Println("Error calling SayHello:", err)
		return
	}

	// Display the received message
	fmt.Println("Server response:", reply)
}
