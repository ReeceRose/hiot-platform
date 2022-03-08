dc-up:
	docker-compose -f docker/docker-compose.yml up

dc-down:
	docker-compose -f docker/docker-compose.yml down

clean:
	docker-compose -f docker/docker-compose.yml down --rmi all -v --remove-orphans