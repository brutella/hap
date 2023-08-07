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

	// disable TCP keepalives
	if tcpconn, ok := con.(*net.TCPConn); ok {
		tcpconn.SetKeepAlive(false)
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
