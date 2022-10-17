package hax

import (
    "bytes"
	"net/http"
	"archive/tar"
    "compress/gzip"

	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func Tarball(b Hax, res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Encoding", "gzip")
	res.Header().Set("Content-Type", "application/tar+gzip")
	var buf bytes.Buffer
	
	gzw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gzw)
	
	header := new(tar.Header)
	header.Name = b.TarballFileName
	header.Size = int64(len(b.TarballFileText))
	header.Mode = int64(0755)
	tw.WriteHeader(header)
	tw.Write([]byte(b.TarballFileText))

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
