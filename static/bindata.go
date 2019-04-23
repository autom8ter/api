package static

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

var _data_dash_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xc4\x40\x0c\x44\xfb\xfb\x0a\xcb\x3d\xec\x85\x74\x68\xd7\x15\x48\x54\x08\x89\x2f\x30\x17\x5f\x6c\x29\x21\xa7\x5d\x23\x8a\xd3\xfe\x3b\x5a\x72\x21\x54\x9e\xf1\x1b\x8f\x1c\xb5\xa3\x27\x2e\xfa\xb1\x70\x1e\x62\xd0\x8e\x0e\x51\x33\x84\x36\x3a\x8a\x36\x8f\x50\xf2\x29\xe1\xf5\x0a\xf7\x6f\x76\xf2\xaf\x2c\x50\x2b\x02\x4f\x9e\xf0\x92\x97\xb3\x4d\x72\x77\x59\x01\xc2\xb7\x0d\xae\x09\x1f\x8e\x47\x04\x15\x1b\xd5\x6f\x26\xd0\xd6\xdd\xd3\x2b\xcf\xf2\x08\xad\xb0\x29\xa8\x35\x06\xed\x57\xf4\x3c\xb3\x4d\x2b\xfb\x95\xff\xe0\xed\xa7\x9e\x22\x83\x9b\x4f\x92\xf0\x65\x99\x05\x41\xb3\x9c\x13\x06\xa4\x66\x63\x60\xda\xeb\xf6\xe8\xbb\xb8\xdb\xe7\x58\xfe\xe2\x03\x17\x0d\x65\xdb\xd2\xc6\xf7\xfb\xc3\x4f\x00\x00\x00\xff\xff\x4c\xad\x8b\xac\x1a\x01\x00\x00")

func data_dash_html() ([]byte, error) {
	return bindata_read(
		_data_dash_html,
		"data/dash.html",
	)
}

var _data_index_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\x73\x2c\x2d\xc9\xcf\xb5\x28\x49\x2d\xb2\xd1\xcf\x30\xb4\xe3\xb2\xc9\x28\x52\xd0\x07\x51\xc6\x76\x36\x89\x0a\x25\x99\x25\x39\xa9\xb6\x4a\x2e\x89\xc5\x19\x49\xf9\x89\x45\x29\x4a\x0a\x19\x45\xa9\x69\xb6\x4a\xfa\x29\x89\xc5\x19\x4a\x76\x70\x71\x1b\xfd\x44\x3b\x1b\xfd\x0c\x63\x3b\x2e\x2e\x40\x00\x00\x00\xff\xff\xa1\xc9\x6e\xd7\x54\x00\x00\x00")

func data_index_html() ([]byte, error) {
	return bindata_read(
		_data_index_html,
		"data/index.html",
	)
}

var _data_settings_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xc9\x30\xb4\x0b\x4e\x2d\x29\xc9\xcc\x4b\x2f\xb6\xd1\xcf\x30\xb4\xe3\xb2\xc9\x28\x52\xd0\x07\x51\xc6\x76\x36\x89\x0a\x25\x99\x25\x39\xa9\xb6\x4a\x1e\xf9\xb9\xa9\x4a\x0a\x19\x45\xa9\x69\xb6\x4a\xfa\x4a\x76\x20\xae\x8d\x7e\xa2\x9d\x8d\x7e\x86\x31\xba\x52\x97\xc4\xe2\x8c\xa4\xfc\xc4\xa2\x14\xb8\xfa\x94\xc4\xe2\x0c\x25\x3b\xb8\x38\x5c\x23\x20\x00\x00\xff\xff\x9c\x7d\x5c\x11\x7c\x00\x00\x00")

func data_settings_html() ([]byte, error) {
	return bindata_read(
		_data_settings_html,
		"data/settings.html",
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
	"data/dash.html":     data_dash_html,
	"data/index.html":    data_index_html,
	"data/settings.html": data_settings_html,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"dash.html":     &_bintree_t{data_dash_html, map[string]*_bintree_t{}},
		"index.html":    &_bintree_t{data_index_html, map[string]*_bintree_t{}},
		"settings.html": &_bintree_t{data_settings_html, map[string]*_bintree_t{}},
	}},
}}
