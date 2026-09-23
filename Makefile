# Each directory with a go.mod is an independent module.
MODULES := $(dir $(wildcard */go.mod))

.PHONY: all build vet fmt fmt-check test tidy

all: fmt-check vet build

build:
	@for m in $(MODULES); do echo "== build $$m"; (cd $$m && go build -o /dev/null ./...) || exit 1; done

vet:
	@for m in $(MODULES); do echo "== vet $$m"; (cd $$m && go vet ./...) || exit 1; done

test:
	@for m in $(MODULES); do echo "== test $$m"; (cd $$m && go test ./...) || exit 1; done

fmt:
	@gofmt -w $(MODULES)

fmt-check:
	@out=$$(gofmt -l $(MODULES)); if [ -n "$$out" ]; then echo "gofmt needed:"; echo "$$out"; exit 1; fi

tidy:
	@for m in $(MODULES); do echo "== tidy $$m"; (cd $$m && go mod tidy) || exit 1; done
