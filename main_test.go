package main

import "testing"

func TestExtractPlaylistID(t *testing.T) {
	tests := []struct {
		name     string
		link     string
		expected string
	}{
		{
			name:     "Valid Playlist Link",
			link:     "https://www.youtube.com/playlist?list=PL972B272B66F2D98B",
			expected: "PL972B272B66F2D98B",
		},
		{
			name:     "Playlist Link with Additional Parameters",
			link:     "https://www.youtube.com/playlist?list=PL972B272B66F2D98B&abc=123",
			expected: "PL972B272B66F2D98B",
		},
		{
			name:     "Invalid Link",
			link:     "https://www.youtube.com/watch?v=abcd1234",
			expected: "", // Expecting an empty string for invalid link
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractPlaylistID(tt.link)
			if result != tt.expected {
				t.Errorf("Got: %s, Expected: %s", result, tt.expected)
			}
		})
	}
}

func TestGenerateXMLFeedLink(t *testing.T) {
	tests := []struct {
		name       string
		playlistID string
		expected   string
	}{
		{
			name:       "Valid Playlist ID",
			playlistID: "PL972B272B66F2D98B",
			expected:   "https://www.youtube.com/feeds/videos.xml?playlist_id=PL972B272B66F2D98B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateXMLFeedLink(tt.playlistID)
			if result != tt.expected {
				t.Errorf("Got: %s, Expected: %s", result, tt.expected)
			}
		})
	}
}
