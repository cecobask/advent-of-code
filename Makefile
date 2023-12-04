.PHONY: *

generate:
	@echo "Generating..."
	@bash hack/generate.sh
	@echo "Done!"

lint:
	@echo "Linting..."
	@golangci-lint run
	@echo "Done!"

lint-fix:
	@echo "Linting and fixing..."
	@golangci-lint run --fix
	@echo "Done!"