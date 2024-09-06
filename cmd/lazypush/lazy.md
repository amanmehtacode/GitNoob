# Lazypush

Lazypush is a command-line tool designed to simplify the process of adding, committing, and pushing changes to a Git repository. It allows users to perform these actions with minimal effort, making it ideal for developers who want to streamline their workflow.

## Features

- **Lazy Add, Commit, and Push**: Automatically stages, commits, and pushes changes to the remote repository.
- **Pull Before Push**: Optionally pull the latest changes from the remote branch before pushing.
- **Verbose Mode**: Enable detailed output for each operation.
- **Custom Commit Messages**: Enter a custom commit message or use a default message.

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/lazypush.git
   cd lazypush
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. Then, run:
   ```bash
   go get ./...
   ```

3. **Build the Application**:
   Compile the application using:
   ```bash
   go build -o lazypush cmd/lazypush/main.go
   ```

## Usage

### Command Line Options

- `-p`, `--pull`: Pull the latest changes from the remote branch before pushing.
- `-v`, `--verbose`: Enable verbose output for detailed logging.

### Example

To use `lazypush`, run the following command in your terminal:
