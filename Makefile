emailfile: main.go
	go build -o emailfile main.go

linux: ALWAYS
	GOOS=linux make

macos: ALWAYS
	GOOS=macos make

clean: ALWAYS
	rm -f emailfile

ALWAYS:
