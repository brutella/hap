package hap

import (
	"net"
)

type listener struct {
	*net.TCPListener
}

func (ln *listener) Accept() (con net.Conn, err error) {
	con, err = ln.AcceptTCP()
	if err != nil {
		return
	}

	conn := newConn(con)
	setConn(conn.RemoteAddr().String(), conn)

	return conn, err
}

func (ln *listener) Close() error {
	return ln.TCPListener.Close()
}

func (ln *listener) Addr() net.Addr {
	return ln.TCPListener.Addr()
}
