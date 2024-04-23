package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	playlistLink := stringPrompt("Please provide a YouTube playlist link: ")

	playlistID := extractPlaylistID(playlistLink)

	if playlistID == "" {
		fmt.Println("Invalid YouTube playlist link.")
		return
	}

	xmlFeedLink := generateXMLFeedLink(playlistID)

	fmt.Println(xmlFeedLink)
}

// The entered string is then trimmed of leading and trailing whitespace before being returned.
func stringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

// Function to extract the playlist ID from the provided link
func extractPlaylistID(link string) string {
	parts := strings.Split(link, "list=")
	if len(parts) < 2 {
		return "" // Invalid link
	}
	playlistID := parts[1]

	// Remove any additional parameters after the playlist ID
	ampersandIndex := strings.Index(playlistID, "&")
	if ampersandIndex != -1 {
		playlistID = playlistID[:ampersandIndex]
	}

	return playlistID
}

// Function to generate the XML feed link
func generateXMLFeedLink(playlistID string) string {
	return fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?playlist_id=%s", playlistID)
}
