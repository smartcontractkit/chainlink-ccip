{ pkgs, keystone ? false }:
with pkgs;
let
  go = go_1_23;

  mkShell' = mkShell.override {
    # The current nix default SDK for macOS fails to compile Go projects, so we use a newer one for now.
    # stdenv = if stdenv.isDarwin then overrideSDK stdenv "11.0" else stdenv;
  };

  # Custom derivation to include only logcli from loki using runCommand
  logcli = pkgs.runCommand "logcli" { buildInputs = [ grafana-loki ]; } ''
    mkdir -p $out/bin
    cp ${grafana-loki}/bin/logcli $out/bin/
  '';

in
mkShell' {
  nativeBuildInputs = [
    go
    libiconv

    curl
    grpcurl

    # Tooling
    actionlint
    gotools
    gopls
    delve
    github-cli
    jq
    gomplate
    go-task
    kind

    # Include only logcli instead of the full loki package
    logcli

    # Deployment
    awscli2
    devspace
    kubectl
    kubernetes-helm
    mkcert

    # gofuzz
  ] ++ lib.optionals stdenv.isLinux [
    # Some dependencies needed for node-gyp on pnpm install
    pkg-config
    libudev-zero
    libusb1
  ];

  GOROOT = "${go}/share/go";
  shellHook = ''
    # Exit on unhandled errors
    set -e

    # Some useful custom aliases
    alias k=kubectl
    alias kgp="kubectl get pod"
    alias kge="kubectl get events --sort-by=.metadata.creationTimestamp"
    alias kexec="kubectl exec -it"

    # Used by build scripts in CHAINLINK_CODE_DIR to configure a crib build environment
    export IS_CRIB=true

    # Find the root of the git repository
    repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")
    export GOBIN=$(go env GOPATH)/bin
    export PATH=$PATH:$repo_root/scripts:$GOBIN

    # If the solana binary isn't found, install it
    if ! command -v solana &>/dev/null; then
      echo "Installing solana-cli..."
      sh -c "$(curl -sSfL https://release.anza.xyz/v2.2.0/install)"
    fi

    export PATH="$HOME/.local/share/solana/install/active_release/bin:$PATH"


    if [ "$CRIB_CI_ENV" = "true" ] && [ "$CLI_CHANGED" != "true" ]; then
      # in CI, download the CLI from the corresponding GH release if the CLI hasn't changed
      task fetch-cli
    else
      # when CLI has changed or when running in local, build the CLI from source
      task build
    fi

    # Set up crib CLI using task
    task setup

    # Check if the current directory matches either 'cre-dev' or 'cre-enterprise-sandbox'
    # under the 'deployments' directory and build the CRE cli
    if [[ "$PWD" == $repo_root/deployments/cre-dev || "$PWD" == $repo_root/deployments/cre-enterprise-sandbox ]]; then
      task cre-cli:clone
      task cre-cli:build
    fi

    # Sourcing the .env file as the last step
    if [ -f ".env" ]; then
      set -a
      source .env
      set +a
    fi

    # Prevent errors from exiting the shell from this point on
    set +e
  '';
}
