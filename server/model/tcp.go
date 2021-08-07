package model

import "net"

type LongConnServer struct {
	LogId       string
	Address     string         //IP地址
	Host        string         //端口号
	MaxConnSize int32          //最大连接数
	ClientConns []*net.TCPConn //现有链接用户数
}
