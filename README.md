# SSH Honeypot

A simple SSH honeypot written in Go that simulates an SSH server to log attempted login credentials from potential attackers. This tool is designed for educational and research purposes to understand unauthorized access attempts.

## Features

- Simulates an Ubuntu OpenSSH server
- Logs connection attempts with source IP addresses
- Captures usernames and passwords from login attempts
- Runs on a specified TCP port (default: 2222)
- Handles multiple concurrent connections

## Prerequisites

- Go 1.15 or higher
- Administrative privileges (for binding to ports)

## Installation

1. Clone this repository
2. Navigate to the project directory
3. Build the executable:
```bash
go build -o ssh-honeypot
```

## Usage

1. Run the honeypot:
```bash
./ssh-honeypot
```

The server will start listening on port 2222 by default.

To change the port, modify the `address` variable in the `main()` function.

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

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
