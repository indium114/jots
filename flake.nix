{
  description = "jots devshell and package";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "jots-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
            just
          ];
        };

        packages.jots = pkgs.buildGoModule {
          pname = "jots";
          version = "2026.03.08-a";

          src = self;

          vendorHash = "sha256-Mmrdb/M0gigRFXV3cn1Or8b2MI9PD0DWyVXtf/twrWw=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A minimal CLI journaling tool";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.jots = {
          type = "app";
          program = "${self.packages.${system}.jots}/bin/jots";
        };
      });
}
