package main

import "rekompose.com/engine/mime"

func main() {
	mime.Parse([]byte("invalid"))
}
