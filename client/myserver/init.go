package myserver

import (
	"ButterHost69/PKr-client/encrypt"
	"ButterHost69/PKr-client/models"
	pb "ButterHost69/PKr-client/myserver/pb"
	"ButterHost69/PKr-client/utils"
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type InitServer struct {
	pb.UnimplementedInitConnectionServer
	OTP pb.OTP
}

var (
	verficatonOTP   int32
	My_Username     string
	CONNECTION_SLUG string

	VERIFY_IP          string
	DIAL_CONNECTION_IP string
)

const (
	COMMAND_CONNECTION_PORT = 8069
	PUBLIC_KEYS_PATH        = "tmp/mykeys/publickey.pem"
	PRIVATE_KEYS_PATH       = "tmp/mykeys/privatekey.pem"
)

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

// Set Workspace Folders

func setWorkSpaceFolders() {
	var opt int
	fmt.Println("1. Use Existing Workspace")
	fmt.Println("2. Create New Workspace  ")
	fmt.Scanln(&opt)
	fmt.Println("\n")

	switch opt {
	case 1:
		var workspaceName string
		fmt.Println(" Enter Existing Workspace Name: ")
		fmt.Scanln(&workspaceName)

		models.AddNewConnectionToTheWorkspace(workspaceName, CONNECTION_SLUG)
	case 2:
		var workspaceName string
		var workspacePath string

		fmt.Print(" Enter NEW Workspace Name: ")
		fmt.Scanln(&workspaceName)

		fmt.Print(" Enter Workspace Path: ")
		fmt.Scanln(&workspacePath)

		if err := models.CreateNewWorkspace(workspaceName, workspacePath, CONNECTION_SLUG); err != nil {
			fmt.Println("error occured in Creating New Workspace")
			fmt.Println(err)
			return
		}

		fmt.Println("NEW Workspace Created Successfully !!")

	}
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

// gRPC ~ GETWorkspaceinfo

// func (is *InitServer) GETWorkspaceinfo(ctx context.Context, in *pb.Workspaces) (*pb.Workspaces, error) {
// 	p, _ := peer.FromContext(ctx)
// 	incommingIP := p.Addr.String()
// 	fmt.Printf("Esatablishing Connection...\n")
// 	if incommingIP != VERIFY_IP {
// 		fmt.Println(" Init Ip and Incomming IPs Do not match...")
// 		return nil, errors.New("init ip and incomming ip's do not match")
// 	}

// 	for connection
// }

// func getWorkspaceInfoRequest(ctx context.Context, c pb.InitConnectionClient) string {

// }

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

// gRPC ~ Exchnage Certificate Implementation

func (is *InitServer) ExchangeCertificates(ctx context.Context, in *pb.Certificate) (*pb.CertificateResponse, error) {
	p, _ := peer.FromContext(ctx)
	incommingIP := p.Addr.String()
	fmt.Printf("Esatablishing Connection...\n")
	if incommingIP != VERIFY_IP {
		fmt.Println(" Init Ip and Incomming IPs Do not match...")
		return nil, errors.New("init ip and incomming ip's do not match")
	}

	myPrivateKey := loadPrivateKey()
	password, err := encrypt.DecryptData(myPrivateKey, in.ConnectionPassword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil
	}

	fmt.Printf("Your Connection password: %s\n", password)

	if err := utils.StoreInitPublicKeys(CONNECTION_SLUG, in.PublicKey); err != nil {
		fmt.Println("Error Occured In Storing Connection's Key")
		return nil, nil
	}

	fmt.Println("Keys Have Been Stored ...")

	models.AddConnectionInUserConfig(CONNECTION_SLUG, password, incommingIP)
	setWorkSpaceFolders()
	return &pb.CertificateResponse{
		CommandConnectionPort: 8069,
	}, nil
}

func sendCertificateRequest(ctx context.Context, c pb.InitConnectionClient) string {
	var password string
	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	encypPass, _ := encrypt.EncryptData(password, loadPublicOthersKey("tmp/connections/"+CONNECTION_SLUG+"/publickey.pem"))
	response, err := c.ExchangeCertificates(
		ctx,
		&pb.Certificate{
			ConnectionPassword: encypPass,
			PublicKey:          loadPublicKey(),
		},
	)
	if err != nil {
		fmt.Println("error in recieving Command Connection Port Number")
		fmt.Println(err.Error())

		return ""
	}

	cmdConnectionPort := response.CommandConnectionPort
	fmt.Printf("Command Connection Port: %d\n", cmdConnectionPort)

	models.AddConnectionInUserConfig(CONNECTION_SLUG, password, DIAL_CONNECTION_IP)
	setWorkSpaceFolders()
	return string(cmdConnectionPort)
}

func loadPrivateKey() string {
	// file, err := os.OpenFile(KEYS_PATH, os.O_RDONLY, 0444)
	// if err != nil {
	// 	fmt.Println("error in loading public key")
	// 	fmt.Println(err.Error())

	// 	return ""
	// }
	key, err := os.ReadFile(PRIVATE_KEYS_PATH)
	if err != nil {
		fmt.Println("error in reading public key")
		fmt.Println(err.Error())

		return ""
	}
	return string(key)
}

func loadPublicOthersKey(fp string) string {
	// file, err := os.OpenFile(KEYS_PATH, os.O_RDONLY, 0444)
	// if err != nil {
	// 	fmt.Println("error in loading public key")
	// 	fmt.Println(err.Error())

	// 	return ""
	// }
	key, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println("error in reading public key")
		fmt.Println(err.Error())

		return ""
	}
	return string(key)
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

func loadPublicKey() string {
	// file, err := os.OpenFile(KEYS_PATH, os.O_RDONLY, 0444)
	// if err != nil {
	// 	fmt.Println("error in loading public key")
	// 	fmt.Println(err.Error())

	// 	return ""
	// }
	key, err := os.ReadFile(PUBLIC_KEYS_PATH)
	if err != nil {
		fmt.Println("error in reading public key")
		fmt.Println(err.Error())

		return ""
	}
	return string(key)
}

func (is *InitServer) VerifyOTP(ctx context.Context, in *pb.OTP) (*pb.OTPResponse, error) {
	if in.Otp == verficatonOTP {
		p, _ := peer.FromContext(ctx)
		ip := p.Addr.String()

		fmt.Println("IP : ", ip)
		connectionSlug := utils.CreateSlug()
		CONNECTION_SLUG = connectionSlug
		fmt.Printf("%v is now recognized as user: %v \n", ip, in.Username)
		fmt.Printf("The Connection Slug is: %s\n", connectionSlug)
		VERIFY_IP = ip
		fmt.Printf("Establishing Connection...\n")

		return &pb.OTPResponse{
			IfOtpCorrect:   true,
			ConnectionSlug: connectionSlug,
			PublicKey:      loadPublicKey(),
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

func sendOTPRequest(ctx context.Context, c pb.InitConnectionClient) bool {
	var otp int32
	fmt.Print("Enter OTP: ")
	fmt.Scan(&otp)

	response, err := c.VerifyOTP(
		ctx,
		&pb.OTP{
			Username: "Testing",
			// TODO : REMOVE THE IP ADDRESS AS IP CAN BE
			//GET FROM CTX ON THE RECIEVER END
			IpAddress: "Hello", // p.LocalAddr.String()
			Otp:       otp,
		},
	)
	if err != nil {
		fmt.Println("error in recieving OTP ~")
		fmt.Println(err.Error())

		return false
	}

	ifOTPCorrect := response.IfOtpCorrect
	slug := response.ConnectionSlug
	key := response.PublicKey

	if ifOTPCorrect {
		fmt.Printf("Your Connection Slug: %s\n", slug)
		// fmt.Printf("Recievers Public Key: %s\n", key)
	}

	// If err Than Return Back to the parent function
	CONNECTION_SLUG = slug
	if err := utils.StoreInitPublicKeys(slug, key); err != nil {
		fmt.Println("Error Occured In Storing Connection's Key")
		return ifOTPCorrect
	}

	fmt.Println("Keys Have Been Stored ...")
	return ifOTPCorrect
}

func (s *Sender) closeGRPCInitConnectionSender() {
	fmt.Println("~ Closing GRPC Init Connection Sender ...")
	s.wg.Done()
	s.GRPCConnection.Close()
}

func (s *Sender) DialGRPCInitConnection() {
	var err error

	DIAL_CONNECTION_IP = s.TARGET_DOMAIN + s.TARGET_PORT
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
		ifnot := sendOTPRequest(ctx, c)
		if ifnot {
			fmt.Println("Connecting....")
			break
		}
	}

	sendCertificateRequest(ctx, c)
}
