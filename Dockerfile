FROM golang:1.25.7-trixie AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go


FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /
COPY --from=builder /app/main /main
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/main"]