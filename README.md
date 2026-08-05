# Zenith

A lightweight container manager written in Go with a simple CLI and REST API support (coming soon).

## Features

- **Lightweight** - Minimal dependencies
- **Simple CLI** - Easy-to-use command-line interface
- **Container Management** - Create, start, stop containers
- **REST API** (Coming Soon)

## Installation

```bash
git clone https://github.com/x87w/Zenith.git
cd Zenith
go build -o zenith main.go
```

## Quick Start

```bash
./zenith --help
```

## Downloading Ubuntu Server ISO

#### Download

```bash
# Ubuntu Server 24.04 LTS
wget https://releases.ubuntu.com/24.04/ubuntu-24.04-live-server-amd64.iso

# Verify (optional)
wget https://releases.ubuntu.com/24.04/SHA256SUMS
sha256sum -c SHA256SUMS 2>&1 | grep ubuntu-24.04-live-server-amd64.iso
```

Other versions: https://releases.ubuntu.com/

#### Using with Zenith

```bash
# Create a container from ISO
./zenith create my-ubuntu --iso ./ubuntu-24.04-live-server-amd64.iso --disk-size 20G

# Start the container
./zenith start my-ubuntu

# List containers
./zenith list
```

#### Disk Management

```bash
# Create a disk image
./zenith disk create my-disk.qcow2 --size 30G

# Convert ISO to disk
./zenith disk convert ubuntu-24.04-live-server-amd64.iso my-disk.qcow2
```

## Storage

Specify the image directory:

```bash
./zenith --image-dir ~/zenith-images/
```

## Configuration

```bash
# Command-line
./zenith --image-dir /path/to/images --debug

# Environment variables
export ZENITH_IMAGE_DIR=~/zenith-images
export ZENITH_DEBUG=1
```

## Development

```bash
# Build
go build -o zenith main.go

# Tests
go test ./...

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o zenith main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push and open a Pull Request

## Support

Open an [issue](https://github.com/x87w/Zenith/issues) on GitHub

---

**Status:** Under active development. REST API coming soon!