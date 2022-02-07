package hap

import (
	"net"
)

type Listener struct {
	*net.TCPListener
}

func (ln *Listener) Accept() (con net.Conn, err error) {
	con, err = ln.AcceptTCP()
	if err != nil {
		return
	}

	conn := &Conn{Conn: con}
	SetConn(conn.RemoteAddr().String(), conn)

	return conn, err
}

func (ln *Listener) Close() error {
	return ln.TCPListener.Close()
}

func (ln *Listener) Addr() net.Addr {
	return ln.TCPListener.Addr()
}
