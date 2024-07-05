package network

import (
	"bufio"
	"fmt"
	"io"
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

	reader := bufio.NewReader(conn)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}

	message := string(content)

	fmt.Printf("Recieved: %q\n", message)

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
