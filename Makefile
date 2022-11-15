MOCKS_DIR=$(CURDIR)/testutil/mocks/
MOCKGEN_REPO=github.com/golang/mock/mockgen
MOCKGEN_VERSION=v1.6.0
MOCKGEN_CMD=go run ${MOCKGEN_REPO}@${MOCKGEN_VERSION}

test:
	go test ./...

mock-gen:
	mkdir -p $(MOCKS_DIR)
	$(MOCKGEN_CMD) -source=client/interface.go -package mocks -destination $(MOCKS_DIR)/client.go
