clean:
	sudo rm -rf ~/boo
	sudo mkdir -p ~/boo

build:
	go build -o boo ./cmd/app

install: clean build
	sudo mv boo ~/boo/boo
	sudo cp ./.env ~/boo/.env
	sudo cp -r ./assets/ ~/boo/assets/