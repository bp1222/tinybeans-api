.PHONY: all
all: clean
	swagger-cli bundle api/tinybeans.yaml --outfile bundle.yaml --type yaml && \
	openapi-generator generate -g go -o go -i bundle.yaml -p packageName=tinybeans --git-user-id bp1222 --git-repo-id=tinybeans-api/go

.PHONY: clean
clean:
	rm -fR go
