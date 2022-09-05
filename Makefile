# bin/yaccgo generate go grammar.y grammar.go
TOP = $(shell pwd)
grammar: exprgram, gram
	@echo "Generating grammar.go"	
gram:
	@bin/yaccgo generate  go parser/grammar.y parser/grammar.go
exprgram:
	@bin/yaccgo generate  go expr/expr.y expr/grammar.go
build:
	go fmt $(TOP)/...
	go build -o bin/Decartes ./Decartes/*.go
test:
	go test ./...
todo:
	@grep -rnw './' -e 'TODO:'|grep -v grep