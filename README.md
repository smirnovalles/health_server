# Health Check Server

## Description

Health Check Server provides an RPC API to query the server's status. The main route `/health_check` executes the external script `health_check.sh` and returns the result in JSON format.

---

## Table of Contents

- [Description](#Description)
- [Requirements](#Requirements)
- [Installation and Setup](#Installation-and-Setup)
  - [Step 1: Install Go](#Step-1-Install-Go)
  - [Step 2: Build the Application](#Step-2-Build-the-Application)
  - [Step 3: Configure systemd](#Step-3-Configure-systemd)
  - [Step 4: Start the Service](#Step-4-Start-the-Service)
  - [Step 5: Configure Firewall](#Step-5-Configure-Firewall)
- [Usage](#Usage)
- [License](#License)

---

## Requirements

- Ubuntu 18.04 or higher
- Go 1.16 or higher
- systemd

---

## Installation and Setup

### Step 1: Install Go

1. Update the package list and install Go:
    ```bash
    sudo apt update
    sudo apt install golang-go -y
    ```

### Step 2: Build the Application

1. Create a working directory and navigate to it:
    ```bash
    mkdir -p ~/health_server
    cd ~/health_server
    ```

2. Save the server code to `main.go`:
    ```bash
    nano main.go
    ```
    Paste the code and save the file (Ctrl + O, Enter, Ctrl + X).

3. Build the application:
    ```bash
    go build -o health_server main.go
    ```

4. Move the executable and script to `/usr/local/bin`:
    ```bash
    sudo mv health_server /usr/local/bin/
    sudo mv path/to/health_check.sh /usr/local/bin/health_check.sh
    sudo chmod +x /usr/local/bin/health_check.sh
    ```

### Step 3: Configure systemd

1. Create the systemd service file:
    ```bash
    sudo nano /etc/systemd/system/health_server.service
    ```
2. Insert the following content, replacing parameters as necessary:
    ```ini
    [Unit]
    Description=Health Check Server
    After=network.target

    [Service]
    ExecStart=/usr/local/bin/health_server -port=8080 -script=/usr/local/bin/health_check.sh
    Restart=always
    User=www-data
    Group=www-data
    Environment=PATH=/usr/bin:/usr/local/bin
    Environment=GO_ENV=production
    WorkingDirectory=/usr/local/bin

    [Install]
    WantedBy=multi-user.target
    ```

### Step 4: Start the Service

1. Reload systemd to apply changes:
    ```bash
    sudo systemctl daemon-reload
    ```

2. Start the service and ena

Chat AI Bot - Chat GPT | Midjourney | Claude | Gemini, [30.09.2024 20:14]
ble it to run on boot:
    ```bash
    sudo systemctl start health_server
    sudo systemctl enable health_server
    ```

3. Check the service status:
    ```bash
    sudo systemctl status health_server
    ```

### Step 5: Configure Firewall

1. Allow access to the chosen port (e.g., 8080):
    ```bash
    sudo ufw allow 8080
    sudo ufw reload
    ```

---

## Usage

After successful installation and service startup, you can check the server status by sending an HTTP GET request to the `/health_check` route. For example:

```bash
curl http://localhost:8080/health_check
```

The response will be in JSON format:

```json
{
  "msg": "message"
}
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
