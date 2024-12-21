docker_up_redis:
	docker-compose -f docker-compose.yaml up -d --build --force-recreate --remove-orphans redis

docker_exec_redis:
	docker exec -it env_redis redis-cli
