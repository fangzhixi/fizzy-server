package server

import (
	"errors"
	"fmt"
	"net"
)

const (
	TCP  string = "tcp"
	TCP4 string = "tcp4"
	TCP6 string = "tcp6"
)

type LongConnServer struct {
	LogId       string
	Address     string         //IP地址
	Host        string         //端口号
	MaxConnSize int32          //最大连接数
	clientConns []*net.TCPConn //现有链接用户数

}

//NewServer mean create a new server that can keep a long time to connection the client
func NewLongConnServer(logId, address, host string, maxClinetLongConnection ...int32) (server *LongConnServer, err error) {

	var maxConnSize int32 = 0
	if logId == "" || address == "" || host == "" {
		return nil, errors.New("logId or address or host should not be null")
	} else if len(maxClinetLongConnection) > 0 {
		maxConnSize = maxClinetLongConnection[0]
	}
	return &LongConnServer{LogId: logId, Address: address, Host: host, MaxConnSize: maxConnSize}, nil
}

func (s *LongConnServer) createTcpListering() error {
	var (
		network = "tcp"
		host    = "127.0.0.1:5050"
	)
	tcpAddress, err := net.ResolveTCPAddr(network, host)
	if err != nil {
		fmt.Printf("%s net.ResolveTcpAddr function error: %v\n", s.LogId, err)
		return err
	}
	tcpListener, err := net.ListenTCP(network, tcpAddress)
	if err != nil {
		fmt.Printf("%s net.ListenTCP function error: %v\n", s.LogId, err)
		return err
	}
	defer tcpListener.Close()
	for {
		fmt.Printf("%s blocking... until listen client get in\n", s.LogId)
		clientConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Printf("%s net.Accept function error: %v\n", s.LogId, err)
			continue
		}

		fmt.Printf("%s When listen to the client access, open a coroutines to processing work\n", s.LogId)
		fmt.Printf("%s add the client in slice first\n", s.LogId)
		s.clientConns = append(s.clientConns, clientConn)
		go s.tcpPipe(clientConn)
	}
}

func (s *LongConnServer) tcpPipe(clientConn *net.TCPConn) {
	ipAddress := clientConn.RemoteAddr().String()
	defer func() {
		clientConn.Close()
		_ = s.deleteConn(clientConn)
	}()
	fmt.Printf("%s %s", s.LogId, ipAddress)
}

func (s *LongConnServer) addClientConn(tcpConn *net.TCPConn) error {
	if tcpConn == nil {
		return errors.New("TCPConn should not be empty")
	}
	s.clientConns = append(s.clientConns, tcpConn)
	return nil
}

func (s *LongConnServer) deleteConn(tcpConn *net.TCPConn) error {
	if tcpConn == nil {
		return errors.New("TCPConn should not be empty")
	}
	for index, item := range s.clientConns {
		if tcpConn == item {
			s.clientConns = append(s.clientConns[:index], s.clientConns[index+1:]...)
			fmt.Printf("%s client %d was delete successful", s.LogId, index)
			break
		}
	}
	return nil
}
