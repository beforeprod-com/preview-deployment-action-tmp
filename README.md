# Werft Deployment GitHub Action

A powerful GitHub Action for automated deployments using Werft. This action provides a streamlined way to deploy applications across different platforms, with built-in support for Go applications and extensibility for other platforms.

## Features

- 🚀 Automated deployments through Werft
- 🔧 Multi-platform support (Go and more)
- 🐳 Containerized execution for consistency
- 🔒 Secure credential handling
- 🧪 Built-in testing capabilities
- 📦 Minimal container size using Alpine Linux

## Quick Start

Add this action to your GitHub Actions workflow:

```yaml
name: Deploy
on:
  push:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Deploy with Werft
        uses: your-org/werft-deployment-action@v1
        with:
          platform: GO
          build-folder: ./build
        env:
          WERFT_TOKEN: ${{ secrets.WERFT_TOKEN }}
```

## Development Setup

### Prerequisites

- Go 1.24.1 or later
- Docker
- GitHub CLI (optional, for testing)

### Project Structure

```
.
├── .github/           # GitHub Action specific files
│   ├── actions/      # Action implementation
│   ├── workflows/    # CI/CD workflows
│   └── README.md     # User-facing documentation
├── bin/              # Binary files
│   └── shpr         # Werft CLI binary
├── go-test-app/      # Test application for development
├── Dockerfile        # Container definition for the action
└── entrypoint.sh     # Action entrypoint script
```

### Local Development

1. Clone the repository:
   ```bash
   git clone https://github.com/your-org/werft-deployment-action.git
   cd werft-deployment-action
   ```

2. Build the test application:
   ```bash
   cd go-test-app
   go build -o build/app main.go
   ```

3. Test the action locally:
   ```bash
   # Using act (recommended)
   act -j deploy

   # Or manually using Docker
   docker build -t werft-deployment-action .
   docker run -e INPUT_PLATFORM=GO -e INPUT_BUILD_FOLDER=./go-test-app/build werft-deployment-action
   ```

## Configuration

### Input Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `platform` | Yes | Target platform (e.g., GO) |
| `build-folder` | Yes | Path to the build output directory |
| `additional-args` | No | Additional arguments for the Werft CLI |

### Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `WERFT_TOKEN` | Yes | Authentication token for Werft |
| `WERFT_ENDPOINT` | No | Custom Werft endpoint URL |

## Testing

The repository includes a comprehensive test suite:

1. **Unit Tests**: Run with `go test ./...`
2. **Integration Tests**: Use the test application in `go-test-app/`
3. **End-to-End Tests**: Available in `.github/workflows/test.yml`

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and coding standards
- Add tests for new features
- Update documentation as needed
- Keep commits clean and well-documented

## Troubleshooting

### Common Issues

1. **Action fails to start**
   - Check if the Werft CLI binary is present in `bin/shpr`
   - Verify the entrypoint script has execute permissions
   - Check GitHub Actions logs for detailed error messages

2. **Deployment fails**
   - Verify Werft credentials are correctly set
   - Check the build folder path is correct
   - Ensure the platform is supported
   - Review Werft deployment logs

3. **Container issues**
   - Verify Docker is running
   - Check container logs for detailed error messages
   - Ensure sufficient disk space is available

## Maintenance

### Updating Dependencies

1. Update Go version in workflows if needed
2. Update Alpine base image in Dockerfile
3. Update Werft CLI binary in `bin/shpr`
4. Run tests to ensure compatibility

### Version Management

The action follows semantic versioning:
- Major version: Breaking changes
- Minor version: New features
- Patch version: Bug fixes

## Support

- GitHub Issues: [Report bugs](https://github.com/your-org/werft-deployment-action/issues)
- Documentation: [User Guide](.github/README.md)
- Community: [Discussions](https://github.com/your-org/werft-deployment-action/discussions)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Werft team for their excellent deployment platform
- GitHub Actions community
- Contributors and maintainers
