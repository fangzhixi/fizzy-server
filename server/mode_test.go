package server_test

import (
	"bufio"
	"errors"
	"fmt"
	"learn/long_connection/server"
	"net"
	"testing"
)

var connSlice []*net.TCPConn

// 创建TCP长连接服务
func createTcp() {
	tcpAdd, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5050") //解析tcp服务
	if err != nil {
		fmt.Println("net.ResolveTCPAddr error:", err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAdd) //监听指定TCP服务
	if err != nil {
		fmt.Println("net.ListenTCP error:", err)
		return
	}
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP() //阻塞，当有客户端连接时，才会运行下面
		if err != nil {
			fmt.Println("tcpListener error :", err)
			continue
		}
		fmt.Println("A client connected:", tcpConn.RemoteAddr().String())
		boradcastMessage(tcpConn.RemoteAddr().String() + "进入房间" + "\n") //当有一个客户端进来之时，广播某某进入房间
		connSlice = append(connSlice, tcpConn)
		// 监听到被访问时，开一个协程处理
		go tcpPipe(tcpConn)
	}
}

// 对客户端作出反应
func tcpPipe(conn *net.TCPConn) {
	ipAddress := conn.RemoteAddr().String()
	fmt.Println("ipAddress:", ipAddress)
	defer func() {
		fmt.Println("disconnected:", ipAddress)
		conn.Close()
		deleteConn(conn)
		boradcastMessage(ipAddress + "离开了房间" + "\n")
	}()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n') //读取直到输入中第一次发生 ‘\n’
		//因为按强制退出的时候，他就先发送换行，然后在结束
		if err != nil || message == "\n" {
			return
		}
		message = ipAddress + "说：" + message
		if err != nil {
			fmt.Println("topPipe:", err)
			return
		}
		// 广播消息
		fmt.Println(ipAddress, "说：", message)
		err = boradcastMessage(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// 广播数据
func boradcastMessage(message string) error {
	b := []byte(message)
	for i := 0; i < len(connSlice); i++ {
		fmt.Println(connSlice[i])
		_, err := connSlice[i].Write(b)
		if err != nil {
			fmt.Println("发送给", connSlice[i].RemoteAddr().String(), "数据失败"+err.Error())
			continue
		}
	}
	return nil
}

// 移除已经关闭的客户端
func deleteConn(conn *net.TCPConn) error {
	if conn == nil {
		fmt.Println("conn is nil")
		return errors.New("conn is nil")
	}
	for i := 0; i < len(connSlice); i++ {
		if connSlice[i] == conn {
			connSlice = append(connSlice[:i], connSlice[i+1:]...)
			break
		}
	}
	return nil
}

func TestLongConnectionServer(t *testing.T) {
	fmt.Println("服务端")
	createTcp()
	// data := []string{"a","b"}
	// data = append(data[:1],data[2:]...)  //测试data[2:]...会不会因为超过范围报错
	// fmt.Println(data)
}

func TestServer(t *testing.T) {
	fmt.Println("创建新服务端...")
	server := server.NewLongConnServer()
	err := server.createTcpListering()
	if err != nil {
		t.Fatal(err)
	}
}
