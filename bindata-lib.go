// +build wflib

package writefreely

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _schema_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x5f\x6f\xa3\x38\x10\x7f\xef\xa7\xf0\xdb\xa6\x52\x23\x6d\x7a\xdd\xaa\xba\xd3\x3e\x64\x53\x76\x2f\xba\x94\xee\x25\x44\xba\x7d\x02\x03\x93\xd4\xaa\xb1\x91\x6d\x92\xe6\xdb\x9f\x8c\x49\x08\x86\x24\xd0\xdb\x3b\x71\x7d\x2a\xcc\x6f\x8c\xfd\x9b\x3f\x9e\x99\x0c\x87\x57\xc3\x21\x7a\xc4\x0a\x87\x58\xc2\xaf\x28\xd8\x0a\xa2\x60\x25\x00\xe8\x2e\xb8\x1a\x0e\xaf\xb4\x78\xf8\xce\x3f\xad\xac\xf5\x3d\x1c\x52\x40\x52\x89\x2c\x52\x99\x00\xb4\xe2\x02\xa9\xfc\x5d\x80\xa3\x08\xa4\x54\xfc\x15\x98\x34\xdf\x9b\xcc\x9d\xb1\xe7\x20\x6f\xfc\x65\xe6\xa0\xe9\x57\xe4\x3e\x7b\xc8\xf9\x6b\xba\xf0\x16\x16\x1a\x0d\xae\x10\x0a\xf2\x87\x00\x85\x84\x61\xb1\x1b\x8c\xee\xaf\x73\x05\x77\x39\x9b\xdd\x68\x71\x26\x41\xf8\x24\x0e\x10\x61\x6a\x60\x0b\x65\x16\xf3\x00\x29\xc2\x76\x5a\x3a\x2a\xa5\xe8\xd1\xf9\x3a\x5e\xce\x3c\xf4\xe1\xe3\x87\x1c\xc9\x19\xf8\x8a\x24\xd0\x0e\x1d\x09\xc0\x0a\xe2\x00\xc5\x58\x81\x56\xab\x43\x27\xcb\xf9\xdc\x71\x3d\xdf\x9b\x3e\x39\x0b\x6f\xfc\xf4\x3d\x57\x84\xb7\x94\x08\x90\x47\x8a\x7b\x7c\xf5\x40\x78\x0d\x4c\x05\x68\x83\x45\xf4\x82\xc5\xe0\xf6\xd3\xa7\xeb\x1a\xf2\xfb\x7c\xfa\x34\x9e\xff\x40\x7f\x38\x3f\xd0\xa0\xa0\xe9\xfa\xea\x1a\x39\xee\xb7\xa9\xeb\x7c\x9e\x32\xc6\x1f\xbf\x94\xfb\xf9\x7d\x3c\x5f\x38\xde\x67\x8a\x15\x61\xa3\xdf\xfe\x75\xb3\xa7\x69\xc4\x99\xd2\xa7\xb8\x6c\xf4\x12\x6b\x4c\xae\xcd\xb9\x3f\xfa\x2f\xb6\x4d\x0f\xd0\x04\x62\x92\x25\x0a\xde\x54\x7e\xb8\xf1\xc4\x73\xe6\x68\xe1\x78\x28\x53\xab\x07\x34\x79\x9e\xcd\xf4\x17\xf5\x83\x1f\x12\x66\x79\x4d\x1a\xbf\xcb\x80\x55\xce\x49\xdc\x2b\xc2\x13\xb2\x16\x58\x11\xde\x18\x68\x16\xc0\x10\xbd\x01\x21\x09\x67\x26\x78\x46\x23\x8b\x69\x03\x6f\x64\x29\x97\x0b\x90\x19\x55\x01\xca\x4d\xb0\x97\xf4\x85\x8f\x88\x53\x0a\x91\x3e\x2c\x56\x4a\x90\x30\x53\xd0\x22\xff\x34\x6a\x19\xae\x4a\xd1\xc9\x74\x73\xd0\x29\xdd\x77\x74\xfb\x60\x81\x36\x98\x66\x60\x85\x76\xdd\x7f\x93\xf0\xae\xe2\xc2\x49\x78\x57\xf3\xe2\xaa\x33\x56\xf7\x77\x73\xb4\x99\xde\xf8\x68\xb9\xc5\x57\xd8\x75\xb2\x46\x8e\x6f\x6d\x87\x34\x0b\x29\x89\xfc\x57\xd8\x05\x28\xa4\x3c\xb4\xa4\x82\x6c\xb0\x82\x13\xe2\x73\xa4\xf6\x90\xc8\x14\x4b\xb9\xe5\x22\xee\xc4\x66\xa9\xd4\x9e\xd2\x42\x25\x40\xb9\xd7\xde\x7f\xbc\xfe\x3f\xb3\x26\x20\x26\x02\x22\xd5\x89\xb5\x52\xc9\xb0\x96\x0a\xd8\xf8\x98\x12\x2c\x8f\xc2\xfd\xa3\x45\x4c\xc0\x60\x7b\x11\x54\x65\xef\x68\xdd\x1e\x52\xd7\x89\x32\x79\x74\xa1\x5b\x5e\x85\xc6\x4b\xef\xd9\x9f\xba\x93\xb9\xf3\xe4\xb8\x9e\xc9\x9f\x0d\x3c\xb5\x4f\x8d\xb5\x4a\x4a\x11\x45\x7f\x4e\xa6\x0d\x62\x90\x91\x20\xa9\xca\x2f\xcb\xc3\xfe\xee\x3b\xed\xaf\x5a\x99\xaa\x1d\x05\x5f\xbe\x00\x14\x17\xa8\x79\x9b\x7f\xa4\xb8\x51\x5b\xaf\x9c\xab\xae\xb8\x48\xf0\x51\xc9\xf8\x50\x2f\x18\x4d\xe6\x8b\x76\x8d\x35\xae\xa9\x82\xb7\xec\x4c\x35\xbd\x21\xb0\xf5\x23\x9e\xe9\xe2\xab\x41\x5e\xaf\x8d\xf4\xdb\xa5\x3b\xfd\x73\xe9\xe4\x2f\xf7\xf6\x1d\x04\x3d\xf3\xee\x94\xcb\x36\xa9\xc0\xc0\x4a\x8f\x2e\x9c\xc0\xee\x39\x68\xb6\xb6\x7c\xb8\x66\x88\x84\xc7\x64\xb5\xf3\x8b\xd6\xc6\xd4\xb9\xb7\x0d\x38\xed\x07\x3e\x4e\x53\xc0\x02\xb3\x08\x0a\xe8\x5d\x53\x67\xc2\xb8\x48\x4c\x73\x42\x31\x5b\x67\x78\xbd\x47\x37\xad\x2b\x14\xad\x38\xc1\x4f\xf0\x94\xda\x12\xcd\x97\x4a\xfd\x4b\x84\x31\x88\xfd\x94\x4b\x62\xa2\xeb\xe8\x8b\x4b\x77\x31\xfd\xe6\x3a\x8f\x0d\x8b\xef\x1b\x30\x5d\x95\x4a\x85\x93\xb4\x6d\x07\x76\xa8\xfc\x3b\x6b\x5e\x70\x7f\x3b\xdd\xfc\x93\xec\x70\xe8\x71\xba\x25\x82\x8e\xe1\x48\x62\xdf\x38\x6b\xbd\x78\xcc\xdf\xd7\x14\x4a\xa3\x0f\xca\xff\x6f\x0e\x6b\xe7\x98\xc2\x73\x0a\xd4\xde\x8f\x6e\x7a\xd5\x2b\x09\x48\xb8\x82\x15\xa7\x94\x6f\x5b\xc4\x7d\x15\x7e\xb2\x64\xaa\xf5\x4f\x46\xcf\xaf\x4c\x28\x6a\xa0\xd3\xa3\x84\xcb\x25\xbe\xf5\x81\x9e\xf1\xab\xb7\xd5\xae\xce\xb7\xf0\xf5\x21\x40\x7e\x75\x77\xe7\xf6\x6c\x1f\x70\x39\x3e\x8c\xc5\x0f\x1e\xdf\x7f\xb6\x3b\x51\x6d\xd7\x66\xc7\xec\x35\x16\x67\x91\xe2\x86\x8a\xd3\x56\x21\x2c\xe4\x6f\xe7\x00\xf2\x05\x0b\x88\xfd\x4b\xb8\xcb\xb6\xb1\xe2\x6f\x50\x6e\xaf\x37\x76\xd1\x24\x77\x99\x3d\x58\x78\x63\x9d\xb3\xe3\xcd\x86\x79\xc3\xfd\xdd\x7f\x34\x6e\xd8\x6f\xac\x97\x83\x06\xbd\x39\xc2\x36\xa4\x99\xf7\x8a\xd8\x2a\xe7\x6c\x8a\xab\x75\x4e\x7d\x44\x86\xdf\x74\x42\x90\x01\x92\x09\xa6\xf4\x64\x2d\x74\x36\xc9\xb7\x99\x0a\x13\x86\x23\x45\x36\xcd\xf3\xe9\x3e\xd1\xde\xd2\xd1\x3b\x76\x86\x5a\x85\xe1\x04\xde\xdd\x1c\x5e\x1a\x66\x54\x57\x32\x7c\x1d\x16\x32\x8f\xf5\x75\x20\xc1\x84\xe6\x5b\x2a\x7e\x9d\x68\x9c\xd3\xbf\xfb\xd7\x82\xcb\x59\xb0\xa4\x65\x50\xfe\xdf\xab\x28\x94\x26\xce\xe2\x53\x61\x78\x90\x17\xee\x90\x3f\xf9\x27\xc3\xf1\xe4\x7d\xdf\xfa\xcc\x7f\x07\x00\x00\xff\xff\xbe\x79\x68\xa8\x10\x1b\x00\x00")

func schema_sql() ([]byte, error) {
	return bindata_read(
		_schema_sql,
		"schema.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"schema.sql": schema_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"schema.sql": &_bintree_t{schema_sql, map[string]*_bintree_t{
	}},
}}
