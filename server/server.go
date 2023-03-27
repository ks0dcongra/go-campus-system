package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

// 服务器证书的使用
func main() {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}
	// 需要双向认证，服务器也需要验证客户端的真实性。
	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	// Package crypto/x509提供了证书管理的相关操作。
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}
	// 我们创建的服务器私钥和pem文件中得到证书cert，并且生成一个tls.Config对象。这个对象有多个字段可以设置，本例中我们使用它的默认值。
	// 因为需要验证客户端，我们需要额外配置下面两个字段：ClientAuth與ClientCAs
	config := &tls.Config{
		// Note: if there are multiple Certificates, and they don't have the
		// optional field Leaf set, certificate selection will incur a significant
		// per-handshake performance cost.
		Certificates: []tls.Certificate{cert},
		// ClientAuth determines the server's policy for
		// TLS Client Authentication. The default is NoClientCert.
		ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientCAs defines the set of root certificate authorities
		// that servers use if required to verify a client certificate
		// by the policy in ClientAuth.
		ClientCAs: clientCertPool,
	}
	// tls.Listen开始监听客户端的连接，accept后得到一个net.Conn，后续处理和普通的TCP程序一样。
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
