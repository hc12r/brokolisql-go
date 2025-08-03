# GitHub Actions for BrokoliSQL-Go

This document explains the GitHub Actions workflows set up for the BrokoliSQL-Go project and how to use them.

## Available Workflows

### 1. Go Tests (`tests.yml`)

This workflow runs the test suite and generates code coverage reports.

- **Triggers**: Runs on pushes to `main` branch and on pull requests
- **Features**:
  - Runs all tests with verbose output
  - Generates code coverage reports
  - Uploads coverage data to Codecov

### 2. Code Quality (`code-quality.yml`)

This workflow checks code quality using linters and formatters.

- **Triggers**: Runs on pushes to `main` branch and on pull requests
- **Features**:
  - Runs golangci-lint to check for code issues
  - Verifies that all code is properly formatted with gofmt

### 3. Cross-Platform Tests (`cross-platform.yml`)

This workflow ensures the code works across different operating systems and Go versions.

- **Triggers**: Runs on pushes to `main` branch and on pull requests
- **Features**:
  - Tests on Ubuntu, Windows, and macOS
  - Tests with Go 1.24+

### 4. Release Automation (`release.yml`)

This workflow automates the release process when a new version is tagged.

- **Triggers**: Runs when a tag starting with 'v' is pushed (e.g., v1.0.0)
- **Features**:
  - Uses GoReleaser to build binaries for multiple platforms
  - Creates a GitHub Release with the built binaries
  - Generates release notes

### 5. Dependency Updates (`dependencies.yml`)

This workflow keeps dependencies up-to-date.

- **Triggers**: Runs weekly on Mondays and can be manually triggered
- **Features**:
  - Updates all Go dependencies to their latest versions
  - Creates a pull request with the updates

### 6. Security Scan (`security.yml`)

This workflow scans the code for security vulnerabilities.

- **Triggers**: Runs on pushes to `main` branch, on pull requests, and weekly on Sundays
- **Features**:
  - Uses Gosec to scan for security issues in the code
  - Uses govulncheck to check for vulnerabilities in the code
  - Uses Nancy to scan dependencies for known vulnerabilities

## How to Use These Workflows

### Running Workflows Manually

Some workflows can be triggered manually from the GitHub Actions tab:

1. Go to the "Actions" tab in your GitHub repository
2. Select the workflow you want to run
3. Click the "Run workflow" button
4. Select the branch and click "Run workflow"

### Creating a Release

To create a new release:

1. Create and push a new tag with a version number:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
2. The release workflow will automatically build the binaries and create a GitHub Release

### Customizing Workflows

You can customize these workflows by editing the YAML files in the `.github/workflows/` directory:

- **Changing Go versions**: Update the `go-version` field in the workflow files
- **Adding more operating systems**: Add to the `os` list in the matrix strategy
- **Modifying test commands**: Change the `run` commands in the test steps
- **Adjusting schedule**: Modify the `cron` expressions in the `schedule` section

## Workflow Status Badges

You can add status badges to your README.md to show the current status of your workflows:

```markdown
![Tests](https://github.com/yourusername/brokolisql-go/actions/workflows/tests.yml/badge.svg)
![Code Quality](https://github.com/yourusername/brokolisql-go/actions/workflows/code-quality.yml/badge.svg)
![Security Scan](https://github.com/yourusername/brokolisql-go/actions/workflows/security.yml/badge.svg)
```

## Troubleshooting

If a workflow fails, you can:

1. Check the workflow run in the GitHub Actions tab for detailed logs
2. Look for specific error messages in the failed step
3. Make necessary code or configuration changes
4. Push the changes to trigger the workflow again

For more complex issues, refer to the [GitHub Actions documentation](https://docs.github.com/en/actions) or the documentation for the specific action being used.