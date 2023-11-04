FROM golang:1.21-alpine AS BUILDER
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/server

FROM scratch
COPY --from=BUILDER /bin/server /bin/
EXPOSE 8081
ENTRYPOINT ["/bin/server", "grpc"]