ClientMainFiles = client/main.go client/users.go client/styles.go

run-client : $(ClientMainFiles)
			@go run $(ClientMainFiles)

compile-proto :
			protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative myserver/protofiles/init.proto