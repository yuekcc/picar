resouce:
	./buildres

get:
	gopm get -v -g

all: resouce
	go build

release: resouce
	go build -ldflags -w

install: resouce
	go install

clean:
	rm -f picar