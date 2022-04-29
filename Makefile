.PHONY: default clean help

default: help

build: ## Build all probes to bin/ directory
	@echo "Building all probes to directory bin/..."
	@mkdir -p bin
	@go build -o bin ./cmd/...

clean: ## Clean files from bin/ directory
	@echo "Cleaning /bin..."
	@rm -rf bin

rebuild: clean build ## Clean and then build binaries


help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo ''
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)