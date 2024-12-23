{
  description = "Ah... To be naive about Nix complexity again...";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

    # For simpler multi-system outputs
    flake-utils.url = "github:numtide/flake-utils";

    pyproject-nix = {
      url = "github:pyproject-nix/pyproject.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    uv2nix = {
      url = "github:pyproject-nix/uv2nix";
      inputs.pyproject-nix.follows = "pyproject-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    pyproject-build-systems = {
      url = "github:pyproject-nix/build-system-pkgs";
      inputs.pyproject-nix.follows = "pyproject-nix";
      inputs.uv2nix.follows = "uv2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils, uv2nix, pyproject-nix
    , pyproject-build-systems, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        # Import pkgs for this system
        pkgs = import nixpkgs { inherit system; };

        # Load a uv2nix workspace from the current directory
        workspace = uv2nix.lib.workspace.loadWorkspace { workspaceRoot = ./.; };

        # Generate an overlay from uv.lock
        overlay = workspace.mkPyprojectOverlay {
          # Prefer binary wheels as a package source
          sourcePreference = "wheel";
        };

        # Additional overlays or build fixups
        pyprojectOverrides = _final: _prev:
          {
            # Put your override definitions here...
          };

        # Example picking a specific Python
        python = pkgs.python312;

        # Create your Python package set using pyproject.nix
        pythonSet = (pkgs.callPackage pyproject-nix.build.packages {
          inherit python;
        }).overrideScope (pkgs.lib.composeManyExtensions [
          pyproject-build-systems.overlays.default
          overlay
          pyprojectOverrides
        ]);
      in {
        # 1) Our default package: a virtual environment with no optional deps
        packages.default =
          pythonSet.mkVirtualEnv "leetcode" workspace.deps.default;

        # 2) Devshell: 'impure' approach, just uses uv & python, doesn't rely on uv2nix
        devShells.default = pkgs.mkShell {
          packages = [ python pkgs.uv pkgs.go pkgs.git ];
          shellHook = ''
            unset PYTHONPATH
            export UV_PYTHON_DOWNLOADS=never
          '';
        };

        # 3) Devshell: 'uv2nix' approach that uses purely Nix-based local editable packages
        devShells.uv2nix = let
          # Make local packages editable with an overlay
          editableOverlay = workspace.mkEditablePyprojectOverlay {
            root = "$REPO_ROOT";
            # members = [ "leetcode" ]; # Optionally only make these editable
          };

          # Override pythonSet to apply "editable" overlay
          editablePythonSet = pythonSet.overrideScope editableOverlay;

          # Build an editable-mode venv with all optional deps
          virtualenv = editablePythonSet.mkVirtualEnv "leetcode-dev-env"
            workspace.deps.all;
        in pkgs.mkShell {
          packages = [ virtualenv pkgs.uv pkgs.go pkgs.git ];
          shellHook = ''
            unset PYTHONPATH
            export UV_NO_SYNC=1
            export UV_PYTHON_DOWNLOADS=never
            export REPO_ROOT=$(git rev-parse --show-toplevel)
          '';
        };
      });
}
