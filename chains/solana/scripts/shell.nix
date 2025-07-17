{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [

    # Rust
    pkgs.rustc
    pkgs.cargo

    # Networking & Node
    pkgs.curl
    pkgs.nodejs
    # pkgs.pnpm # Uncomment if needed

    # Common build dependencies
    pkgs.openssl
    pkgs.pkg-config
    pkgs.zlib

    # Misc tools
    pkgs.git
    pkgs.jq
  ];

  shellHook = ''
    echo "🔧 Dev environment ready"
    cargo install solana-verify
    cargo install --git https://github.com/coral-xyz/anchor anchor-cli --locked
    export PATH="$HOME/.cargo/bin:$PATH"
  '';
}
