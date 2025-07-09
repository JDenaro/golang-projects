# Multi-stage build for smaller final image
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o bootstrap \
    ./cmd/lambda

# Final stage
FROM public.ecr.aws/lambda/provided:al2

# Copy the binary from builder stage
COPY --from=builder /app/bootstrap ${LAMBDA_TASK_ROOT}

# Create non-root user for security
RUN groupadd -r lambda && useradd -r -g lambda lambda
USER lambda

# Set the CMD to your handler
CMD ["bootstrap"] 