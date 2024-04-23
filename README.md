### YouTube Playlist to RSS Converter

This Go program converts a YouTube playlist link into an RSS feed link.

### Usage

#### Prerequisites

- Go installed on your system.

#### Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/VBeshta/yt_playlist_to_rss.git
   cd yt_playlist_to_rss
2. Build the program:
    ```bash
    make build
    ```
3. Run the program:
    ```bash
    make run
    ```
    Or run the binary directly:
    ```bash
    ./bin/yt_playlist_to_rss
    ```
4. Enter the YouTube playlist link when prompted.

#### Example
```
Please provide a YouTube playlist link:  https://www.youtube.com/playlist?list=PL972B272B66F2D98B

https://www.youtube.com/feeds/videos.xml?playlist_id=PL972B272B66F2D98B
```