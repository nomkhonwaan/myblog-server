// Code generated by go-bindata.
// sources:
// pkg/graphql/schema/interfaces/node.graphql
// pkg/graphql/schema/query.graphql
// pkg/graphql/schema/schema.graphql
// pkg/graphql/schema/types/page_info.graphql
// pkg/graphql/schema/types/post.graphql
// pkg/graphql/graphiql/graphiql.html
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

var _pkgGraphqlSchemaInterfacesNodeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\xcc\x53\xc8\x4f\xca\x4a\x4d\x2e\x51\x28\xcf\x2c\xc9\x50\x48\xcc\x53\xf0\x74\xe1\xca\xcc\x2b\x49\x2d\x4a\x4b\x4c\x4e\x55\xf0\xcb\x4f\x49\x55\xa8\xe6\x52\x50\x50\x56\x08\xc9\x48\x55\xc8\x4c\x51\xc8\x4f\x53\x28\xc9\x48\x85\x6a\xd2\xe3\x52\x50\xc8\x4c\xb1\x52\xf0\x74\x51\xe4\xaa\xe5\x02\x04\x00\x00\xff\xff\xb0\xa5\x30\xed\x4e\x00\x00\x00")

func pkgGraphqlSchemaInterfacesNodeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaInterfacesNodeGraphql,
		"pkg/graphql/schema/interfaces/node.graphql",
	)
}

func pkgGraphqlSchemaInterfacesNodeGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaInterfacesNodeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/interfaces/node.graphql", size: 78, mode: os.FileMode(420), modTime: time.Unix(1520819826, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x5d\x6e\xdb\x40\x0c\x84\xdf\xf7\x14\x53\xe8\xa1\x2e\xe0\x13\xe8\xad\xb5\xda\xc2\x40\xd1\xd6\x71\x0e\xe0\x95\x96\x92\x36\x90\x48\x85\x4b\xd9\x10\x02\xdf\x3d\xb0\xec\xf8\x27\x8f\x9c\xe1\x7e\x33\x0b\x66\x78\x6e\x09\xaf\x23\xe9\x04\x15\x31\x48\x8d\x7e\x2a\x3b\x69\xbe\x26\xfc\x56\x3f\xb4\x9b\x3f\x88\x6c\xa4\xb5\xaf\xc8\xd9\x34\x10\x36\xf3\xf6\x9b\x03\x32\x64\xf8\x45\x56\xb5\x94\xe0\x19\x52\xbe\x50\x65\x68\xe2\x9e\x18\xd1\x12\xd6\xc5\xbc\xc4\x12\x68\x11\x43\x8e\x75\xf1\xe5\x5b\x8e\xbf\x12\x68\xd6\x9f\xc8\x46\x65\x78\x0c\x63\xd9\xc5\xd4\x52\xc0\x20\xc9\x50\xab\xf4\xb7\xf7\x57\xf3\xbf\x24\xbb\xc3\x9c\xc6\x47\x4c\x17\xd3\xfc\x81\x47\x5c\xfa\xcc\x48\x0b\x5f\x1b\x69\x8e\xad\x69\xe4\x66\x89\x3a\x6a\xb2\x1c\x6b\xb6\x25\x4a\xaa\x45\xe9\xe6\x75\xfe\x6a\x89\x06\xd2\x1f\xd3\x39\xf9\xdf\x69\xb8\xb4\x58\x09\x33\x55\x16\x85\xdd\xd1\xb9\xec\xa4\xa5\x58\x76\x84\x10\xf5\xac\x27\x44\xc6\xa1\x8d\x55\x0b\x93\x33\xe8\xae\x6f\x34\xea\x13\x0e\x2d\x31\x06\x95\x7d\x0c\x14\xe0\x39\x60\x77\x49\xdc\xc1\x6b\x33\xf6\xc4\xe6\x88\xc7\x1e\x73\x76\xf1\xc1\x9e\x2f\xf1\x7d\xbb\x72\x40\xf1\x73\xbb\x72\x47\xf7\x1e\x00\x00\xff\xff\x83\x1b\xb4\x95\xd5\x01\x00\x00")

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

	info := bindataFileInfo{name: "pkg/graphql/schema/query.graphql", size: 469, mode: os.FileMode(420), modTime: time.Unix(1520908942, 0)}
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

	info := bindataFileInfo{name: "pkg/graphql/schema/schema.graphql", size: 26, mode: os.FileMode(420), modTime: time.Unix(1520478991, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaTypesPage_infoGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcf\xbd\x4a\xc7\x40\x10\x04\xf0\xfe\x9e\x62\xe4\xdf\x4a\x1e\x20\x8d\xa0\x95\x8d\x04\x2c\xac\x37\x97\xcd\xdd\xa1\xd9\x0d\x7b\x7b\x7e\x20\xbe\xbb\x9c\x82\x82\x44\xc4\x72\x77\x98\x1f\xcc\x09\xd7\xb2\xaa\x6d\xe4\x45\x05\x34\x6b\x73\xec\x94\x8a\x7c\x3e\x8a\x80\x10\x55\x84\x63\xbf\x87\xe0\x2f\x3b\x63\xa2\xc4\xbd\x86\xd7\x00\x9c\x70\x97\x59\xbe\x4a\x92\xb0\xaa\x3d\x91\x2d\xf5\x1c\x64\x0c\xcf\x6c\x8c\x4d\x8d\x51\x9c\xb7\x7a\x11\x80\x4c\xf5\x86\x9f\xbd\x3b\x23\x2e\x55\x1f\x98\xe4\x2c\x1c\x62\x33\xc5\xfb\xbf\xb5\xc9\xf8\xb1\x68\xab\xff\x14\x3d\x33\x62\xb3\xaa\x06\xd7\xbe\xd3\x8b\x34\x1e\x02\x50\x9d\xcc\xaf\x3e\xa2\x11\xb7\x6e\x45\xd2\xb1\xf6\x3d\xf6\x77\x8c\x65\xf9\x41\xbd\x85\xf7\x00\x00\x00\xff\xff\x92\xd2\xc8\xaf\x79\x01\x00\x00")

func pkgGraphqlSchemaTypesPage_infoGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlSchemaTypesPage_infoGraphql,
		"pkg/graphql/schema/types/page_info.graphql",
	)
}

func pkgGraphqlSchemaTypesPage_infoGraphql() (*asset, error) {
	bytes, err := pkgGraphqlSchemaTypesPage_infoGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/schema/types/page_info.graphql", size: 377, mode: os.FileMode(420), modTime: time.Unix(1520909643, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlSchemaTypesPostGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xd1\x6e\xdb\x3a\x0c\x7d\xd7\x57\x9c\xe0\x3e\xdc\x0f\xe8\xd3\xf5\x5b\x1a\xa7\xb7\x01\xda\x2d\x68\x92\xa7\xa2\x58\x55\x8b\xb1\x85\xda\x92\x27\xd1\x0b\x82\xa1\xff\x3e\x48\xb2\xab\x34\x41\x81\x6d\x2f\x31\x49\x1d\x1e\x1e\x91\x54\xfe\xc1\x1c\xbd\xf5\x0c\x32\xac\xf9\x28\xf8\xd8\x13\xd6\x21\xa0\xbb\xbe\xa5\x8e\x0c\x7b\x7c\xb1\x8a\xf0\x53\x00\x13\xfa\x5f\x8f\x55\x29\x00\xad\x0a\xac\xca\x99\xf8\x78\xc4\x9a\x5b\x12\x48\xdf\x02\x1b\x76\xda\xd4\xe7\x20\xdf\x0e\xb5\x40\xfc\x9c\x43\x0c\xe4\x8b\xb7\xed\xc0\x84\x56\x9b\x57\xb0\x05\x37\x14\x13\x05\x62\xe8\x53\x52\x96\x3c\xf8\x40\x1b\x8d\x22\xde\x64\x13\xed\x73\x68\x65\x0d\x93\x61\x68\x83\xdb\xed\xfd\x1d\xf6\xd6\x75\x32\xf0\x37\xdc\xb5\x13\xff\xe7\x39\xf7\xd2\xbd\x2a\x7b\x30\x39\xaf\x1b\x23\x1f\x73\x4f\xb2\xe5\xc0\x8d\x75\x31\x9a\xcc\x02\x3b\x4f\x6e\x76\x01\x64\x59\xfb\x18\x0b\x46\x81\xc7\xad\xac\x67\x4f\x17\xa8\x4a\x32\xd5\xd6\x69\x4a\xd8\xec\x16\x78\x5c\x24\xe7\x38\x7b\xca\x3d\x7d\xb8\x59\x5c\x5d\x5d\xfd\x07\x25\x99\x46\xd5\xe0\x26\xfe\xa4\xde\xe2\x20\x3d\x2a\x47\x92\x49\x09\x4c\xd6\x9c\x2f\xc7\xf3\x7b\x54\x43\xaf\x46\xaa\xd1\xca\x54\x7f\xc8\xd4\x0f\x2f\xad\xf6\x4d\xe4\x7a\xb7\x4f\xd8\xde\x84\x08\x8d\xa9\xac\x31\x54\xb1\xb6\x26\xac\x8c\x44\xab\x3d\xc3\xee\xa1\x99\x3a\x9f\x37\x7b\x91\x61\x69\xa7\x57\x26\x55\x9e\x12\xb5\x0a\x13\xee\x65\xad\x4d\x0c\x86\xaa\xb2\xa6\x00\x2b\xb0\x1e\xad\xf7\x7d\x9a\xaa\x90\xaa\xe3\x28\xe2\xb7\xc0\x63\x28\xb5\x54\x35\x3d\x8d\xf2\x4c\x3c\x09\xcc\xf2\x44\x69\x96\x15\xb0\xa3\xa0\x6d\x43\x51\x34\xc6\x46\x90\x51\xa1\x42\x34\x55\x1d\x9e\x96\xb1\x8a\xd2\x76\x4f\x32\xaa\xc1\x79\xeb\x42\x0f\x31\x78\xba\xb8\x40\x3a\xce\xa3\x8c\x9a\xbe\x3a\x45\xc1\x85\xed\x03\xcc\xc7\xec\xef\x03\xb9\x63\x08\xe6\x06\x86\x41\x78\xa1\x4d\x3f\x70\xac\x19\xf3\xa2\x56\xa5\x5d\xba\x47\x91\xc8\xca\xc9\x17\xc0\x5e\x53\xab\x8a\x9c\x70\x13\xfc\x54\x78\x6d\xbd\xd7\x2f\x2d\x8d\xef\x94\x7c\xd0\x7b\x68\x74\xd5\x84\x01\xec\x75\xcb\xe4\xce\xeb\xe3\xd0\x90\x41\xef\xec\x0f\xad\x48\x41\xe2\x39\x25\x3f\x43\xba\x7a\x08\x7f\x54\x82\xcc\xd0\x9d\x3c\xf9\xa8\xb0\x7c\x98\xdf\x6c\x05\xb0\xde\x5d\xdf\xad\x36\xb7\xcb\xf2\x4c\x41\x54\xf9\x37\xf5\x0d\x9e\x6d\xb8\xd6\xf5\x31\x2b\x40\x96\x90\xaf\x1c\x65\x6c\xee\x76\xff\x0b\x60\xf1\xb0\x9c\x6f\x97\xe5\xb7\x79\x90\xb4\x5b\x97\xd9\x79\xd7\x17\xdc\x37\xf1\x2b\x00\x00\xff\xff\x30\x7c\x27\x0e\x98\x05\x00\x00")

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

	info := bindataFileInfo{name: "pkg/graphql/schema/types/post.graphql", size: 1432, mode: os.FileMode(420), modTime: time.Unix(1520844220, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgGraphqlGraphiqlGraphiqlHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x6f\xd3\x30\x14\x3f\x77\x9f\xc2\x58\x42\x4a\xa5\xcd\x6e\x06\xda\x21\x4d\x7b\x18\x1b\x08\x34\xd8\x06\x5c\x38\x7a\xf6\x4b\xed\xe1\xd8\xd9\xb3\xd3\xad\x9a\xf6\xdd\x51\x9c\x64\x74\x45\x02\x84\xc6\xa9\x7e\xaf\xef\xf7\xc7\x7e\x3f\xa5\x7c\x71\x72\xfe\xe6\xeb\xb7\x8b\x53\xa2\x63\x6d\x97\x7b\x65\xff\x33\x29\x35\x08\xb5\xdc\x9b\x4c\x4a\x6b\xdc\x77\xa2\x11\xaa\x05\xd5\x31\x36\xa1\xe0\x5c\x2a\x77\x1d\x98\xb4\xbe\x55\x95\x15\x08\x4c\xfa\x9a\x8b\x6b\x71\xc7\xad\xb9\x0a\x7c\x85\xa2\xd1\xe6\xc6\xf2\x19\xcb\x73\x96\xe7\x8f\x0d\x56\x1b\xc7\x64\x08\x94\x20\xd8\x05\x0d\x71\x63\x21\x68\x80\x48\x09\x4f\x5a\x41\xa2\x69\x22\x09\x28\xff\x5a\x0c\xc2\xd1\x41\x83\xbe\x36\x01\xf8\x6b\x96\xb3\x7c\xbb\xc3\x44\x1b\x7d\x52\xbd\x0e\x74\x59\xf2\x9e\xff\x5f\xa5\x2a\x88\x52\xf3\x43\x36\x63\xaf\xfa\xf3\xb3\x31\x23\x08\x19\x79\x7e\xc4\x0e\xd9\x8c\xb7\xb5\xea\x1b\xac\x41\xaf\x5a\x19\x8d\x77\xcf\xab\x74\xa0\x7c\xfd\x8b\x5a\xd7\xfc\x1f\x8a\xbf\x4f\xc3\x8e\x42\xc9\x87\xdc\x95\x57\x5e\x6d\x48\x4a\xc8\x82\xde\x1a\x15\x75\x41\xf2\xd9\xec\xe5\x9c\x68\x30\x2b\x1d\xc7\xaa\x16\xb8\x32\xae\x20\xb3\x39\xf1\x6b\xc0\xca\xfa\xdb\x82\x68\xa3\x14\xb8\x39\x4d\x96\x95\x59\x13\xa3\x16\x74\x94\xa5\x23\xeb\x16\xd1\x5a\xcf\xe9\xf2\xcc\x0b\x65\xdc\x8a\x31\x56\x72\x65\xd6\x5b\xf7\xed\x8e\x93\xaa\x75\xe9\x61\x48\x22\xba\x3c\x7b\xdb\x25\x00\x30\x1b\xca\x0b\x81\xa2\x0e\x53\x72\xdf\xcd\x4e\x10\x62\x8b\x8e\xa4\x94\x64\xb4\xbf\xf2\x8d\xa5\xfb\xc3\xdf\x93\x1a\xa2\xf6\xaa\x20\xb4\xf1\x21\xd2\xfd\xbe\xd9\x5d\xb9\x20\x1f\xbe\x9c\x7f\x62\x21\xa2\x71\x2b\x53\x6d\x76\xe8\x87\x49\x89\xa0\xc0\x45\x23\x6c\x28\x08\x35\x4e\xda\x56\xc1\x40\xf3\x30\x65\x51\x83\xcb\x1e\xfd\x66\x08\xa1\xf1\x2e\xc0\x68\x6e\x74\x37\xf6\x59\x84\xbb\x98\x4d\xe7\x7f\x80\x1f\x7b\xb5\x79\xa4\x88\xb8\x19\x8f\x23\x5d\x32\xde\x08\x0c\xf0\x14\xd2\xf3\x4e\x1e\x88\x14\x51\x6a\x92\x01\xa2\xc7\xe9\x2e\x7a\x1b\x32\x22\x06\x43\xa9\x4c\xc5\xe7\x2e\xa8\x27\xe7\x1f\x19\x82\x53\x80\x59\x1a\x48\x4d\x26\x11\x44\x84\x53\x0b\x35\xb8\x98\xbd\x4b\xcb\xbe\x3c\xdb\x27\xf7\x55\xbf\xa7\x62\x67\x6f\x0f\xc3\x5b\x2a\x2f\xdb\x0e\xc2\x56\x10\x07\xf4\xf1\xe6\xbd\xca\x7e\xe6\x65\xda\xcd\x25\x13\x4f\x82\xda\x6d\x6b\xb9\x57\xf2\xf4\xc1\xfc\x11\x00\x00\xff\xff\xf3\x8d\xa6\xf7\x47\x05\x00\x00")

func pkgGraphqlGraphiqlGraphiqlHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pkgGraphqlGraphiqlGraphiqlHtml,
		"pkg/graphql/graphiql/graphiql.html",
	)
}

