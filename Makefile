PROJECT := matrix

all: build

.PHONY: build
build:    ## Build with native env.
	@./scripts/build.sh ${PROJECT}

.PHONY: help
help:     ## Show this help.
	@echo "Makefile Help Menu:\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: clean
clean:    ## Clean native build cache.
	@rm -rf bin
	@echo "clean [ ok ]"