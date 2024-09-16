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

var (
    pushAfterCommit bool
    verboseMode     bool
    interactive     bool
    s               *spinner.Spinner
    green           = color.New(color.FgGreen, color.Bold).SprintFunc()
    red             = color.New(color.FgRed, color.Bold).SprintFunc()
    yellow          = color.New(color.FgYellow, color.Bold).SprintFunc()
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "autocommit",
        Short: "Automatically commit changes with a generated message",
        Run:   autoCommit,
    }

    rootCmd.Flags().BoolVarP(&pushAfterCommit, "push", "p", false, "Push after committing")
    rootCmd.Flags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")
    rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", true, "Run in interactive mode")

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Failed to execute command: %v", err)
    }
}

func autoCommit(cmd *cobra.Command, args []string) {
    // Check if there are any changes to commit
    if !hasUnstagedChanges() {
        fmt.Println(green("✓ No changes to commit. You're all caught up! 🎉"))
        return
    }

    // Get commit message from user
    commitMessage := getCommitMessage()

    // Stage and commit changes
    fmt.Println(yellow("→ Staging and committing changes..."))
    if err := commitChanges(commitMessage); err != nil {
        logError("Error committing changes", err)
        return
    }

    // Push changes to remote if the flag is set or if confirmed in interactive mode
    if pushAfterCommit || (interactive && confirmPush()) {
        fmt.Println(yellow("→ Pushing changes to remote..."))
        if err := pushChanges(); err != nil {
            logError("Failed to push changes", err)
            return
        }
    }

    fmt.Println(green("✓ Changes have been committed and pushed successfully! 🚀"))
}

func commitChanges(commitMessage string) error {
    startSpinner("Committing changes")
    defer stopSpinner()

    commitOutput, err := runCommandWithOutput("git", "commit", "-am", commitMessage)
    if err != nil {
        return fmt.Errorf("error committing changes: %w", err)
    }
    printFormattedOutput(commitOutput)
    return nil
}

func pushChanges() error {
    startSpinner("Pushing changes to remote")
    defer stopSpinner()

    pushOutput, err := runCommandWithOutput("git", "push", "origin", currentBranch())
    if err != nil {
        return fmt.Errorf("failed to push changes: %w", err)
    }
    printFormattedOutput(pushOutput)
    return nil
}

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

func hasUnstagedChanges() bool {
    output, err := exec.Command("git", "status", "--porcelain").Output()
    if err != nil {
        logError("Error checking git status", err)
        os.Exit(1)
    }
    return len(output) > 0
}

func runCommandWithOutput(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

func confirmPush() bool {
    var confirm string
    fmt.Print(yellow("Do you want to push the changes to remote? (y/n): "))
    fmt.Scanln(&confirm)
    return strings.ToLower(confirm) == "y"
}

func logError(message string, err error) {
    log.Printf("%s %s: %v", red("✗"), message, err)
}

func startSpinner(message string) {
    s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
    s.Suffix = " " + message
    s.Start()
}

func stopSpinner() {
    s.Stop()
}

func printFormattedOutput(output string) {
    lines := strings.Split(strings.TrimSpace(output), "\n")
    for _, line := lines {
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

func currentBranch() string {
    output, err := exec.Command("git", "branch", "--show-current").Output()
    if err != nil {
        logError("Error getting current branch", err)
        os.Exit(1)
    }
    return strings.TrimSpace(string(output))
}