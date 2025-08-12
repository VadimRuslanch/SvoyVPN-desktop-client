param()
$ErrorActionPreference = "Stop"
protoc --go_out=core --go_opt=paths=source_relative `
       --go-grpc_out=core --go-grpc_opt=paths=source_relative `
       core/api/control.proto core/api/profile.proto
Write-Host "Protobuf generated."
