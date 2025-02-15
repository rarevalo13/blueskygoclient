package cmd

import (
	"blueskyClient/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// define styles
var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("12")).
	PaddingLeft(4).
	Width(40)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long: `Log into Bluesky Social with Username and Password
	Example blueskyClient login -u <username> -p <password>`,
	Run: login,
}

func init() {
	rootCmd.AddCommand(loginCmd)

}

func login(cmd *cobra.Command, args []string) {

	var username string
	var password string
	huh.NewInput().
		Title("Username: ").
		Value(&username).
		Run()
	huh.NewInput().
		Title("Password").EchoMode(huh.EchoModePassword).
		Value(&password).
		Run()

	rootURL := "https://bsky.social"
	createSession := "/xrpc/com.atproto.server.createSession"
	jsonBody := []byte(fmt.Sprintf(`{"identifier": "%s", "password": "%s"}`, username, password))
	req, err := http.NewRequest(http.MethodPost, rootURL+createSession, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("There has been an error %s", err)

	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := io.ReadAll(resp.Body)
	var session utils.BluSkySession

	err = json.Unmarshal(body, &session)
	if err != nil {
		log.Printf("There has been an error %s", err)
	}

	fmt.Println(style.Render("Login Successful"))
	utils.WriteToEnv(session)
}
