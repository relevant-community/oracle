# Compile
FROM golang:alpine AS builder
WORKDIR /src/app/
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
# COPY ./scripts/* /home/nonroot/scripts
RUN for bin in cmd/*; do CGO_ENABLED=0 go build -o=/usr/local/bin/$(basename $bin) ./cmd/$(basename $bin); done
# RUN apk add --no-cache bash

# Add to a distroless container
FROM gcr.io/distroless/base:debug
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --chown=nonroot ./scripts/* /home/nonroot/scripts/
USER nonroot:nonroot
# RUN sh /home/nonroot/scripts/init.sh val 400token 4000000stake
CMD ["oracled"]
