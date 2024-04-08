FROM golang:1.22

RUN useradd -u 10001 appuser
USER appuser

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY controllers controllers
COPY models models
COPY helpers helpers
COPY database.db database.db
COPY main.go .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
