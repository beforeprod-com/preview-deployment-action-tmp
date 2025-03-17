# Werft Deployment GitHub Action Development

This repository contains the source code for the Werft Deployment GitHub Action. For usage instructions, please see [.github/README.md](.github/README.md).

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

1. Clone the repository
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

### Testing

The repository includes a test application in `go-test-app/` for development and testing purposes. The test workflow in `.github/workflows/test.yml` demonstrates how to use the action.

### Building the Action

The action is containerized using Alpine Linux for minimal size. The build process:

1. Copies the entrypoint script and Werft CLI binary
2. Sets up the necessary environment
3. Configures the container to run the deployment process

### Contributing

1. Create a feature branch
2. Make your changes
3. Test locally using the test application
4. Submit a pull request

### CI/CD

The repository uses GitHub Actions for CI/CD:
- Automated testing on push
- Action validation
- Documentation updates

## Troubleshooting

### Common Issues

1. **Action fails to start**
   - Check if the Werft CLI binary is present in `bin/shpr`
   - Verify the entrypoint script has execute permissions

2. **Deployment fails**
   - Verify Werft credentials are correctly set
   - Check the build folder path is correct
   - Ensure the platform is supported

3. **Container issues**
   - Verify Docker is running
   - Check container logs for detailed error messages

## Maintenance

### Updating Dependencies

- Update Go version in workflows if needed
- Update Alpine base image in Dockerfile
- Update Werft CLI binary in `bin/shpr`

### Version Management

The action follows semantic versioning. Update version tags when:
- Adding new features (minor version)
- Fixing bugs (patch version)
- Making breaking changes (major version)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
