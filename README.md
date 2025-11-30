# Commands

## Docker
```bash
docker compose up -d      # PostgreSQL 실행
docker compose down       # PostgreSQL 종료
docker compose down -v    # PostgreSQL 볼륨 제거 및 종료
```

## DB Migration (Goose)
```bash
goose -s create <name> sql   # Migration 생성
goose up                     # Migration 적용
goose down                   # Migration 롤백
```

## SQLC Generate
```bash
sqlc generate
```

## Run Server
```bash
go run ./cmd
```
