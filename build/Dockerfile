# Build the binary
ARG GO_VERSION
FROM rickroche/spin-it:dev-latest as builder
WORKDIR /home/go/apps
COPY --chown=go:go . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main cmd/main.go

# Build the secure runtime
FROM scratch
# Import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /home/go/apps/api/swagger /home/go/apps/api/swagger/
COPY --from=builder /home/go/apps/third_party/swaggerui /home/go/apps/third_party/swaggerui/
COPY --from=builder /home/go/apps/main /home/go/apps/main
# Use an unprivileged user.
USER go
WORKDIR /home/go/apps
ENTRYPOINT ["/home/go/apps/main"]
