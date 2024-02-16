package main

import (
	"fmt"
	"log"
	"net"
)

const (
	PORT 	= 	":3000"
	DOMAIN 	= 	"localhost"
	TYPE	=	"tcp"
)

type Server struct {
	ADDR		string
	DOM			string
	TYPE		string

	ListofUsers	[]net.Conn
}

func (s *Server) InitServer(port string, domain string, typeCon string){
	s.ADDR 	= 	port
	s.DOM	=	domain
	s.TYPE	= 	typeCon
}

func handleClient(conn net.Conn) {
	fmt.Println("~ A Client Connected")
	conn.Write([]byte("~ Connected To Server"))
}

func main() {
	fmt.Println("!! Starting Server")
	var server Server
	server.InitServer(PORT, DOMAIN, TYPE)

	listner, err := net.Listen(server.TYPE, server.DOM + server.ADDR)
	if err != nil {
		log.Fatal(err)
	}
	// defer listner.Close()

	for {
		
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		server.ListofUsers = append(server.ListofUsers, conn)
		
		
		go handleClient(conn)
	}
}
