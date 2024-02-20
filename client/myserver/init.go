package myserver

import (
	pb "ButterHost69/PKr-client/myserver/pb"
	"ButterHost69/PKr-client/utils"
	"context"
	"fmt"
	"net"

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

func (l *Listener) closeGRPCInitConnection() {
	fmt.Println("~ Closing GRPC Init Connection ...")
	l.wg.Done()
	l.listn.Close()
}

func (l *Listener) StartGRPCInitConnection() {
	defer l.closeGRPCInitConnection()

	var err error
	l.listn, err = net.Listen(l.CONTYPE, l.DOMAIN+l.PORT)

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
