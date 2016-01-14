GO = go
PREFIX := /usr/local/bin

all: mqttpub

mqttpub: golang-mqttpub
	mv $^ $@

golang-mqttpub: main.go
	$(GO) build 

install: mqttpub
	cp $^ ${PREFIX}/.
clean:
	rm -f mqttpub
	
