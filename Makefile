# Название вашего бинарного файла
BINARY_PATH=bin
BINARY_NAME=msh
# Директория для установки
INSTALL_DIR=/usr/local/bin

.PHONY: install uninstall

# Команда по умолчанию
.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building..."
	go build -o $(BINARY_PATH)/$(BINARY_NAME) cmd/main.go

.PHONY: run
run: build
	@echo "Running..."
	./$(BINARY_PATH)/$(BINARY_NAME)

.PHONY: test
test:
	@echo "Testing..."
	go test ./...

.PHONY: lint
lint:
	@echo "Linting..."
	golangci-lint run

.PHONY: clean
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_PATH)/$(BINARY_NAME)
