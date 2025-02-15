package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	publicBlueskyBaseURL = "https://public.api.bsky.app"
)

// maybe add the handle to the .env as well
func WriteToEnv(session BluSkySession) {
	tokens := []map[string]string{
		{
			"key":   "ACCESS_TOKEN",
			"value": fmt.Sprintf("%s", session.AccessJwt),
		},
		{
			"key":   "REFRESH_TOKEN",
			"value": fmt.Sprintf("%s", session.RefreshJwt),
		},
	}
	f, err := os.OpenFile(".env", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening %s, not found, creating it.", err)

	}
	defer f.Close()
	for _, token := range tokens {
		f.WriteString(token["key"] + "=" + token["value"] + "\n")
	}

}

func Refresh(refreshToken string) {
	var refreshSession BluSkySession
	rootURL := "https://bsky.social"
	createSession := "/xrpc/com.atproto.server.refreshSession"
	req, err := http.NewRequest(http.MethodPost, rootURL+createSession, nil)
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", refreshToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("There has been an error %s", err)

	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &refreshSession)
	if err != nil {
		log.Printf("There has been an error %s", err)
	}

	fmt.Println(refreshSession)
	//utils.WriteToEnv(session)
}

// maybe rework this to read from a .env
func GetUserHandle(handle string) string {
	getProfileURL, err := url.Parse(publicBlueskyBaseURL + "/xrpc/app.bsky.actor.getProfile")
	if err != nil {
		log.Printf("Error parsing URL", err)
	}
	var user Profile
	params := url.Values{}
	params.Add("actor", handle)
	getProfileURL.RawQuery = params.Encode()
	req, err := http.NewRequest(http.MethodGet, getProfileURL.String(), nil)

	if err != nil {
		log.Printf("Error getting profile: %s ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("There has been an error %s", err)

	}
	defer resp.Body.Close()

	// Read and print the response
	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Printf("There has been an error %s", err)
	}

	return user.Handle
}
