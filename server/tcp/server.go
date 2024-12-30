package tcp

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/30 20:27:10
* @description:
**/

type Server struct {
	Name           string
	Addr           string
	MaxConnections int
	inShutdown     atomic.Bool
	activeConns    map[Connection]struct{}
	onShutdown     []func()
}

var ErrServerClosed = errors.New("tcp: server closed")

func (srv *Server) serve(ln net.Listener) error {
	l := onceCloseListener{Listener: ln}
	defer l.Close()

	for {

	}
}

func (srv *Server) shuttingDown() bool {
	return srv.inShutdown.Load()
}

type onceCloseListener struct {
	net.Listener
	once sync.Once
	err  error
}

func (ocl *onceCloseListener) Close() error {
	ocl.once.Do(ocl.close)
	return ocl.err
}

func (ocl *onceCloseListener) close() {
	ocl.err = ocl.Listener.Close()
}
