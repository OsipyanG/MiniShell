# Название вашего бинарного файла
BINARY_NAME=bin/msh

# Команда по умолчанию
.PHONY: all
all: build

# Сборка приложения
.PHONY: build
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) cmd/main.go

# Запуск приложения
.PHONY: run
run: build
	@echo "Running..."
	./$(BINARY_NAME)

# Запуск тестов
.PHONY: test
test:
	@echo "Testing..."
	go test ./...

# Линтинг кода (необходимо предварительно установить golint)
.PHONY: lint
lint:
	@echo "Linting..."
	golangci-lint run

# Очистка
.PHONY: clean
clean:
	@echo "Cleaning..."
	go clean
	rm -f $(BINARY_NAME)
