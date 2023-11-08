.PHONY: docker
docker:
	# 清理上次编译的东西
	@rm main || true

	@docker rmi -f basic_project:v1
	@export GOPROXY=https://goproxy.io
	@go mod tidy
	# 编译
	@go build -o main .
	# 打包镜像

	@docker build -t basic_project:v1 .
