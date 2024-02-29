package myserver

import (
	"ButterHost69/PKr-client/encrypt"
	"ButterHost69/PKr-client/models"
	"ButterHost69/PKr-client/myserver/pb"
	"context"
	"fmt"
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
func (l *Listener) ListenCommandConnection() {

}

// func (c *CmdServer) ExecuteCommand(ctx context.Context, in *pb.CommandRequest) (*pb.CommandResponse, error){

// }

// -------------------------------------------------------
// ---------------       Command Request      ------------
func (s *Sender) closeGRPCCmdConnectionSender() {
	fmt.Println("~ Closing GRPC Command Connection Sender ...")
	s.wg.Done()
	s.GRPCConnection.Close()
}

func (s *Sender) DialCommandConnection(conn models.Connections) {
	var err error
	CMD_DIAL_CONNECTION_IP = s.TARGET_DOMAIN
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
	publicKey := "tmp/connections/" + CONNECTION_SLUG + "/publickey.pem"
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

	return response.Ack
}
