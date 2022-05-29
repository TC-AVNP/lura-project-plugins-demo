.PHONY: run build-all

build-plugins:
	@go build -buildmode=plugin -o ./plugins/shared-objects/request-example.so  ./plugins/request-example/main.go
	@go build -buildmode=plugin -o ./plugins/shared-objects/inject-header-example.so  ./plugins/inject-header-example/main.go

# compile plugins with delve flags to make VScode debugging possible
dev:
	@go build -gcflags='all=-N -l' -buildmode=plugin -o ./plugins/shared-objects/request-example.so  ./plugins/request-example/main.go
	@go build -gcflags='all=-N -l' -buildmode=plugin -o ./plugins/shared-objects/inject-header-example.so  ./plugins/inject-header-example/main.go



build-lura:
	@go build -o services/lura/cmd/lura services/lura/cmd/main.go

build-all: build-plugins build-lura

run-lura: build-all
	@cd services/lura/cmd/ && ./lura

clean:
	rm plugins/shared-objects/*.so
	rm services/lura/cmd/lura