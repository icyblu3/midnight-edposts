package main

import (
	"fmt"
	"log"
	"net"
	"encoding/hex"
)

func main() {
	// - Create TCP socket
	// - Establish connection to server
	serverConn, err := net.Dial("tcp4", net.JoinHostPort("localhost", "16800"))
	if err != nil {
		log.Fatalln(err)
	}
	// Close conn when done
	defer serverConn.Close()
	
	for {
		/* TODO: should probably figure out what stuff to send to the server here */
		// Server expects a 32byte-key from the host
		inBytes, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
		_, err := serverConn.Write(inBytes)
        if err != nil {
                log.Fatalln("Error sending guess:  ", err)
        }
		GetServerReply(serverConn)
		break
	}
}

func GetServerReply(serverConn net.Conn) {
	buf := make([]byte, 8) 
	// Server will always send a reply with 8 bytes
	_, err := serverConn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))
}