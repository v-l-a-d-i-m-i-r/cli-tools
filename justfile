set dotenv-load := true

# Show available recipes
default:
  @just --list

# Initialize the project
init:
  @[ -f .env ] || cp example.env .env;
  @go mod tidy;
  @git config core.hooksPath .githooks;

# Run all tests
test:
  @go test ./...;

# Run linter
lint:
  @go tool golangci-lint run;

# Build the uuidv7 binary
build-uuidv7:
  @CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/uuidv7 ./cmd/uuidv7;

# Run the uuidv7 command directly
run-uuidv7:
  @go run -race ./cmd/uuidv7;

# Install the uuidv7 binary to the specified install directory
install-uuidv7: build-uuidv7
  @mkdir -p $INSTALL_DIR
  @cp ./bin/uuidv7 $INSTALL_DIR/uuidv7;

# Build the work-item-id binary
build-work-item-id:
  @CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/work-item-id ./cmd/work-item-id;

# Run the work-item-id command directly
run-work-item-id:
  @go run -race ./cmd/work-item-id;

# Install the work-item-id binary to the specified install directory
install-work-item-id: build-work-item-id
  @mkdir -p $INSTALL_DIR
  @cp ./bin/work-item-id $INSTALL_DIR/work-item-id;

# Build all binaries
build-all: build-uuidv7 build-work-item-id build-docker-hub-image-tags

# Install all binaries
install-all: install-uuidv7 install-work-item-id install-docker-hub-image-tags

# Build the docker-hub-image-tags binary
build-docker-hub-image-tags:
  @CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/docker-hub-image-tags ./cmd/docker-hub-image-tags;

# Run the docker-hub-image-tags command directly (e.g. just run-docker-hub-image-tags library node)
run-docker-hub-image-tags namespace repository:
  @go run -race ./cmd/docker-hub-image-tags {{namespace}} {{repository}};

# Install the docker-hub-image-tags binary to the specified install directory
install-docker-hub-image-tags: build-docker-hub-image-tags
  @mkdir -p $INSTALL_DIR
  @cp ./bin/docker-hub-image-tags $INSTALL_DIR/docker-hub-image-tags;
