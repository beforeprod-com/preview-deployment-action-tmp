# Beforeprod.com Deployment GitHub Action

This GitHub Action enables automated deployment of applications to beforeprod.com, a platform for preview deployments. It provides a streamlined way to deploy applications as part of your CI/CD pipeline.

## Features

- Automated deployment to beforeprod.com platform
- Support for Go applications
- Preview URL generation for deployed applications
- Simple configuration through GitHub Actions workflow
- Follows GitHub Actions standards for custom actions

## Usage

> **Important**: Before using this action, ensure that your repository has the necessary permissions enabled. Go to your repository's Settings > Actions > Workflow Permissions and enable "Read and Write permissions". This is required for the action to function properly, particularly for features like updating Pull Request bodies.

To use this action in your GitHub workflow, add the following to your workflow file:

```yaml
- name: beforeprod.com Deployment
  uses: ./.github/actions
  id: beforeprod
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

- `BP_USER`: beforeprod.com platform username (required)
- `BP_PASSWORD`: beforeprod.com platform password (required)

### Outputs

- `url`: The URL of your deployed preview application
- `time`: Timestamp of the deployment

## Example Workflow

Here's a complete example of how to use this action in your workflow:

```yaml
name: Deploy to beforeprod.com

on: [push]

jobs:
  deploy:
    runs-on: ubuntu-latest
    name: Deploy application to beforeprod.com
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.24.1

      - name: Build
        run: env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./build/app ./main.go

      - name: Deploy to beforeprod.com
        uses: ./.github/actions
        id: beforeprod
        with:
          build_folder: './build'
          platform: 'GO'
        env:
          BP_USER: ${{ secrets.BP_USER }}
          BP_PASSWORD: ${{ secrets.BP_PASSWORD }}

      - name: Get Preview URL
        run: echo "Preview URL: ${{ steps.beforeprod.outputs.url }}"
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

This project is licensed under the Apache 2.0 License - see the LICENSE file for details.
