# Build in go env
FROM golang:1.17.2 AS builder
WORKDIR /go/src/
# Go mod will be cached in this way (if mod/sum is not modified)
COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
# Copy remaining source files
COPY . .
# go-sqlite need cgo_enable and some other args: https://7thzero.com/blog/golang-w-sqlite3-docker-scratch-image
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o ./launchserver ./cmd/main.go

# Start fresh from a smaller image
# debug in 'alpine' if needed and deploy in 'scratch'
FROM scratch
# FROM alpine
WORKDIR /app/
# if certificates needed
# RUN apk add --update --no-cache ca-certificates git
COPY --from=builder /go/src/launchserver ./
# RUN ["chmod", "+x", "./launchserver"]
CMD ["./launchserver"]
