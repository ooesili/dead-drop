package internal

import (
	"io"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/ooesili/dead-drop/internal/bindata/assets"
	"github.com/ooesili/dead-drop/internal/config"
)

//go:generate go-bindata -nometadata -debug -o bindata/assets/bindata.go -pkg assets -prefix ../assets ../assets
var AssetFS = &assetfs.AssetFS{
	Asset:     assets.Asset,
	AssetDir:  assets.AssetDir,
	AssetInfo: assets.AssetInfo,
}

var ConfigSchema = config.Schema{
	"PORT":      config.Optional("3000"),
	"BIND_ADDR": config.Optional("0.0.0.0"),
}

type Dropper interface {
	Drop(filename string, file io.Reader)
}

type Renderer interface {
	Render(template string, out io.Writer, request *http.Request) error
}
