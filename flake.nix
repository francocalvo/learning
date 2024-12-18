{
  description = "Basic flake for learning repository";

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
        # Create empty list if system is aarch64-darwin, else use mit-scheme
        pkg_mitScheme =
          if system == "aarch64-darwin" then [ ] else [ pkgs.mit-scheme ];
      in {

        devShells.default = pkgs.mkShell {
          nativeBuildInputs = (with pkgs; [
            nixpkgs-fmt
            lazygit
            git
            nodePackages.markdownlint-cli
            python312Packages.mkdocs-material
          ]);
        };
      });
}
