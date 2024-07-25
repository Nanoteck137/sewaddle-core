{
  description = "Core library for Sewaddle, and tools";

  inputs = {
    nixpkgs.url      = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url  = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        overlays = [];
        pkgs = import nixpkgs {
          inherit system overlays;
        };

        # program = pkgs.buildGoModule {
        #   pname = "";
        #   version = self.shortRev or "dirty";
        #   src = ./.;
        #   vendorHash = "";
        # };
      in
      {
        # packages.default = program;
        # packages. = program;

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
          ];
        };
      }
    );
}
