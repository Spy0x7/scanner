# Port Scanner

A Go-based Port Scanner, this tool is a lightweight yet powerful solution designed to swiftly identify open ports on specified IP addresses or hosts. Utilizing concurrent routines, it ensures efficient and accurate port identification, offering customizable port range options or default common ports for diverse scanning needs. Supporting TCP and UDP protocols, it provides a clear command-line interface and configurable output, simplifying port analysis for users at all expertise levels.

## Installation

To make use of this port scanner:

1. **Clone the repository:**
    ```bash
    git clone https://github.com/Spy0x7/scanner.git
    cd scanner
    ```

2. **Build the program:**
    ```bash
    go build -o scanner main.go
    ```

## Usage

Once installed, execute the port scanner using the following command format:

```bash
scanner -i <target_IP>
