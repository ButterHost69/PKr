package myserver

import (
	"net"
	"sync"

	"google.golang.org/grpc"
)

type Listener struct {
	DOMAIN  string
	PORT    string
	CONTYPE string

	listn net.Listener
	wg    *sync.WaitGroup
}

type Sender struct {
	TARGET_DOMAIN	string
	TARGET_PORT		string
	TARGET_CONNTYPE	string

	Connection		net.Conn
	GRPCConnection	*grpc.ClientConn
	wg 				*sync.WaitGroup
}

func InitListener(d string, p string, ct string) Listener {
	return Listener{
		DOMAIN:  d,
		PORT:    p,
		CONTYPE: ct,
		// wg:      wg,
	}
}
	
func InitWGListener(d string, p string, ct string, wg *sync.WaitGroup) Listener {
	return Listener{
		DOMAIN:  d,
		PORT:    p,
		CONTYPE: ct,
		wg:      wg,
	}
}

func InitSender(d string, p string, ct string, wg *sync.WaitGroup) Sender {
	return Sender{
		TARGET_DOMAIN:  d,
		TARGET_PORT:    p,
		TARGET_CONNTYPE: ct,
		wg:      wg,
	}
}