package log

import (
	"context"
	"log"
)

type StdOut struct{}

func (s StdOut) Info(_ context.Context, msg interface{}) {
	log.Printf("[INFO] %s\n", msg)
}

func (s StdOut) Warn(_ context.Context, msg interface{}) {
	log.Printf("[WARN] %s\n", msg)
}

func (s StdOut) Error(_ context.Context, msg interface{}) {
	log.Printf("[ERROR] %s\n", msg)
}

func (s StdOut) Fatal(_ context.Context, msg interface{}) {
	log.Fatalf("[FATAL] %s\n", msg)
}
