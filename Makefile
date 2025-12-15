BINARY := hycli
BUILD_DIR := bin
CMD := ./cmd/hycli
GO := go

.PHONY: build clean
build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY) $(CMD)

clean:
	rm -rf $(BUILD_DIR)
