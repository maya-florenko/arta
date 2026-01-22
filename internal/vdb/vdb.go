package vdb

import (
	"github.com/qdrant/go-client/qdrant"
)

func main() {
	c, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
}
