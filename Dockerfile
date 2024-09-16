FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./dist/dcp-broadcaster cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=build ./dist/ ./

CMD ["./dcp-broadcaster"]
