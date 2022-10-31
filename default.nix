{ jacobi ? import
    (fetchTarball {
      name = "jpetrucciani-2022-10-31";
      url = "https://nix.cobi.dev/x/02c19fb4ae64983ec48fd8c536c178ace7270549";
      sha256 = "0gn78c9kvwy6dismcwix829ch98dvhxbci3rgpv0kdkjpax9n290";
    })
    { }
}:
let
  name = "caddy-hax";
  tools = with jacobi;
    let
      run-hax = pog {
        name = "run-hax";
        description = "run caddy with the hax plugin in watch mode against the caddyfile in the conf dir";
        script = h: with h; ''
          ${xcaddy}/bin/xcaddy run --config ./conf/Caddyfile --watch "$@"
        '';
      };
      run = pog {
        name = "run";
        description = "run run-hax, restarting when go files are changed";
        script = h: with h; ''
          ${findutils}/bin/find . -iname '*.go' | ${entr}/bin/entr -rz ${run-hax}/bin/run-hax
        '';
      };
    in
    {
      cli = [
        nixpkgs-fmt
      ];
      go = [
        go_1_19
        go-tools
        gopls
      ];
      scripts = [
        xcaddy
        run-hax
        run
        (writeShellScriptBin "test_actions" ''
          export DOCKER_HOST=$(${jacobi.docker-client}/bin/docker context inspect --format '{{.Endpoints.docker.Host}}')
          ${jacobi.act}/bin/act --container-architecture linux/amd64 -r --rm
        '')
      ];
    };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
