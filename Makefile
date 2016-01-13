GO=go

all: mqttpub

mqttpub: golang-mqttpub
	mv $^ $@

golang-mqttpub: main.go
	$(GO) build 

clean:
	rm -f mqttpub
	
