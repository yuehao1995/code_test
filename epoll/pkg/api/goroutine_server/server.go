/**
 * @author zhangyuehao
 * @date 2019-03-13 15:24
 */

package goroutine_server

import (
	"io"
	"log"
	"net"
	"os"
	"syscall"
	"time"

	"github.com/rcrowley/go-metrics"
)

var (
	opsRate = metrics.NewRegisteredMeter("ops", nil)
)

func Run() {
	setLimit()
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	ln, err := net.Listen("tcp", ":37137")
	if err != nil {
		panic(err)
	}

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
		ln.Close()
	}()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Fatalf("accept temp err: %v", ne)
				continue
			}
			log.Fatalf("accept err: %v", e)
			return
		}
		go handleConn(conn)
		connections = append(connections, conn)
		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}

}
func handleConn(conn net.Conn) {
	io.CopyN(conn, conn, 8)
	opsRate.Mark(1)
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
