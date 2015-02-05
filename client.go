package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
)

func main() {

	config := &tls.Config{
		ServerName: "my-tls-server",
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "10.0.2.15:443", config)
	if err != nil {
		panic(err)
	}
	err = conn.Handshake()
	if err != nil {
		fmt.Printf("Failed handshake:%v\n", err)
		return
	}

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Printf("Failed receiving data:%v\n", err)
	}

	conn.Close()
}
