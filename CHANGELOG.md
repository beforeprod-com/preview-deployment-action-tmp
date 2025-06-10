# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Added test workflows for PR deployment and cleanup
- Added support for Go 1.24.1 in test workflow
- Added Code of Conduct
- Added comprehensive publishing guide (removed)
- Added workflow permissions requirement documentation
- Updated shpr binary to v0.0.37 with improved deployment features

### Changed
- Updated test workflow to use actions/checkout@v4 and actions/setup-go@v4
- Optimized workflow to only run on pull request events (opened, synchronize, reopened)
- Improved action and step naming for better clarity and user experience
- Updated documentation to reflect PR-only deployment strategy
- Improved attribution and branding
- Updated README with new features and improvements
- Enhanced PR comment functionality using pull request review API

### Fixed
- Fixed Docker action configuration for proper input handling
- Fixed URL capture and output from shpr app start command
- Fixed environment variable handling in actions
- Fixed build and deployment configuration for Go applications

### Removed
- Removed local shpr binary build in favor of distribution from werft repo releases (still local but put in from releases)
- Removed redundant Docker-related files and configurations

## [0.0.1] - 2024-04-28

### Added
- Preview URL logging when no PR exists for direct branch deployments
- Automatic PR description updates with deployment URLs
- Friendly footer with BeforeProd call-to-action in PR descriptions
- Support for composite action approach
- Improved error handling for URL extraction
- PR cleanup action for automatic deployment cleanup when PRs are closed
- Automatic extraction of deployment URLs from PR descriptions for cleanup
- New action directory structure for better organization and maintainability
- Separate cleanup workflow triggered on PR close events
- Automatic deployment URL extraction from PR descriptions for cleanup
- Independent binary management for each action
- Environment variables (BP_USER and BP_PASSWORD) for cleanup action

### Changed
- Improved PR description updates with retry mechanism and robust error handling
- Improved logging to show preview URL when no PR exists
- Converted from Docker-based to composite action
- Renamed step IDs from "shpr" to "beforeprod" for consistency
- Updated documentation to reflect new features and approach
- Improved workflow structure by removing redundant steps
- Reorganized actions into separate directories for better maintainability
- Updated workflow references to use new action directory structure
- Modified workflow triggers to separate deployment and cleanup events
- Updated README with comprehensive documentation for both actions
- Renamed deployment action to preview_app for better clarity
- Moved binary to action-specific directories for better version control

### Removed
- Docker-based execution approach
- Redundant URL and time output steps from workflow
- Dockerfile and entrypoint.sh as they're no longer needed

## [0.0.0] - 2024-04-14

### Added
- Initial release of the BeforeProd GitHub Action
- Support for Go applications
- Basic deployment functionality
- URL extraction from deployment output
- Secure credential handling through GitHub Secrets
