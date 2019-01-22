GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PROTOC=protoc
PBDIR=pb
PROTO_TARGET=protos/practice.proto
BINDIR=bin

install-pkgs:
	$(GOGET) -u google.golang.org/grpc \
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go

grpc-build:
	@if [ ! -d $(PBDIR) ]; \
		then mkdir $(PBDIR); \
		fi
	$(PROTOC) -I=protos $(PROTO_TARGET) --go_out=plugins=grpc:$(PBDIR)

build: grpc-build
	@if [ ! -d $(BINDIR) ]; \
		then mkdir $(BINDIR); \
		fi
	$(GOBUILD) -o $(BINDIR)/main

client-build: grpc-build
	@if [ ! -d $(BINDIR) ]; \
		then mkdir $(BINDIR); \
		fi
	$(GOBUILD) -o $(BINDIR)/client client/client.go


clean:
	rm -rf $(PBDIR)
	rm -rf $(BINDIR)
