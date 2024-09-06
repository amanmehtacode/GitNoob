package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey"
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
	blue   = color.New(color.FgBlue, color.Bold).SprintFunc()
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
		logError("Failed to execute command", err)
		os.Exit(1)
	}
}

// autoBranch orchestrates the branch creation and commit process
func autoBranch(cmd *cobra.Command, args []string) {
	printLogo()

	if !hasUnstagedChanges() {
		fmt.Println(green("✓ No changes to commit. You're all caught up! 🎉"))
		return
	}

	commitMessage := getCommitMessage()
	branchName := generateBranchName(commitMessage)

	if cfg.Interactive {
		branchName = promptForBranchName(branchName)
	}

	if createBranch(branchName) {
		commitChanges(commitMessage)
		if cfg.PushAfterCommit || (cfg.Interactive && confirmPush()) {
			pushChanges(branchName)
		}
		printSummary(branchName, commitMessage)
	}
}

// printLogo displays the tool's logo
func printLogo() {
	logo := `
   _____         __        ____                        __  
  /  _  \  __ __|  |_ ___ |    |   ____________ _____  |  | 
 /  /_\  \|  |  |   _| _ \|    |  /  _ \_  __ \/     \ |  | 
/    |    |  |  |  |_  __/|    |_(  <_> |  | \/  Y Y  \|  |__
\____|__  |____/|____|___ |_______ \____/|__|  |__|_|  |____/
        \/                \/                         \/      
`
	fmt.Println(yellow(logo))
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

// getCommitMessage prompts the user for a commit message
func getCommitMessage() string {
	var commitMessage string
	prompt := &survey.Input{
		Message: "Enter commit message:",
		Help:    "This message will be used for both the commit and to generate the branch name.",
	}
	survey.AskOne(prompt, &commitMessage, survey.WithValidator(survey.Required))
	return commitMessage
}

// promptForBranchName allows the user to modify the suggested branch name
func promptForBranchName(suggestion string) string {
	var branchName string
	prompt := &survey.Input{
		Message: "Enter branch name:",
		Default: suggestion,
		Help:    "You can modify the suggested branch name or accept the default.",
	}
	survey.AskOne(prompt, &branchName)
	return branchName
}

// createBranch creates a new Git branch
func createBranch(branchName string) bool {
	startSpinner("Creating new branch")
	defer stopSpinner()

	if err := runCommand("git", "checkout", "-b", branchName); err != nil {
		logError("Failed to create new branch", err)
		return false
	}
	fmt.Println(green("✓ Created new branch: " + branchName))
	return true
}

// commitChanges stages and commits the changes
func commitChanges(commitMessage string) {
	startSpinner("Staging and committing changes")
	defer stopSpinner()

	commitOutput, err := runCommandWithOutput("git", "commit", "-am", commitMessage)
	if err != nil {
		logError("Error committing changes", err)
		return
	}
	printFormattedOutput(commitOutput)
}

// pushChanges pushes the new branch to the remote repository
func pushChanges(branchName string) {
	startSpinner("Pushing new branch to remote")
	defer stopSpinner()

	pushOutput, err := runCommandWithOutput("git", "push", "-u", "origin", branchName)
	if err != nil {
		logError("Failed to push new branch", err)
		return
	}
	printFormattedOutput(pushOutput)
}

// confirmPush asks the user for confirmation to push the new branch
func confirmPush() bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: "Do you want to push the new branch to remote?",
		Default: false,
	}
	survey.AskOne(prompt, &confirm)
	return confirm
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
	logVerbose(fmt.Sprintf("Running command: %s %s", name, strings.Join(args, " ")))
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// runCommand executes a command without capturing output
func runCommand(name string, args ...string) error {
	logVerbose(fmt.Sprintf("Running command: %s %s", name, strings.Join(args, " ")))
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// printFormattedOutput formats and prints the output of git commands
func printFormattedOutput(output string) {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "[") {
			fmt.Println(green(line))
		} else if strings.Contains(line, "|") {
			parts := strings.SplitN(line, "|", 2)
			fmt.Printf("%s|%s\n", yellow(parts[0]), blue(parts[1]))
		} else {
			fmt.Println(line)
		}
	}
}

// logVerbose logs a message if verbose mode is enabled
func logVerbose(message string) {
	if cfg.VerboseMode {
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

// printSummary displays a summary of the actions taken
func printSummary(branchName, commitMessage string) {
	fmt.Println(blue("\nSummary:"))
	fmt.Printf("%s %s\n", yellow("Branch:"), branchName)
	fmt.Printf("%s %s\n", yellow("Commit:"), commitMessage)
	if cfg.PushAfterCommit {
		fmt.Printf("%s %s\n", yellow("Pushed:"), "Yes")
	} else {
		fmt.Printf("%s %s\n", yellow("Pushed:"), "No")
	}
}
