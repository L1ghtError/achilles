.PHONY: all clean

all: bin/main bin/load

bin/main: cmd/ahlc/ahlc.go
	go build -o bin/ cmd/ahlc/ahlc.go

bin/load: cmd/lhlc/lhlc.go
	go build -o bin/ cmd/lhlc/lhlc.go

clean:
	rm -rf bin
