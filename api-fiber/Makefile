# Makefile
.EXPORT_ALL_VARIABLES:	


GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

update: 
	@echo "\033[0;33m################################## update ##################################\033[0m"
	@rm -f go.*
	go mod init github.com/jeffotoni/gogrpc.palestra/api.fiber
	GOOS=linux go build -o ecommerce.fiber main.go

deploy:
	@echo "\033[0;33m################################## build statik ##################################\033[0m"
	statik -src=html/
	@echo "\033[0;32mstatik generated\033[0m  \033[0;33m[ok]\033[0m "
	@echo "--------------------------------------------------------------------------------------"
	CGO_ENABLED=0 GOOS=linux go build -o ecommerce.fiber main.go
	upx ecommerce.fiber
	#upx --brute ecommerce.fiber
	docker build --no-cache -f Dockerfile -t gcr.io/infra-admin-204520/ecommerce.fiber:latest .
	docker push gcr.io/infra-admin-204520/ecommerce.fiber:latest	
	@echo "\033[0;32mecommerce.fiber generated\033[0m \033[0;33m[ok]\033[0m \n"
