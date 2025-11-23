FROM golang:1.24.10-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY . .
RUN go build -o main ./cmd/api
EXPOSE 8000
CMD sh -c "echo && \
           sleep 5 && \
           goose -dir ./migrations postgres 'user=postgres password=root dbname=app host=postgres sslmode=disable' up && \
           ./main"