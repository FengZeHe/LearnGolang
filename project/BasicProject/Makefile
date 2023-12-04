.PHONY: docker
docker:
	# 清理上次编译的可执行文件
	@rm main || true

	@docker rmi -f basic_project:v1

	@docker build -t basic_project:v1 .

.PHONY: clean
clean:
	@rm main || true
	@echo "clean success"

