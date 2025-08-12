package main

import (
    "log"
    "github.com/you/vadim-desktop/core/internal/ipc"
)

func main() {
    s, err := ipc.NewServer()
    if err != nil { log.Fatalf("ipc: %v", err) }
    log.Println("vadimd: starting gRPC server…")
    if err := s.Serve(); err != nil { log.Fatalf("serve: %v", err) }
}

