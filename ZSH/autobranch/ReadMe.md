# AutoBranch Script

The **AutoBranch** script simplifies the process of creating and pushing Git branches by interactively prompting users for branch type and name. It ensures a clean working directory, provides the option to push the newly created branch to a remote repository, and includes a verbose mode for additional logging.

## Features
- **Interactive Branch Creation**: Automatically asks for branch type and name, ensuring a consistent naming format.
- **Branch Safety**: Prevents branch creation if there are unstaged changes or if the branch already exists.
- **Optional Remote Push**: Prompts the user to push the new branch to the remote repository.
- **Verbose Mode**: Provides additional logging to track each step of the process.

## Installation

1. Clone this repository or download the script directly.
2. Ensure the script is executable:

    ```bash
    chmod +x autobranch
    ```

3. Optionally, move the script to a directory in your `PATH` for easier access:

    ```bash
    mv autobranch /usr/local/bin/
    ```

## Usage

To use **AutoBranch**, simply run the script from your terminal:

```bash
./autobranch [-v]
```

### Flags:
- `-v` or `--verbose`: Enables verbose mode to display detailed information about each step.

### Example Workflow

1. Run the script:
    ```bash
    ./autobranch
    ```
   
2. Enter the branch type when prompted:
    ```
    Enter branch type (e.g., feature, bugfix, hotfix): feature
    ```

3. Enter the branch name:
    ```
    Enter branch name: new-cool-feature
    ```

4. The script creates and switches to the new branch `feature/new-cool-feature`:
    ```
    Branch 'feature/new-cool-feature' created and checked out successfully.
    ```

5. You will then be prompted to push the branch to the remote repository:
    ```
    Would you like to push the new branch to remote? (y/n): y
    ```

6. If the push is successful, you'll see a confirmation:
    ```
    Branch 'feature/new-cool-feature' pushed to remote successfully.
    ```

## Error Handling

- The script ensures that your working directory is clean before creating a branch. If there are unstaged changes, it will exit with a message:
    ```
    Please commit or stash your changes before creating a new branch.
    ```

- If a branch with the same name already exists, the script will notify you and exit:
    ```
    Branch 'feature/new-cool-feature' already exists. Aborting.
    ```

## Time Savings
By eliminating manual branch creation and push commands, **AutoBranch** can save developers approximately 1-2 minutes per branch creation. It also reduces errors caused by inconsistent branch naming.