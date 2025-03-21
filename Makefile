

# docker-compose up -d
# docker-compose -f docker-compose.yaml down


run:
	@echo "运行"
	@go run cmd/server/server.go


# docker desktop 直接敲也行

container-console-mysql:
	@echo "打开 MySQL 容器控制台"
	@echo "输入 mysql --host=127.0.0.1 --port=3306 --user=root --password=my-secret-pw"
	docker exec -it clwy-api-mysql sh


container-console-redis:
	@echo "打开 Redis 容器控制台"
	@echo "输入 redis-cli"
	docker exec -it clwy-api-redis sh


test:
	@echo "跑一遍测试"
	go test -count=1 -short -v ./...


.PHONY: run container-console-mysql container-console-redis

