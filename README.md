# 🏏 Turf Slot Booking System

> A scalable backend system for booking box cricket turf slots — built incrementally from a clean monolith to a distributed microservices architecture using **Go**.

This project is a **learning-by-building** journey. The goal is not just "a booking app" — it's to demonstrate real backend engineering: concurrency control, transactional safety, distributed architecture, async event processing, and production-grade observability.

---

## Table of Contents

- [Project Overview](#project-overview)
- [High-Level Architecture](#high-level-architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Phased Roadmap](#phased-roadmap)
- [Go Learning Track](#go-learning-track)
- [Setup & Run](#setup--run)
- [Commit Convention](#commit-convention)

---

## Project Overview

**Turfs** are box cricket grounds available for hourly booking. Users browse available turfs, view open time slots, and book them. Turf owners manage their grounds and slot schedules.

### Core Problems This Project Solves

| Problem                                        | Engineering Skill                                           |
| ---------------------------------------------- | ----------------------------------------------------------- |
| Two users booking the same slot simultaneously | Transactional concurrency control (`SELECT ... FOR UPDATE`) |
| Slot held during payment but not yet confirmed | Distributed temporary reservation (Redis TTL)               |
| Notifying users on booking confirmation        | Async event-driven processing (Kafka)                       |
| Internal service-to-service calls              | gRPC + Protobuf contracts                                   |
| Observing system health in production          | Prometheus metrics + Grafana dashboards                     |
| Deploying reliably                             | Docker Compose → Kubernetes                                 |

---

## High-Level Architecture

### Phase 1 — Clean Monolith

```
┌──────────────┐
│   Frontend   │  (Angular/React — later)
└──────┬───────┘
       │ HTTP
┌──────▼───────┐
│  API Layer   │  Go + Gin
│  (Handlers)  │
└──────┬───────┘
       │
┌──────▼───────┐
│   Service    │  Business Logic
│    Layer     │
└──────┬───────┘
       │
┌──────▼───────┐
│  Repository  │  Database Access
│    Layer     │
└──────┬───────┘
       │
┌──────▼───────┐     ┌───────────┐
│  PostgreSQL  │     │   Redis   │
└──────────────┘     └───────────┘
```

### Phase 3 — Microservices (Target)

```
┌────────────┐
│ API Gateway│
└─────┬──────┘
      │
┌─────▼──────┬──────────────┬───────────────┐
│Auth Service│Booking Service│Notification   │
│            │               │Service        │
└─────┬──────┴───────┬───────┴───────┬───────┘
      │    gRPC      │               │
      └──────────────┘               │
                                     │ Kafka Consumer
┌────────────┐  ┌─────────┐  ┌──────▼──────┐
│ PostgreSQL │  │  Redis  │  │   Kafka     │
└────────────┘  └─────────┘  └─────────────┘
```

---

## Tech Stack

| Category                    | Technology              |
| --------------------------- | ----------------------- |
| Language                    | Go 1.22+                |
| Web Framework               | Gin                     |
| Database                    | PostgreSQL 16           |
| Cache / Temp Reservation    | Redis 7                 |
| Message Broker              | Apache Kafka            |
| Inter-service Communication | gRPC + Protobuf         |
| Authentication              | JWT + bcrypt            |
| API Documentation           | Swagger (swaggo/swag)   |
| Logging                     | Zap (structured JSON)   |
| Metrics                     | Prometheus + Grafana    |
| Containerization            | Docker + Docker Compose |
| Orchestration               | Kubernetes (Phase 4)    |
| CI/CD                       | GitHub Actions          |
| Migration Tool              | golang-migrate          |
| Testing                     | Go testing + testify    |

---

## Project Structure

```
turf-booking-system/
├── cmd/
│   └── server/
│       └── main.go                 # Application entrypoint
│
├── internal/                       # Private application code
│   ├── auth/
│   │   ├── handler/
│   │   │   └── auth_handler.go     # HTTP handlers for auth
│   │   ├── service/
│   │   │   └── auth_service.go     # Auth business logic
│   │   ├── repository/
│   │   │   └── auth_repository.go  # DB queries for users
│   │   └── model/
│   │       └── user.go             # User domain model
│   │
│   ├── turf/
│   │   ├── handler/
│   │   │   └── turf_handler.go
│   │   ├── service/
│   │   │   └── turf_service.go
│   │   ├── repository/
│   │   │   └── turf_repository.go
│   │   └── model/
│   │       └── turf.go
│   │
│   ├── booking/
│   │   ├── handler/
│   │   │   └── booking_handler.go
│   │   ├── service/
│   │   │   └── booking_service.go
│   │   ├── repository/
│   │   │   └── booking_repository.go
│   │   └── model/
│   │       ├── slot.go
│   │       └── booking.go
│   │
│   ├── middleware/
│   │   ├── auth.go                 # JWT auth middleware
│   │   ├── logging.go              # Request logging
│   │   └── ratelimit.go            # Rate limiting
│   │
│   └── config/
│       └── config.go               # App configuration loader
│
├── pkg/                            # Shared/public utilities
│   ├── database/
│   │   └── postgres.go             # DB connection helper
│   ├── cache/
│   │   └── redis.go                # Redis connection helper
│   └── response/
│       └── response.go             # Standardized API responses
│
├── migrations/
│   ├── 000001_create_users.up.sql
│   ├── 000001_create_users.down.sql
│   ├── 000002_create_turfs.up.sql
│   ├── 000002_create_turfs.down.sql
│   ├── 000003_create_slots.up.sql
│   ├── 000003_create_slots.down.sql
│   ├── 000004_create_bookings.up.sql
│   └── 000004_create_bookings.down.sql
│
├── configs/
│   └── config.yaml                 # Environment config
│
├── scripts/
│   └── seed.sql                    # Seed data for dev
│
├── docs/                           # Swagger generated docs
│
├── docker-compose.yml
├── Dockerfile
├── .env.example
├── .gitignore
├── Makefile
└── go.mod
```

---

## Database Schema

### Entity Relationship

```
┌──────────┐       ┌──────────┐       ┌──────────┐       ┌──────────┐
│  Users   │       │  Turfs   │       │  Slots   │       │ Bookings │
├──────────┤       ├──────────┤       ├──────────┤       ├──────────┤
│ id (PK)  │◄──┐   │ id (PK)  │◄──┐   │ id (PK)  │◄──┐   │ id (PK)  │
│ name     │   │   │ name     │   │   │ turf_id  │───┘   │ slot_id  │───┐
│ email    │   │   │ location │   │   │ date     │       │ user_id  │───┤
│ phone    │   │   │ owner_id │───┘   │ start    │       │ status   │   │
│ password │   │   │ sport    │       │ end      │       │ payment  │   │
│ role     │   │   │ price/hr │       │ status   │       │ amount   │   │
│ created  │   │   │ created  │       │ created  │       │ created  │   │
└──────────┘   │   └──────────┘       └──────────┘       └──────────┘   │
               │                                                         │
               └─────────────────────────────────────────────────────────┘
```

### Tables Detail

```sql
-- Users: both players and turf owners
CREATE TABLE users (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    phone       VARCHAR(15),
    password_hash TEXT NOT NULL,
    role        VARCHAR(20) NOT NULL DEFAULT 'player',  -- 'player' | 'owner' | 'admin'
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

-- Turfs: box cricket grounds
CREATE TABLE turfs (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(200) NOT NULL,
    location    TEXT NOT NULL,
    city        VARCHAR(100) NOT NULL,
    owner_id    UUID NOT NULL REFERENCES users(id),
    sport_type  VARCHAR(50) DEFAULT 'cricket',  -- 'cricket' | 'football' | 'badminton'
    price_per_hour DECIMAL(10,2) NOT NULL,
    amenities   TEXT[],                          -- {'floodlights','changing_room','parking'}
    is_active   BOOLEAN DEFAULT true,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

-- Slots: available time windows per turf per day
CREATE TABLE slots (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    turf_id     UUID NOT NULL REFERENCES turfs(id),
    date        DATE NOT NULL,
    start_time  TIME NOT NULL,
    end_time    TIME NOT NULL,
    status      VARCHAR(20) DEFAULT 'available',  -- 'available' | 'reserved' | 'booked'
    created_at  TIMESTAMP DEFAULT NOW(),
    UNIQUE(turf_id, date, start_time)
);

-- Bookings: links user to slot
CREATE TABLE bookings (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slot_id         UUID NOT NULL REFERENCES slots(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    status          VARCHAR(20) DEFAULT 'pending',   -- 'pending' | 'confirmed' | 'cancelled'
    payment_status  VARCHAR(20) DEFAULT 'unpaid',    -- 'unpaid' | 'paid' | 'refunded'
    amount          DECIMAL(10,2) NOT NULL,
    booked_at       TIMESTAMP DEFAULT NOW(),
    cancelled_at    TIMESTAMP,
    UNIQUE(slot_id)  -- one booking per slot
);
```

### Key Design Decisions

| Decision                                     | Why                                                             |
| -------------------------------------------- | --------------------------------------------------------------- |
| `UNIQUE(slot_id)` on bookings                | Prevents double booking at the DB level                         |
| `UNIQUE(turf_id, date, start_time)` on slots | No duplicate slot creation                                      |
| UUID primary keys                            | Safe for distributed systems later                              |
| Separate `status` and `payment_status`       | Booking state and payment state are independent concerns        |
| `role` on users                              | Supports player, owner, and admin flows                         |
| `amenities` as `TEXT[]`                      | PostgreSQL array — avoids a separate join table for simple data |

---

## API Endpoints

### Auth

| Method | Endpoint              | Description              | Auth |
| ------ | --------------------- | ------------------------ | ---- |
| POST   | `/api/v1/auth/signup` | Register new user        | No   |
| POST   | `/api/v1/auth/login`  | Login, returns JWT       | No   |
| GET    | `/api/v1/auth/me`     | Get current user profile | Yes  |

### Turfs

| Method | Endpoint            | Description               | Auth        |
| ------ | ------------------- | ------------------------- | ----------- |
| POST   | `/api/v1/turfs`     | Create turf (owner only)  | Yes (owner) |
| GET    | `/api/v1/turfs`     | List turfs (with filters) | No          |
| GET    | `/api/v1/turfs/:id` | Get turf details          | No          |
| PUT    | `/api/v1/turfs/:id` | Update turf               | Yes (owner) |
| DELETE | `/api/v1/turfs/:id` | Soft delete turf          | Yes (owner) |

### Slots

| Method | Endpoint                  | Description                    | Auth        |
| ------ | ------------------------- | ------------------------------ | ----------- |
| POST   | `/api/v1/turfs/:id/slots` | Create slots for a turf        | Yes (owner) |
| GET    | `/api/v1/turfs/:id/slots` | List available slots (by date) | No          |
| PUT    | `/api/v1/slots/:id`       | Update slot                    | Yes (owner) |

### Bookings

| Method | Endpoint                      | Description     | Auth |
| ------ | ----------------------------- | --------------- | ---- |
| POST   | `/api/v1/bookings`            | Book a slot     | Yes  |
| GET    | `/api/v1/bookings`            | My bookings     | Yes  |
| GET    | `/api/v1/bookings/:id`        | Booking details | Yes  |
| PUT    | `/api/v1/bookings/:id/cancel` | Cancel booking  | Yes  |

---

## Phased Roadmap

### Phase 1A — Hello Go ✦ Language Foundations (build tiny things first)

> Spend 3-5 days here. Don't touch the project yet. Write throwaway `.go` files.

| Step | Mini Exercise                                                      | Go Concepts                                                                |
| ---- | ------------------------------------------------------------------ | -------------------------------------------------------------------------- |
| 0.1  | Print "Hello Turf" + accept CLI args                               | `fmt`, `os.Args`, `func main`, packages                                    |
| 0.2  | Struct for `Turf` with methods, print as JSON                      | structs, methods, pointer vs value receivers, `encoding/json`, struct tags |
| 0.3  | Read a config file (YAML/env)                                      | `os`, `io`, file I/O, error handling (`if err != nil`)                     |
| 0.4  | HTTP server that returns JSON                                      | `net/http`, `http.HandlerFunc`, `json.Marshal`                             |
| 0.5  | Interface exercise: `Storage` interface with in-memory + file impl | interfaces, implicit satisfaction, polymorphism                            |
| 0.6  | Goroutine: simulate 3 users booking same slot concurrently         | `go func()`, `sync.Mutex`, race conditions, `go run -race`                 |
| 0.7  | Channels: producer sends slots, consumer books them                | channels, `select`, buffered channels, deadlocks                           |
| 0.8  | Context: HTTP handler with timeout                                 | `context.WithTimeout`, `context.WithCancel`, propagation                   |

**Why this matters**: You'll hit every Go gotcha (nil pointers, goroutine leaks, interface confusion) in isolation — not buried in project code where it's harder to debug.

### Phase 1B — Clean Monolith ✦ Foundation

| Step | What You Build                           | Go Concepts You Reinforce                                            | Checkpoint                                             |
| ---- | ---------------------------------------- | -------------------------------------------------------------------- | ------------------------------------------------------ |
| 1.1  | Project init + Gin health endpoint       | modules, `go mod`, packages, third-party deps                        | `curl localhost:8080/health` returns `{"status":"ok"}` |
| 1.2  | Config loader (env vars → struct)        | `os.LookupEnv`, struct composition, `error`                          | App reads `.env` and prints config on startup          |
| 1.3  | PostgreSQL via Docker Compose            | Docker basics, `docker-compose up`                                   | `psql` connects to the DB                              |
| 1.4  | Database connection + health check       | `database/sql`, connection pooling, `defer db.Close()`               | `/health` also reports DB status                       |
| 1.5  | Migrations (create all 4 tables)         | SQL DDL, `golang-migrate`                                            | Tables exist in DBeaver                                |
| 1.6  | User repository (Create + GetByEmail)    | interfaces, pointer receivers, `sql.Row.Scan`, error wrapping        | Unit test passes                                       |
| 1.7  | Auth service (signup with bcrypt)        | `bcrypt.GenerateFromPassword`, dependency injection via constructors | Postman: signup works                                  |
| 1.8  | Auth service (login + JWT)               | `jwt-go`, token claims, `time.Now().Add()`                           | Postman: login returns token                           |
| 1.9  | Auth middleware                          | `gin.HandlerFunc`, `c.Set`/`c.Get`, `context.Value`                  | Protected routes reject bad tokens                     |
| 1.10 | Turf CRUD (full handler→service→repo)    | complete layered flow, request binding, validation                   | All turf endpoints work via Postman                    |
| 1.11 | Slot management                          | `time.Time`, `time.Parse`, query building, date filtering            | Create/list slots by turf+date                         |
| 1.12 | Booking API (basic — no concurrency yet) | transactions (`sql.Tx`), `Begin`/`Commit`/`Rollback`                 | Book a slot, verify status changes                     |
| 1.13 | Graceful shutdown                        | `os.Signal`, `signal.Notify`, `http.Server.Shutdown`                 | Ctrl+C drains in-flight requests                       |
| 1.14 | Unit tests for service layer             | `testing`, `testify`, table-driven tests, mocks                      | `go test ./...` passes                                 |
| 1.15 | Swagger docs                             | `swaggo/swag` annotations                                            | `/swagger/index.html` renders                          |
| 1.16 | Dockerize the Go app                     | multi-stage `Dockerfile`, `docker-compose` with app                  | `docker-compose up` runs everything                    |

### Phase 2 — Real Engineering ✦ Concurrency & Caching

| Step | What You Build                          | Go Concepts You Learn                                | Checkpoint                                           |
| ---- | --------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| 2.1  | Solve double booking (row locking)      | `SELECT ... FOR UPDATE`, `sql.TxOptions{Isolation}`  | Run 10 concurrent booking requests → only 1 succeeds |
| 2.2  | Redis integration + connection pooling  | `go-redis`, connection options, `context` with Redis | Redis `PING` in health check                         |
| 2.3  | Temporary slot reservation (5-min hold) | Redis `SET EX`, TTL, `GET`+`DEL` pattern             | Slot auto-releases after 5 min                       |
| 2.4  | Rate limiting middleware                | `time.Ticker`, token bucket, Redis `INCR`+`EXPIRE`   | 429 after exceeding limit                            |
| 2.5  | Background worker: expire held slots    | `goroutine`, `time.NewTicker`, `context.WithCancel`  | Expired reservations auto-cleanup                    |
| 2.6  | Worker pool: batch slot generation      | `channels`, `sync.WaitGroup`, worker pool pattern    | Owner creates week of slots in one call              |
| 2.7  | Integration tests with Testcontainers   | `testcontainers-go`, real DB in tests, `t.Cleanup`   | Tests spin up Postgres+Redis in Docker               |
| 2.8  | Load testing with `hey` or `k6`         | observe system under concurrent load                 | Report: X bookings/sec, 0 double bookings            |

### Phase 3 — Microservices ✦ Distribution

| Step | What You Build                        | Go Concepts You Learn                                       | Checkpoint                                   |
| ---- | ------------------------------------- | ----------------------------------------------------------- | -------------------------------------------- |
| 3.1  | Structured logging (Zap) — add FIRST  | `zap.Logger`, `zap.Field`, structured JSON logs             | All requests logged with request_id, user_id |
| 3.2  | Extract Auth into standalone service  | separate `go.mod`, service boundaries, `net/http` client    | Auth service runs on :8081                   |
| 3.3  | gRPC: Auth service contract           | protobuf definitions, `protoc` code gen, gRPC server/client | Booking service validates tokens via gRPC    |
| 3.4  | Kafka: publish booking events         | `confluent-kafka-go`, producer, serialization               | Events appear in Kafka topic                 |
| 3.5  | Notification service (Kafka consumer) | consumer groups, offset management, retry logic             | Email/log on booking confirmation            |
| 3.6  | Dead-letter queue + retry patterns    | error handling in consumers, exponential backoff            | Failed events land in DLQ                    |
| 3.7  | API Gateway                           | `net/http/httputil.ReverseProxy`, routing, auth forwarding  | Single entry point routes to services        |

### Phase 4 — Production Engineering ✦ Observability & Deployment

| Step | What You Build                       | Go Concepts You Learn                                    | Checkpoint                                    |
| ---- | ------------------------------------ | -------------------------------------------------------- | --------------------------------------------- |
| 4.1  | Prometheus metrics                   | `prometheus/client_golang`, histograms, counters         | `/metrics` endpoint exposes data              |
| 4.2  | Grafana dashboards                   | visualization, alerting rules                            | Dashboard shows booking latency, error rate   |
| 4.3  | Distributed tracing (Jaeger)         | OpenTelemetry SDK, trace propagation, spans              | Traces show request flow across services      |
| 4.4  | Health checks (liveness + readiness) | separate `/healthz` and `/readyz` with dependency checks | Probes report DB/Redis/Kafka status           |
| 4.5  | GitHub Actions CI/CD                 | workflow YAML, test → build → push Docker image          | PR triggers test, merge triggers deploy       |
| 4.6  | Kubernetes deployment                | Deployments, Services, Ingress, ConfigMaps, Secrets      | `kubectl get pods` shows all services running |

---

## Go Learning Track

### Before You Start Coding (Phase 1A — 3 to 5 days)

These are the exercises in Phase 1A above. Do them as throwaway `.go` files in a `playground/` folder:

```
□  0.1 — fmt, os.Args, func main, packages
□  0.2 — structs, methods, pointer receivers, JSON encoding
□  0.3 — file I/O, os.Open, error handling
□  0.4 — net/http server, HandlerFunc, JSON response
□  0.5 — interfaces, implicit implementation, polymorphism
□  0.6 — goroutines, sync.Mutex, race detector
□  0.7 — channels, select, buffered channels
□  0.8 — context.WithTimeout, context.WithCancel
```

**Key rule**: If you get stuck on a Go concept while building the project, STOP and write a tiny isolated program to understand it. Then come back.

### Concept Progression (mapped to project phases)

| Phase            | Go Concepts Introduced                                                                                           | Mental Model                             |
| ---------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| 1A (playground)  | All core syntax: types, structs, interfaces, pointers, goroutines, channels, context                             | "How does Go work?"                      |
| 1B (monolith)    | `database/sql`, `http.Handler`, DI via constructors, `defer`, `error` wrapping, struct tags, middleware chaining | "How do real Go apps work?"              |
| 2 (concurrency)  | `sync.WaitGroup`, worker pools, `time.Ticker`, `context.WithCancel`, `select` with real use cases                | "How do Go apps handle concurrency?"     |
| 3 (distribution) | Multi-module, gRPC/protobuf, Kafka producers/consumers, `io.Reader`/`io.Writer`, structured logging              | "How do Go services talk to each other?" |
| 4 (production)   | OpenTelemetry, `embed`, build tags, custom metrics, `expvar`                                                     | "How do Go apps behave in production?"   |

### Go Gotchas You'll Hit (and when)

| Gotcha                                     | When You'll Hit It                                 | Fix                                                                |
| ------------------------------------------ | -------------------------------------------------- | ------------------------------------------------------------------ |
| Nil pointer dereference                    | Phase 1B step 1.6 (repository returns `nil`)       | Always check `if result == nil` before using                       |
| Goroutine leak                             | Phase 2 step 2.5 (worker never exits)              | Always use `context.Done()` or close channels                      |
| Race condition                             | Phase 2 step 2.1 (double booking)                  | Use `go test -race`, then `sync.Mutex` or DB locks                 |
| Interface confusion (`*T` vs `T` receiver) | Phase 1B step 1.6 (repo doesn't satisfy interface) | Pointer receivers → only `*T` satisfies; value receivers → both do |
| `defer` in a loop                          | Phase 1B step 1.12 (DB connections in loop)        | Extract loop body to a function, or close manually                 |
| Shadowed `err` with `:=`                   | Everywhere                                         | Use `=` when `err` already exists in scope                         |
| JSON field not exported                    | Phase 1A step 0.2                                  | Struct fields must be `Uppercase` + use `json:"lowercase"` tag     |

### Resources (pick ONE per topic, don't hoard)

| Topic        | Resource                                                                               |
| ------------ | -------------------------------------------------------------------------------------- |
| Go basics    | [Go by Example](https://gobyexample.com) — do every example                            |
| Go tour      | [tour.golang.org](https://go.dev/tour/) — interactive, 1 hour                          |
| Web apps     | [Let's Go by Alex Edwards](https://lets-go.alexedwards.net/) — best Go web book        |
| Concurrency  | [Go Concurrency Patterns (Rob Pike talk)](https://www.youtube.com/watch?v=f6kdp27TYZs) |
| Testing      | [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)                    |
| Architecture | [Standard Go Project Layout](https://github.com/golang-standards/project-layout)       |

### Additional Topics (my additions to your plan)

| Topic                                                       | Why It Matters                                        |
| ----------------------------------------------------------- | ----------------------------------------------------- |
| **Graceful shutdown** (`os.Signal`, `http.Server.Shutdown`) | Production servers must drain in-flight requests      |
| **Database connection pooling** (`sql.DB.SetMaxOpenConns`)  | Prevents connection exhaustion under load             |
| **Request validation** (go-playground/validator)            | Never trust client input — validate at the boundary   |
| **Pagination** (cursor-based vs offset)                     | Real APIs need pagination; cursor-based scales better |
| **Idempotency keys** on booking creation                    | Prevents duplicate bookings from network retries      |
| **Health checks** (liveness + readiness)                    | Required for Kubernetes; good practice everywhere     |
| **Configuration management** (Viper)                        | Environment-specific config without code changes      |
| **Error types and sentinel errors**                         | Clean error handling across layers                    |
| **Table-driven tests**                                      | Idiomatic Go testing pattern — interviewers love this |
| **Makefile** for common commands                            | One `make run`, `make test`, `make migrate-up`        |

---

## Setup & Run

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Postman](https://www.postman.com/) or Thunder Client (VS Code extension)
- [DBeaver](https://dbeaver.io/) or pgAdmin (optional, for DB inspection)

### Quick Start

```bash
# Clone the repo
git clone https://github.com/<your-username>/turf-booking-system.git
cd turf-booking-system

# Start infrastructure (Postgres + Redis)
docker-compose up -d

# Run database migrations
make migrate-up

# Start the server
make run

# Health check
curl http://localhost:8080/api/v1/health
```

### Environment Variables

Copy `.env.example` to `.env` and fill in values:

```env
# Server
SERVER_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=turf_user
DB_PASSWORD=turf_password
DB_NAME=turf_booking
DB_SSLMODE=disable

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379

# JWT
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRY=24h
```

---

## Commit Convention

Use conventional commits to show progression:

```
feat: initialize project with Gin health endpoint
feat: add PostgreSQL with Docker Compose
feat: create database migrations for core tables
feat: implement user signup and login with JWT
feat: add CRUD for turfs
feat: implement slot management APIs
feat: add booking with transaction-based concurrency control
feat: integrate Redis for temporary slot reservation
feat: add background worker for booking expiry
feat: introduce Kafka for async notifications
feat: extract auth into standalone gRPC service
feat: add Prometheus metrics and Grafana dashboard
feat: dockerize all services with CI/CD pipeline
```

---

## License

MIT

---

> **Remember**: Don't try to build the perfect architecture. Build incrementally.
> `working system → scalable system → distributed system → observable system`
> That progression itself is what makes this project impressive.
