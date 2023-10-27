package network

import (
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
