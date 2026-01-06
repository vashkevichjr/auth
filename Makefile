LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(LOCAL_BIN) v2.7.2


lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml