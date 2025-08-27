{
  description = "CCIPv1.7 developer dependencies shell";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f {
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      devShells = forEachSystem ({ pkgs }: {
        default = pkgs.mkShell {
        shellHook = ''
          echo "Installing CCIPv17 CLI"
          cd cmd/ccip && go install -ldflags="-X main.Version=1.0.0" .
          cd -
          [ -f .envrc ] && source .envrc && echo "Loaded .envrc file"
          echo "📚 Run 'ccip' to check the CLI docs"
          echo "💻 Run 'ccip sh' to use an interactive mode"
          '';
          packages = [
          pkgs.air
          pkgs.postgresql_16
          pkgs.llvmPackages.libcxxClang
          pkgs.clang
          pkgs.minio-client
          pkgs.kubernetes-helm
          pkgs.kubectl
          pkgs.actionlint
          pkgs.golangci-lint
          pkgs.shellcheck
          pkgs.jq
          pkgs.awscli2
          pkgs.kubefwd
          pkgs.nodejs
          pkgs.go
          pkgs.just
          pkgs.kind
          pkgs.k9s
          ];
        };
      });
    };
}