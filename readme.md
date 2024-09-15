
# LazyRepo

LazyRepo is a powerful command-line tool written in Go that automates the process of setting up a new Git repository both locally and on GitHub. It streamlines the process of initializing a local repository, creating a remote repository on GitHub, and pushing the initial commit, saving you time and effort.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Example](#example)
- [License](#license)
- [Contributing](#contributing)
- [Acknowledgements](#acknowledgements)

## Features

- **Local Repository Initialization**: LazyRepo initializes a local Git repository in a directory of your choice.
- **GitHub Repository Creation**: It creates a remote repository on GitHub using your personal access token.
- **Initial Commit Push**: LazyRepo adds the remote repository and pushes the initial commit, which includes a `README.md` file.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (1.16 or later)
- Git
- A GitHub account and a personal access token with repository creation permissions.

## Installation

Follow these steps to install LazyRepo:

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/lazyrepo.git
    cd lazyrepo
    ```

2. Build the project:

    ```sh
    go build -o lazyrepo ./cmd/lazyrepo
    ```

3. Move the binary to a directory in your PATH:

    ```sh
    mv lazyrepo /usr/local/bin/
    ```

## Usage

Here's how to use LazyRepo:

1. Set the `GITHUB_TOKEN` environment variable with your GitHub personal access token:

    ```sh
    export GITHUB_TOKEN=your_github_token
    ```

2. Optionally, set the `GITHUB_USERNAME` environment variable with your GitHub username:

    ```sh
    export GITHUB_USERNAME=your_github_username
    ```

3. Run the `lazyrepo` command with the desired repository name:

    ```sh
    lazyrepo setup-repo <repository-name>
    ```

    Example:

    ```sh
    lazyrepo setup-repo my-new-repo
    ```

4. Use the `-v` or `--verbose` flag for verbose output:

    ```sh
    lazyrepo setup-repo my-new-repo -v
    ```

## Project Structure

The project has the following structure:

```
lazyrepo/
├── cmd/
│   └── lazyrepo/
│       └── main.go
├── go.mod
├── go.sum
└── README.md
```

- `cmd/lazyrepo/main.go`: The main entry point for the LazyRepo command-line tool.
- `go.mod`: The Go module file, which includes the project dependencies.
- `go.sum`: The Go checksum file, which ensures the integrity of the project dependencies.
- `README.md`: This file, which provides documentation for the project.

## Configuration

LazyRepo uses the following environment variables:

- `GITHUB_TOKEN`: Your GitHub personal access token. This is required to create a new repository on GitHub.
- `GITHUB_USERNAME`: Your GitHub username. This is optional but recommended for a smoother experience.

## Example

Here's an example of how to use LazyRepo:

```sh
export GITHUB_TOKEN=your_github_token
export GITHUB_USERNAME=your_github_username
lazyrepo setup-repo my-new-repo -v
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! If you have a feature request, bug report, or want to improve the documentation, please open an issue or submit a pull request.

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) for providing a powerful library for creating command-line interfaces.
- [Viper](https://github.com/spf13/viper) for handling configuration with ease.
- [go-git](https://github.com/go-git/go-git) for offering a highly extensible Git implementation in pure Go.
