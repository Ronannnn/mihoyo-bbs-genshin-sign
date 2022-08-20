package main

import (
	"context"
	"mihoyo-bbs-genshin-sign/server"
)

func main() {
	server.NewServer(context.Background())
}
