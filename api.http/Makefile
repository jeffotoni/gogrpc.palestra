# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update: 
	@echo "\033[0;33m################################## update ##################################\033[0m"
	#@rm -f go.*
	#go mod init github.com/jeffotoni/gogrpc.palestra/api.http
	GOOS=linux go build -o api.http main.go

deploy:
	@echo "--------------------------------------------------------------------------------------"
	CGO_ENABLED=0 GOOS=linux go build -o api.http main.go
	upx api.http
	#upx --brute api.http
	docker build --no-cache -f Dockerfile -t jeffotoni/api.http:latest .
	#docker push jeffotoni/api.http:latest
	@echo "\033[0;32mapi.http generated\033[0m \033[0;33m[ok]\033[0m \n"
