# ---------- base ----------
FROM golang:1.25-alpine AS base
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# ---------- dev ----------
FROM base AS dev
RUN apk add --no-cache bash curl
RUN go install github.com/air-verse/air@latest
COPY . .
CMD ["air", \
     "--build.cmd", "go build -o ./tmp/main .", \
     "--build.bin", "./tmp/main", \
     "--build.delay", "1000ms", \
     "--build.exclude_dir", "tmp,vendor,testdata", \
     "--build.include_ext", "go,gohtml,html", \
     "--build.stop_on_error", "false", \
     "--misc.clean_on_exit", "true", \
     "--log.main_only", "true"]

# ---------- builder ----------
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/main .

# ---------- prod ----------
FROM alpine:3.19 AS prod
RUN apk --no-cache add ca-certificates wget
RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup
WORKDIR /app
COPY --from=builder --chown=appuser:appgroup /app/main /app/main
COPY --from=builder --chown=appuser:appgroup /app/public /app/public
COPY --from=builder --chown=appuser:appgroup /app/templates /app/templates
USER appuser
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=30s --retries=3 \
  CMD ["wget","-q","--spider","http://127.0.0.1:8080/health"]
CMD ["./main"]