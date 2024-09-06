package main

//
//
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

var (
	pullBeforePush bool
	verboseMode    bool
	s              *spinner.Spinner
	green          = color.New(color.FgGreen).SprintFunc()
	red            = color.New(color.FgRed).SprintFunc()
	yellow         = color.New(color.FgYellow).SprintFunc()
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "lazypush",
		Short: "A tool to lazily add, commit, and push changes to a Git repository",
		Run:   lazyPush,
	}

	rootCmd.Flags().BoolVarP(&pullBeforePush, "pull", "p", false, "Pull before pushing")
	rootCmd.Flags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")

	if err := rootCmd.Execute(); err != nil {
		logError("Failed to execute command", err)
		os.Exit(1)
	}
}

func lazyPush(cmd *cobra.Command, args []string) {
	if !hasUnstagedChanges() {
		fmt.Println("No changes to commit. You're all caught up! 🎉")
		return
	}

	if pullBeforePush {
		logVerbose("Pulling latest changes from the remote branch...")
		if err := runCommand("git", "pull", "origin", currentBranch()); err != nil {
			logError("Merge conflict or error occurred during pull. Please resolve manually", err)
			os.Exit(1)
		}
	}

	commitMessage := fmt.Sprintf("Auto commit on %s", time.Now().Format(time.RFC1123))

	logVerbose("Staging changes...")
	if err := runCommand("git", "add", "."); err != nil {
		logError("Error staging changes", err)
		return
	}

	logVerbose("Committing changes...")
	if err := runCommand("git", "commit", "-m", commitMessage); err != nil {
		logError("Error committing changes", err)
		return
	}

	// Capture commit details
	commitDetails := fmt.Sprintf("[main %s] %s\n", getLastCommitHash(), commitMessage)
	fmt.Println(commitDetails)

	logVerbose("Pushing changes to the remote branch...")
	if err := pushChanges(); err != nil {
		logError("Failed to push changes", err)
		return
	}

	fmt.Printf("%s Changes have been committed and pushed successfully!\n", green("✓"))
	fmt.Printf("%s You're on fire! Keep up the great work! 🔥\n", yellow("→"))
}

func getCommitMessage() string {
	fmt.Print("Enter commit message (leave empty for default): ")
	reader := bufio.NewReader(os.Stdin)
	commitMessage, _ := reader.ReadString('\n')
	commitMessage = strings.TrimSpace(commitMessage)

	if commitMessage == "" {
		commitMessage = fmt.Sprintf("Auto commit on %s - Lazy but productive! 😎", time.Now().Format(time.RFC1123))
	}

	return commitMessage
}

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

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w (stderr: %s)", err, stderr.String())
	}
	return nil
}

func hasUnstagedChanges() bool {
	output, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		logError("Error checking git status", err)
		os.Exit(1)
	}
	return len(output) > 0
}

func currentBranch() string {
	output, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		logError("Error getting current branch", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

func logVerbose(message string) {
	if verboseMode {
		fmt.Printf("%s %s\n", yellow("→"), message)
	}
}

func logError(message string, err error) {
	fmt.Printf("%s %s", red("✗"), message)
	if err != nil {
		fmt.Printf(": %v", err)
	}
	fmt.Println()
}

func startSpinner(message string) {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Start()
}

func stopSpinner() {
	s.Stop()
}

func getLastCommitHash() string {
	output, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		logError("Error getting last commit hash", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}
