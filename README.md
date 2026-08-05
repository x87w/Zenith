# Zenith

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.13+-green.svg)](https://golang.org)

A lightweight, simple container manager written in Go. Zenith provides essential container lifecycle management capabilities with an elegant CLI interface and a powerful REST API (coming soon).

## Features

- **Lightweight Design** - Minimal dependencies for fast deployment
- **Simple CLI** - Intuitive command-line interface for container operations
- **Container Management** - Create, start, stop, and manage containers
- **REST API** (Coming Soon) - Programmatic access to container operations
- **Cross-Platform** - Runs on Linux, macOS, and other Unix-like systems

## Getting Started

### Prerequisites

- Go 1.13 or higher
- Linux kernel with container support

### Installation

Clone the repository:

```bash
git clone https://github.com/x87w/Zenith.git
cd Zenith
```

Build from source:

```bash
go build -o zenith main.go
```

### Quick Start

Run Zenith:

```bash
./zenith
```

## Downloading and Using ISO Images

### Ubuntu Server

#### Download the ISO

Download the latest Ubuntu Server LTS:

```bash
# Download Ubuntu Server 24.04 LTS (recommended)
wget https://releases.ubuntu.com/24.04/ubuntu-24.04-live-server-amd64.iso

# Or using curl
curl -O https://releases.ubuntu.com/24.04/ubuntu-24.04-live-server-amd64.iso
```

Other Ubuntu versions available at: https://releases.ubuntu.com/

#### Verify the Download

```bash
# Download the SHA256 checksum file
wget https://releases.ubuntu.com/24.04/SHA256SUMS
```

#### Using the ISO with Zenith

Once you've downloaded an Ubuntu Server ISO, you can use it with Zenith:

```bash
# Create a container from the ISO image
zenith create my-ubuntu --iso ./ubuntu-24.04-live-server-amd64.iso

# Specify disk size
zenith create my-ubuntu --iso ./ubuntu-24.04-live-server-amd64.iso --disk-size 20G

# Start the container
zenith start my-ubuntu
```

#### Disk Image Management

Zenith uses `qcow2` format for virtual disk images:

```bash
# Create a new disk image
zenith disk create my-disk.qcow2 --size 30G

# Convert from ISO to disk
zenith disk convert ubuntu-24.04-live-server-amd64.iso my-disk.qcow2

# List disk images
zenith disk list
```

#### Storage Location

Downloaded ISOs and disk images should be organized:

```
~/zenith-images/
├── iso/
│   └── ubuntu-24.04-live-server-amd64.iso
└── disks/
    ├── my-ubuntu.qcow2
    └── backup-disk.qcow2
```

Specify the image directory when running Zenith:

```bash
zenith --image-dir ~/zenith-images/
```

## Usage

### CLI Commands

Zenith provides a simple command-line interface for container management:

```bash
# List all containers
zenith list

# Create a new container
zenith create [container-name] [image]

# Start a container
zenith start [container-id]

# Stop a container
zenith stop [container-id]

# Remove a container
zenith rm [container-id]

# View container logs
zenith logs [container-id]
```

## Project Structure

### Source Code
```
Zenith/
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── internal/
│   └── cli/            # CLI command implementations
├── pkg/                # Reusable packages and utilities
└── README.md           # This file
```

### User Data (Optional)
```
~/.zenith/ or ~/zenith-images/
├── iso/                # Downloaded ISO images
│   └── ubuntu-24.04-live-server-amd64.iso
├── disks/              # Virtual disk images
│   ├── my-ubuntu.qcow2
│   └── backup-disk.qcow2
└── config/             # Configuration files
```

## Architecture

### Internal CLI Package
The `internal/cli` package contains the command-line interface implementation, handling user input and delegating operations to the core management layer.

### Package Structure
The `pkg` directory houses reusable components for container operations, including container lifecycle management, configuration handling, and utility functions.

## API Reference (Coming Soon)

A comprehensive REST API is under development. This will enable:

- **Programmatic container management** via HTTP endpoints
- **Container state queries** and monitoring
- **Container lifecycle operations** through standard HTTP methods
- **Integration** with orchestration platforms and monitoring tools

Stay tuned for API documentation and endpoints.

## Development

### Building the Project

```bash
# Build the binary
go build -o zenith main.go

# Run tests (if available)
go test ./...

# Build for specific OS
GOOS=linux GOARCH=amd64 go build -o zenith main.go
```

### Project Dependencies

Dependencies are managed using Go modules. View dependencies with:

```bash
go list -m all
```

Update dependencies:

```bash
go get -u
```

## Configuration

Configuration options can typically be set via:
- Command-line flags
- Environment variables
- Configuration files (if supported)

### Common Configuration Options

```bash
# Set custom image directory
zenith --image-dir /path/to/images

# Set storage location for containers
zenith --storage-path /var/zenith/storage

# Enable debug logging
zenith --debug

# Specify configuration file
zenith --config ~/.zenith/config.yaml
```

### Environment Variables

```bash
# Set default image directory
export ZENITH_IMAGE_DIR=~/zenith-images

# Set storage path
export ZENITH_STORAGE=/var/zenith/storage

# Enable debug mode
export ZENITH_DEBUG=1
```

Refer to the CLI help for detailed configuration options:

```bash
./zenith --help
./zenith disk --help
./zenith container --help
```

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add your feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## Roadmap

- [x] Basic CLI implementation
- [ ] REST API endpoints
- [ ] Advanced container networking
- [ ] Volume management
- [ ] Environment variable support
- [ ] Docker image compatibility
- [ ] Comprehensive logging and monitoring

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or suggestions:
- Open an [issue](https://github.com/x87w/Zenith/issues) on GitHub
- Check existing issues for solutions

## Acknowledgments

- Built with Go for simplicity and performance
- Inspired by modern container management best practices

---

**Note:** This project is actively under development. The REST API is coming soon with significant new capabilities. Feedback and contributions are appreciated!