all: bifs
	go build

bifs:
	cd web/assets && make

get:
	gopm get -v -g

test:
	./picar webui

install: resouce
	go install

clean:
	rm -f picar