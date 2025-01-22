.PHONY: all clean

all: bin/main bin/load

bin/main: cmd/main/main.go
	go build -o bin/main cmd/main/main.go

bin/load: cmd/load/main.go
	go build -o bin/load cmd/load/main.go

clean:
	rm -rf bin
