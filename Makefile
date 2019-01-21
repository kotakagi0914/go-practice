GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PROTOC=protoc
PBDIR=pb
PROTO_TARGET=protos/practice.proto

install-pkg:
	$(GOGET) -u google.golang.org/grpc \
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go
	mkdir $(PBDIR)

grpc-build:
	$(PROTOC) -I=protos $(PROTO_TARGET) --go_out=plugins=grpc:$(PBDIR)

clean:
	rm -rf $(PBDIR)
