clean:
	sudo rm -rf ~/boo
	sudo mkdir -p ~/boo

build:
	go build -o boo-app ./cmd/app

install: clean build
	sudo mv boo-app ~/boo/boo-app
	sudo cp ./.env ~/boo/.env
	sudo cp -r ./assets/ ~/boo/assets/