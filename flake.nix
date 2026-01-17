{
  description = "Flake to build nix-style-search";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=25.11";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        appName = "nix-style-search";
      in
      {
        packages = {
          ${appName} = pkgs.buildGoModule rec {
            pname = appName;
            version = "1.1.2";
            src = ./.;

            subPackages = [ "cmd/nix-style-search" ];

            vendorHash = "sha256-NICwZgVyJ9q9Eg0b7vU0QJSkyAaK+BW0fPFGZXtDkbk=";

            meta = with pkgs.lib; {
              description = "A stylish (community-driven) search tool for Nix packages";
              license = licenses.asl20;
              maintainers = [ "snowflake-hd" ];
            };
          };

          default = self.packages.${system}.${appName};
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.${appName}}/bin/${appName}";
        };

        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.golangci-lint
            pkgs.git
          ];
        };
      }
    );
}