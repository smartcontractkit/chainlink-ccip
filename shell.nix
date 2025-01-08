{ stdenv, pkgs, lib }:
let
  version = "v2.0.18";
  getBinDerivation =
    {
      name,
      filename,
      sha256,
    }:
    pkgs.stdenv.mkDerivation rec {
      inherit name;
      url = "https://github.com/anza-xyz/agave/releases/download/${version}/${filename}";
      src = pkgs.fetchzip {
        inherit url sha256;
      };

      installPhase = ''
        mkdir -p $out/bin
        ls -lah $src
        cp -r $src/bin/* $out/bin
      '';
    };

  # It provides two derivations, one for x86_64-linux and another for aarch64-apple-darwin.
  # Each derivation downloads the corresponding Solana release.

  # The SHA256 hashes below are automatically updated by action.(dependency-updates.yml)
  # The update script(./scripts/update-solana-nix-hashes.sh) looks for the BEGIN and END markers to locate the lines to modify.
  # Do not modify these markers or the lines between them manually.
  solanaBinaries = {
    x86_64-linux = getBinDerivation {
      name = "solana-cli-x86_64-linux";
      filename = "solana-release-x86_64-unknown-linux-gnu.tar.bz2";
      ### BEGIN_LINUX_SHA256 ###
      sha256 = "sha256-3FW6IMZeDtyU4GTsRIwT9BFLNzLPEuP+oiQdur7P13s=";
      ### END_LINUX_SHA256 ###
    };
    aarch64-apple-darwin = getBinDerivation {
      name = "solana-cli-aarch64-apple-darwin";
      filename = "solana-release-aarch64-apple-darwin.tar.bz2";
      ### BEGIN_DARWIN_SHA256 ###
      sha256 = "sha256-6VjycYU0NU0evXoqtGAZMYGHQEKijofnFQnBJNVsb6Q=";
      ### END_DARWIN_SHA256 ###
    };
  };
in
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    (rust-bin.stable.latest.default.override { extensions = ["rust-src"]; })
    # lld_11
    llvm_12
    stdenv.cc.cc.lib
    pkg-config
    openssl

    # Solana
    # solana.solana-full
    # spl-token-cli
    # anchor

    # Golang
    # Keep this golang version in sync with the version in .tool-versions please
    go_1_23
    gopls
    delve
    golangci-lint
    gotools
    kubernetes-helm

    # NodeJS + TS
    nodePackages.typescript
    nodePackages.typescript-language-server
    nodePackages.npm
    nodePackages.pnpm
    # Keep this nodejs version in sync with the version in .tool-versions please
    nodejs-18_x
    (yarn.override { nodejs = nodejs-18_x; })
    python3
    # ] ++ lib.optionals stdenv.isLinux [
    #   # ledger specific packages
    #   libudev-zero
    #   libusb1
    # ];
     ] ++ pkgs.lib.optionals pkgs.stdenv.isLinux [
    libudev-zero
    libusb1
    solanaBinaries.x86_64-linux
  ] ++ pkgs.lib.optionals (pkgs.stdenv.isDarwin && pkgs.stdenv.hostPlatform.isAarch64) [
    solanaBinaries.aarch64-apple-darwin
  ];
  RUST_BACKTRACE = "1";

  LD_LIBRARY_PATH = lib.makeLibraryPath [pkgs.zlib stdenv.cc.cc.lib]; # lib64

  # Avoids issues with delve
  CGO_CPPFLAGS="-U_FORTIFY_SOURCE -D_FORTIFY_SOURCE=0";

  # shellHook = ''
  #   # install gotestloghelper
  #   go install github.com/smartcontractkit/chainlink-testing-framework/tools/gotestloghelper@latest
  # '';
  shellHook = ''
    echo "===================================================="
    echo "Welcome to the combined development shell."
    echo "Current environment: $(uname -a)"
    echo "You are using the package for ${pkgs.stdenv.hostPlatform.system}."
    echo "----------------------------------------------------"
    echo "Solana CLI information:"
    solana --version || echo "Solana CLI not available."
    solana config get || echo "Solana config not available."
    echo "----------------------------------------------------"

    # Install gotestloghelper
    go install github.com/smartcontractkit/chainlink-testing-framework/tools/gotestloghelper@latest
  '';
}
