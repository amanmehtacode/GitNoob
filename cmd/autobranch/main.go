package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Configuration structure to hold command-line flags
type Config struct {
	PushAfterCommit bool
	VerboseMode     bool
	Interactive     bool
}

// Global variables
var (
	cfg Config
	s   *spinner.Spinner
	// Color functions for output
	green  = color.New(color.FgGreen, color.Bold).SprintFunc()
	red    = color.New(color.FgRed, color.Bold).SprintFunc()
	yellow = color.New(color.FgYellow, color.Bold).SprintFunc()
)

// Main function to set up the CLI
func main() {
	rootCmd := &cobra.Command{
		Use:   "autobranch",
		Short: "Automatically create and manage Git branches based on commit messages",
		Run:   autoBranch,
	}

	// Command-line flags
	rootCmd.Flags().BoolVarP(&cfg.PushAfterCommit, "push", "p", false, "Push the new branch after committing")
	rootCmd.Flags().BoolVarP(&cfg.VerboseMode, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolVarP(&cfg.Interactive, "interactive", "i", true, "Run in interactive mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

// autoBranch orchestrates the branch creation and commit process
func autoBranch(cmd *cobra.Command, args []string) {
	if !hasUnstagedChanges() {
		fmt.Println(green("✓ No changes to commit. You're all caught up! 🎉"))
		return
	}

	commitMessage := getCommitMessage()
	branchName := generateBranchName(commitMessage)

	if cfg.Interactive {
		branchName = promptForBranchName(branchName)
	}

	if err := createBranch(branchName); err != nil {
		logError("Failed to create branch", err)
		return
	}

	if err := commitChanges(commitMessage); err != nil {
		logError("Error committing changes", err)
		return
	}

	if cfg.PushAfterCommit || (cfg.Interactive && confirmPush()) {
		if err := pushChanges(branchName); err != nil {
			logError("Failed to push changes", err)
			return
		}
	}

	fmt.Println(green("✓ Changes have been committed and pushed successfully! 🚀"))
}

// createBranch creates a new Git branch
func createBranch(branchName string) error {
	startSpinner("Creating new branch")
	defer stopSpinner()

	if err := runCommand("git", "checkout", "-b", branchName); err != nil {
		return fmt.Errorf("failed to create new branch: %w", err)
	}
	fmt.Println(green("✓ Created new branch: " + branchName))
	return nil
}

// commitChanges stages and commits the changes
func commitChanges(commitMessage string) error {
	startSpinner("Staging and committing changes")
	defer stopSpinner()

	commitOutput, err := runCommandWithOutput("git", "commit", "-am", commitMessage)
	if err != nil {
		return fmt.Errorf("error committing changes: %w", err)
	}
	printFormattedOutput(commitOutput)
	return nil
}

// pushChanges pushes the new branch to the remote repository
func pushChanges(branchName string) error {
	startSpinner("Pushing new branch to remote")
	defer stopSpinner()

	pushOutput, err := runCommandWithOutput("git", "push", "-u", "origin", branchName)
	if err != nil {
		return fmt.Errorf("failed to push new branch: %w", err)
	}
	printFormattedOutput(pushOutput)
	return nil
}

// printFormattedOutput formats and prints the output of git commands
func printFormattedOutput(output string) {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "[") {
			fmt.Println(green(line))
		} else if strings.Contains(line, "|") {
			parts := strings.SplitN(line, "|", 2)
			fmt.Printf("%s|%s\n", yellow(parts[0]), green(parts[1]))
		} else {
			fmt.Println(line)
		}
	}
}

// getCommitMessage prompts the user for a commit message
func getCommitMessage() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow("Enter commit message: "))
	commitMessage, _ := reader.ReadString('\n')
	return strings.TrimSpace(commitMessage)
}

// promptForBranchName allows the user to modify the suggested branch name
func promptForBranchName(suggestion string) string {
	var branchName string
	fmt.Printf(yellow("Enter branch name (default: %s): "), suggestion)
	reader := bufio.NewReader(os.Stdin)
	branchName, _ = reader.ReadString('\n')
	branchName = strings.TrimSpace(branchName)
	if branchName == "" {
		return suggestion
	}
	return branchName
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

// runCommandWithOutput executes a command and returns its output as a string
func runCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// runCommand executes a command without capturing output
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// confirmPush asks the user for confirmation to push the new branch
func confirmPush() bool {
	var confirm string
	fmt.Print(yellow("Do you want to push the new branch to remote? (y/n): "))
	fmt.Scanln(&confirm)
	return strings.ToLower(confirm) == "y"
}

// logError logs an error message
func logError(message string, err error) {
	log.Printf("%s %s: %v", red("✗"), message, err)
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

// generateBranchName creates a branch name based on the commit message
func generateBranchName(commitMessage string) string {
	branchName := strings.ToLower(commitMessage)
	branchName = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return '-'
	}, branchName)
	branchName = strings.Trim(branchName, "-")
	if len(branchName) > 50 {
		branchName = branchName[:50]
	}
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("%s-%s", branchName, timestamp)
}
