// Code generated by go-bindata.
// sources:
// sql/generic.sql
// DO NOT EDIT!

package dao

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

var _sqlGenericSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xc1\x4a\xc3\x40\x10\xc6\xf1\xfb\x3e\xc5\x1c\x12\x68\x0a\x3d\xd4\xa3\xe2\x21\x76\x27\xba\x90\x66\x25\x99\xe8\x79\x6b\x47\x08\x6c\x52\xe9\x4e\x04\x09\x79\x77\xe9\xa1\x26\xd6\x8b\xbd\x0e\xfc\xfe\xf3\xad\x56\xd0\xb9\x96\x6f\xe1\x91\x25\xf5\x5e\x55\x98\xe3\x86\x60\x09\x59\x69\xb7\x10\x0f\xe2\x76\x9e\x47\xb0\x59\x56\x21\x41\xb4\x86\xdc\x6c\x0d\x41\x74\x73\xa7\xd4\xdc\x1a\x1d\xce\xb6\xd9\x5f\x85\xe9\x20\xce\x17\x7d\xbb\xe3\xa3\x7d\xc7\x4e\x1a\x69\xf8\xa7\xb5\xb1\x75\x41\x8b\x65\xf2\xbb\x78\xf1\xfb\xe1\xcb\xe8\x33\x88\x87\xb7\x83\xef\xdb\x2e\x8c\x17\x23\x5e\x9f\xb0\xc4\xd3\xb6\x7b\x88\xd6\xf3\x80\x66\xcf\xc2\x4a\x63\x8e\x84\xff\x45\xa6\x0b\x7c\x14\x65\x8a\x0a\x4b\x02\x53\x90\x9d\xd0\x62\xda\x90\xc0\x4b\x9a\xd7\x58\x9d\x6e\x9f\xce\xf7\x1c\xc6\x64\x9e\xa9\x3f\xf6\x4e\x58\xd5\xcf\x3a\x25\x9c\x0a\xf1\x10\x58\xc2\x9f\xf7\xdf\x01\x00\x00\xff\xff\xb9\x2d\x15\xa1\xac\x01\x00\x00")

func sqlGenericSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlGenericSql,
		"sql/generic.sql",
	)
}

func sqlGenericSql() (*asset, error) {
	bytes, err := sqlGenericSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/generic.sql", size: 428, mode: os.FileMode(420), modTime: time.Unix(1511875491, 0)}
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
	"sql/generic.sql": sqlGenericSql,
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
	"sql": &bintree{nil, map[string]*bintree{
		"generic.sql": &bintree{sqlGenericSql, map[string]*bintree{}},
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
