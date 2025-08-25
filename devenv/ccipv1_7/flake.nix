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
              # install the CLI
              echo "Building CCIPv17 CLI"
              go install cmd/ccip.go
              # Load .envrc file if it exists
              [ -f .envrc ] && source .envrc && echo "Loaded .envrc file" || echo "No .envrc file found, skipping"
              # Set MinIO alias if minio-client is available and alias doesn't exist
              if command -v mc >/dev/null 2>&1; then
                if ! mc alias list | grep -q "minio"; then
                  echo "Setting up MinIO client alias..."
                  mc alias set minio http://localhost:9000 mYrAnD0mAcc3ssK3y2024 s3cR3tK3yL0ngEn0ughF0rMinI0Testing2024!
                  echo "MinIO alias 'minio' configured successfully"
                fi
              fi
            '';
          packages = [
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