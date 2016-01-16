GO = go
GMESSAGE := "Default Message"
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
	
push:
	git commit -m ${GMESSAGE}
	git status | grep 'modified:' | awk 'BEGIN{FS=":";}{print $2;}' | xargs git add
	git push origin master
