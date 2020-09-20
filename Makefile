install:
	go build -o /usr/local/bin/go-cli . && echo "go-cli installed!"
uninstall:
	rm /usr/local/bin/go-cli && echo "Successfully uninstalled go-cli"