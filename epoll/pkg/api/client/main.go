/**
 * @author zhangyuehao
 * @date 2019-03-13 14:31
 */

package main

import (
	"flag"
	"fmt"
	"github.com/eechains/epoll/pkg/core/logger"
	"net"
	"time"
)

var (
	log         = logger.InitLogger()
	ip          = flag.String("ip", "127.0.0.1", "evio_server IP")
	connections = flag.Int("conn", 1, "number of tcp connections")
)

func main() {
	flag.Parse()
	addr := *ip + ":36137"
	log.Infof("连接到 %s", addr)
	var conns []net.Conn
	for i := 0; i < *connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 1*time.Second)
		if err != nil {
			fmt.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(time.Millisecond)
	}
	defer func() {
		for _, c := range conns {
			c.Close()
		}
	}()
	log.Infof("完成初始化 %d 连接", len(conns))
	tts := time.Second
	if *connections > 100 {
		tts = time.Millisecond * 5
	}
	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(tts)
			conn := conns[i]
			conn.Write([]byte("hello world\r\n"))
		}
	}

}
