yd: cli.go core.go
	go build -ldflags "-s -w" -o yd cli.go core.go

clean:
	rm yd

install:
	chmod +x yd
	cp yd ~/.local/bin/
