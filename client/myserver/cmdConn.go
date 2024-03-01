package myserver

import (
	"ButterHost69/PKr-client/encrypt"
	"ButterHost69/PKr-client/models"
	"ButterHost69/PKr-client/myserver/pb"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	CMD_DIAL_CONNECTION_IP string
)

type CmdServer struct{
	pb.UnimplementedCmdConnectionServer
}

// -------------------------------------------------------
// ---------------     Command Listener    ------------
func (l *Listener) closeGRPCCmdConnectionListener(){
	fmt.Println("~ Closing GRPC Cmd Connection ...")
	// l.wg.Done()
	l.listn.Close()
}

func (l *Listener) StartGRPCCmdConnection() {
	var err error
	l.listn, err = net.Listen(l.CONTYPE, l.DOMAIN + l.PORT)

	if err != nil {
		fmt.Printf("error occured in listening to %s:%s\n", l.DOMAIN, l.PORT)
		fmt.Println(err)
		defer l.closeGRPCCmdConnectionListener()
		return 
	}

	fmt.Println("Starting CMD Connection Server...")
	g := grpc.NewServer()
	pb.RegisterCmdConnectionServer(g, &CmdServer{})

	go func() {
		if err := g.Serve(l.listn); err != nil {
			fmt.Println("error could not start grpc Sever")
		}
	}()
}

func (c *CmdServer) ExecuteCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error){
	pass, err := encrypt.DecryptData(in.ConnPassword)
	if err != nil {
		fmt.Println("error occured in decrypting password")
		fmt.Println(err)
		return &pb.CommandResponse{
			Ack: false,
			RsPort: err.Error(),
		}, err
	}
	ifConnExits := models.ValidateConnection(in.ConnSlug,pass)
	if !ifConnExits {
		return &pb.CommandResponse{
			Ack: ifConnExits,
			RsPort: "0",
		}, nil
	}
	
	return &pb.CommandResponse{
		Ack: true,
		RsPort: "0",
	}, nil
}

// -------------------------------------------------------
// ---------------       Command Request      ------------
func (s *Sender) closeGRPCCmdConnectionSender() {
	fmt.Println("~ Closing GRPC Command Connection Sender ...")
	
	s.GRPCConnection.Close()
}

func (s *Sender) DialCommandConnection(conn models.Connections) {
	var err error
	CMD_DIAL_CONNECTION_IP = s.TARGET_DOMAIN + ":" + s.TARGET_PORT
	s.GRPCConnection, err = grpc.Dial(CMD_DIAL_CONNECTION_IP, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error occured in dialing grpc for cmd conncection")
		fmt.Printf("target Ip: %v", s.TARGET_DOMAIN)
		fmt.Println(err)
		return
	}
	defer s.closeGRPCCmdConnectionSender()

	c := pb.NewCmdConnectionClient(s.GRPCConnection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	ifReqSent := sendExecuteCommandRequest(ctx, c, conn)
	if !ifReqSent {
		fmt.Println("error: could not send request to get files")
		fmt.Println("please try again...")
		return
	}

}

func sendExecuteCommandRequest(ctx context.Context, c pb.CmdConnectionClient, conn models.Connections) bool {
	publicKey := "tmp/connections/" + conn.ConnectionSlug + "/publickey.pem"
	pemBlock := encrypt.GetPublicKey(publicKey)
	encryptPass, err := encrypt.EncryptData(conn.Password, pemBlock)
	if err != nil {
		fmt.Println("error in encrypting data")
		fmt.Println(err)
		return false
	}
	response, err := c.ExecuteCommand(
		ctx,
		&pb.CommandRequest{
			CommandType:  "GET",
			ConnSlug:     conn.ConnectionSlug,
			ConnPassword: encryptPass,
		},
	)
	if err != nil {
		fmt.Println("error in recieving Command Connection Request")
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(response.RsPort)
	return response.Ack
}
