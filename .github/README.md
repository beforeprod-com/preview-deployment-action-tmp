# Werft Deployment GitHub Action

This GitHub Action enables automated deployment of applications to beforeprod.com, a platform for preview deployments. It provides a streamlined way to deploy applications as part of your CI/CD pipeline.

## Features

- Automated deployment to beforeprod.com platform
- Support for Go applications
- Preview URL generation for deployed applications
- Simple configuration through GitHub Actions workflow
- Follows GitHub Actions standards for custom actions

## Usage

To use this action in your GitHub workflow, add the following to your workflow file:

```yaml
- name: Werft Deployment
  uses: ./.github/actions
  id: werft
  with:
    build_folder: './path/to/build'
    platform: 'GO'  # Currently supports GO platform
  env:
    BP_USER: ${{ secrets.BP_USER }}
    BP_PASSWORD: ${{ secrets.BP_PASSWORD }}
```

### Inputs

- `build_folder`: The path to your built application files
- `platform`: The platform of your application (currently supports 'GO')

### Environment Variables

- `BP_USER`: Werft platform username (required)
- `BP_PASSWORD`: Werft platform password (required)

### Outputs

- `url`: The URL of your deployed preview application
- `time`: Timestamp of the deployment

## Example Workflow

Here's a complete example of how to use this action in your workflow:

```yaml
name: Deploy to Werft

on: [push]

jobs:
  deploy:
    runs-on: ubuntu-latest
    name: Deploy application to Werft
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.24.1

      - name: Build
        run: env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./build/app ./main.go

      - name: Deploy to Werft
        uses: ./.github/actions
        id: werft
        with:
          build_folder: './build'
          platform: 'GO'
        env:
          BP_USER: ${{ secrets.BP_USER }}
          BP_PASSWORD: ${{ secrets.BP_PASSWORD }}

      - name: Get Preview URL
        run: echo "Preview URL: ${{ steps.werft.outputs.url }}"
```

## Development

This action is built using:
- Alpine Linux as the base container
- Shell script for the entrypoint
- Docker for containerization
- Follows GitHub Actions standards for custom actions

### Local Development

1. Clone the repository
2. Make your changes in the `.github/actions` directory
3. Test locally using `act` or by pushing to a test repository
4. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
