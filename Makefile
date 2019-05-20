BINARY=main
MAIN_FILE=main.go

# Basic go commands
GOCMD=go
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

deps:
	$(GOGET) github.com/go-chi/chi

run:
	$(GORUN) $(MAIN_FILE)

build:
	$(GOBUILD) $(MAIN_FILE)

exec:
	./$(BINARY)

test:
	$(GOTEST) ./...