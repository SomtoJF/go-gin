
run: run-db run-server

run-db:
	docker-compose -f ./compose.yml up -d

run-server:
	CompileDaemon -command="./go-gin"