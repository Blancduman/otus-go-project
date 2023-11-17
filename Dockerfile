FROM golang:1.21-alpine AS BUILDER
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/twirler

FROM scratch
COPY --from=BUILDER /bin/twirler /bin/
EXPOSE 8081
ENTRYPOINT ["/bin/twirler", "grpc"]