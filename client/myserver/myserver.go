package myserver

import (
	"fmt"
	"net"
	"sync"
)

type MyServer struct {
	Listener Reciever
	Dialer   Sender
}

type Reciever struct {
	PORT   string
	DOMAIN string
	TYPE   string

	Listener    net.Listener
	Connections []net.Conn

	wg *sync.WaitGroup
}

type Sender struct {
	TARGET_PORT   string
	TARGET_DOMAIN string
	TARGET_TYPE   string

	Connection net.Conn
}

func InitListener(p string, d string, t string, wg *sync.WaitGroup) Reciever {
	var l Reciever

	l.PORT = p
	l.DOMAIN = d
	l.TYPE = t
	l.wg = wg

	return l
}

func (recv *Reciever) CloseListener() {
	fmt.Println("~ Closing Listener")
	recv.Listener.Close()
	recv.wg.Done()
}

func (recv *Reciever) StartListener() {
	var err error

	recv.Listener, err = net.Listen(recv.TYPE, recv.DOMAIN+recv.PORT)
	if err != nil {
		fmt.Printf("Could Not Listen To Connections at %s:%s\n", recv.DOMAIN, recv.PORT)
		panic(" Stoping PKr ")
	}
	defer recv.CloseListener()

	fmt.Printf("~ Listener Running on %s%s\n", recv.DOMAIN, recv.PORT)
	for {
		connection, err := recv.Listener.Accept()
		if err != nil {
			fmt.Println("Could Not Accept A Connection :")
			fmt.Println(err.Error())
		}

		recv.Connections = append(recv.Connections, connection)
		go handleInitConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 2048)
		len, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("~ Error Reading From Connection")
			break
		}

		fmt.Print(string(buffer[:len]))
	}
}

func InitSender(p string, d string, t string) *Sender {
	var sender Sender
	var err error
	sender.TARGET_PORT = p
	sender.TARGET_DOMAIN = d
	sender.TARGET_TYPE = t

	sender.Connection, err = net.Dial(sender.TARGET_TYPE, sender.TARGET_DOMAIN+sender.TARGET_PORT)
	if err != nil {
		fmt.Println("error dialing connection")
		fmt.Println("please check if entered port, domain, type is correct...")
		return nil
	}

	return &sender
}

func (s *Sender) SendBytes(data []byte) error {
	_, err := s.Connection.Write(data)
	return err
}
