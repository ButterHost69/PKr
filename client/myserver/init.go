package myserver

import (
	pb "ButterHost69/PKr-client/myserver/pb"
	"ButterHost69/PKr-client/utils"
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

type InitServer struct {
	pb.UnimplementedInitConnectionServer
	OTP pb.OTP
}

var (
	verficatonOTP int32
)

func (is *InitServer) VerifyOTP(ctx context.Context, in *pb.OTP) (*pb.OTPResponse, error) {
	if in.Otp == verficatonOTP {
		fmt.Printf("%v is now recognized as a user: %v \n", in.IpAddress, in.Username)
		return &pb.OTPResponse{
			IfOtpCorrect: true,
		}, nil
	}

	fmt.Println("~ Incorrect OTP Entered Closing the Connecion")

	return &pb.OTPResponse{
		IfOtpCorrect: false,
	}, nil
}

func (l *Listener) closeGRPCInitConnectionListener() {
	fmt.Println("~ Closing GRPC Init Connection ...")
	l.wg.Done()
	l.listn.Close()
}

func (l *Listener) StartGRPCInitConnection() {

	var err error
	l.listn, err = net.Listen(l.CONTYPE, l.DOMAIN+l.PORT)
	defer l.closeGRPCInitConnectionListener()

	if err != nil {
		fmt.Println("error occured at starting listener")
		fmt.Println(err.Error())
		return
	}

	verficatonOTP = utils.CreateOTP(5)
	fmt.Printf(" Your OTP is: %v\n", verficatonOTP)

	g := grpc.NewServer()
	pb.RegisterInitConnectionServer(g, &InitServer{})
	

	if err := g.Serve(l.listn); err != nil {
		fmt.Println("error could not start grpc Sever")
		return
	}

}

func sendOTP(ctx context.Context, c pb.InitConnectionClient) bool {
	var otp int32
	fmt.Print("Enter OTP: ")
	fmt.Scan(&otp)

	response, err := c.VerifyOTP(
		ctx,
		&pb.OTP{
			Username:  "Testing",
			IpAddress: "temp",
			Otp:       otp,
		},
	)
	if err != nil {
		fmt.Println("error in recieving OTP ~")
		fmt.Println(err.Error())

		return false
	}

	return response.IfOtpCorrect
}

func (s *Sender) closeGRPCInitConnectionSender() {
	fmt.Println("~ Closing GRPC Init Connection Sender ...")
	s.wg.Done()
	s.GRPCConnection.Close()
}

func (s *Sender) DialGRPCInitConnection() {
	var err error
	
	s.GRPCConnection, err = grpc.Dial(s.TARGET_DOMAIN+s.TARGET_PORT, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error in Dialing Connection to: %s:%s\nPlease Check IF The IP and PORT is Entered Correctly or not...\n", s.TARGET_DOMAIN, s.TARGET_PORT)
		fmt.Println(err.Error())
		return
	}
	defer s.closeGRPCInitConnectionSender()

	c := pb.NewInitConnectionClient(s.GRPCConnection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	for {
		ifnot := sendOTP(ctx, c)
		if ifnot {
			fmt.Println("Connecting....")
			break
		}
	}
}
