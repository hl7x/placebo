package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func SendClient(ip string, port string, data string) error {
	
	serverAddr := ip + ":" + port

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return err
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

func RequestHandler(conn net.Conn) {
	
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil { 
		fmt.Println("Reading Error: ", err)
	} else {
		fmt.Printf("Message received: %v\n", message)
	}

}

func ListenClient(port string) error {
	
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Printf("Listening on TCP port %v\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go RequestHandler(conn)
	}

	return nil
}
