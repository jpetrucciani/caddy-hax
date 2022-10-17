{ jacobi ? import
    (
      fetchTarball {
        name = "jpetrucciani-2022-10-11";
        url = "https://github.com/jpetrucciani/nix/archive/4d57dd53857d172cea6f52dce7c5bec425db4550.tar.gz";
        sha256 = "0zlg2inn2nizsfazxwcbrmjkrmryf924pk5r568cln3rwkg6081l";
      }
    )
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
      ];
    };

  env = jacobi.enviro {
    inherit name tools;
  };
in
env
