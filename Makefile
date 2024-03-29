up:
	@docker-compose up -d


down:
	@docker-compose down


download-bin:
	# Use this to download only the kafka/bin folder from github.
	@svn checkout https://github.com/apache/kafka/trunk/bin

cli:
	@docker exec -it `docker ps --filter name=kafka-1 -q` bash
