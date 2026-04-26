run: 
	go run .

build:
	go build -o aj

addbin:
	sudo cp aj /usr/local/bin

adduserbin:
	cp aj ~/.local/bin