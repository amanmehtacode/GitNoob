# GitNoob

GitNoob is a collection of command-line tools designed to simplify and automate various Git and GitHub operations. It includes the following tools:

- **autobranch**: Automatically creates and manages Git branches based on commit messages.
- **autocommit**: Automatically commits changes with a generated message.
- **automerge**: Automatically merges all branches into the main branch.
- **deleterepo**: Deletes a GitHub repository.
- **lazypush**: Simplifies the process of adding, committing, and pushing changes to a Git repository.
- **lazyrepo**: Sets up a new Git repository with a predefined structure and publishes it to GitHub.
- **newrepo**: Creates a new Git repository and publishes it to GitHub.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/GitNoob.git
    cd GitNoob
    ```

2. Build the tools:

    ```sh
    go build -o autobranch ./cmd/autobranch
    go build -o autocommit ./cmd/autocommit
    go build -o automerge ./cmd/automerge
    go build -o deleterepo ./cmd/deleterepo
    go build -o lazypush ./cmd/lazypush
    go build -o lazyrepo ./cmd/lazyrepo
    go build -o newrepo ./cmd/newrepo
    ```

3. Move the binaries to a directory in your PATH:

    ```sh
    mv autobranch /usr/local/bin/
    mv autocommit /usr/local/bin/
    mv automerge /usr/local/bin/
    mv deleterepo /usr/local/bin/
    mv lazypush /usr/local/bin/
    mv lazyrepo /usr/local/bin/
    mv newrepo /usr/local/bin/
    ```

## Usage

### autobranch

Automatically creates and manages Git branches based on commit messages.

```sh
autobranch
```

### autocommit

Automatically commits changes with a generated message.

```sh
autocommit
```

### automerge

Automatically merges all branches into the main branch.

```sh
automerge
```

### deleterepo

Deletes a GitHub repository.

```sh
deleterepo --name <repository-name>
```

### lazypush

Simplifies the process of adding, committing, and pushing changes to a Git repository.

```sh
lazypush
```

### lazyrepo

Sets up a new Git repository with a predefined structure and publishes it to GitHub.

```sh
lazyrepo --name <repository-name>
```

### newrepo

Creates a new Git repository and publishes it to GitHub.

```sh
newrepo --name <repository-name>
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) for CLI framework.
- [Viper](https://github.com/spf13/viper) for configuration management.
- [go-git](https://github.com/go-git/go-git) for Git library.
- [spinner](https://github.com/briandowns/spinner) for terminal spinner.
- [color](https://github.com/fatih/color) for colorized output.
more detailed and read the file names to make up the commands and fucntionality that i am gonna be adding later
# GitNoob

GitNoob is a comprehensive suite of command-line tools designed to simplify and automate various Git and GitHub operations. It's built with Go and leverages several powerful libraries to provide a streamlined and efficient workflow for developers.

## Tools

The GitNoob suite includes the following tools:

### autobranch

The `autobranch` tool automatically creates and manages Git branches based on commit messages. It's designed to streamline the process of branching in Git, making it easier to manage multiple lines of development.

Command: `autobranch`

### autocommit

The `autocommit` tool automatically stages and commits changes with a generated message. This is useful for quickly committing changes without having to manually stage files or write a commit message.

Command: `autocommit`

### automerge

The `automerge` tool automatically merges all branches into the main branch. This can be useful for consolidating changes from multiple branches.

Command: `automerge`

### deleterepo

The `deleterepo` tool deletes a GitHub repository. This can be useful for cleaning up old or unnecessary repositories.

Command: `deleterepo --name <repository-name>`

### lazypush

The `lazypush` tool simplifies the process of adding, committing, and pushing changes to a Git repository. It combines these three operations into a single command, making it easier to quickly save and push changes.

Command: `lazypush`

### lazyrepo

The `lazyrepo` tool sets up a new Git repository with a predefined structure and publishes it to GitHub. This can be useful for quickly setting up new projects.

Command: `lazyrepo --name <repository-name>`

### newrepo

The `newrepo` tool creates a new Git repository and publishes it to GitHub. This can be useful for quickly creating new repositories without having to manually create them on GitHub.

Command: `newrepo --name <repository-name>`

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/GitNoob.git
    cd GitNoob
    ```

2. Build the tools:

    ```sh
    go build -o autobranch ./cmd/autobranch/main.go
    go build -o autocommit ./cmd/autocommit/main.go
    go build -o automerge ./cmd/automerge/main.go
    go build -o deleterepo ./cmd/deleterepo/main.go
    go build -o lazypush ./cmd/lazypush/main.go
    go build -o lazyrepo ./cmd/lazyrepo/main.go
    go build -o newrepo ./cmd/newrepo/main.go
    ```

3. Move the binaries to a directory in your PATH:

    ```sh
    mv autobranch /usr/local/bin/
    mv autocommit /usr/local/bin/
    mv automerge /usr/local/bin/
    mv deleterepo /usr/local/bin/
    mv lazypush /usr/local/bin/
    mv lazyrepo /usr/local/bin/
    mv newrepo /usr/local/bin/
    ```

## Usage

Each tool in the GitNoob suite can be used as a standalone command-line tool. For detailed usage instructions, refer to the individual tool sections above.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have a feature request, bug report, or want to improve the documentation or code.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Cobra](https://github.com/spf13/cobra) for providing a powerful library for creating command-line interfaces.
- [Viper](https://github.com/spf13/viper) for handling configuration with ease.
- [go-git](https://github.com/go-git/go-git) for offering a highly extensible Git implementation in pure Go.
- [spinner](https://github.com/briandowns/spinner) for terminal spinner.
- [color](https://github.com/fatih/color) for colorized output.
