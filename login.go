package main

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

// Struct to hold GitHub credentials
type GitHubCredentials struct {
	Username    string `survey:"username"`
	Password    string `survey:"password"`
	Email       string `survey:"email"`
	AccessToken string `survey:"access_token"`
}

// Function to prompt for GitHub credentials
func promptForGitHubCredentials() (*GitHubCredentials, error) {
	credentials := &GitHubCredentials{}

	// Define the questions to ask
	questions := []*survey.Question{
		{
			Name:   "username",
			Prompt: &survey.Input{Message: "Enter your GitHub username:"},
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Enter your GitHub password:"},
		},
		{
			Name:   "email",
			Prompt: &survey.Input{Message: "Enter your GitHub email:"},
		},
		{
			Name:   "access_token",
			Prompt: &survey.Input{Message: "Enter your GitHub access token:"},
		},
	}

	// Ask the questions
	if err := survey.Ask(questions, credentials); err != nil {
		return nil, err
	}

	return credentials, nil
}

func main() {
	// Prompt for GitHub credentials
	credentials, err := promptForGitHubCredentials()
	if err != nil {
		log.Fatalf("Failed to get GitHub credentials: %v", err)
	}

	// Display the entered credentials (for demonstration purposes)
	fmt.Printf("GitHub Credentials:\nUsername: %s\nEmail: %s\nAccess Token: %s\n",
		credentials.Username, credentials.Email, credentials.AccessToken)
	// Note: Do not log or display passwords in production code
}
