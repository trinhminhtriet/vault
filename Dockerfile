FROM golang:1.24-bullseye AS builder

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
	upx-ucl

WORKDIR /build

COPY . .

# Build
RUN go mod download && go mod tidy
RUN CGO_ENABLED=0 go build \
	-o ./dist/vault \
	&& upx-ucl --best --ultra-brute ./dist/vault

# final stage
FROM scratch

ARG APPLICATION="vault"
ARG DESCRIPTION="Vault is a CLI tool for securely decrypting encrypted data using RSA private key."
ARG PACKAGE="trinhminhtriet/vault"

LABEL org.opencontainers.image.ref.name="${PACKAGE}" \
	org.opencontainers.image.authors="Triet Trinh <contact@trinhminhtriet.com>" \
	org.opencontainers.image.documentation="https://github.com/${PACKAGE}/README.md" \
	org.opencontainers.image.description="${DESCRIPTION}" \
	org.opencontainers.image.licenses="MIT" \
	org.opencontainers.image.source="https://github.com/${PACKAGE}"

COPY --from=builder /build/dist/vault /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/vault"]