func pkgGraphqlGraphiqlGraphiqlHtml() (*asset, error) {
	bytes, err := pkgGraphqlGraphiqlGraphiqlHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/graphql/graphiql/graphiql.html", size: 1351, mode: os.FileMode(420), modTime: time.Unix(1520478991, 0)}
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
	"pkg/graphql/schema/interfaces/node.graphql": pkgGraphqlSchemaInterfacesNodeGraphql,
	"pkg/graphql/schema/query.graphql": pkgGraphqlSchemaQueryGraphql,
	"pkg/graphql/schema/schema.graphql": pkgGraphqlSchemaSchemaGraphql,
	"pkg/graphql/schema/types/page_info.graphql": pkgGraphqlSchemaTypesPage_infoGraphql,
	"pkg/graphql/schema/types/post.graphql": pkgGraphqlSchemaTypesPostGraphql,
	"pkg/graphql/graphiql/graphiql.html": pkgGraphqlGraphiqlGraphiqlHtml,
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
				"graphiql.html": &bintree{pkgGraphqlGraphiqlGraphiqlHtml, map[string]*bintree{}},
			}},
			"schema": &bintree{nil, map[string]*bintree{
				"interfaces": &bintree{nil, map[string]*bintree{
					"node.graphql": &bintree{pkgGraphqlSchemaInterfacesNodeGraphql, map[string]*bintree{}},
				}},
				"query.graphql": &bintree{pkgGraphqlSchemaQueryGraphql, map[string]*bintree{}},
				"schema.graphql": &bintree{pkgGraphqlSchemaSchemaGraphql, map[string]*bintree{}},
				"types": &bintree{nil, map[string]*bintree{
					"page_info.graphql": &bintree{pkgGraphqlSchemaTypesPage_infoGraphql, map[string]*bintree{}},
					"post.graphql": &bintree{pkgGraphqlSchemaTypesPostGraphql, map[string]*bintree{}},
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

