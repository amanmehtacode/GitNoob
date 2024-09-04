# GitNoob
**GitNoob** is a Python library designed to simplify common Git commands for new developers. It provides easy-to-use command-line tools to streamline Git workflows and help you manage your version control more efficiently.

## Installation
Since GitNoob is currently not available on PyPI, you can install it directly from the GitHub repository. Clone the repository and install it locally:

```bash
git clone https://github.com/amanmehtacode/GitNoob.git
cd GitNoob
pip install .
```

## Commands

### `lazypush`
The `lazypush` command simplifies the process of committing and pushing changes to a Git repository. It includes options to pull the latest changes before pushing and supports verbose output to help you understand what's happening under the hood.

#### Usage
```bash
lazypush [options]
```

#### Options
- `-p`, `--pull`: Pull the latest changes from the remote branch before pushing.
- `-v`, `--verbose`: Enable verbose output to show detailed information about the actions being performed.

#### Examples
1. **Basic Usage:**
   ```bash
   lazypush
   ```
   Stages all changes, commits them with a user-provided message, and pushes them to the current branch.

2. **Pull Before Push:**
   ```bash
   lazypush -p
   ```
   Pulls the latest changes from the remote branch before pushing the committed changes.

3. **Verbose Mode:**
   ```bash
   lazypush -v
   ```
   Runs the command in verbose mode, providing detailed messages about the process.

4. **Pull and Verbose Mode:**
   ```bash
   lazypush -p -v
   ```
   Combines pulling the latest changes with verbose mode for an in-depth look at the operations.

## Under Development
GitNoob is a work in progress. The `lazypush` command is just the beginning. Future updates will include more commands and features designed to make Git management easier and more intuitive. Keep an eye on the [GitHub repository](https://github.com/amanmehtacode/GitNoob) for new releases and updates.

## Contributing
We welcome contributions! If you have ideas for new features or improvements, or if you find a bug, please open an issue or submit a pull request on our [GitHub repository](https://github.com/yourusername/GitNoob).

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact
For any questions or feedback, please reach out to:
- **Author:** Aman Mehta
