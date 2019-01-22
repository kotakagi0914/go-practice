GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SRCDIR=src
PROTOC=protoc
PBDIR=$(SRCDIR)/pb
PROTO_TARGET=$(SRCDIR)/protos/practice.proto
BINDIR=bin

install-pkgs:
	$(GOGET) -u google.golang.org/grpc \
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go

grpc-build:
	@if [ ! -d $(PBDIR) ]; \
		then mkdir $(PBDIR); \
		fi
	$(PROTOC) -I=$(SRCDIR)/protos $(PROTO_TARGET) --go_out=plugins=grpc:$(PBDIR)

build:
	@if [ ! -d $(BINDIR) ]; \
		then mkdir $(BINDIR); \
		fi
	$(GOBUILD) -o $(BINDIR)/server src/main.go

client-build:
	@if [ ! -d $(BINDIR) ]; \
		then mkdir $(BINDIR); \
		fi
	$(GOBUILD) -o $(BINDIR)/client $(SRCDIR)/client/client.go


clean:
	rm -rf $(PBDIR)
	rm -rf $(BINDIR)
