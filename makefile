PROJECT=GO-micro-service
GO=go
execPing_SRC=github.com/HarleyWaagmeester/server/execPing
server_SRC=github.com/HarleyWaagmeester/server/server

all: execPing $(server_SRC)
	$(GO) install $(server_SRC)

execPing: $(execPing_SRC)
	$(GO) install $(execPing_SRC)
