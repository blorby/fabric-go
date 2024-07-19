# Fabric Go

Fabric Go is an open-source framework for augmenting humans using AI, implemented in Go. It's a Go-based version of the original Fabric project, designed to provide efficient and scalable AI-powered text processing capabilities.

## Features

- CLI-based interaction
- Multiple AI processing patterns
- Extensible architecture for adding new patterns
- Configuration management
- Easy setup and usage

## Installation

1. Ensure you have Go 1.16 or later installed on your system.
2. Clone this repository:

```bash
# Clone Fabric to your computer
git clone https://github.com/blorby/fabric-go.git
```

3. Navigate to the project directory:

```bash
cd fabric-go
```

4. Build the project:

```bash
go build ./cmd/fabric
```

## Usage

### Setup

Before first use, run the setup command:
```bash
./fabric --setup
```
This will create necessary directories and copy initial patterns.

### List Available Patterns

To see all available processing patterns:
```bash
./fabric --list
```

### Run a Pattern

To process text using a specific pattern:
```bash
./fabric --pattern <pattern_name> --text "Your input text here"
```

Or use input from a file or pipe:
```bash
cat input.txt | ./fabric --pattern <pattern_name>
```

## Adding New Patterns

1. Create a new directory under `patterns/` with your pattern name.
2. Add a `system.md` file in this directory with the pattern's instructions.
3. Implement any necessary Go code to support the pattern.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Original Fabric project creators and contributors
- Go community for excellent libraries and tools
