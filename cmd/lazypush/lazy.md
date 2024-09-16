# Lazypush

Lazypush is a command-line tool designed to simplify the process of adding, committing, and pushing changes to a Git repository. It allows users to perform these actions with minimal effort, making it ideal for developers who want to streamline their workflow.

## Features

- **Lazy Add, Commit, and Push**: Automatically stages, commits, and pushes changes to the remote repository.
- **Pull Before Push**: Optionally pull the latest changes from the remote branch before pushing.
- **Verbose Mode**: Enable detailed output for each operation.
- **Custom Commit Messages**: Enter a custom commit message or use a default message.
- **Conflict Resolution**: Attempts to resolve conflicts by pulling and rebasing before pushing again.
- **Colorized Output**: Uses color-coded output for better readability.
- **Progress Spinner**: Displays a spinner during long-running operations.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/amanmehtacode/lazypush.git
   cd lazypush
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. Then, run:
   ```bash
   go get github.com/briandowns/spinner
   go get github.com/fatih/color
   go get github.com/spf13/cobra
   ```

3. **Build the Application**:
   Compile the application using:
   ```bash
   go build -o lazypush cmd/lazypush/main.go
   ```

## Usage

### Basic Command

To use `lazypush`, run the following command in your Git repository:
