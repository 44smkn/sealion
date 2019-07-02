start:
	docker-compose up --build -d mysql 
	sleep 30s
	cat ./_sql/create_table.sql | docker exec -i mysql mysql -u root -pmochoten sealion 
down:
	docker-compose down
build:
	wire gen
	go build -o sealion