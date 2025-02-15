package cmd

import (
	"blueskyClient/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "feed",
	Long:  `show the last 10 posts from your Bluesky Social feed.`,
	Run:   feed,
}

func init() {
	rootCmd.AddCommand(feedCmd)
	feedCmd.Flags().StringP("limit", "l", "50", "number of posts to dispay")
	feedCmd.Flags().StringP("algorithm", "a", "chronological", "the algorithm used to display")

}

func feed(cmd *cobra.Command, args []string) {
	var feed utils.TimelineResponse
	limit, _ := cmd.Flags().GetString("limit")
	feedURL, err := url.Parse("https://velvetfoot.us-east.host.bsky.network/xrpc/app.bsky.feed.getTimeline")
	if err != nil {
		log.Printf("There was an err parsing the URL: %v", err)
	}
	params := url.Values{}
	params.Set("limit", limit)
	params.Set("algorithm", "chronological")
	feedURL.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", feedURL.String(), nil)
	if err != nil {
		log.Printf("Error getting feed: %s ", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("There has been an error %s", err)

	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &feed)
	if err != nil {
		log.Printf("There has been an error %s", err)
	}

	for i := 0; i < len(feed.Feed); i++ {
		fmt.Println(feed.Feed[i].Post.Record)
		fmt.Println(feed.Feed[i].Post.Author)
		fmt.Println(feed.Feed[i].Post.LikeCount)
	}
}
