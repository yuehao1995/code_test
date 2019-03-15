/**
 * @author zhangyuehao
 * @date 2019-03-13 14:31
 */

package main

import (
	"flag"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "evio_server IP")
	connections = flag.Int("conn", 1, "number of tcp connections")
	startMetric = flag.String("sm", time.Now().Format("2006-01-02T15:04:05 -0700"), "start time point of all clients")
	opsRate     = metrics.NewRegisteredTimer("ops", nil)
)

func main() {
	flag.Parse()

	setLimit()

	go func() {
		startPoint, err := time.Parse("2006-01-02T15:04:05 -0700", *startMetric)
		if err != nil {
			panic(err)
		}
		time.Sleep(startPoint.Sub(time.Now()))

		metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	}()

	addr := *ip + ":36137"
	log.Printf("连接到 %s", addr)
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
	log.Printf("完成初始化 %d 连接", len(conns))
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

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
}
