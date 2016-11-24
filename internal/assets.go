package internal

import (
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/ooesili/dead-drop/internal/bindata/assets"
)

//go:generate go-bindata -nometadata -debug -o bindata/assets/bindata.go -pkg assets -prefix ../assets ../assets
var AssetFS = &assetfs.AssetFS{
	Asset:     assets.Asset,
	AssetDir:  assets.AssetDir,
	AssetInfo: assets.AssetInfo,
}
