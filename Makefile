
build:
	docker build -t clean-architetris:$$(git branch | grep "^*" | awk '{print $$2}') -f Dockerfile.dev .

run:
	docker-compose up --build

test:
	for pkg in $$(go list ./...); do go test $$pkg; done

.PHONY: build test
