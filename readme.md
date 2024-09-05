# GitNoob

GitNoob is a suite of Git tools designed to streamline your development workflow and make Git operations more user-friendly, especially for beginners.

## Commands

### 1. AutoBranch
**Usage**: `autobranch <type> <name>`
**Example**: `autobranch feature my-feature`

**Background Operations**:
- Checks if the current branch is clean
- Creates a new branch with the format `<type>/<name>`
- Switches to the newly created branch

**Time Savings**: Eliminates manual branch creation and naming inconsistencies. Saves ~1-2 minutes per branch creation.

### 2. Git Stash Manager
**Usage**: 
- `stash-manager list`
- `stash-manager apply <index>`
- `stash-manager delete <index>`

**Background Operations**:
- List: Shows all stashes with their indices and descriptions
- Apply: Applies the stash at the specified index
- Delete: Removes the stash at the specified index

**Time Savings**: Simplifies stash management, saving ~2-3 minutes when working with multiple stashes.

### 3. Git Diff Formatter
**Usage**: `gitdiff --highlight-additions`

**Background Operations**:
- Runs `git diff` with custom formatting
- Applies color highlighting to additions/deletions
- Optionally shows only additions or deletions

**Time Savings**: Improves diff readability, potentially saving 5-10 minutes during code reviews.

### 4. AutoCommit
**Usage**: `autocommit`

**Background Operations**:
- Prompts for a commit message
- Stages all changes
- Commits with a formatted message

**Time Savings**: Streamlines the commit process, saving 1-2 minutes per commit.

### 5. Git Sync All
**Usage**: `git-sync --all`

**Background Operations**:
- Identifies all Git repositories in a directory
- Pulls latest changes for each repository
- Optionally pushes local changes

**Time Savings**: Significant time-saver for projects with multiple repositories, potentially saving 10-15 minutes per sync operation.

### 6. Git Snapshot
**Usage**: `gitsnapshot`

**Background Operations**:
- Captures current branch name
- Saves commit history
- Records file changes

**Time Savings**: Useful for quick state capture without committing, saving 3-5 minutes compared to manual documentation.

### 7. Interactive Git Cleanup
**Usage**: `git-cleanup`

**Background Operations**:
- Lists merged local and remote branches
- Prompts for branch deletion
- Performs safe deletion of selected branches

**Time Savings**: Simplifies branch management, potentially saving 10-15 minutes during cleanup sessions.

### 8. Commit Diary
**Usage**: `commit-diary generate`

**Background Operations**:
- Retrieves commit history for the day
- Parses commit messages for tags
- Generates a formatted Markdown log

**Time Savings**: Automates daily progress tracking, saving 10-15 minutes of manual logging.

### 9. Git Preflight
**Usage**: `git-preflight`

**Background Operations**:
- Runs predefined checks (e.g., tests, linting)
- Blocks push if checks fail
- Provides detailed error output

**Time Savings**: Catches issues before they reach the remote repository, potentially saving hours of debugging and hotfixing.

### 10. AutoRebase
**Usage**: `autorebase`

**Background Operations**:
- Fetches latest remote changes
- Attempts to rebase current branch
- Resolves conflicts if possible

**Time Savings**: Simplifies keeping branches up-to-date, saving 5-10 minutes per rebase operation.

### 11. LazyPush
**Usage**: `lazypush -p -v`

**Background Operations**:
- Checks for unstaged changes
- Optionally pulls latest changes
- Stages all changes
- Prompts for commit message (or uses default)
- Commits and pushes changes
- Handles push failures with pull and retry

**Time Savings**: Automates the entire commit and push process, saving 3-5 minutes per operation.

### 12. SmartPush
**Usage**: `smartpush`

**Background Operations**:
- Similar to LazyPush, but with advanced conflict resolution
- Attempts to auto-resolve conflicts during pull
- Retries push after conflict resolution

**Time Savings**: Handles complex push scenarios, potentially saving 10-15 minutes when dealing with conflicts.

## Installation

[Provide installation instructions here]

## Contributing

[Provide contribution guidelines here]

## License

[Specify the license information here]

