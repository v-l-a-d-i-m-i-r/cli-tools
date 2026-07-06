set dotenv-load := true

default:
  @just --list

init:
  @[ -f .env ] || cp example.env .env;
  @go mod tidy;

build-uuidv7:
  @go build -o ./bin/uuidv7 ./cmd/uuidv7;

run-uuidv7:
  @go run -race ./cmd/uuidv7;

install-uuidv7: build-uuidv7
  @mkdir -p $INSTALL_DIR
  @cp ./bin/uuidv7 $INSTALL_DIR/uuidv7;
