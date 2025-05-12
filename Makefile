# Определите переменные
PROTOC=protoc
PROTO_PATH=api
GO_OUT=.
GRPC_GATEWAY_OUT=.
BINARY_NAME=my_server

# Цель по умолчанию
all: generate build

# Генерация кода для gRPC и grpc-gateway из order.proto
generate:
	$(PROTOC) -I$(PROTO_PATH) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) --grpc-gateway_out=$(GRPC_GATEWAY_OUT) \
	$(PROTO_PATH)/annotations.proto \
	$(PROTO_PATH)/field_behavior.proto \
	$(PROTO_PATH)/http.proto \
	$(PROTO_PATH)/httpbody.proto \
	$(PROTO_PATH)/order.proto

# Сборка бинарника
build:
	go build -o $(BINARY_NAME) .

# Запуск сервера
run:
	./$(BINARY_NAME)

# Очистка
clean:
	rm -f $(BINARY_NAME) *.pb.go *.gw.go

.PHONY: all generate build run clean