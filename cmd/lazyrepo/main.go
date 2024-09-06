// main.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var verboseMode bool

	var rootCmd = &cobra.Command{
		Use:   "setup-repo",
		Short: "Setup a new Git repository both locally and on GitHub",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			repoName := args[0]
			token := viper.GetString("github_token")

			if token == "" {
				fmt.Println("Error: GITHUB_TOKEN environment variable is not set.")
				os.Exit(1)
			}

			if verboseMode {
				fmt.Printf("Creating local repository directory: %s\n", repoName)
			}

			// Create local repository
			if err := createLocalRepo(repoName); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Create remote repository on GitHub
			repoURL, err := createGitHubRepo(repoName, token)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			if verboseMode {
				fmt.Printf("Adding remote repository: %s\n", repoURL)
			}

			// Add remote and push
			if err := addRemoteAndPush(repoName, repoURL); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Local repository initialized and pushed to GitHub repository successfully.")
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verboseMode, "verbose", "v", false, "Enable verbose output")
	viper.BindEnv("github_token", "GITHUB_TOKEN")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createLocalRepo(name string) error {
	if err := os.Mkdir(name, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	os.Chdir(name)

	repo, err := git.PlainInit(".", false)
	if err != nil {
		return fmt.Errorf("failed to initialize repository: %w", err)
	}

	f, err := os.Create("README.md")
	if err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	defer f.Close()

	repo.CommitObject
	_, err = exec.Command("git", "add", "README.md").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stage README.md: %w", err)
	}

	_, err = exec.Command("git", "commit", "-m", "Initial commit").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to commit README.md: %w", err)
	}

	return nil
}

func createGitHubRepo(name, token string) (string, error) {
	// Create the GitHub repository
	cmd := exec.Command("curl", "-s", "-H", "Authorization: token "+token, "-H", "Accept: application/vnd.github.v3+json", "-d", fmt.Sprintf("{\"name\":\"%s\",\"private\":false}", name), "https://api.github.com/user/repos")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to create GitHub repository: %w", err)
	}

	// Check for success
	if !strings.Contains(string(output), "\"full_name\":") {
		return "", fmt.Errorf("GitHub repository creation failed: %s", output)
	}

	return fmt.Sprintf("https://github.com/%s/%s.git", viper.GetString("github_username"), name), nil
}

func addRemoteAndPush(repoName, repoURL string) error {
	_, err := exec.Command("git", "remote", "add", "origin", repoURL).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add remote repository: %w", err)
	}

	_, err = exec.Command("git", "push", "-u", "origin", "master").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to push to remote repository: %w", err)
	}

	return nil
}
