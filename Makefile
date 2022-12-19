# Colors
RESET=\033[0m
GREEN=\033[32m
YELLOW=\033[33m
CYAN=\033[36m

.PHONY: all
all: help

# Cite: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Display this help page
	@grep -E '^[a-zA-Z0-9/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | 		\
	sort | 															\
	awk ' 															\
		BEGIN {FS = ":.*?## "; print "${GREEN}content-api help:\n"} \
		{printf "${CYAN}%-30s${RESET} %s\n", $$1, $$2} 				\
	  	END {print "\n${YELLOW}See README.md for more information${RESET}\n"}'

.PHONY: bundle
bundle: ## Builds the yaml bundle
	npx --yes @apidevtools/swagger-cli bundle api/tinybeans.yaml --outfile bundle.yaml --type yaml

.PHONY: gen
gen: bundle ## Generate client library
	npx --yes @openapitools/openapi-generator-cli generate --generator-key client
	GOSIMPORTS=$GOPATH/bin/gosimports
	$GOSIMPORTS -l -w ./go

.PHONY: setup
setup: ## Intalls all the tools
	npm install

