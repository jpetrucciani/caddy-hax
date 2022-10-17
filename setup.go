package hax

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

var (
	_ caddy.Provisioner           = (*Hax)(nil)
	_ caddy.Validator             = (*Hax)(nil)
	_ caddyfile.Unmarshaler       = (*Hax)(nil)
	_ caddyhttp.MiddlewareHandler = (*Hax)(nil)
)

func init() {
	caddy.RegisterModule(Hax{})
	httpcaddyfile.RegisterHandlerDirective("hax", parseCaddyfile)
}

type Hax struct {
	// options
	EnableTarball   bool   `json:"enable_tarball"`
	TarballFileName string `json:"tarball_file_name"`
	TarballFileText string `json:"tarball_file_text"`

	// globals
	log *zap.Logger
}

func (Hax) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.hax",
		New: func() caddy.Module { return new(Hax) },
	}
}

func (b *Hax) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.NextArg() // skip block beginning: "hax"

	for d.NextBlock(0) {
		var err error
		switch d.Val() {
		case "enable_tarball":
			b.EnableTarball = true
		case "tarball_file_name":
			err = parseStringArg(d, &b.TarballFileName)
		case "tarball_file_text":
			err = parseStringArg(d, &b.TarballFileText)
		default:
			err = d.Errf("not a valid hax option")
		}
		if err != nil {
			return d.Errf("Error parsing %s: %s", d.Val(), err)
		}
	}

	return nil
}

func (b *Hax) Provision(ctx caddy.Context) (err error) {
	b.log = ctx.Logger(b)

	// parse env vars
	replacer := caddy.NewReplacer()
	b.TarballFileName = replacer.ReplaceAll(b.TarballFileName, "")
	b.log.Info(
		"BOOT",
	)
	return nil
}

func (b *Hax) Validate() error {
	return nil
}
