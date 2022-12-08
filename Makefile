up:
	@docker-compose up -d


down:
	@docker-compose down


download-bin:
	# Use this to download only the kafka/bin folder from github.
	@svn checkout https://github.com/apache/kafka/trunk/bin


create-topics:
	@# To access,
	@# $ docker exec -it $(docker ps --filter name=kafka-1 -q) bash
	@# $ cd opt/bitnami/kafka
	@./bin/kafka-topics.sh --create --topic my-topic --bootstrap-server localhost:9092 --partitions 3 --replication-factor 1
