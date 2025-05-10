# Haiku JSON Server (Go)

A simple HTTP server written in Go that returns a random haiku in JSON format, depending on the browser's language (Polish/English).

## 📦 Requirements

- Ubuntu (e.g. a VPS with SSH access)
- Go installed on the server
- Files: `main.go`, `haiku_pl.json`, `haiku_en.json`

---

## 🚀 Installation on Ubuntu Server

### 1. Connect to the server

```bash
ssh user@your-server
```

### 2. Install Go (if not already installed)

```bash
sudo apt update
sudo apt install golang-go
```

### 3. Upload project files to the server

From your local machine:

```bash
scp main.go haiku_*.json user@your-server:/home/user/haiku-server/
```

Or use Git if hosted in a repository.

### 4. Build the application

```bash
cd /home/user/haiku-server
go build -o haiku-server main.go
```

### 5. Run the server (on port 8080)

```bash
./haiku-server
```

The server will be accessible at `http://localhost:8080/haiku`

---

## ⚙️ Optional: Run as a systemd service

1. Create a systemd unit file:

```bash
sudo nano /etc/systemd/system/haiku.service
```

Paste the following content:

```ini
[Unit]
Description=Haiku Go Web Server
After=network.target

[Service]
ExecStart=/home/user/haiku-server/haiku-server
WorkingDirectory=/home/user/haiku-server
Restart=always
User=www-data

[Install]
WantedBy=multi-user.target
```

2. Enable and start the service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable haiku.service
sudo systemctl start haiku.service
```

---

## 🌐 Optional: Use Nginx as a reverse proxy

1. Install Nginx:

```bash
sudo apt install nginx
```

2. Create a config file at `/etc/nginx/sites-available/haiku`:

```nginx
server {
    listen 80;
    server_name yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

3. Enable and restart Nginx:

```bash
sudo ln -s /etc/nginx/sites-available/haiku /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

---

## 🔒 HTTPS (Let's Encrypt)

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com
```

---

## 📎 API Endpoint

**GET** `/haiku`

Returns a haiku in the following format:

```json
{
  "haiku": "Cisza nad stawem\nksiężyc odbity w wodzie\nżaba skacze — plask!"
}
```

The correct file (`haiku_pl.json` or `haiku_en.json`) is selected based on the `Accept-Language` header.

---

## 📄 License

MIT
