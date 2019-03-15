build:
	GOOS=linux GOARCH=amd64 go build -o ./api/client/bin/goroutine_server
	chmod -R 777 ./api/client/bin/goroutine_server
epoll:
	GOOS=linux GOARCH=amd64 go build -o ./api/client/bin/epoll_server
	chmod -R 777 ./api/client/bin/epoll_server
epoll_throughts:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/epoll_throughts_server
	chmod -R 777 ./api/epoll_client/bin/epoll_throughts_server
goroutine_throughts:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/goroutine_throughts_server
	chmod -R 777 ./api/epoll_client/bin/goroutine_throughts_server
multiple_server:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/multiple_server
	chmod -R 777 ./api/epoll_client/bin/multiple_server
server_workerpool:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/server_workerpool
	chmod -R 777 ./api/epoll_client/bin/server_workerpool
io_intensive_goroutine:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/io_intensive_goroutine
	chmod -R 777 ./api/epoll_client/bin/io_intensive_goroutine
io_intensive_epoll_server:
	GOOS=linux GOARCH=amd64 go build -o ./api/epoll_client/bin/io_intensive_epoll_server
	chmod -R 777 ./api/epoll_client/bin/io_intensive_epoll_server