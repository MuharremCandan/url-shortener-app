# Use a minimal base image with Go 1.21.4

FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go



# Use a minimal base image for the final image
FROM scratch
WORKDIR /app
COPY config.yml /app/
COPY --from=builder /app/main .


# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]