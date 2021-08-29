package server_test

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"testing"
)

func TestLongConnectionServerMyself(t *testing.T) {
	fmt.Println("服务端")
	server, err := NewLongConnServer("123", "127.0.0.1", "5050")
	if err != nil {
		t.Fatal(err)
	}
	err = server.CreateTcpListering()
	if err != nil {
		t.Fatal(err)
	}
	// data := []string{"a","b"}
	// data = append(data[:1],data[2:]...)  //测试data[2:]...会不会因为超过范围报错
	// fmt.Println(data)
}

type LongConnServer struct {
	LogId       string
	Address     string //IP地址
	Host        string //端口号
	MaxConnSize int32  //最大连接数
	clientConns []*net.TCPConn
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

func (s *LongConnServer) CreateTcpListering() error {
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
		s.addClientConn(clientConn)
		go s.tcpPipe(clientConn)
	}
}

// is capable to deal with the tcp connection
//tips: using this function should be create a runtinue
func (s *LongConnServer) tcpPipe(clientConn *net.TCPConn) {
	ipAddress := clientConn.RemoteAddr().String()
	defer func() {
		clientConn.Close()
		_ = s.deleteConn(clientConn)
		fmt.Printf("%s %s has left\n", s.LogId, ipAddress)
	}()
	fmt.Printf("%s %s come in\n", s.LogId, ipAddress)
	reader := bufio.NewReader(clientConn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		message = fmt.Sprintf("%s say: %s", ipAddress, message)
		fmt.Print(message)
		s.broadcast(message)
	}
}

//add new tcp connection to the slice
func (s *LongConnServer) addClientConn(tcpConn *net.TCPConn) error {
	if tcpConn == nil {
		return errors.New("TCPConn should not be empty")
	}
	s.clientConns = append(s.clientConns, tcpConn)
	return nil
}

//broadcast information to other clients
func (s *LongConnServer) broadcast(message string) error {
	var msgBytes = []byte(message)
	for _, client := range s.clientConns {
		_, err := client.Write(msgBytes)
		if err != nil {
			fmt.Printf("sent to %v was failed, err: %v", client.RemoteAddr(), err)
		}
	}
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
			fmt.Printf("%s client %d was delete successful\n", s.LogId, index)
			break
		}
	}
	return nil
}
