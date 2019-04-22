start:
	docker-compose up -d mysql
	docker exec mysql mysql -u kenji -pkenji sealion < ./_sql/create_table.sql