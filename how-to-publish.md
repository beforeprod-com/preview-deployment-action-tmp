# How to Publish the BeforeProd GitHub Action

This document outlines the process for publishing the BeforeProd GitHub Action to the public GitHub Marketplace while maintaining private development.

## Prerequisites

- A private GitHub repository for development
- A public GitHub repository for the published action
- GitHub account with access to both repositories
- Git installed locally

## Setup Process

### 1. Create Public Repository

1. Create a new public repository on GitHub (e.g., `beforeprod/deploy-action`)
2. Initialize it with a README.md
3. Add the MIT or Apache 2.0 license
4. Add a CODE_OF_CONDUCT.md

### 2. Configure Local Development Environment

```bash
# Clone your private repository
git clone <private-repo-url>
cd <repo-name>

# Add public repository as a remote
git remote add public <public-repo-url>

# Verify remotes
git remote -v
```

## Publishing Process

### 1. Prepare Release

1. Create a new release branch from main:
   ```bash
   git checkout main
   git pull
   git checkout -b release/vX.Y.Z
   ```

2. Update version numbers and documentation:
   - Update version in action.yml
   - Update CHANGELOG.md
   - Review and update README.md
   - Remove any private/sensitive information

3. Commit changes:
   ```bash
   git add .
   git commit -m "Prepare release vX.Y.Z"
   ```

### 2. Create Release

1. Create and push the tag:
   ```bash
   git tag vX.Y.Z
   git push public release/vX.Y.Z
   git push public vX.Y.Z
   ```

2. Create GitHub Release:
   - Go to your public repository on GitHub
   - Click "Releases"
   - Click "Draft a new release"
   - Select the tag you just created
   - Fill in release title and description
   - Check "Publish this Action to the GitHub Marketplace"
   - Fill in required marketplace information:
     - Action name
     - Description
     - Icon (recommended: upload-cloud)
     - Color (recommended: blue)
     - Categories
   - Click "Publish release"

## Versioning Guidelines

Follow semantic versioning (MAJOR.MINOR.PATCH):

- MAJOR: Breaking changes
- MINOR: New features, backward compatible
- PATCH: Bug fixes, backward compatible

Example: v1.2.3
- 1: Major version
- 2: Minor version
- 3: Patch version

## Maintenance Workflow

### Regular Development

1. Work in private repository:
   ```bash
   git checkout -b feature/your-feature
   # Make changes
   git add .
   git commit -m "Add your feature"
   git push origin feature/your-feature
   ```

2. Create pull request to main branch
3. Review and merge
4. Run tests and CI

### Release Process

1. Create release branch:
   ```bash
   git checkout main
   git pull
   git checkout -b release/vX.Y.Z
   ```

2. Prepare release:
   - Update version numbers
   - Update documentation
   - Remove private information
   - Run final tests

3. Create and push release:
   ```bash
   git add .
   git commit -m "Prepare release vX.Y.Z"
   git tag vX.Y.Z
   git push public release/vX.Y.Z
   git push public vX.Y.Z
   ```

4. Create GitHub Release in the public repository

## Best Practices

1. **Documentation**:
   - Keep README.md up to date
   - Document all inputs and outputs
   - Provide usage examples
   - Include troubleshooting guide

2. **Testing**:
   - Run all tests before release
   - Test with different GitHub workflows
   - Verify all inputs and outputs

3. **Security**:
   - Remove any sensitive information
   - Sanitize logs and outputs
   - Review dependencies for vulnerabilities

4. **Version Management**:
   - Use semantic versioning
   - Keep CHANGELOG.md updated
   - Document breaking changes

## Troubleshooting

### Common Issues

1. **Release not appearing in Marketplace**:
   - Verify you checked "Publish this Action to the GitHub Marketplace"
   - Check that all required marketplace information is filled
   - Ensure the repository is public

2. **Version tag issues**:
   - Verify tag format (vX.Y.Z)
   - Check if tag already exists
   - Ensure tag is pushed to public repository

3. **Authentication issues**:
   - Verify GitHub token permissions
   - Check repository access
   - Ensure proper SSH keys are set up

## Support

For issues with the publishing process:
1. Check this document
2. Review GitHub's documentation
3. Contact the development team

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Creating a GitHub Action](https://docs.github.com/en/actions/creating-actions)
- [Publishing Actions in GitHub Marketplace](https://docs.github.com/en/actions/creating-actions/publishing-actions-in-github-marketplace)
- [Semantic Versioning](https://semver.org/)
