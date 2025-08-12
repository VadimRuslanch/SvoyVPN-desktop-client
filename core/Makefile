.PHONY: proto core
proto:
	../scripts/generate_protos.sh

core:
	GOOS=darwin GOARCH=amd64 go build -o bin/vadimd-darwin-amd64 ./cmd/vadimd
	GOOS=windows GOARCH=amd64 go build -o bin/vadimd-windows-amd64.exe ./cmd/vadimd
