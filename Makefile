APP_NAME := haiku-server
BUILD_DIR := .
SERVICE_NAME := haiku
PORT := 22283

.PHONY: build run install setup-nginx service restart clean

## Compile the Go application
build:
	go build -o $(APP_NAME) main.go

## Run the server directly (on port 8080)
run: build
	PORT=$(PORT) ./$(APP_NAME)

## Install Go and Nginx on Ubuntu
install:
	sudo apt update
	sudo apt install -y golang-go nginx

## Configure Nginx reverse proxy (requires sudo)
setup-nginx:
	sudo cp nginx.conf /etc/nginx/sites-available/$(SERVICE_NAME)
	sudo ln -sf /etc/nginx/sites-available/$(SERVICE_NAME) /etc/nginx/sites-enabled/
	sudo nginx -t && sudo systemctl restart nginx

## Create a systemd service for running the server in the background
service:
	echo "[Unit]\n\
Description=Haiku Go Web Server\n\
After=network.target\n\n\
[Service]\n\
ExecStart=$(BUILD_DIR)/$(APP_NAME)\n\
WorkingDirectory=$(BUILD_DIR)\n\
Restart=always\n\
User=www-data\n\n\
[Install]\n\
WantedBy=multi-user.target" | sudo tee /etc/systemd/system/$(SERVICE_NAME).service > /dev/null

	sudo systemctl daemon-reload
	sudo systemctl enable $(SERVICE_NAME)
	sudo systemctl start $(SERVICE_NAME)

## Restart the systemd service
restart:
	sudo systemctl restart $(SERVICE_NAME)

## Clean up the compiled binary
clean:
	rm -f $(APP_NAME)
