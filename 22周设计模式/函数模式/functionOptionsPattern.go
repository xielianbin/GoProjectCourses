
package main

import "fmt"

// Server配置结构体
type Server struct {
	host     string
	port     int
	protocol string
	timeout  int
	maxConns int
}

// Option函数类型
type Option func(*Server)

// 定义各种配置选项函数
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.protocol = protocol
	}
}

func WithTimeout(timeout int) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConns(maxConns int) Option {
	return func(s *Server) {
		s.maxConns = maxConns
	}
}

// 创建Server实例，应用所有选项
func NewServer(opts ...Option) *Server {
	// 设置默认值
	s := &Server{
		host:     "localhost",
		port:     8080,
		protocol: "http",
		timeout:  30,
		maxConns: 100,
	}

	// 应用所有选项
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func main() {
	// 使用默认配置创建服务器
	server1 := NewServer()
	fmt.Printf("%+v\n", server1)

	// 使用自定义配置创建服务器
	server2 := NewServer(
		WithHost("example.com"),
		WithPort(443),
		WithProtocol("https"),
		WithTimeout(60),
		WithMaxConns(500),
	)
	fmt.Printf("%+v\n", server2)
}
