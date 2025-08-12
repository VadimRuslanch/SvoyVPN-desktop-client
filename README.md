# Вадим — Decktop (starter)

Стартовый каркас: Go Core + Electron UI + gRPC.

## Быстрый старт

### Зависимости
- Go 1.22+
- Node 20+ (npm)
- protoc 3.21+
- плагин для Go: 
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
  Убедитесь, что `$GOPATH/bin` в PATH.

### Генерация protobuf
```bash
./scripts/generate_protos.sh
# Windows (PowerShell):
./scripts/generate_protos.ps1
```

### Сборка демона
```bash
(cd core && go build ./cmd/vadimd)
```

### UI (dev)
```bash
(cd ui && npm i && npm run dev)
# в другом окне:
(cd ui && npm run electron)
```

Подробнее см. комментарии в коде.
