package model

import (
	"errors"
	"fmt"
	"net"

	"github.com/fangzhixi/fizzy-server/server/data"
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

func (s *LongConnServer) CreateTcpListering(network data.Network, ipAddress, port string) error {

	address := fmt.Sprintf("%s:%s", ipAddress, port)
	fmt.Printf("%s listening address is %s", s.LogId, address)
	tcpAddress, err := net.ResolveTCPAddr(string(network), address)
	if err != nil {
		fmt.Printf("%s net.ResolveTcpAddr function error: %v\n", s.LogId, err)
		return err
	}
	tcpListener, err := net.ListenTCP(string(network), tcpAddress)
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
		s.addClientConn(clientConn)
		go s.TcpPipe(clientConn)
	}
}

// is capable to deal with the tcp connection
//tips: using this function should be create a runtinue
func (s *LongConnServer) TcpPipe(clientConn *net.TCPConn) {
	ipAddress := clientConn.RemoteAddr().String()
	defer func() {
		clientConn.Close()
		_ = s.deleteConn(clientConn)
	}()
	fmt.Printf("%s %s", s.LogId, ipAddress)
}

//add new tcp connection to the slice
func (s *LongConnServer) addClientConn(tcpConn *net.TCPConn) error {
	if tcpConn == nil {
		return errors.New("TCPConn should not be empty")
	}
	s.clientConns = append(s.clientConns, tcpConn)
	return nil
}

//delete disconnected records in slice
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
