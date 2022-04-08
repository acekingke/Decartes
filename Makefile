# bin/yaccgo generate go grammar.y grammar.go
TOP = $(shell pwd)
grammar:
	@echo "Generating grammar.go"
	@bin/yaccgo generate go parser/grammar.y parser/grammar.go

build:
	go fmt $(TOP)/...
	go build -o bin/Decartes ./Decartes/*.go
