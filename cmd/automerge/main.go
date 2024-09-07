package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	verboseMode bool
	s           *spinner.Spinner
	green       = color.New(color.FgGreen, color.Bold).SprintFunc()
	red         = color.New(color.FgRed, color.Bold).SprintFunc()
	yellow      = color.New(color.FgYellow, color.Bold).SprintFunc()
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "automerge",
		Short: "Automatically merge all branches into the main branch",
		Run:   autoMerge,
	}

	rootCmd.Flags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func autoMerge(cmd *cobra.Command, args []string) {
	branches, err := getBranches()
	if err != nil {
		log.Fatalf("Failed to get branches: %v", err)
	}

	if err := checkoutBranch("main"); err != nil {
		log.Fatalf("Failed to checkout main branch: %v", err)
	}

	for _, branch := range branches {
		if branch == "main" {
			continue
		}

		if err := mergeBranch(branch); err != nil {
			logError(fmt.Sprintf("Failed to merge branch %s", branch), err)
			if isMergeConflict(err) {
				handleMergeConflict(branch)
			}
			continue
		}
	}

	fmt.Println(green("✓ All branches have been merged into main successfully! 🚀"))
}

func getBranches() ([]string, error) {
	output, err := runCommandWithOutput("git", "branch", "--list")
	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	branches := strings.Fields(output)
	return branches, nil
}

func checkoutBranch(branch string) error {
	startSpinner(fmt.Sprintf("Checking out branch %s", branch))
	defer stopSpinner()

	if err := runCommand("git", "checkout", branch); err != nil {
		return fmt.Errorf("failed to checkout branch %s: %w", branch, err)
	}
	fmt.Println(green(fmt.Sprintf("✓ Checked out branch: %s", branch)))
	return nil
}

func mergeBranch(branch string) error {
	startSpinner(fmt.Sprintf("Merging branch %s", branch))
	defer stopSpinner()

	mergeMessage := fmt.Sprintf("Merging branch %s into main", branch)
	if err := runCommand("git", "merge", "--no-ff", "-m", mergeMessage, branch); err != nil {
		return fmt.Errorf("failed to merge branch %s: %w", branch, err)
	}
	fmt.Println(green(fmt.Sprintf("✓ Merged branch: %s", branch)))
	return nil
}

func runCommandWithOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
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

func isMergeConflict(err error) bool {
	return strings.Contains(err.Error(), "CONFLICT")
}

func handleMergeConflict(branch string) {
	fmt.Println(red(fmt.Sprintf("Merge conflict detected in branch %s. Aborting merge.", branch)))
	if err := runCommand("git", "merge", "--abort"); err != nil {
		logError("Failed to abort merge", err)
	}
	fmt.Println(yellow("Merge aborted. Please resolve conflicts manually."))
}
