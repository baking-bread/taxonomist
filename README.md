# Taxonomist

[![Build](https://github.com/baking-bread/taxonomist/actions/workflows/default.yaml/badge.svg?branch=main)](https://github.com/baking-bread/taxonomist/actions/workflows/default.yaml)

A simple CLI tool for generating consistent, memorable names for your resources, perfect for cloud infrastructure, development environments, or any naming needs.

## Features

- Generate consistent, memorable names
- Multiple output formats (kebab, camel, snake case)
- Customizable with prefixes and sufixes
- Cross-platform support (Windows, Linux, macOS)
- Configurable through YAML (Provide your own words)
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
taxonomist             # Generates a single name
taxonomist -n 5        # Generates 5 names
taxonomist -a 2        # Uses 2 adjectives


```

Supported Flags

```bash
  -a, --adjectives int     Number of adjectives to use in the name (default 1)
  -c, --config string      Path to the configuration file (optional)
  -n, --count int          Number of names to generate (default 1)
  -d, --debug              Enable debug logging
  -f, --format string      Output format (kebab, camel, snake) (default "kebab")
  -h, --help               help for taxonomist
  -p, --prefix string      Prefix to add to generated names
  -e, --separator string   Separator to use between prefix, generated name, and sufix (default "-")
  -s, --sufix string       sufix to add to generated names
```

### Configuration

Default configuration can be overridden using a YAML file:

```yaml
adjectives:
  - custom
  - words
nouns:
  - here
```
