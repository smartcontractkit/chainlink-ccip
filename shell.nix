{ pkgs }:
with pkgs;
let
  go = go_1_22;
  nodejs = nodejs-18_x;
  nodePackages = pkgs.nodePackages.override { inherit nodejs; };

  mkShell' = mkShell.override {
    # The current nix default sdk for macOS fails to compile go projects, so we use a newer one for now.
    stdenv = if stdenv.isDarwin then overrideSDK stdenv "11.0" else stdenv;
  };
in
mkShell' {
  nativeBuildInputs = [
    go

    curl
    nodejs
    nodePackages.pnpm

    go-mockery

    # tooling
    gotools
    gopls
    delve
    golangci-lint
    github-cli
    jq
    gomplate

    kind

    # deployment
    awscli2
    devspace
    kubectl
    kubernetes-helm

    # gofuzz
  ] ++ lib.optionals stdenv.isLinux [
    # some dependencies needed for node-gyp on pnpm install
    pkg-config
    libudev-zero
    libusb1
  ];
  GOROOT = "${go}/share/go";

  shellHook = ''
    # Some useful custom aliases
    alias k=kubectl
    alias kgp="kubectl get pod"
    alias kge="kubectl get events --sort-by=.metadata.creationTimestamp"
    alias kexec="kubectl exec -it"

    # Used by build scripts in CHAINLINK_CODE_DIR to configure a crib build environment
    export IS_CRIB=true

    # Find the root of the git repository
    repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")
    export PATH=$PATH:$repo_root/scripts

    # install changesets (no nix package available atm)
    pnpm install

    # build crib-cli and make it available in PATH
    echo -n "Building crib-cli... "
    (cd $repo_root/cli && go build -o ./dist/crib-cli .) && echo "Done." || echo "Failed to build crib-cli. Please post this error message to #project-crib." >&2
    export PATH=$PATH:$repo_root/cli/dist
  '';
}
