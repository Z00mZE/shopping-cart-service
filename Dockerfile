FROM golang:latest AS Builder
RUN mkdir /app
ADD ./commands /app
WORKDIR /app
RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest AS production
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
CMD ["/app/main"]