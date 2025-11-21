{
  description = "A presence setting application for multiple platforms";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "wms";
          version = "0.3.2";
          src = ./.;

          subPackages = [ "wms" ];

          vendorHash = "sha256-U2WpMg9gvSJvSQHhuvS4ixHmTwuANK0/m+ReTjHJL2o=";

          nativeBuildInputs = [ ];
          buildInputs = [ ];
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            go-task
            dbus
            docker
            gcc
            pkg-config
            libcap
          ];
        };
      }
    );
}
