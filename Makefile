run:
	go run .
# unit-tests:
# 	go test ./...
# functional-tests:
# 	go test ./functional_tests/transformer_test.go
build:
	docker build . -t go-fit-tracker 