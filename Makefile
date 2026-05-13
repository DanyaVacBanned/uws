
BUILD_DIR := ./bin


build: $(BUILD_DIR)
	go build -o $(BUILD_DIR)/uws main.go

$(BUILD_DIR):
	@if [ ! -d "$@" ]; then \
			echo "Build directory does not exists. Creating it.."; \
			mkdir "$@"; \
	fi
