# BeforeProd GitHub Action

This GitHub Action automatically deploys your application to [BeforeProd](https://beforeprod.com) and updates your PR description with the deployment URL. It also includes automatic cleanup of deployments when PRs are closed.

## Features

- 🚀 Automatic deployments to BeforeProd
- 📝 Automatic PR description updates with deployment URLs
- 🔄 Automatic cleanup of deployments when PRs are closed
- 🛠️ Support for both Go and JavaScript applications
- 🔒 Secure credential handling
- 📦 Independent binary management for each action (currently because of KISS)
- 📊 Preview URL logging for direct branch deployments (no PR required)
- 🔄 Robust error handling with automatic retries for PR updates

## Usage

### Deployment Action

The deployment action (`.github/actions/preview_app`) handles the deployment of your application. Add the following to your workflow file (e.g., `.github/workflows/deploy.yml`):

```yaml
name: BeforeProd preview app action
on:
  pull_request:
    types: [opened, synchronize, reopened]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Create a preview app on beforeprod.com
        uses: ./.github/actions/preview_app
        with:
          platform: 'JS'  # or 'GO' for Go applications
          build_folder: './build'  # path to your build artifacts
```

> **Note**: The action will automatically update PR descriptions with deployment URLs. The action only runs on pull request events (opened, updated, or reopened) to ensure deployments are only created when needed.

### Cleanup Action

The cleanup action (`.github/actions/cleanup`) automatically removes deployments when PRs are closed. Add this to a separate workflow file (e.g., `.github/workflows/cleanup.yml`):

```yaml
name: Cleanup PR Deployments
on:
  pull_request:
    types: [closed]

jobs:
  cleanup:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: Cleanup deployment
        uses: ./.github/actions/cleanup
        env:
          BP_USER: ${{ secrets.BP_USER }}
          BP_PASSWORD: ${{ secrets.BP_PASSWORD }}
```

## Action Structure

The action is organized into two main components:

1. **Preview App Action** (`.github/actions/preview_app/`)
   - Handles the deployment of your application
   - Updates PR descriptions with deployment URLs
   - Triggered on `pull_request` events (opened, synchronize, reopened)
   - Contains its own copy of the BeforeProd CLI binary

2. **Cleanup Action** (`.github/actions/cleanup/`)
   - Automatically cleans up deployments when PRs are closed
   - Removes the deployment from BeforeProd
   - Triggered on `pull_request` events with type `closed`
   - Contains its own copy of the BeforeProd CLI binary

Each action maintains its own copy of the BeforeProd CLI binary, allowing for independent version management and updates.

## Inputs

### Preview App Action Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `platform` | The platform your application runs on (`JS` or `GO`) | Yes | `JS` |
| `build_folder` | The folder containing your build artifacts | Yes | `./build` |

### Cleanup Action Inputs

The cleanup action doesn't require any inputs as it automatically extracts the necessary information from the PR description.

## Outputs

### Preview App Action Outputs

| Output | Description |
|--------|-------------|
| `url` | The URL where your application is deployed |
| `time` | The timestamp when the deployment was completed |

## Example

See the [example workflows](.github/workflows/) for complete examples of how to use both actions.

## License

This software is proprietary and confidential. Unauthorized copying, distribution, or use is strictly prohibited. All rights reserved.
