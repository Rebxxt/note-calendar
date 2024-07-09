FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o bin/proto ./cmd

FROM alpine

COPY --from=builder /app/bin/notes /notes

CMD ["/notes"]
