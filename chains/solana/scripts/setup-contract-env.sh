#!/usr/bin/env bash
set -euo pipefail

ANCHOR_VERSION="0.29.0"
SOLANA_VERSION="1.17.25"
ANCHOR_GO_VERSION="v0.2.3"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info()  { echo -e "${GREEN}[INFO]${NC} $*"; }
warn()  { echo -e "${YELLOW}[WARN]${NC} $*"; }
error() { echo -e "${RED}[ERROR]${NC} $*"; }

check_command() {
    if ! command -v "$1" &>/dev/null; then
        error "$1 is not installed."
        echo "  Install it from: $2"
        return 1
    fi
}

# ─── Prerequisites ──────────────────────────────────────────────────────────────

info "Checking prerequisites..."

check_command rustc "https://www.rust-lang.org/tools/install" || exit 1
check_command cargo "https://www.rust-lang.org/tools/install" || exit 1
check_command go    "https://go.dev/doc/install"              || exit 1

# ─── AVM / Anchor ───────────────────────────────────────────────────────────────

if ! command -v avm &>/dev/null; then
    info "Installing AVM (Anchor Version Manager)..."
    cargo install --git https://github.com/coral-xyz/anchor avm --force
fi

CURRENT_ANCHOR=$(anchor --version 2>/dev/null | awk '{print $2}' || echo "none")
if [[ "$CURRENT_ANCHOR" != "$ANCHOR_VERSION" ]]; then
    info "Switching Anchor to ${ANCHOR_VERSION} (current: ${CURRENT_ANCHOR})..."
    avm install "$ANCHOR_VERSION" 2>/dev/null || true
    avm use "$ANCHOR_VERSION"
else
    info "Anchor ${ANCHOR_VERSION} already active."
fi

# ─── Solana CLI ──────────────────────────────────────────────────────────────────

CURRENT_SOLANA=$(solana --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' || echo "none")
if [[ "$CURRENT_SOLANA" != "$SOLANA_VERSION" ]]; then
    info "Installing Solana CLI ${SOLANA_VERSION} (current: ${CURRENT_SOLANA})..."
    sh -c "$(curl -sSfL https://release.anza.xyz/v${SOLANA_VERSION}/install)"

    # Ensure solana is on PATH for the rest of this script
    export PATH="$HOME/.local/share/solana/install/active_release/bin:$PATH"
else
    info "Solana CLI ${SOLANA_VERSION} already installed."
fi

# ─── anchor-go ───────────────────────────────────────────────────────────────────

info "Installing anchor-go ${ANCHOR_GO_VERSION}..."
# anchor-go v0.2.3 has an old golang.org/x/net dep that fails to link with Go 1.22+.
# Build from source with updated deps as a workaround.
if anchor-go -src /dev/null 2>&1 | grep -q "src"; then
    info "anchor-go (v0.2.3-compatible) already installed."
else
    ANCHOR_GO_TMP="$(mktemp -d)"
    trap "rm -rf '$ANCHOR_GO_TMP'" EXIT
    git clone --depth 1 --branch "${ANCHOR_GO_VERSION}" \
        https://github.com/gagliardetto/anchor-go.git "$ANCHOR_GO_TMP" 2>/dev/null
    pushd "$ANCHOR_GO_TMP" >/dev/null
    go get golang.org/x/net@latest 2>/dev/null
    go mod tidy 2>/dev/null
    go install .
    popd >/dev/null
    rm -rf "$ANCHOR_GO_TMP"
    info "anchor-go installed from source with patched dependencies."
fi

# ─── Verify ─────────────────────────────────────────────────────────────────────

echo ""
info "Environment setup complete. Versions:"
echo "  anchor : $(anchor --version 2>/dev/null)"
echo "  solana : $(solana --version 2>/dev/null)"
echo "  rustc  : $(rustc --version 2>/dev/null)"
echo "  go     : $(go version 2>/dev/null)"
echo ""
info "You can now build contracts with:"
echo "  cd chains/solana/contracts && anchor build"
