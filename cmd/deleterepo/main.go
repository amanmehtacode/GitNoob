package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Configuration structure to hold command-line flags
type Config struct {
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
	rootCmd.Flags().BoolVarP(&cfg.VerboseMode, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolVarP(&cfg.Interactive, "interactive", "i", true, "Run in interactive mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func deleteRepo(cmd *cobra.Command, args []string) {
	username, token, err := getGitHubCredentials()
	if err != nil {
		log.Fatalf(red("Failed to get GitHub credentials: %v"), err)
	}

	repoNames, err := listRepositories(username, token)
	if err != nil {
		log.Fatalf(red("Failed to list repositories: %v"), err)
	}

	if len(repoNames) == 0 {
		log.Fatalf(red("No repositories found for user %s"), username)
	}

	selectedRepo := ""
	if cfg.Interactive {
		if err := survey.AskOne(&survey.Select{
			Message: "Select a repository to delete:",
			Options: repoNames,
		}, &selectedRepo); err != nil {
			log.Fatalf(red("Failed to select repository: %v"), err)
		}
	} else {
		selectedRepo = args[0] // Use the first argument if not interactive
	}

	localRepos, err := listLocalRepositories()
	if err != nil {
		log.Fatalf(red("Failed to list local repositories: %v"), err)
	}

	if len(localRepos) == 0 {
		log.Fatalf(red("No local repositories found."))
	}

	selectedLocalRepo := ""
	if err := survey.AskOne(&survey.Select{
		Message: "Select a local repository to delete:",
		Options: localRepos,
	}, &selectedLocalRepo); err != nil {
		log.Fatalf(red("Failed to select local repository: %v"), err)
	}

	if confirmDeletion(selectedLocalRepo) {
		if err := deleteLocalRepo(selectedLocalRepo); err != nil {
			log.Fatalf(red("Failed to delete local repository: %v"), err)
		}

		// Ask if the user wants to delete the GitHub repo
		if confirmDeletion("the GitHub repository '" + selectedLocalRepo + "'") {
			if err := deleteGitHubRepo(selectedLocalRepo, username, token); err != nil {
				log.Fatalf(red("Failed to delete GitHub repository: %v"), err)
			}
		} else {
			fmt.Println(yellow("Deletion from GitHub cancelled."))
		}
	} else {
		fmt.Println(yellow("Local repository deletion cancelled."))
	}
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

func listRepositories(username, token string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(username, token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list repositories: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list repositories: %s", resp.Status)
	}

	var repos []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, repo.Name)
	}

	return repoNames, nil
}

func confirmDeletion(repoName string) bool {
	var confirm string
	fmt.Print(yellow("Are you sure you want to delete the repository '" + repoName + "'? This action cannot be undone. (y/N): "))
	fmt.Scanln(&confirm)
	return strings.ToLower(confirm) == "y"
}

func deleteLocalRepo(repoName string) error {
	localPath := fmt.Sprintf("./%s", repoName) // Adjust the path as necessary
	if err := os.RemoveAll(localPath); err != nil {
		return fmt.Errorf("failed to delete local repository: %w", err)
	}
	fmt.Println(green("✓ Local repository deleted successfully!"))
	return nil
}

func listLocalRepositories() ([]string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var repos []string
	for _, file := range files {
		if file.IsDir() {
			repos = append(repos, file.Name())
		}
	}
	return repos, nil
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
