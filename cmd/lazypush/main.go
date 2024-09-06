package main

// Import necessary packages for file I/O, executing commands, and CLI functionality
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Global variables for configuration and styling
var (
	pullBeforePush bool
	verboseMode    bool
	s              *spinner.Spinner
	green          = color.New(color.FgGreen).SprintFunc()
	red            = color.New(color.FgRed).SprintFunc()
	yellow         = color.New(color.FgYellow).SprintFunc()
)

// main is the entry point of the application
// It sets up the CLI command and executes it
func main() {
	var rootCmd = &cobra.Command{
		Use:   "lazypush",
		Short: "A tool to lazily add, commit, and push changes to a Git repository",
		Run:   lazyPush,
	}

	// Set up command-line flags
	rootCmd.Flags().BoolVarP(&pullBeforePush, "pull", "p", false, "Pull before pushing")
	rootCmd.Flags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		logError("Failed to execute command", err)
		os.Exit(1)
	}
}

// lazyPush is the main function that orchestrates the git operations
func lazyPush(cmd *cobra.Command, args []string) {
	// Check if there are any changes to commit
	if !hasUnstagedChanges() {
		fmt.Println(green("✓ No changes to commit. You're all caught up! 🎉"))
		return
	}

	// Pull latest changes if the flag is set
	if pullBeforePush {
		fmt.Println(yellow("→ Pulling latest changes..."))
		if err := runCommand("git", "pull", "--rebase", "origin", currentBranch()); err != nil {
			logError("Merge conflict or error occurred during pull. Please resolve manually", err)
			os.Exit(1)
		}
	}

	// Get commit message from user
	commitMessage := getCommitMessage()

	// Stage and commit changes
	fmt.Println(yellow("→ Staging and committing changes..."))
	commitOutput, err := runCommandWithOutput("git", "commit", "-am", commitMessage)
	if err != nil {
		logError("Error committing changes", err)
		return
	}
	printFormattedOutput(commitOutput)

	// Push changes to remote
	fmt.Println(yellow("→ Pushing changes to remote..."))
	pushOutput, err := runCommandWithOutput("git", "push", "origin", currentBranch())
	if err != nil {
		logError("Failed to push changes", err)
		return
	}
	printFormattedOutput(pushOutput)

	fmt.Println(green("✓ Changes have been committed and pushed successfully! 🚀"))
}

// printFormattedOutput formats and prints the output of git commands
func printFormattedOutput(output string) {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "[") {
			fmt.Println(green(line))
		} else if strings.Contains(line, ":|") {
			parts := strings.Split(line, ":|")
			fmt.Printf("%s %s\n", yellow(parts[0]+":"), parts[1])
		} else {
			fmt.Println(line)
		}
	}
}

// getCommitMessage prompts the user for a commit message or uses a default one
func getCommitMessage() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow("Enter commit message (leave empty for default): "))
	commitMessage, _ := reader.ReadString('\n')
	commitMessage = strings.TrimSpace(commitMessage)

	if commitMessage == "" {
		commitMessage = fmt.Sprintf("Auto commit on %s", time.Now().Format(time.RFC1123))
	}

	return commitMessage
}

// runCommandWithOutput executes a command and returns its output as a string
func runCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// pushChanges attempts to push changes, handling potential conflicts
func pushChanges() error {
	startSpinner("Pushing changes")
	defer stopSpinner()

	err := runCommand("git", "push", "origin", currentBranch())
	if err != nil {
		fmt.Println("Initial push failed. Trying to pull the latest changes and push again...")
		if err := runCommand("git", "pull", "--rebase", "origin", currentBranch()); err != nil {
			return fmt.Errorf("pull (rebase) failed: %w", err)
		}
		if err := runCommand("git", "push", "origin", currentBranch()); err != nil {
			return fmt.Errorf("push failed again: %w", err)
		}
		fmt.Println("Changes have been committed and pushed successfully after resolving conflicts.")
	}
	return nil
}

// runCommand executes a command and captures its stderr output
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w (stderr: %s)", err, stderr.String())
	}
	return nil
}

// hasUnstagedChanges checks if there are any unstaged changes in the repository
func hasUnstagedChanges() bool {
	output, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		logError("Error checking git status", err)
		os.Exit(1)
	}
	return len(output) > 0
}

// currentBranch gets the name of the current git branch
func currentBranch() string {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		logError("Error getting current branch", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

// logVerbose logs a message if verbose mode is enabled
func logVerbose(message string) {
	if verboseMode {
		fmt.Printf("%s %s\n", yellow("→"), message)
	}
}

// logError logs an error message
func logError(message string, err error) {
	fmt.Printf("%s %s", red("✗"), message)
	if err != nil {
		fmt.Printf(": %v", err)
	}
	fmt.Println()
}

// startSpinner starts a spinner with a given message
func startSpinner(message string) {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Start()
}

// stopSpinner stops the current spinner
func stopSpinner() {
	s.Stop()
}

// getLastCommitHash retrieves the hash of the last commit
func getLastCommitHash() string {
	output, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		logError("Error getting last commit hash", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}
