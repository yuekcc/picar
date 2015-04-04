all:
	go build

get:
	gopm get -v -g

release:
	go build -ldflags -w

clean:
	rm -f picar
