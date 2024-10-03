{
  description = "CRIB development shell";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    nur.url = "github:nix-community/NUR";
    goreleaser-nur.url = "github:goreleaser/nur";
  };

  outputs = inputs@{ self, nixpkgs, flake-utils, goreleaser-nur, nur, ... }:

    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; 
          config = { allowUnfree = true; }; 
          overlays = [
            (final: prev: {
              nur = import nur
                {
                  pkgs = prev;
                  repoOverrides = {
                    goreleaser = import goreleaser-nur { pkgs = prev; };
                  };
                };
            })
          ];
         };
      in {
        devShells.default = import ./shell.nix { inherit pkgs; };
        formatter = pkgs.nixpkgs-fmt;
      });
}
