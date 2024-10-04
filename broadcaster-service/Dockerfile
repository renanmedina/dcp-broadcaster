FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./dist/dcp-broadcaster main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build ./app/dist/ ./
COPY --from=build ./app/db/ ./db/

CMD ["./dcp-broadcaster"]
