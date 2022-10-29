{ jacobi ? import
    (fetchTarball {
      name = "jpetrucciani-2022-10-29";
      url = "https://nix.cobi.dev/x/226c57d8dceeb0556b5405ccc674f24e2c97307b";
      sha256 = "068i7zpvgk9lydbhksyqr58vildjwmp01rhvd2r9fh61sbplpbmj";
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
