# Tech Stack

| Tool | Role |
|------|------|
| **Go** | 백엔드 언어 |
| **Chi** | HTTP 라우터, 미들웨어 |
| **PostgreSQL** | 데이터베이스 |
| **pgx/v5** | Go용 PostgreSQL 드라이버 |
| **sqlc** | SQL을 타입 안전한 Go 코드로 생성 |
| **goose** | 데이터베이스 마이그레이션 관리 |
| **Docker** | PostgreSQL 컨테이너 실행 |

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
