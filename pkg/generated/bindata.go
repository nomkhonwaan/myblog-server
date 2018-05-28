// Code generated by go-bindata.
// sources:
// pkg/graphql/schema/query.graphql
// pkg/graphql/schema/schema.graphql
// pkg/graphql/schema/types/post.graphql
// pkg/graphql/schema/types/tag.graphql
// pkg/graphql/graphiql/index.html
// DO NOT EDIT!

package generated

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

var _pkgGraphqlSchemaQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xc1\x4e\xf3\x30\x10\x84\xef\xfb\x14\xf3\xab\x87\xbf\x7d\x85\xdc\x68\x83\x50\x24\x10\x94\xfa\x86\x38\x98\x66\xed\x58\x4a\xec\xb0\x5e\x1f\x22\xd4\x77\x47\x75\x91\x28\xbd\x79\xf4\xcd\x37\xf2\xae\x60\x06\xc6\x67\x61\x59\x20\x29\x29\x92\xc3\xd3\xb2\x1d\x93\xff\x9f\xf1\x20\x76\x1e\xf6\x8f\x08\x51\x59\x9c\x3d\x32\xe9\x32\x33\xf6\xb5\xfd\x45\xc0\x0a\xaf\xac\x45\x22\x2c\xe6\xf2\x31\x86\x3c\x70\x8f\x97\x94\x15\x4e\xd2\x84\xa0\x19\x5d\x4b\xf8\x85\x67\xb6\x0e\x7d\x83\xae\xfd\xb7\x69\x6a\x95\xfe\xee\x18\xeb\x6f\x64\xb5\xfe\x4a\x31\xd6\xdf\x18\x63\xc8\xf5\xdb\xc6\xfa\x7c\xa9\xe7\x35\x01\x40\x72\x2e\xb3\x36\xe8\xa2\xd6\x3c\x86\x29\x5c\xc5\x24\x3d\xcb\x76\xa9\x93\xcf\x97\x77\x17\xe7\x72\x86\x9b\x06\x6f\xc6\xfa\x77\x3a\x11\x71\x2c\x13\x7e\x78\x1b\x84\x8f\x1a\x52\xac\xd7\xdf\x1d\x76\x04\xb4\xf7\x87\x1d\x9d\xe8\x3b\x00\x00\xff\xff\xc9\xe1\x5b\xc8\x49\x01\x00\x00")

func pkgGraphqlSchemaQueryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaQueryGraphql,
		"pkg/graphql/schema/query.graphql",
	)
}

func pkgGraphqlSchemaQueryGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaQueryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/query.graphql", size: 329, mode: os.FileMode(420), modTime: time.Unix(1521812750, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x3b\x44\xa1\x54\x1a\x00\x00\x00")

func pkgGraphqlSchemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaSchemaGraphql,
		"pkg/graphql/schema/schema.graphql",
	)
}

func pkgGraphqlSchemaSchemaGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/schema.graphql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1521643244, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaTypesPostGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\x41\x6b\x02\x31\x10\x85\xef\xf9\x15\x4f\x7a\xe8\x0f\xf0\xd4\xbd\x49\xa5\x54\xa8\x50\xac\xb7\xd2\xc3\xd4\x8d\xbb\xc1\x6c\xb2\x6c\x5e\x10\x29\xfd\xef\x65\x63\x56\x5d\x8b\xd0\x5e\x92\xc7\xcc\x7b\xdf\x30\x73\x87\x19\x5e\x7d\x20\xb4\xa3\xe1\x41\xf1\xd0\xea\x63\xe1\x4b\x01\x43\xf7\x3e\x60\x31\x57\x80\x29\x0b\x2c\xe6\x13\x05\x8c\x9b\x34\xb4\x5a\xe1\xf8\x17\x78\x63\x67\x5c\x35\x51\x63\x53\xb0\xb1\x52\x48\xdf\x2f\x8b\x35\x6e\x07\x7a\xb0\x36\x21\xf9\x15\x52\x6d\x30\x5e\xa3\x28\x8c\xa1\x87\x25\x71\x6b\xe2\xc6\x3b\x6a\x47\x18\x87\xe7\xf5\xf2\x05\x5b\xdf\x35\xd2\xa3\x6b\x36\xf6\x06\xfa\x22\xb3\x94\x6e\x57\xfa\xbd\x3b\xe7\x9a\x5c\x19\x67\x57\x9a\xb1\x73\x10\x58\x13\x08\xbf\x3d\x5d\x45\xfa\x7d\x29\x55\x28\xf0\xbe\x96\xea\x23\x8f\x72\x58\x3d\x3d\x4e\xa7\xd3\x07\x94\x42\x9d\xe9\x60\x9d\x9e\x7c\xfd\xbd\x04\x6c\x3a\x2d\xd4\xa5\xc2\xa0\x66\xbc\xde\xf4\xaf\xa8\xd8\x96\x19\x95\xd5\x19\xf5\x4f\x52\x1b\x3f\xad\x09\x75\x62\x9d\xf4\x05\xed\x5b\xfd\x04\x00\x00\xff\xff\x36\x83\x0e\x97\x53\x02\x00\x00")

func pkgGraphqlSchemaTypesPostGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaTypesPostGraphql,
		"pkg/graphql/schema/types/post.graphql",
	)
}

func pkgGraphqlSchemaTypesPostGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaTypesPostGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/types/post.graphql", size: 595, mode: os.FileMode(420), modTime: time.Unix(1521643251, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaTypesTagGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\xc1\x0a\x82\x40\x10\xc6\xf1\xfb\x3c\xc5\x17\x1d\x7a\x87\xbd\x15\x52\x08\x41\x07\xeb\x01\xa4\x9d\x6c\x48\x47\x59\xc7\x83\x84\xef\x1e\xbb\x18\x95\x9d\x86\x99\xff\x6f\x61\xd7\xd8\xe2\x5c\x56\x60\x35\xb1\x91\x6c\xec\x38\xed\x4f\x02\xe6\xb6\xe9\x91\x67\x04\x88\x77\xc8\xb3\x15\xfd\x14\x2d\x1b\x26\xa4\xe1\x50\x58\x10\xad\x16\xa2\xaf\x87\x8a\x90\xc6\x9f\xa8\x45\x1f\xb0\x16\x76\x97\x3e\x72\x42\x3a\x7d\xdc\x44\xc4\x3a\x34\xb1\x9d\x82\xe7\xb0\x1b\xf7\xc2\xb5\x4f\xdf\x2b\x8e\x97\x43\x04\xa2\xdd\x60\x5f\x22\x4f\x7b\x14\xb7\x68\xdd\xf2\x31\x01\x5e\x02\x5f\x4d\x5a\x75\x98\x4b\xf6\xbe\xd0\x44\xaf\x00\x00\x00\xff\xff\x5f\x30\xb4\x13\x14\x01\x00\x00")

func pkgGraphqlSchemaTypesTagGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaTypesTagGraphql,
		"pkg/graphql/schema/types/tag.graphql",
	)
}

func pkgGraphqlSchemaTypesTagGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaTypesTagGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/types/tag.graphql", size: 276, mode: os.FileMode(420), modTime: time.Unix(1521812750, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlGraphiqlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x6f\xd3\x30\x14\x3f\x77\x9f\xc2\x58\x42\x4a\xa5\xcd\x6e\x06\xda\x21\x4d\x7b\x18\x1b\x08\x34\xd8\x06\x5c\x38\x7a\xf6\x4b\xed\xe1\xd8\xd9\xb3\xd3\xad\x9a\xf6\xdd\x51\x9c\x64\x74\x45\x02\x84\xc6\xa9\x7e\xaf\xef\xf7\xc7\x7e\x3f\xa5\x7c\x71\x72\xfe\xe6\xeb\xb7\x8b\x53\xa2\x63\x6d\x97\x7b\x65\xff\x33\x29\x35\x08\xb5\xdc\x9b\x4c\x4a\x6b\xdc\x77\xa2\x11\xaa\x05\xd5\x31\x36\xa1\xe0\x5c\x2a\x77\x1d\x98\xb4\xbe\x55\x95\x15\x08\x4c\xfa\x9a\x8b\x6b\x71\xc7\xad\xb9\x0a\x7c\x85\xa2\xd1\xe6\xc6\xf2\x19\xcb\x73\x96\xe7\x8f\x0d\x56\x1b\xc7\x64\x08\x94\x20\xd8\x05\x0d\x71\x63\x21\x68\x80\x48\x09\x4f\x5a\x41\xa2\x69\x22\x09\x28\xff\x5a\x0c\xc2\xd1\x41\x83\xbe\x36\x01\xf8\x6b\x96\xb3\x7c\xbb\xc3\x44\x1b\x7d\x52\xbd\x0e\x74\x59\xf2\x9e\xff\x5f\xa5\x2a\x88\x52\xf3\x43\x36\x63\xaf\xfa\xf3\xb3\x31\x23\x08\x19\x79\x7e\xc4\x0e\xd9\x8c\xb7\xb5\xea\x1b\xac\x41\xaf\x5a\x19\x8d\x77\xcf\xab\x74\xa0\x7c\xfd\x8b\x5a\xd7\xfc\x1f\x8a\xbf\x4f\xc3\x8e\x42\xc9\x87\xdc\x95\x57\x5e\x6d\x48\x4a\xc8\x82\xde\x1a\x15\x75\x41\xf2\xd9\xec\xe5\x9c\x68\x30\x2b\x1d\xc7\xaa\x16\xb8\x32\xae\x20\xb3\x39\xf1\x6b\xc0\xca\xfa\xdb\x82\x68\xa3\x14\xb8\x39\x4d\x96\x95\x59\x13\xa3\x16\x74\x94\xa5\x23\xeb\x16\xd1\x5a\xcf\xe9\xf2\xcc\x0b\x65\xdc\x8a\x31\x56\x72\x65\xd6\x5b\xf7\xed\x8e\x93\xaa\x75\xe9\x61\x48\x22\xba\x3c\x7b\xdb\x25\x00\x30\x1b\xca\x0b\x81\xa2\x0e\x53\x72\xdf\xcd\x4e\x10\x62\x8b\x8e\xa4\x94\x64\xb4\xbf\xf2\x8d\xa5\xfb\xc3\xdf\x93\x1a\xa2\xf6\xaa\x20\xb4\xf1\x21\xd2\xfd\xbe\xd9\x5d\xb9\x20\x1f\xbe\x9c\x7f\x62\x21\xa2\x71\x2b\x53\x6d\x76\xe8\x87\x49\x89\xa0\xc0\x45\x23\x6c\x28\x08\x35\x4e\xda\x56\xc1\x40\xf3\x30\x65\x51\x83\xcb\x1e\xfd\x66\x08\xa1\xf1\x2e\xc0\x68\x6e\x74\x37\xf6\x59\x84\xbb\x98\x4d\xe7\x7f\x80\x1f\x7b\xb5\x79\xa4\x88\xb8\x19\x8f\x23\x5d\x32\xde\x08\x0c\xf0\x14\xd2\xf3\x4e\x1e\x88\x14\x51\x6a\x92\x01\xa2\xc7\xe9\x2e\x7a\x1b\x32\x22\x06\x43\xa9\x4c\xc5\xe7\x2e\xa8\x27\xe7\x1f\x19\x82\x53\x80\x59\x1a\x48\x4d\x26\x11\x44\x84\x53\x0b\x35\xb8\x98\xbd\x4b\xcb\xbe\x3c\xdb\x27\xf7\x55\xbf\xa7\x62\x67\x6f\x0f\xc3\x5b\x2a\x2f\xdb\x0e\xc2\x56\x10\x07\xf4\xf1\xe6\xbd\xca\x7e\xe6\x65\xda\xcd\x25\x13\x4f\x82\xda\x6d\x6b\xb9\x57\xf2\xf4\xc1\xfc\x11\x00\x00\xff\xff\xf3\x8d\xa6\xf7\x47\x05\x00\x00")

func pkgGraphqlGraphiqlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlGraphiqlIndexHtml,
		"pkg/graphql/graphiql/index.html",
	)
}

func pkgGraphqlGraphiqlIndexHtml() (*asset, error) {
	bytes, err := pkgGraphqlGraphiqlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/graphiql/index.html", size: 1351, mode: os.FileMode(420), modTime: time.Unix(1521643244, 0)}
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
	"pkg/graphql/schema/query.graphql": pkgGraphqlSchemaQueryGraphql,
	"pkg/graphql/schema/schema.graphql": pkgGraphqlSchemaSchemaGraphql,
	"pkg/graphql/schema/types/post.graphql": pkgGraphqlSchemaTypesPostGraphql,
	"pkg/graphql/schema/types/tag.graphql": pkgGraphqlSchemaTypesTagGraphql,
	"pkg/graphql/graphiql/index.html": pkgGraphqlGraphiqlIndexHtml,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"graphql": &bintree{nil, map[string]*bintree{
			"graphiql": &bintree{nil, map[string]*bintree{
				"index.html": &bintree{pkgGraphqlGraphiqlIndexHtml, map[string]*bintree{}},
			}},
			"schema": &bintree{nil, map[string]*bintree{
				"query.graphql": &bintree{pkgGraphqlSchemaQueryGraphql, map[string]*bintree{}},
				"schema.graphql": &bintree{pkgGraphqlSchemaSchemaGraphql, map[string]*bintree{}},
				"types": &bintree{nil, map[string]*bintree{
					"post.graphql": &bintree{pkgGraphqlSchemaTypesPostGraphql, map[string]*bintree{}},
					"tag.graphql": &bintree{pkgGraphqlSchemaTypesTagGraphql, map[string]*bintree{}},
				}},
			}},
		}},
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

