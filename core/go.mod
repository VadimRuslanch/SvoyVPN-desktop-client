module github.com/you/vadim-desktop/core

go 1.22

require (
    google.golang.org/grpc v1.66.0
    google.golang.org/protobuf v1.34.2
    go.uber.org/zap v1.27.0
    gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

replace github.com/you/vadim-desktop/core => .