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
    goreleaser

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

    # Find the root of the git repository
    repo_root=$(git rev-parse --show-toplevel 2>/dev/null || echo ".")
    export PATH=$PATH:$repo_root/scripts

    # install changesets (no nix package available atm)
    pnpm install

    # build crib CLI and make it available in PATH
    echo -n "Building crib CLI... "
    (cd $repo_root/cli && go build -o ./dist/crib .) && echo "Done." || { echo "Failed to build crib CLI. Please post this error message to #project-crib." >&2; exit 1; }
    export PATH=$PATH:$repo_root/cli/dist

    # crib init will make sure everything else is set up prior to running any devspace commands
    crib init --write-config || exit $?

    # sourcing the .env file as the last step
    [ -f ".env" ] && export $(cat .env | grep -v ^# | xargs)
  '';
}
