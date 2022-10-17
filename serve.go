package hax

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"net/http"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func Tarball(b Hax, res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "application/tar+gzip")
	var buf bytes.Buffer

	gzw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gzw)

	replacer := caddyhttp.NewTestReplacer(req)
	text := replacer.ReplaceAll(b.TarballFileText, "")

	header := new(tar.Header)
	header.Name = b.TarballFileName
	header.Size = int64(len(text))
	header.Mode = int64(0755)
	tw.WriteHeader(header)
	tw.Write([]byte(text))

	tw.Close()
	gzw.Close()
	res.Write(buf.Bytes())
}

func (b Hax) ServeHTTP(res http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	b.log.Info(
		"HAX",
		zap.String("ip", req.RemoteAddr),
		zap.String("url", req.URL.String()),
		zap.String("user-agent", req.Header.Get("User-Agent")),
	)

	if b.EnableTarball {
		Tarball(b, res, req)
	}

	return next.ServeHTTP(res, req)
}
