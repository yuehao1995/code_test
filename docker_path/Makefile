build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/dockerpath
	chmod -R 777 ./bin/dockerpath
	docker build -t dockerpath .
run:
	docker run -e Config_Path=/opt/eechains/dockerpath dockerpath