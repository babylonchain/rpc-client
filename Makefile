MOCKS_DIR=$(CURDIR)/testutil/mocks
MOCKGEN_REPO=github.com/golang/mock/mockgen
MOCKGEN_VERSION=v1.6.0
MOCKGEN_CMD=go run ${MOCKGEN_REPO}@${MOCKGEN_VERSION}

# Update changelog vars
ifneq (,$(SINCE_TAG))
	sinceTag := --since-tag $(SINCE_TAG)
endif
ifneq (,$(UPCOMING_TAG))
	upcomingTag := --future-release $(UPCOMING_TAG)
endif

.PHONY: test
test:
	go test ./...

.PHONY: update-changelog
update-changelog:
	@echo ./scripts/update_changelog.sh $(sinceTag) $(upcomingTag)
	./scripts/update_changelog.sh $(sinceTag) $(upcomingTag)
