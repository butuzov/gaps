# ----- Dependencies -----------------------------------------------------------
FROM golang:1.26-alpine AS deps
WORKDIR /build
COPY	go.mod go.sum .
RUN 	--mount=type=cache,target=/go/pkg/mod \
			go mod download -x

# ----- Builder ----------------------------------------------------------------

FROM deps AS builder
COPY     . .
RUN		 go build -trimpath -o bin/gaps ./

# ----- Runner -----------------------------------------------------------------

FROM alpine:3.23.3 AS base
WORKDIR    /
COPY       --from=builder /build/bin/gaps gaps
WORKDIR    /app
VOLUME     /app
ENTRYPOINT [ "/gaps" ]
