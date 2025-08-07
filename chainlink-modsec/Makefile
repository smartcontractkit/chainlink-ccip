.PHONY: tidy

tidy:
	@echo "Tidying go modules in subdirectories..."
	@for dir in integration-tests libmodsec evmimpls; do \
		echo "Tidying $$dir"; \
		(cd $$dir && go mod tidy); \
	done
