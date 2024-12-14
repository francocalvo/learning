{
  description = "Nix Flake for AoC 2024 with Go";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";

    # flake-utils
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true; # Propietary software
        };
      in {

        devShells.default = pkgs.mkShell {
          nativeBuildInputs = (with pkgs; [ nixpkgs-fmt lazygit go git ])
            ++ (with pkgs.nodePackages; [ markdownlint-cli ]);
        };
      });
}

