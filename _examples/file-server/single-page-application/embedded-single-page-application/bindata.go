// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package main generated by go-bindata.// sources:
// public/app.js
// public/app2/index.html
// public/css/main.css
// public/index.html
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _publicAppJs = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xcf\xcc\x4b\xc9\x2f\xd7\x4b\xcc\x49\x2d\x2a\xd1\x50\x4a\x2c\x28\xd0\xcb\x2a\x56\xc8\xc9\x4f\x4c\x49\x4d\x51\x48\x2b\xca\xcf\x55\x88\x51\xd2\x57\xd2\xb4\x06\x04\x00\x00\xff\xff\xa9\x06\xf7\xa3\x27\x00\x00\x00"

func publicAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicAppJs,
		"public/app.js",
	)
}

func publicAppJs() (*asset, error) {
	bytes, err := publicAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app.js", size: 39, mode: os.FileMode(438), modTime: time.Unix(1595516291, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicApp2IndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x41\x0a\xc2\x40\x0c\x45\xf7\x81\xdc\xe1\x9f\xc0\xd0\xae\x43\xc0\x9d\xd7\xa8\x4c\x24\x85\xd4\x09\x32\x0b\xbd\xbd\xb4\xd6\xe5\x87\xf7\xde\xd7\x18\x5b\x1a\x13\x93\x86\x2f\xcd\x98\x00\x40\xc7\x3a\xd2\xed\x5a\x85\x59\xe5\x37\x98\x54\x4e\x84\x49\xef\xbd\x7d\xfe\x70\x4c\x86\x9b\x67\x76\x3c\x5e\x7d\xc3\x52\x35\xcb\xfa\x6c\xfe\xbe\xec\x71\xa8\xc4\x74\xd8\xa7\x73\x84\xf6\xd7\x6f\x00\x00\x00\xff\xff\xfd\x28\x92\x95\x7c\x00\x00\x00"

func publicApp2IndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicApp2IndexHtml,
		"public/app2/index.html",
	)
}

func publicApp2IndexHtml() (*asset, error) {
	bytes, err := publicApp2IndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/app2/index.html", size: 124, mode: os.FileMode(438), modTime: time.Unix(1572105320, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicCssMainCss = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\xe5\x52\x50\x50\x50\x48\x4a\x4c\xce\x4e\x2f\xca\x2f\xcd\x4b\xd1\x4d\xce\xcf\xc9\x2f\xb2\x52\x48\xca\x49\x4c\xce\xb6\xe6\xe5\xaa\xe5\xe5\x02\x04\x00\x00\xff\xff\x03\x25\x9c\x89\x29\x00\x00\x00"

func publicCssMainCssBytes() ([]byte, error) {
	return bindataRead(
		_publicCssMainCss,
		"public/css/main.css",
	)
}

func publicCssMainCss() (*asset, error) {
	bytes, err := publicCssMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/css/main.css", size: 41, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\x41\x0e\xc2\x20\x10\x45\xf7\x24\xdc\xe1\xa7\x07\x80\x74\x3f\xb2\x76\xe9\xc2\x0b\x60\x41\xc1\x50\x21\xc0\x42\xd3\xf4\xee\x06\x4a\x97\x93\xf7\x66\xde\x90\xab\x6b\x50\x9c\x71\x46\xce\x6a\xa3\x38\x03\x00\xaa\xbe\x06\xab\xb6\x0d\xe2\xa6\x5f\x56\xdc\xdb\x88\x7d\x27\x79\x00\xce\x48\x0e\x9d\x33\x7a\x44\xf3\x3b\x17\xdd\xac\x70\xb5\x21\x44\x3c\x73\x5c\xe1\x3f\xc6\x7e\x45\x6b\x80\xa4\x9b\xbb\x3f\xcc\xb2\x64\x9f\x2a\x4a\x5e\x2e\x93\xd4\x29\x89\x77\x99\x14\x40\xf2\x00\xbd\x31\x2e\xf7\x5c\xfb\xf3\x1f\x00\x00\xff\xff\x25\xe9\x37\x57\xae\x00\x00\x00"

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 174, mode: os.FileMode(438), modTime: time.Unix(1565946441, 0)}
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
	"public/app.js":          publicAppJs,
	"public/app2/index.html": publicApp2IndexHtml,
	"public/css/main.css":    publicCssMainCss,
	"public/index.html":      publicIndexHtml,
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
	"public": {nil, map[string]*bintree{
		"app.js": {publicAppJs, map[string]*bintree{}},
		"app2": {nil, map[string]*bintree{
			"index.html": {publicApp2IndexHtml, map[string]*bintree{}},
		}},
		"css": {nil, map[string]*bintree{
			"main.css": {publicCssMainCss, map[string]*bintree{}},
		}},
		"index.html": {publicIndexHtml, map[string]*bintree{}},
	}},
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
