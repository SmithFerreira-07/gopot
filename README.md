# SSH Honeypot

A simple SSH honeypot written in Go that simulates an SSH server to log attempted login credentials from potential attackers. This tool is designed for educational and research purposes to understand unauthorized access attempts.

## Features

- Simulates an Ubuntu OpenSSH server
- Logs connection attempts with source IP addresses
- Captures usernames and passwords from login attempts
- Runs on a specified TCP port (default: 2222)
- Handles multiple concurrent connections
- Dockerized for easy deployment

## Prerequisites

- Docker
- OR
  - Go 1.15 or higher
  - Administrative privileges (for binding to ports)

## Installation & Usage

### Using Docker (Recommended)

1. Build the Docker image:
```bash
docker build -t ssh-honeypot .
```

2. Run the container:
```bash
docker run -p 2222:2222 ssh-honeypot
```

### Manual Installation

1. Clone this repository
2. Navigate to the project directory
3. Build and run:
```bash
go build -o honeypot main.go
./honeypot
```

## Output

The honeypot logs:
- Incoming connection IP addresses
- Attempted usernames and passwords
- Connection errors

Example log output:
```
2024/01/01 12:00:00 Connection received from: 192.168.1.100:54321
2024/01/01 12:00:05 Credentials from 192.168.1.100:54321 - Username: admin, Password: password123
```

## Security Considerations

- This is a basic honeypot implementation intended for educational purposes
- Do not run on production systems
- Consider legal implications before deployment
- Monitor system resources as the honeypot accepts unlimited connections

## Disclaimer

This tool should only be used in controlled environments for educational or research purposes. The user assumes all responsibility for its use and must comply with all applicable laws and regulations.

## License

MIT License

Copyright (c) 2024
