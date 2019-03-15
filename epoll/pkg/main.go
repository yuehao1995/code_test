/**
 * @author zhangyuehao
 * @date 2019-03-13 14:24
 */

package main

import "github.com/eechains/epoll/pkg/api/io_intensive_goroutine"

func main() {
	//goroutine_server.Run()
	//epoll_server.Run()
	//epoll_server_throughts.Run()
	//multiple_server.Run()
	//server_workerpool.Run()
	//io_intensive_epoll_server.Run()
	io_intensive_goroutine.Run()
}
