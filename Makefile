go-run-dev:
	DB_USERNAME=postgres DB_PASS=p4ssw0rd DB_NAME=grocery DB_HOST=localhost REDIS_HOST=localhost REDIS_PORT=6379 go run main.go

go-run-test:
	go test -v tests/category_test.go -args DB_USERNAME=postgres DB_PASS=p4ssw0rd DB_NAME=grocery DB_HOST=localhost REDIS_HOST=localhost REDIS_PORT=6379