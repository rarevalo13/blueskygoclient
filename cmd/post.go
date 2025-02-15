package cmd

import (
	"blueskyClient/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "post",
	Long:  `post to your Bluesky Social`,
	Run:   post,
}

func init() {
	rootCmd.AddCommand(postCmd)

}

func post(cmd *cobra.Command, args []string) {
	postURL := "https://bsky.social/xrpc/com.atproto.repo.createRecord"
	handle := utils.GetUserHandle("ronarevalo.com")
	postTime := time.Now().UTC().Format(time.RFC3339)
	var newPost = utils.Post{
		Repo:       handle,
		Collection: "app.bsky.feed.post",
		Record:     utils.Record{CreatedAt: postTime},
	}
	postfield := huh.NewText().
		Title("Whats up?").
		CharLimit(300).
		Value(&newPost.Record.Text)

	if err := postfield.Run(); err != nil {
		fmt.Println("Error running text field:", err)
		return
	}

	data, err := json.Marshal(&newPost)
	if err != nil {
		log.Printf("There is an error with the post ", err)
	}
	req, err := http.NewRequest(http.MethodPost, postURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to post:  %s", err)

	}
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Println(string(body))
	if result["error"] == "ExpiredToken" || result["error"] == "InvalidToken" {
		utils.Refresh(os.Getenv("REFRESH_TOKEN"))
	}
	defer resp.Body.Close()

	fmt.Printf("Post Successful: %s\n", newPost.Record.Text)

}
