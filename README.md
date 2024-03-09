## TELSSE

This project implements a simple telnet + server-sent events (SSE) application with a terminal-like web interface,
utilizing TailwindCSS for styling and Golang for the backend. The application serves a static website and maintains a telnet
service for receiving messages, which are then broadcasted in real-time to connected web clients via SSE.

### Features

- Static website serving using Golang
- Real-time message broadcasting using Server-Sent Events (SSE)
- Terminal-like interface for message display using TailwindCSS
- Telnet service for receiving messages

### Prerequisites

To run this project, you need to have the following installed:

- Golang (version 1.15 or later)
- Node.js and npm (for TailwindCSS)
- `godotenv` package for Golang

### Installation

1. Clone the repository:

```bash
git clone https://github.com/nkoster/telsse
cd telsse
```

2. Install the Go dependencies:

```bash
go mod tidy
```

3. Run the Golang server:

```bash
go run main.go
```

or build a binary, and run the binary `telsse`

```bash
CGO_ENABLED=0 go build -ldflags="-extldflags=-static"
./telsse
```
Make sure you have configured your `.env` file. See below in the _Configuration_ section.

### Usage

- Open a web browser and navigate to `http://localhost:8080` to view the SSE chat interface. Replace `8080` with the actual port number specified in your `.env` file.
- Use a telnet client to connect to the telnet service and send messages. Replace `5023` with the actual port number specified in your `.env` file:

```bash
telnet localhost 5023
```
or
```bash
tail -f /var/log/nginx/access.log | telnet localhost 5023
```

- Messages sent via telnet will appear in real-time on the web interface.

### Configuration

Create a `.env` file in the root directory of your project with the following content:

```
UI=./ui
HTTP_PORT=8080
TELNET_PORT=5023
LOG=false         # when true, telnet lines also appear in the server log
```

### Contributing

Contributions to this project are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

### License

This project is licensed under the MIT License.
