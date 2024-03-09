# Terminal Style SSE Chat

This project demonstrates a simple implementation of a server-sent events (SSE) application with a terminal-like web interface,
utilizing TailwindCSS for styling and Golang for the backend. The application serves a static website and maintains a telnet
service for receiving messages, which are then broadcasted in real-time to connected web clients via SSE.

## Features

- Static website serving using Golang
- Real-time message broadcasting using Server-Sent Events (SSE)
- Terminal-like interface for message display using TailwindCSS
- Telnet service for receiving messages

## Prerequisites

To run this project, you need to have the following installed:

- Golang (version 1.15 or later)
- Node.js and npm (for TailwindCSS)
- `godotenv` package for Golang

## Installation

1. Clone the repository:

```bash
git clone https://github.com/nkoster/telsse
cd telsse
```

2. Install TailwindCSS dependencies (if you plan to customize TailwindCSS):

```bash
npm install
```

3. Build the TailwindCSS file (optional, if you've customized Tailwind):

```bash
npx tailwindcss build src/styles.css -o static/css/tw.min.css
```

4. Run the Golang server:

```bash
go run server.go
```

## Usage

- Open a web browser and navigate to `http://localhost:8080` to view the SSE chat interface.
- Use a telnet client to connect to the telnet service and send messages. Replace `TELNET_PORT` with the actual port number specified in your `.env` file:

```bash
telnet localhost TELNET_PORT
```

- Messages sent via telnet will appear in real-time on the web interface.

## Configuration

Create a `.env` file in the root directory of your project with the following content:

```
UI=./ui
HTTP_PORT=8080
TELNET_PORT=5023
LOG=false         # when true, telnet lines also appear in the server log
```

## Contributing

Contributions to this project are welcome. Please feel free to fork the repository, make changes, and submit pull requests.

## License

This project is licensed under the MIT License.
