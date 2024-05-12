sqlc:
	sqlc generate

run:
	go run main.go

deploy-infra:
	docker-compose -f ./deployment/docker-compose.yaml up

down-infra:
	docker-compose -f ./deployment/docker-compose.yaml down




