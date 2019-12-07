package server

import (
	"net"
)

type Pool struct {
	Links map[uint16]net.Conn
	LinkNum uint16
	LinkMax uint16
}

func NewPool() *Pool {
	return &Pool{
		Links : make(map[uint16]net.Conn),
		LinkNum : 0,
		LinkMax : 1024,
	}
}

func (this *Pool) Set(conn net.Conn) uint16{
	this.LinkNum++
	this.Links[this.LinkNum] = conn
	return this.LinkNum
}

func (this *Pool) Get(index uint16) net.Conn{
	return this.Links[index]
}

func (this *Pool) Del(index uint16) {
	this.Links[index] = nil
}