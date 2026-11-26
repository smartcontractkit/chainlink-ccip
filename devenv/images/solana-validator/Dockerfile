# Agave no longer publishes docker images: `release` wraps the x86_64 release binaries, `source`
# compiles Agave for other architectures.

ARG AGAVE_VERSION=4.4.1

FROM debian:bookworm-slim AS release-download
ARG AGAVE_VERSION
RUN apt-get update && \
    apt-get install -y --no-install-recommends bzip2 ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*
RUN curl -sSfL "https://github.com/anza-xyz/agave/releases/download/v${AGAVE_VERSION}/solana-release-x86_64-unknown-linux-gnu.tar.bz2" | tar -xj && \
    mkdir /out && \
    cp solana-release/bin/solana-test-validator solana-release/bin/solana /out/

FROM debian:bookworm AS source-build
ARG AGAVE_VERSION
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        build-essential ca-certificates clang cmake curl libclang-dev libssl-dev libudev-dev pkg-config protobuf-compiler && \
    rm -rf /var/lib/apt/lists/*
# the toolchain comes from agave's rust-toolchain.toml
RUN curl -sSf https://sh.rustup.rs | sh -s -- -y --profile minimal --default-toolchain none
ENV PATH="/root/.cargo/bin:${PATH}"
WORKDIR /agave
RUN curl -sSfL "https://codeload.github.com/anza-xyz/agave/tar.gz/refs/tags/v${AGAVE_VERSION}" | tar -xz --strip-components=1
RUN --mount=type=cache,target=/root/.cargo/registry \
    --mount=type=cache,target=/agave/target \
    cargo build --release --bin solana-test-validator --bin solana && \
    mkdir /out && \
    cp target/release/solana-test-validator target/release/solana /out/

FROM debian:bookworm-slim AS source
COPY --from=source-build /out/ /usr/local/bin/
# fail during build if a shared library is missing
RUN solana-test-validator --version && solana --version

FROM debian:bookworm-slim AS release
COPY --from=release-download /out/ /usr/local/bin/
# fail during build if a shared library is missing
RUN solana-test-validator --version && solana --version
