{
  description = "aj - a shell tool that automatically fixes mistakes in your last command";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    {
      overlays.default = final: prev: {
        aj = final.callPackage ./package.nix { };
      };
    }
    // flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = rec {
          aj = pkgs.callPackage ./package.nix { };
          default = aj;
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            wgo
            gopls
            delve
            gotools
            golangci-lint
          ];
        };
      }
    );
}
