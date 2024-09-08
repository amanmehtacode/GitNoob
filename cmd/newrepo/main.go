package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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
		Use:   "newrepo",
		Short: "Create a new Git repository with a predefined structure and publish it to GitHub",
		Run:   createNewRepo,
	}

	// Command-line flags
	rootCmd.Flags().StringVarP(&cfg.RepoName, "name", "n", "", "Name of the new repository")
	rootCmd.Flags().BoolVarP(&cfg.VerboseMode, "verbose", "v", true, "Enable verbose output")
	rootCmd.Flags().BoolVarP(&cfg.Interactive, "interactive", "i", true, "Run in interactive mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func createNewRepo(cmd *cobra.Command, args []string) {
	if cfg.Interactive {
		cfg.RepoName = promptForInput("Enter repository name: ")
	}

	repoPath := filepath.Join(".", cfg.RepoName)

	if err := os.Mkdir(repoPath, 0755); err != nil {
		log.Fatalf(red("Failed to create directory: %v"), err)
	}

	if err := initGitRepo(repoPath); err != nil {
		log.Fatalf(red("Failed to initialize Git repository: %v"), err)
	}

	if err := createReadme(repoPath); err != nil {
		log.Fatalf(red("Failed to create README.md: %v"), err)
	}

	if err := createGitignore(repoPath); err != nil {
		log.Fatalf(red("Failed to create .gitignore: %v"), err)
	}

	if err := createBoilerplate(repoPath); err != nil {
		log.Fatalf(red("Failed to create boilerplate structure: %v"), err)
	}

	if err := initialCommit(repoPath); err != nil {
		log.Fatalf(red("Failed to make initial commit: %v"), err)
	}

	username, token, err := getGitHubCredentials()
	if err != nil {
		log.Printf("Error getting GitHub credentials: %v", err)
		token = promptForInput("Enter your GitHub token: ")
		if err := exec.Command("git", "config", "--global", "github.token", token).Run(); err != nil {
			log.Fatalf(red("Failed to save GitHub token: %v"), err)
		}
	}

	if err := createGitHubRepo(cfg.RepoName, username, token); err != nil {
		log.Fatalf(red("Failed to create GitHub repository: %v"), err)
	}

	if err := pushToGitHub(repoPath, username, cfg.RepoName); err != nil {
		log.Fatalf(red("Failed to push to GitHub: %v"), err)
	}

	fmt.Println(green("✓ New Git repository created and published to GitHub successfully! 🚀"))
}

func promptForInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow(prompt))
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getGitHubCredentials() (string, string, error) {
	usernameCmd := exec.Command("git", "config", "--global", "user.name")
	usernameOutput, err := usernameCmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get GitHub username: %w", err)
	}
	username := strings.TrimSpace(string(usernameOutput))

	tokenCmd := exec.Command("git", "config", "--global", "github.token")
	tokenOutput, err := tokenCmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return "", "", fmt.Errorf("failed to get GitHub token: %w\nError output: %s", err, exitErr.Stderr)
		}
		return "", "", fmt.Errorf("failed to get GitHub token: %w", err)
	}
	token := strings.TrimSpace(string(tokenOutput))

	if token == "" {
		return "", "", fmt.Errorf("GitHub token is empty. Please set it using 'git config --global github.token YOUR_TOKEN'")
	}

	return username, token, nil
}

func initGitRepo(repoPath string) error {
	startSpinner("Initializing Git repository")
	defer stopSpinner()

	if err := runCommand(repoPath, "git", "init"); err != nil {
		return fmt.Errorf("failed to initialize Git repository: %w", err)
	}
	fmt.Println(green("✓ Initialized Git repository"))
	return nil
}

func createReadme(repoPath string) error {
	startSpinner("Creating README.md")
	defer stopSpinner()

	readmePath := filepath.Join(repoPath, "README.md")
	readmeContent := fmt.Sprintf("# %s\n\nThis is the README file for the %s repository.", cfg.RepoName, cfg.RepoName)

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	fmt.Println(green("✓ Created README.md"))
	return nil
}

func createGitignore(repoPath string) error {
	startSpinner("Creating .gitignore")
	defer stopSpinner()

	gitignorePath := filepath.Join(repoPath, ".gitignore")
	gitignoreContent := "node_modules/\n.DS_Store\n"

	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	fmt.Println(green("✓ Created .gitignore"))
	return nil
}

func createBoilerplate(repoPath string) error {
	startSpinner("Creating boilerplate structure")
	defer stopSpinner()

	dirs := []string{"src", "bin", "pkg", "cmd", "internal", "configs", "scripts", "build", "deploy", "test", "docs"}
	for _, dir := range dirs {
		if err := os.Mkdir(filepath.Join(repoPath, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}
	fmt.Println(green("✓ Created boilerplate structure"))
	return nil
}

func initialCommit(repoPath string) error {
	startSpinner("Making initial commit")
	defer stopSpinner()

	if err := runCommand(repoPath, "git", "add", "."); err != nil {
		return fmt.Errorf("failed to add files to Git: %w", err)
	}

	commitMessage := fmt.Sprintf("Initial commit for %s", cfg.RepoName)
	if err := runCommand(repoPath, "git", "commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("failed to commit files: %w", err)
	}
	fmt.Println(green("✓ Made initial commit"))
	return nil
}

func createGitHubRepo(repoName, username, token string) error {
	startSpinner("Creating GitHub repository")
	defer stopSpinner()

	url := "https://api.github.com/user/repos"
	repo := map[string]string{
		"name": repoName,
	}
	repoJSON, _ := json.Marshal(repo)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(repoJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.SetBasicAuth(username, token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create GitHub repository: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create GitHub repository: %s", resp.Status)
	}

	fmt.Println(green("✓ Created GitHub repository"))
	return nil
}

func pushToGitHub(repoPath, username, repoName string) error {
	startSpinner("Pushing to GitHub")
	defer stopSpinner()

	remoteURL := fmt.Sprintf("https://github.com/%s/%s.git", username, repoName)
	if err := runCommand(repoPath, "git", "remote", "add", "origin", remoteURL); err != nil {
		return fmt.Errorf("failed to add remote: %w", err)
	}

	if err := runCommand(repoPath, "git", "push", "-u", "origin", "master"); err != nil {
		return fmt.Errorf("failed to push to GitHub: %w", err)
	}

	fmt.Println(green("✓ Pushed to GitHub"))
	return nil
}

func runCommand(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, output)
	}
	if cfg.VerboseMode {
		fmt.Println(string(output))
	}
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
