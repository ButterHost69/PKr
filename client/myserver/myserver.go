package myserver

import (
	"net"
	"sync"
)

type Listener struct {
	DOMAIN  string
	PORT    string
	CONTYPE string

	listn	net.Listener
	wg 		*sync.WaitGroup
}



func InitListener(d string, p string, ct string, wg *sync.WaitGroup) Listener {
	return Listener{
		DOMAIN: d,
		PORT: p,
		CONTYPE: ct,
		wg: wg,
	}	
}