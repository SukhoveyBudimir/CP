up:
	migrate -database postgresql://postgres:123@localhost:49153/model?sslmode=disable  -path ./db up
down:
	migrate -database postgresql://postgres:123@localhost:49153/model?sslmode=disable  -path ./db down
drop:
	migrate -database postgresql://postgres:123@localhost:49153/model?sslmode=disable  -path ./db drop
gen:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     proto/model.proto

clean:
	rm proto/*.go
run:
	go run main.go