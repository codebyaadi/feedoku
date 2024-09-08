# To learn more about how to use Nix to configure your environment
# see: https://developers.google.com/idx/guides/customize-idx-env
{ pkgs, ... }: {
  # Which nixpkgs channel to use.
  channel = "stable-23.11"; # or "unstable"
  # Use https://search.nixos.org/packages to find packages
  packages = [
    # pkgs.go
    pkgs.air
    pkgs.go
    pkgs.sqlc
    pkgs.goose
    pkgs.redis
  ];
  # Sets environment variables in the workspace
  env = {};
  services.docker.enable = true;
  idx = {
    # Search for the extensions you want on https://open-vsx.org/ and use "publisher.id"
    extensions = [
      "golang.go"
      "rangav.vscode-thunder-client"
    ];
    workspace = {
      # Runs when a workspace is (re)started
      onStart= {
        run-redis-server = "redis-server";
        # run-server = "air";
      };
      # To run something each time the workspace is first created, use the `onStart` hook
    };
  };
}
