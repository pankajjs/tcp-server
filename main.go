package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	// start the process
	// listens to a port for tcp protocol
	listener, err:= net.Listen("tcp", ":8081") 

	if err != nil{ 
		log.Fatal(err)
	}

	for {
		fmt.Println("Waiting for client to connect")
		conn, err:= listener.Accept() // blocking call

		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println("Client connected")
		fmt.Print(conn)
		// create a go routine/thread for each request to handle concurrent requests.
		go do(conn)
	}
}

func do(conn net.Conn) {
	buff:= make([]byte, 1024)

	_, err:= conn.Read(buff) // blocking call
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Processing request....")
	time.Sleep(8*time.Second)
	
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n")) // blocking call
	conn.Close()
}