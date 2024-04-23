build:
	go build -o bin/yt_playlist_to_rss

run: build
	./bin/yt_playlist_to_rss

test: 
	go test -v ./...