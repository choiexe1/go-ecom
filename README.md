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

---

# Go 백엔드 작성 흐름

```
SQL → sqlc generate → Domain Type → Repository → Service → Handler → Router
```

## 순서

### 1. SQL 쿼리 추가
`internal/adapters/postgresql/sqlc/queries.sql`
```sql
-- name: CreateProduct :one
INSERT INTO products (name, price_cents, quantity)
VALUES ($1, $2, $3) RETURNING *;
```

### 2. sqlc generate 실행
```bash
sqlc generate
```

### 3. 도메인 타입 추가
`internal/{도메인}/types.go` - DB 타입과 분리
```go
type CreateProductParams struct {
    Name       string `json:"name"`
    PriceCents int32  `json:"priceCents"`
    Quantity   int32  `json:"quantity"`
}
```

### 4. Repository 인터페이스 추가
`internal/{도메인}/repository.go`
```go
Create(ctx context.Context, params CreateProductParams) (Product, error)
```

### 5. PostgreSQL 구현
`internal/{도메인}/postgres/repository.go` - 도메인 타입 ↔ sqlc 타입 변환

### 6. Service 추가
`internal/{도메인}/service.go` - 비즈니스 로직

### 7. Handler 추가
`internal/{도메인}/handlers.go` - 컨트롤러랑 동일, 요청/응답 처리, Service 호출

### 8. 라우터 연결
`cmd/api.go`
```go
r.Post("/products", productHandler.CreateProduct)
```

### 9. 빌드 확인
```bash
go build ./...
```
