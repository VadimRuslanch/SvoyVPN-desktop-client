package ipc

import (
    "net"
    "os"
    "runtime"
    "google.golang.org/grpc"
    ctrl "github.com/you/vadim-desktop/core/api/control"
    "github.com/you/vadim-desktop/core/internal/app"
)

type Server struct { grpcServer *grpc.Server }

func socketPath() string {
    if runtime.GOOS == "windows" {
        return `\\.\pipe\vadim-desktop` // Named Pipe (заменить слушатель ниже)
    }
    return "/tmp/vadim-desktop.sock" // Unix domain socket
}

func NewServer() (*Server, error) {
    gs := grpc.NewServer(
        // TODO: interceptors + проверка локального клиента
    )
    ctrlSvc := app.NewService()
    ctrl.RegisterVadimControlServer(gs, ctrlSvc)
    return &Server{grpcServer: gs}, nil
}

func (s *Server) Serve() error {
    addr := socketPath()
    if runtime.GOOS != "windows" {
        _ = os.Remove(addr)
        l, err := net.Listen("unix", addr)
        if err != nil { return err }
        return s.grpcServer.Serve(l)
    }
    // TODO: Windows Named Pipe listener вместо net.Listen("unix", ...)
    l, err := net.Listen("tcp", "127.0.0.1:50055") // временно для разработки
    if err != nil { return err }
    return s.grpcServer.Serve(l)
}

