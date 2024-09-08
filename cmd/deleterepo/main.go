package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
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
	RepoName    string
	VerboseMode bool
	Interactive bool
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

func main() {
	rootCmd := &cobra.Command{
		Use:   "deleterepo",
		Short: "Delete a GitHub repository",
		Run:   deleteRepo,
	}

	// Command-line flags
	rootCmd.Flags().StringVarP(&cfg.RepoName, "name", "n", "", "Name of the repository to delete")
	rootCmd.Flags().BoolVarP(&cfg.VerboseMode, "verbose", "v", true, "Enable verbose output")
	rootCmd.Flags().BoolVarP(&cfg.Interactive, "interactive", "i", true, "Run in interactive mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func deleteRepo(cmd *cobra.Command, args []string) {
	if cfg.Interactive && cfg.RepoName == "" {
		cfg.RepoName = promptForInput("Enter repository name to delete: ")
	}

	if cfg.RepoName == "" {
		log.Fatalf(red("Repository name is required"))
	}

	username, token, err := getGitHubCredentials()
	if err != nil {
		log.Fatalf(red("Failed to get GitHub credentials: %v"), err)
	}

	if err := deleteGitHubRepo(cfg.RepoName, username, token); err != nil {
		log.Fatalf(red("Failed to delete GitHub repository: %v"), err)
	}

	fmt.Println(green("✓ GitHub repository deleted successfully! 🗑️"))
}

func promptForInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow(prompt))
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getGitHubCredentials() (string, string, error) {
	usernameCmd := exec.Command("git", "config", "--global", "github.user")
	usernameOutput, err := usernameCmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get GitHub username: %w", err)
	}
	username := strings.TrimSpace(string(usernameOutput))

	tokenCmd := exec.Command("git", "config", "--global", "github.token")
	tokenOutput, err := tokenCmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get GitHub token: %w", err)
	}
	token := strings.TrimSpace(string(tokenOutput))

	if token == "" {
		return "", "", fmt.Errorf("GitHub token is empty. Please set it using 'git config --global github.token YOUR_TOKEN'")
	}

	return username, token, nil
}

func deleteGitHubRepo(repoName, username, token string) error {
	startSpinner("Deleting GitHub repository")
	defer stopSpinner()

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", username, repoName)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(username, token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete GitHub repository: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete GitHub repository: %s", resp.Status)
	}

	fmt.Println(green("✓ Deleted GitHub repository"))
	return nil
}

func startSpinner(message string) {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	s.Start()
}

func stopSpinner() {
	s.Stop()
}
