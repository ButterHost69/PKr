package myserver

import (
	"ButterHost69/PKr-client/utils"
	"fmt"
	"net"
)

func (recv *Reciever) StartInitListener() {
	var err error

	recv.Listener, err = net.Listen(recv.TYPE, recv.DOMAIN+recv.PORT)
	if err != nil {
		fmt.Printf("Could Not Listen To Connections at %s:%s\n", recv.DOMAIN, recv.PORT)
		panic(" Stoping PKr ")
	}
	defer recv.CloseListener()

	otp := utils.CreateOTP(5)

	fmt.Printf("~ Listening On  %s%s... \n~ Your OTP is: %d", recv.DOMAIN, recv.PORT, otp)
	connection, err := recv.Listener.Accept()
	if err != nil {
		fmt.Println("Could Not Accept A Connection :")
		fmt.Println(err.Error())
	}
	
	handleInitConnection(connection)

}

func handleInitConnection(conn net.Conn) {
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

func (s *Sender) SendCertificate(data []byte) error {
	_, err := s.Connection.Write(data)
	return err
}
