{
  description = "Flake to build nix-style-search";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=25.11";
  };

  outputs = { self, nixpkgs }: let
    system = "x86_64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
    appName = "nix-style-search";
  in
  {
    packages.${system} = {
      ${appName} = pkgs.buildGoModule rec {
        pname = appName;
        version = "1.0.0";
        src = ./.;

       vendorHash = "sha256-NICwZgVyJ9q9Eg0b7vU0QJSkyAaK+BW0fPFGZXtDkbk=";

        meta = with pkgs.lib; {
          description = "A stylish (community-driven) search tool for Nix packages";
          license = licenses.asl20;
          maintainers = [ "snowflake-hd" ];
        };
      };

      default = self.packages.${system}.${appName};
    };
  };
}
