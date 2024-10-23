
# TinyRP - A Simple Reverse Proxy

TinyRP is a lightweight reverse proxy server built in Go. It routes requests to different backend services based on configuration defined in a YAML file. This project allows you to easily mock server responses using Docker containers.

## Table of Contents

- [Features](#features)
- [Directory Structure](#directory-structure)
- [Configuration](#configuration)
- [Running Mock Servers](#running-mock-servers)
- [Usage](#usage)
- [Logging](#logging)
- [License](#license)

## Features

- Simple reverse proxy functionality.
- Configurable routes defined in a YAML file.
- Easy to run mock servers using Docker.
- Logging of incoming requests and errors.

## Directory Structure
```plaintext
tinyrp
├── data
│   └── config.yaml        # Configuration file defining routes
├── go.mod                 # Go module file
├── go.sum                 # Go module dependencies
├── internal
│   ├── configs            # Configuration handling
│   │   └── config.go      # Configuration struct and loading logic
│   ├── handlers           # Request handlers for routing
│   │   └── handlers.go    # Main handler for proxying requests
│   ├── loggers            # Logging utilities
│   │   └── logger.go      # Logger implementation
│   └── Makefile           # Makefile for build and run tasks
└── tinyrp.go              # Entry point for the application
```

## Configuration

The routes for the mock servers are defined in the `data/config.yaml` file. The `config.Load` function reads this file and initializes the proxy routes accordingly. Each route specifies an endpoint and its corresponding destination URL.

### Example of `config.yaml`:

```yaml
resources:
  - name: images
    endpoint: /server1
    destination_url: "http://localhost:9001"
  - name: users
    endpoint: /server2
    destination_url: "http://localhost:9002"
  - name: api
    endpoint: /server3
    destination_url: "http://localhost:9003"
```

## Running Mock Servers

To run the mock servers defined in `config.yaml`, use the following command:

```bash
make run-containers
```

## Usage

Install Dependencies: Ensure you have Go installed. Then run:

```bash
go mod tidy
```

Run the Server: To start the TinyRP server, run:

```bash
go run tinyrp.go
```

Access the Proxy: The server will listen on the configured address and port as specified in your config.yaml file.

## Logging

The application logs incoming requests and errors using the logging package provided. The log format includes timestamps for easier tracking of events.

