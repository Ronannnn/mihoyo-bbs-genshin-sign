package main

import (
	"context"
	"mihoyo-bbs-genshin-sign/internal/server"
)

func main() {
	server.NewServer(context.Background())
}
