# caddy-hax

[![built in go](https://img.shields.io/badge/built%20in-go-%2301ADD8)](https://go.dev/)

`caddy-hax` is a caddy v2 plugin that contains various useful 'hacks'.

# Installation

This repo uses [nix](https://nixos.org/download.html) + [direnv](https://direnv.net/) to easily and automatically install dependencies and run caddy with this plugin enabled in an easy way. Once both nix and direnv are installed, run `direnv allow` in the root of the project to install all the required dependencies.

# Building

Use [xcaddy](https://github.com/caddyserver/xcaddy) to build, or use nix!

## xcaddy example:

```bash
xcaddy build --output ./caddy --with github.com/jpetrucciani/caddy-hax@main
```

## nix example:

caddy with caddy-hax already included:

```nix
TODO
```

build your own!

```nix
TODO
```

# How to run

There are two ways to run the project.

1. The `run` command which will rebuild the go caddy plugin when files are changed as well as run the `run-hax` command.
1. The `run-hax` command which will run Caddy in watch mode on the Caddyfile in the conf directory.

## Current Hax

The local server runs on `localhost:6420`. Some of the hacks can be run in isolation using different routes.
