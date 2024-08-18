# Auto-generated Dockerfile file.
# See https://gowebly.org for more information.

FROM golang:1.22-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY ./go.mod go.sum ./

# Download dependencies.
RUN go mod download

# Copy your code into the container.
COPY . .

# Set necessary environment variables and build your project.
ENV CGO_ENABLED=0
RUN go build -o server cmd/server/main.go

FROM scratch

# Copy project's binary and templates from /build to the scratch container.
COPY --from=builder /build/server /


# Set entry point.
ENTRYPOINT ["/server"]
