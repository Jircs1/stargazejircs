# docker build . -t publicawesome/stargaze:latest
# docker run --rm -it publicawesome/stargaze:latest /bin/sh
FROM golang:1.21.3-alpine3.17 AS go-builder


RUN set -eux; apk add --no-cache ca-certificates build-base git;

# TARGETPLATFORM build argument provided by buildx: e.g., linux/amd64, linux/arm64. etc.
ARG TARGETPLATFORM=linux/amd64

# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

WORKDIR /code
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
# Download the correct version of libwasmvm for the given platform and verify checksum
RUN case "${TARGETPLATFORM}" in \
      "linux/amd64") \
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvm_muslc.x86_64.a" && \
        WASMVM_CHECKSUM="465e3a088e96fd009a11bfd234c69fb8a0556967677e54511c084f815cf9ce63" \
        ;; \
      "linux/arm64") \
        WASMVM_URL="https://github.com/CosmWasm/wasmvm/releases/download/v1.5.0/libwasmvm_muslc.aarch64.a" && \
        WASMVM_CHECKSUM="2687afbdae1bc6c7c8b05ae20dfb8ffc7ddc5b4e056697d0f37853dfe294e913" \
        ;; \
      *) \
        echo "Unsupported platform: ${TARGETPLATFORM}" ; \
        exit 1 \
        ;; \
    esac && \
    wget "${WASMVM_URL}" -O /lib/libwasmvm_muslc.a && \
    echo "${WASMVM_CHECKSUM}  /lib/libwasmvm_muslc.a" | sha256sum -c

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN  LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true  make build


# --------------------------------------------------------
FROM alpine:3

COPY --from=go-builder /code/bin/starsd /usr/bin/starsd
RUN apk add -U --no-cache ca-certificates
WORKDIR /data
ENV HOME=/data
COPY ./docker/entry-point.sh ./entry-point.sh
# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657


CMD ["starsd", "start", "--pruning", "nothing", "--log_format", "json"]
