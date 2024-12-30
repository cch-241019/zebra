package tcp

import (
	"context"
	"net"
)

/*
* @author: Chen Chiheng
* @date: 2024/12/30 20:38:34
* @description:
**/

type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateClosed
)

type Connection interface {
	Serve(ctx context.Context, rwc net.Conn)
	SetState(state ConnState, runHook bool)
	GetState() ConnState
}

type Conn struct {
	rwc net.Conn
}
