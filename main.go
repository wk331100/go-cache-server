package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	go_cache "github.com/wk331100/go-cache"
	"github.com/wk331100/go-cache-server/module"
	"github.com/wk331100/go-cache-server/types"
)

var server = go_cache.NewCache()

func main() {
	config, err := module.ParseConfig()
	if err != nil {
		log.Println("error parsing config:", err)
		return
	}

	address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	// 监听Redis客户端连接
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {
			log.Fatal("error closing listener:", err)
			return
		}
	}(listener)

	fmt.Println("go cache server listening on ", address)

	// 接受并处理客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

// 处理客户端连接
func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("error closing connection:", err)
			return
		}
	}(conn)

	// 创建读写缓冲区
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// 读取客户端请求
		request, err := module.ReadArray(reader)
		if err != nil {
			log.Println("error reading request:", err)
			return
		}
		if len(request) == 0 {
			continue
		}

		// 处理请求,发送响应给客户端
		response, err := processRequest(request)
		if err != nil {
			module.WriteError(writer, err.Error())
		} else if response == "" {
			module.WriteSimpleString(writer, types.OK)
		} else {
			module.WriteBulkString(writer, response)
		}

		if err = writer.Flush(); err != nil {
			log.Println("error flushing writer:", err)
			return
		}
	}
}

// 处理客户端请求
func processRequest(request []string) (string, error) {
	fmt.Println("request", request)
	// 获取命令和参数
	command := strings.ToUpper(request[0])
	args := request[1:]
	handler, exist := commandList[command]
	if !exist {
		return "", types.ErrInvalidCommand
	}
	return handler.Handle(server, args)
}
