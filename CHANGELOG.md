# Changelog

## [Unreleased]

### Added

- Core name generation functionality with random adjective-noun combinations
- Configuration system with YAML support for word dictionaries
- Command line interface with customization flags:
  - `-c, --config`: Specify configuration file path
  - `-n, --count`: Number of names to generate
  - `-p, --prefix`: Add prefix to generated names
  - `-s, --postfix`: Add postfix to generated names
  - `-e, --separator`: Customize word separator (default: "-")
  - `-a, --adjectives`: Number of adjectives to use (default: 1)
  - `-f, --format`: Output format selection
- Multiple name formatting options:
  - Kebab case: `swift-azure-falcon` (default)
  - Camel case: `swiftAzureFalcon`
  - Snake case: `swift_azure_falcon`
- Environment variable support:
  - `CONFIG_FILE`: Override configuration file path

### Updated

- Name generation to support multiple adjectives
- Output formatting with configurable separators
- Command line interface with improved options
