{ pkgs, keystone ? false }:
with pkgs;
let
  go = go_1_22;
  nodejs = nodejs-18_x;
  nodePackages = pkgs.nodePackages.override { inherit nodejs; };

  mkShell' = mkShell.override {
    # The current nix default SDK for macOS fails to compile Go projects, so we use a newer one for now.
    stdenv = if stdenv.isDarwin then overrideSDK stdenv "11.0" else stdenv;
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

    curl
    nodejs
    nodePackages.pnpm

    go-mockery

    # Tooling
    actionlint
    gotools
    gopls
    delve
    golangci-lint
    github-cli
    jq
    gomplate
    go-task
    yamllint
    shfmt
    shellcheck
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
    export PATH=$PATH:$repo_root/scripts:$GOBIN

    ${lib.optionalString (!keystone) ''
      # Install changesets (no nix package available at the moment)
      pnpm install

      if [ "$CRIB_CI_ENV" = "true" ]; then
        # in CI, download the CLI from the corresponding GH release
        task fetch-cli
      else
        # in local, build the CLI from source
        task build
      fi

      # Set up crib CLI using task
      task setup

      # Sourcing the .env file as the last step
      if [ -f ".env" ]; then
        export $(grep -v '^#' .env | xargs)
      fi
    ''}

    ${lib.optionalString keystone ''
      echo "Welcome to crib, build the crib CLI via \"task build\". It'll be available as \"crib\" in either your \$(go env GOBIN)' or \$(go env GOPATH)/bin directory, depending on if your GOBIN is set."
    ''}
  '';
}
