package server

import "net"

/*
* @author: Chen Chiheng
* @date: 2024/12/30 21:07:50
* @description:
**/

type Server interface {
	Serve(ln net.Listener) error
}
