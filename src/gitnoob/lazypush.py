import subprocess
import sys
import os

def verbose(message, verbose_mode):
    if verbose_mode:
        print(message)

def main():
    # Parse options
    pull_before_push = False
    verbose_mode = False

    args = sys.argv[1:]
    while args and args[0].startswith('-'):
        if args[0] in ['-p', '--pull']:
            pull_before_push = True
        elif args[0] in ['-v', '--verbose']:
            verbose_mode = True
        else:
            print(f"Invalid option: {args[0]}")
            sys.exit(1)
        args.pop(0)

    # Check for unstaged changes
    result = subprocess.run(['git', 'status', '--porcelain'], capture_output=True, text=True)
    if not result.stdout.strip():
        print("No changes to commit.")
        sys.exit(0)

    # Pull the latest changes if the option is enabled
    if pull_before_push:
        verbose("Pulling latest changes from the remote branch...", verbose_mode)
        result = subprocess.run(['git', 'pull', 'origin', subprocess.getoutput('git branch --show-current')])
        if result.returncode != 0:
            print("Merge conflict or error occurred during pull. Please resolve manually.")
            sys.exit(1)

    # Ask for a commit message
    commit_message = input("Enter commit message (leave empty for default): ")

    # Default message if none is provided
    if not commit_message:
        commit_message = f"Auto commit on {subprocess.getoutput('date')}"

    # Add changes to the staging area
    verbose("Staging changes...", verbose_mode)
    subprocess.run(['git', 'add', '.'])

    # Commit with the provided or default message
    verbose("Committing changes...", verbose_mode)
    subprocess.run(['git', 'commit', '-m', commit_message])

    # Push the changes to the current branch
    verbose("Pushing changes to the remote branch...", verbose_mode)
    result = subprocess.run(['git', 'push', 'origin', subprocess.getoutput('git branch --show-current')])
    
    if result.returncode != 0:
        print("Initial push failed. Trying to pull the latest changes and push again...")
        result = subprocess.run(['git', 'pull', '--rebase', 'origin', subprocess.getoutput('git branch --show-current')])
        if result.returncode == 0:
            result = subprocess.run(['git', 'push', 'origin', subprocess.getoutput('git branch --show-current')])
            if result.returncode == 0:
                print("Changes have been committed and pushed successfully after resolving conflicts.")
            else:
                print("Error: Push failed again. Please resolve manually.")
        else:
            print("Error: Pull (rebase) failed. Please resolve manually.")
    else:
        print("Changes have been committed and pushed successfully.")

if __name__ == '__main__':
    main()
