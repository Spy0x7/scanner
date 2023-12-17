# Your Port Scanner

Your Port Scanner is a lightweight and efficient Go-based tool designed to scan for open ports on a specified host or IP address.

## Features

- **Fast and Concurrent:** Employs concurrent routines for quick port scanning without compromising accuracy.
- **Port Range Specification:** Allows users to define a range of ports or use default common ports for scanning.
- **Versatile Protocol Support:** Equipped to scan both TCP and UDP ports for comprehensive analysis.
- **Clear Command-Line Interface:** Simple and intuitive CLI for initiating scans and viewing results.
- **Customizable Output:** Provides clear and configurable output for easy interpretation of scan results.

## Installation

Follow these steps to install and utilize this port scanner:

1. Clone this repository:
    ```bash
    git clone https://github.com/Spy0x7/scanner.git
    cd scanner
    ```

2. Build the program:
    ```bash
    go build -o scanner main.go
    ```

## Usage

Post installation, execute the port scanner using the following command format:

```bash
scanner -i <target_IP>
