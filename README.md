# Taxonomist

A powerful CLI tool for generating consistent, memorable names for your resources, perfect for cloud infrastructure, development environments, or any naming needs.

## Features

- Generate consistent, memorable names
- Multiple output formats (kebab, camel, snake case)
- Customizable with prefixes and postfixes
- Cross-platform support (Windows, Linux, macOS)
- Configurable through YAML
- Perfect for CI/CD pipelines

## Installation

```bash
# Using go install
go install github.com/baking-bread/taxonomist@latest

# Or clone and build
git clone https://github.com/baking-bread/taxonomist.git
cd taxonomist
make build
```

## Usage

Basic usage:

```bash
taxonomist              # Generates a single name
taxonomist -n 5        # Generates 5 names
taxonomist -a 2        # Uses 2 adjectives
```

Advanced options:

```bash
taxonomist -p dev -s prod    # Adds prefix and postfix
taxonomist -f camel         # Uses camelCase format
taxonomist -c custom.yaml   # Uses custom dictionary
```

### Configuration

Default configuration can be overridden using a YAML file:

```yaml
dictionary:
  adjectives:
    - custom
    - words
  nouns:
    - here
```

## Building from Source

Requirements:

- Go 1.22 or higher
- Make

```bash
make deps    # Install dependencies
make build   # Build for current platform
make test    # Run tests
make lint    # Run linters
```

Cross-platform builds:

```bash
make windows  # Build for Windows
make linux    # Build for Linux
make darwin   # Build for macOS
```

## Security

### Vulnerability Scanning

The project uses Trivy for vulnerability scanning:

```bash
# Run security scan locally
make scan

# Or run directly with Trivy
trivy fs --security-checks vuln,config .
```

### Binary Verification

Each release includes SHA256 hash files for binary verification:

```bash
# Verify on Linux/MacOS
sha256sum -c taxonomist.sha256

# Verify on Windows
certutil -hashfile taxonomist.exe | findstr /i /c:"$(type taxonomist.exe.sha256)"
```

## CI/CD Integration

Perfect for GitOps and infrastructure as code:

```bash
taxonomist -n 1 -p prod -f kebab > resource-name.txt
```
