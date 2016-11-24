// Code generated by go-bindata.
// sources:
// ../templates/complete.html.tmpl
// ../templates/index.html.tmpl
// ../templates/layout.html.tmpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _completeHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x4d\x0a\xc3\x20\x10\x46\xf7\x9e\x62\x98\x6d\x69\x25\x7b\xe3\x5d\xac\x4e\x8d\xc5\x9f\xd0\x19\x24\x45\x72\xf7\x92\x64\xd3\xdd\xc7\x07\xef\xf1\xc6\x80\x40\xaf\x54\x09\xd0\xb7\x2a\x54\x05\x61\xdf\x95\x09\xa9\x83\xcf\x8e\x79\x3e\x7f\x97\x2a\x7d\xd0\x2a\x00\xd3\x9e\x6f\xf2\x02\xc1\x89\x9b\x51\x3b\x66\x12\xd6\xb9\xc5\xf6\xe0\x1e\x11\xe4\xbb\xd2\x8c\xa9\xb8\x48\x9a\x7b\xbc\x6d\x25\xa3\x35\xfa\xa2\x4e\xc1\x32\xd9\x7f\xfd\x81\xde\x85\x36\x41\xeb\x5b\x59\x33\x09\x19\x1d\x52\xb7\x46\x2f\x93\x55\xd7\x56\x63\x00\xd5\x70\xa4\xfd\x02\x00\x00\xff\xff\xfc\x83\x48\xc5\xb2\x00\x00\x00")

func completeHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_completeHtmlTmpl,
		"complete.html.tmpl",
	)
}

func completeHtmlTmpl() (*asset, error) {
	bytes, err := completeHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "complete.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x4d\x6f\xa4\x30\x0c\xbd\xf3\x2b\x2c\x5f\x57\x33\xd1\xdc\x81\xe3\x5e\x77\xa5\xf6\x0f\x78\x12\xc3\x64\x1a\x12\x94\x18\x34\x15\xe2\xbf\x57\x21\x41\xed\xa5\x27\xfc\xf1\xde\xf3\xc3\xce\xb6\x81\xe1\xc1\x7a\x06\xd4\xc1\x0b\x7b\x41\xd8\xf7\xa6\x35\x76\x05\xed\x28\xa5\xee\xa8\x93\xf5\x1c\xb1\x6f\x00\xda\x70\x7f\xb2\x16\x30\x24\xd4\xa1\xa2\x94\x58\x92\x72\x61\x0c\xd7\xb4\x8e\x08\xf2\x39\x73\x87\x76\xa2\x91\x55\x5a\xc7\x3f\xaf\xc9\x61\xdf\xaa\xc2\x3a\x04\x1e\xb7\xfe\xa7\x7c\xa6\x5e\x84\x5f\x82\xbd\x61\x32\x60\x62\x98\x5b\x65\xec\xda\xb7\xea\x71\x3b\x18\x43\x88\x13\x90\x16\x1b\x7c\x87\x2a\x03\x10\x26\x96\x47\x30\x1d\xfe\xff\xf7\xf6\x8e\xc0\x5e\x97\xc1\xd3\xe2\xc4\xce\x14\x45\x65\xd2\x25\xbb\x3c\x6c\x03\xb4\x8e\xee\xec\x60\x08\xb1\xc3\xc1\x3a\xc6\xd3\xc0\x7d\x11\x09\x1e\xca\xe7\x32\x5b\xfd\x51\x19\x00\x39\x01\x82\x0c\x2f\x1a\xea\x10\xa9\x82\xd6\xcf\x8b\x80\x35\xa7\x5e\x71\x50\x62\x4f\xd3\x19\x57\x74\x1d\x93\xe1\xcb\xec\x02\x99\x5f\x0c\xd4\xe6\x69\xa1\xa4\x75\x7a\x81\x94\xd6\xb6\x5d\x75\x8a\xc3\x5f\xcb\xce\xec\x7b\x5e\xd3\xf1\xcb\x7d\x53\x96\xd7\x34\x6d\xd2\xd1\xce\x52\x6d\xe5\x0d\xab\x27\xad\x54\xaa\x08\x29\xea\xef\xfb\x4d\x64\xfd\xf5\x99\xf2\xa5\x4a\xbf\x6f\xb6\x0d\xd8\x9b\xfc\x18\xbe\x02\x00\x00\xff\xff\xb0\x1d\x06\xb0\x24\x02\x00\x00")

func indexHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_indexHtmlTmpl,
		"index.html.tmpl",
	)
}

func indexHtmlTmpl() (*asset, error) {
	bytes, err := indexHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x90\xbd\x4e\xc5\x30\x0c\x85\xf7\x3e\x85\xf1\xce\xcd\xca\x90\x74\xe1\x67\x85\xe1\x32\x30\x86\xe6\x40\xae\x48\xd3\x2a\x3e\x03\xd5\xd5\x7d\x77\x54\x02\x3c\x00\x93\x2d\x7f\xc7\x9f\x64\xfb\xab\xbb\xc7\xdb\xe3\xcb\xd3\xbd\x64\xce\x65\x1c\xfc\x5e\xa4\xc4\xfa\x1e\x14\x55\xc7\x41\xc4\x67\xc4\xb4\x37\x22\x7e\x06\xa3\x4c\x39\x36\x03\x83\x3e\x1f\x1f\xae\x6f\xf4\x07\xf1\xc4\x82\x31\x21\x26\x49\x6d\x59\xbd\xeb\x83\xa1\xd3\x72\xaa\x1f\xd2\x50\x82\x1a\xb7\x02\xcb\x00\x55\xb8\xad\x08\x4a\x7c\xd2\x4d\x66\x2a\xb9\xe1\x2d\xa8\x8b\x66\xa0\xb9\x06\x03\x0f\x3b\x18\xff\x2f\xe9\xc9\x3f\x8b\x77\xbf\xd7\xf8\xd7\x25\x6d\x5d\x7c\x3e\x0b\x31\xaf\x25\x12\xa2\xd3\x52\x89\x4a\x95\x83\x5c\x2e\xdf\x1b\x3d\xe8\x5d\xff\xd0\x57\x00\x00\x00\xff\xff\xc6\xfe\x52\xab\x32\x01\x00\x00")

func layoutHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutHtmlTmpl,
		"layout.html.tmpl",
	)
}

func layoutHtmlTmpl() (*asset, error) {
	bytes, err := layoutHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"complete.html.tmpl": completeHtmlTmpl,
	"index.html.tmpl": indexHtmlTmpl,
	"layout.html.tmpl": layoutHtmlTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"complete.html.tmpl": &bintree{completeHtmlTmpl, map[string]*bintree{}},
	"index.html.tmpl": &bintree{indexHtmlTmpl, map[string]*bintree{}},
	"layout.html.tmpl": &bintree{layoutHtmlTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

