// Code generated by go-bindata.
// sources:
// web/static/flipclock.css
// web/static/flipclock.min.js
// web/static/index.html
// DO NOT EDIT!

package web

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _webStaticFlipclockCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xdb\x36\x14\x7e\x2f\xd0\xff\x70\x96\x22\x40\x62\x58\xb6\x6c\xc7\x6d\xa2\xbe\xac\x40\x81\xf5\x61\x6b\x87\xf5\xa1\xdb\x23\x2d\x51\x36\x11\x8a\x14\x48\x2a\x71\x12\xf4\xbf\x0f\xa4\x2e\xd1\x95\xa2\x0d\x2f\xc0\x80\xa0\x45\x11\x91\x87\xe7\xfa\x9d\xc3\x8f\x45\xe6\x13\xf8\x0d\x2b\x50\x3b\x0c\x1b\x9e\x89\x0d\x67\x90\x90\x3d\x61\x10\x0b\x9e\xc0\x4e\xa9\x34\x98\xcf\x8b\x9d\x19\xe1\x30\x99\xbf\x7d\x33\x9f\xc0\x5f\x58\x62\x65\x3e\x66\x31\x25\xa9\x17\x52\x1e\xde\x7a\xf7\x02\xa5\x29\x16\x30\x81\xa7\xb7\x6f\x00\x00\xbc\x7b\xbc\xb9\x25\xca\xdb\xf0\xbd\x27\xc9\x23\x61\xdb\x00\x36\x5c\x44\x58\xe8\xa5\x8f\x85\x50\xc2\x1f\xc7\x24\xe4\x88\x00\xb7\xef\xdb\x0f\x97\x4e\xa2\xf0\x36\x46\x21\xf6\xee\x88\x24\x1b\x42\x89\x7a\x08\x60\x47\xa2\x08\xb3\x86\xa7\x0e\x62\xd2\x45\x8a\x3b\x08\xd9\x25\x7e\xbe\x7d\xd3\x5f\x00\x94\x17\x20\xcc\x84\xe4\x22\x80\x94\x13\xa6\xb0\x30\x3a\x15\xde\x2b\x2f\xc2\x21\x17\x48\x11\xce\x02\x60\x9c\x61\xb3\x13\x72\xaa\x85\xdf\x85\x61\xf8\x11\x86\x55\x07\x3b\x7e\x87\x45\x61\xa0\x38\x12\xc7\xb1\xe5\x48\x46\x73\x69\x4a\xa4\xf2\xa4\x7a\xa0\xb8\xb0\x3a\x78\x64\x16\x52\x8c\x44\x4c\xf6\xc1\x06\xc7\x5c\xe0\xe9\x88\x14\x8a\xd5\xb3\x4b\x4c\x61\xa6\x02\x38\x83\x33\x13\x56\x44\x64\x4a\xd1\x43\x00\x0a\x6d\xa8\x93\xcd\xba\x36\xbd\xa6\x31\xa3\x76\x0e\x27\xf3\x33\x93\x47\xce\x93\x00\x16\xc5\x81\xf9\x04\xfe\x40\x84\x0d\x36\x8b\x39\x13\x73\xed\x32\xe3\x22\x41\x14\x16\x8b\x74\x0f\x67\x5f\x30\xbd\xc3\x8a\x84\x08\xbe\xe2\x0c\x9f\x4d\xa1\x5a\x98\x82\x44\x4c\x7a\x12\x0b\x12\x9b\x10\x4b\x04\x67\x12\x0b\x4f\x62\x8a\x43\x35\x98\xe1\x04\x0b\x12\x91\x2c\xc9\xed\x6a\x7c\x6d\x05\xcf\x58\x94\x1f\x80\x5f\x48\x92\x72\xa1\x10\x53\x46\xb3\x69\x9c\x1d\x8a\xf8\x7d\x00\xbe\xf9\xd3\x12\xd0\x8e\xeb\xd6\xc2\x01\xac\xde\xa7\xfb\xfa\xb6\xc5\x36\x82\xa7\x0a\x3b\xab\xc5\x6a\xb5\x5a\x59\xe0\xf3\x54\xc1\x16\x51\xb2\x65\x01\x84\xb8\x42\x73\xca\x25\xc9\x61\x2c\x30\x45\x8a\xdc\xe5\x50\xbe\x27\x91\xda\x05\xb0\xf0\xfd\x73\xf3\x9d\x20\xb1\x25\x2c\x80\x05\x4e\x2c\x6d\x63\x05\x5b\x1d\x15\x2d\x94\xc1\x7c\x02\x0b\x53\x5f\xbd\xd5\xc6\xdb\x7c\x02\x4b\xb3\xf9\xd3\x45\x6f\x0d\x6f\x85\xa3\xf3\x09\x7c\xbf\xc5\x14\x2b\x3e\x8c\xa1\xb2\xc3\x06\xd2\x11\x53\x8e\x54\x00\x14\xc7\xaa\x91\x8e\x75\xba\xaf\xa7\xeb\xbd\x5f\x7c\xef\x30\xd9\xee\x54\x00\x37\xe5\x42\xad\xca\xd7\x8d\xb5\xfb\x42\x72\xc3\x69\xf4\x31\xef\x71\x86\xbd\xf2\xfc\xf5\x87\x42\xb6\x98\xb9\x02\x45\x24\x93\x01\xbc\x2f\x97\x6b\xe8\x7b\xe7\xfb\xbe\x6d\xa4\x65\x14\x28\xc9\xa3\x7c\xf4\x08\x8b\xf0\x5e\xb7\x58\x33\x68\xb4\x91\x9c\x66\x2a\x0f\x5a\x47\x1b\x80\x9f\x0f\x3d\x9e\x96\x3f\xb6\xa1\x51\xfa\x5a\x2d\xf4\x47\xd0\x3b\x36\x9b\xcd\x60\x77\x3d\x88\x89\x90\xca\x0b\x77\x84\x46\xad\x30\x96\xf6\xd9\x49\x49\x39\xcf\x2b\x68\x6d\xb4\x4c\xbf\xfb\xe5\x2c\x48\xb1\x90\x29\x0e\x35\x06\x02\x58\xfa\x65\xd5\xcc\x15\x36\xb0\x37\xb0\x5c\xa2\xa5\xd3\xfb\xfa\x1e\x88\xa9\x1e\x0d\xe6\x76\xa2\x9d\xe9\x51\xde\x3f\x11\x8e\x51\x46\x95\x7d\x38\xb4\x23\x8e\xc8\xdd\xb1\xe5\x1e\xaa\xf1\xba\xf8\xee\x83\xf3\x73\x30\xc5\xf5\x0a\x66\x35\x53\x1a\x0f\x01\xe8\x91\x2c\x39\x25\x11\x28\x81\x98\x4c\x91\xc0\xce\x51\xcc\xf2\x09\xda\x6e\xd1\x86\xfb\xa3\xb0\x3c\x08\x2d\x11\xb9\x9b\x65\x69\x6e\xb0\x04\x84\xf1\x3b\xe6\x22\xf1\xb8\x20\x79\xfb\xfb\xe7\x35\xdc\x68\x60\x8c\xca\xc8\x51\x11\x3e\x26\x31\xb6\x9d\xb7\xaa\x6b\x8c\xfd\x57\xff\x99\x0d\x29\xc6\xc2\xd5\x55\x51\xf7\x3a\x6e\xaa\x1c\xaf\xad\x35\x59\x75\xc6\x97\x57\x5e\x66\xf9\x10\xeb\xdb\x12\xdb\x0d\xba\xf0\xa7\x50\xfc\x9d\x5d\x5d\x3a\xc6\x18\xf1\x7b\xe6\x52\x49\x7f\xac\x8c\xfe\x48\x0d\x7d\x7b\x01\x7d\x4b\xf5\x8a\xa0\xb9\x52\x9a\xf9\xf8\xf5\x99\x9f\x2f\x7a\x3a\xcd\xdd\xf9\xdf\x10\x11\x3a\xbb\x2d\x19\xa7\xfe\xd2\x59\x22\x8c\x59\xfb\xab\xb7\xcc\x0b\x6b\x99\x97\xe5\x42\x9d\x18\x57\x77\xc1\x33\x2b\xd2\xa3\x61\x99\xee\x9f\x8b\x3f\xc0\x55\x7a\xe0\xa2\x89\x8f\xe5\x7e\xac\xcd\xa9\x0f\x7a\x4e\x39\x37\x7e\x23\x23\x07\x75\x94\x41\x5b\xe3\xf8\x73\x55\x2b\x3e\xfb\xe7\xef\x9f\xfe\xb1\x70\x91\x99\xbe\xa4\x80\x92\xfa\x6e\x4e\xad\x5a\x03\xdd\xc6\xfa\xcc\x1a\x3c\x41\x9b\x83\xea\x4c\xaf\xd3\x7d\xa7\x9f\x3e\xd8\xfb\xa9\xcf\x27\x64\xae\xba\x66\x6f\x21\x46\x92\xe2\x86\x47\x32\x02\x7f\xb6\x96\xf9\x3f\xfa\x1a\x40\xa2\x64\x66\x45\xa3\x39\x4b\x3b\x0b\xd6\x46\x50\x37\x9c\x88\xdc\x91\xa8\x7a\x32\xb4\x48\x5d\x45\x0e\x08\x33\x14\xe6\x99\x23\xd8\x49\xf2\xb2\xcd\xfa\x16\x7e\x3f\xd6\x0a\xeb\x5d\x22\x53\x28\xea\x83\x58\xc4\xd5\x20\x75\x69\xb0\xbf\xd5\x72\x75\xb5\xba\x6a\xf6\x63\xc7\xaf\x92\xa7\xf4\xb7\x78\xab\x89\xca\xcb\xbe\xfd\x84\xe9\x03\xcf\xfa\xb2\x36\x24\xd6\xb6\xe8\xa1\xbe\x46\xd1\x06\x77\x98\x77\xf7\xae\xf1\x16\xb3\xb5\x79\x78\x00\x88\x3c\x14\xef\xba\x6c\xf1\x62\x12\x6c\x28\x2a\xb2\xd2\x18\x2e\x03\x6f\xb8\xc2\x97\x59\x42\x58\xa6\xb0\x1c\xf2\xa9\x32\x76\x6d\x8b\x68\x26\x71\xc8\x59\x34\xaa\xe5\x66\xd1\xaf\x85\xab\x99\xe2\x69\x6d\xd4\xac\x06\xe0\xc3\xd5\x2c\x9f\x25\xcd\xb9\x52\x13\xff\xb5\xec\xc2\x5b\xfc\x10\x0b\x94\x60\x69\xba\xc5\x88\xfb\xe7\xe5\x1b\xa9\xc3\x84\x00\x96\x3d\xbb\x57\xd5\xae\x9e\xec\xbd\xdb\xa5\x55\xdd\xca\x2f\x6c\x92\xbf\xb0\xc1\x17\xb5\x76\xf0\x08\xae\x91\x9c\x9a\x2f\x03\x73\x59\x65\x82\x1d\x30\x98\x47\xc4\xdd\x24\x07\xf1\x69\xce\xb4\xd2\xd9\xe1\x69\x01\x08\xae\x90\xc2\x7f\x5f\xdc\xf8\x11\xde\x5e\xf6\xe6\xd1\x72\xaa\x3c\xd4\x8f\xd7\x5e\x17\x1a\x34\xd0\xcd\xfe\xc0\x91\xb6\x71\x3e\x66\x9a\x1f\x6a\xb8\xf7\x40\xcb\xac\xdd\xe6\x61\x06\x47\xad\x1d\x4c\x6b\xaa\xc7\x96\x1b\x7a\x97\xee\xc8\x1d\x10\x1d\x97\xb2\x22\x76\x79\x00\x64\x8f\x42\xac\x77\x33\x8e\xd9\xae\x13\x23\x08\x74\xc7\x6c\xd7\x7c\x1b\xb5\x5d\xe3\x56\x14\xba\xa2\xb6\x63\x78\xc4\xea\x41\x26\x1d\xec\x0d\xd2\xfb\x21\xde\xdb\xe1\xe2\xf3\x09\x7c\xff\xf2\xe9\xf3\xb7\x1f\xc7\x11\x7c\xdd\x09\x8d\xff\xef\xa8\xb3\x3c\x53\xae\x1c\xa4\xde\x56\xf3\x34\xcc\xd4\x85\xe2\xe9\xb4\x43\xc8\x16\x97\xe0\x9f\x4f\x73\x5a\x64\x72\x70\xd9\xa1\x8c\x25\xfa\x2a\x45\xb9\xe2\xa9\x61\x71\x60\xb4\x9a\x9f\x72\x9e\x31\xcd\x99\x96\x27\x15\x4f\x2f\xb4\xea\xae\xc5\xcb\x86\x8c\x36\x5a\x38\x70\xd9\x35\x5e\xda\x72\x73\xbe\xeb\x3b\x3f\x5d\x1a\x12\x79\x32\x5d\xcf\x51\x55\x59\x73\x8b\xad\x67\xce\xc9\x1d\xbf\x77\x1b\x73\xc3\x92\xa3\x42\xc7\xd2\x8d\x57\x90\xbe\x82\x54\xa3\x6a\x47\x22\x5c\x32\xbe\xd5\x38\x52\x47\xc4\xdd\x24\xcb\x31\xfb\xf9\xdb\x8f\xaf\xc7\x8d\x58\x43\x96\x8f\xc0\x6f\x9e\x8f\x7e\x5c\xfd\x37\xf8\xcd\x91\xd9\x03\xd9\x1e\x60\x8f\xe0\x77\xd4\x79\x67\xfc\x1e\x93\x86\x21\xfc\x1e\xa1\xab\x07\xbf\x6e\xb1\xfd\xdf\x86\xec\x2b\x4c\x5f\x61\xfa\x82\x63\x76\xe9\xf8\xfe\x31\x20\x6f\xd1\x70\x9e\xa2\xd0\xfc\x6a\x8d\xdf\x4b\xbc\xab\xed\xc5\xe0\x83\xe6\x54\x5a\xf9\xe9\x75\x9e\xde\xc9\x4e\x4e\x4d\x1d\x86\x14\x2f\xec\x8a\xfd\xc1\x9c\x9e\x4a\x2b\x3f\xbd\xce\xd3\x2a\xfc\x37\x00\x00\xff\xff\xd4\x59\xd0\x44\x85\x27\x00\x00")

func webStaticFlipclockCssBytes() ([]byte, error) {
	return bindataRead(
		_webStaticFlipclockCss,
		"web/static/flipclock.css",
	)
}

func webStaticFlipclockCss() (*asset, error) {
	bytes, err := webStaticFlipclockCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/flipclock.css", size: 10117, mode: os.FileMode(438), modTime: time.Unix(1467867549, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticFlipclockMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x7b\x4d\x93\x1b\xb9\x79\xff\xfd\x5f\xf5\xff\x0e\x64\x5b\xa1\xba\x97\x60\x4f\x73\x2c\xef\x96\xc9\xc1\xb0\x64\x69\xd7\x5e\x47\xf2\x2a\x96\xec\x54\x42\x33\x5b\xe8\x6e\x34\x89\x55\x4f\x37\xab\x5f\x66\x34\x19\xf2\x60\xbb\xfc\x52\xbe\xe4\x90\x4f\x90\x8a\x93\xad\x8a\x0f\x79\xb1\xb7\x2a\x4e\xec\x5c\x36\xbe\x8f\x6e\xa3\x6f\x93\xc2\x5b\x37\xd0\x00\x39\x1c\x79\x37\xa9\x52\x69\x48\xe0\x01\xf0\xe0\x87\xe7\x1d\xe0\xd1\x3b\xfd\x5e\x92\x92\x75\x94\xe6\xd1\xcb\xde\x71\x30\x7e\x30\x1a\x1f\x8f\xc6\xc7\xbd\x77\x8e\xfe\xff\xff\x3b\x47\x45\xef\x1b\xa8\xc4\x30\xa9\xb3\xa8\x22\x79\xe6\x7a\x57\xdb\x29\x6d\xf1\xf1\xab\x0a\x67\x71\xdb\x81\x40\xe8\x5d\x39\x75\x89\x7b\x65\x55\x90\xa8\x72\xa6\x74\x70\x04\x19\xf1\xba\xc8\xab\xbc\xba\x5c\xcb\x61\x7c\x8a\x8f\x65\x33\xc9\x96\xb0\x1f\xb0\x01\x31\xcc\xf0\x45\xaf\x5a\x91\x72\x1a\xf9\x11\x4a\x53\x37\x06\xc8\x03\xb1\x1f\x76\xd9\x00\x31\x4e\x71\x85\x7b\xc6\x54\x6c\x1e\x0c\x63\x3f\xca\xb3\xb2\x2a\xea\xa8\xca\x0b\x90\xe8\xdf\xd5\x99\x48\xe2\xf6\x8d\x49\x3c\x92\xb8\x94\x0b\xff\xe3\x66\x14\xc9\x96\x9b\x0d\x6b\x53\x27\x82\x89\x67\xd2\xc1\x7e\x00\xb0\x8f\xd6\xeb\xf4\x92\xcd\x02\x50\xb1\xac\xcf\x70\x56\x95\x9e\x64\xdb\x1c\x34\xc5\x69\x89\x7b\x24\x71\xb3\x3a\x4d\xfb\x10\x36\x63\xe6\xc1\xc2\x2b\x70\x55\x17\x99\xab\xb6\x09\x2c\x37\x9b\xc8\xe3\x48\xa9\x9d\x20\xf6\xb6\x53\x3e\xa8\x97\xf8\x28\x8b\x70\x49\xb9\x65\xcc\x24\xf2\xf0\x18\x0f\xfc\x33\x48\xfc\x24\x2f\xde\x47\xd1\x8a\xb7\x8a\x2f\x20\xf1\xc9\xd9\x3a\xc5\x74\x5a\xde\xd1\x7c\x05\x49\x7b\xac\x90\x8e\xaf\xf2\xe7\x55\x41\x37\xcf\xe8\xe4\x37\x90\xf8\xe7\x28\xad\xf1\x47\x89\x22\x2b\xde\x15\x67\xcd\xc9\xc3\x4f\x70\x54\x39\x10\xa2\x59\x32\xc1\x92\xd2\xf5\xb6\x40\x9c\x7e\x02\x42\x0f\x38\x72\xa4\x03\x21\x5d\x2e\x4f\x7a\x89\x4f\x32\x52\x0d\x06\xfc\xaf\xeb\x81\x64\x0b\x74\x51\x83\x57\x7c\x67\x13\x5d\x44\x49\xd2\xe2\xe4\xa7\x38\x5b\x56\xab\xd3\xb1\x77\xc5\x85\x95\x72\x3e\x47\x8b\x29\x49\xdc\x68\x30\xb0\x2c\x1b\x0e\x06\x6e\x3f\x92\x7c\x6e\x36\x51\xcb\x72\x1f\x86\xed\x17\x6f\x30\x38\xfa\x41\x48\x65\xf6\x07\xe1\x91\x5f\xe1\xb2\x72\x43\x8f\xaf\x12\xab\x74\xd3\x50\x95\x44\xda\x8d\x38\x7c\x74\xe8\x66\xd3\x51\x1e\xda\x38\x6d\xba\x61\xc4\x24\x3d\x84\xb1\x5d\xd2\xe4\xf1\xb7\x03\x10\x08\xb7\x20\x3c\xf0\x3c\xc2\x49\x4c\x89\x9b\x53\x65\xbc\xc8\x6f\x5b\x81\x14\x0c\xb7\x52\x6a\x11\xe7\x1f\x1f\xaa\xf1\x9b\x8d\x05\x5f\x3a\xed\x66\xe3\x62\x55\x34\x37\x1b\xec\x4d\x93\xbc\x70\xe9\xf4\x09\xbc\xaa\xf2\xe7\x79\x5d\x44\x78\x42\xf5\x64\x0b\x96\x70\xee\x28\x0a\xe9\x00\x47\x32\xe9\x00\x47\x6c\xd5\x59\x80\x15\x34\x38\x98\x05\x93\xf1\x94\xc0\xe5\x7c\x35\x1c\x2e\xa6\x1e\x9a\x93\x45\x1f\x26\x73\xb2\x18\x0c\x30\x97\x3e\x86\x28\x01\xb4\xa7\x65\x81\xf4\x48\xd6\x43\x1e\x25\xdc\x6c\x2c\x84\x5b\x05\xf7\x2d\x97\x4a\xa8\x98\x4c\xf7\x4a\xe1\x76\xa2\x9c\xbe\xb2\x65\x4d\x95\xbd\xed\x16\x5c\x49\x0d\x9e\x7c\xc4\x4e\x08\x9c\xe3\xa2\x24\x79\x36\x71\xc6\xfe\xd8\x01\x42\x5b\x35\x51\x07\x91\x77\x25\x59\x8e\x39\xcb\xe7\x39\x89\x7b\x01\x84\x1c\xdd\xe6\x88\xe6\xf1\x62\x30\x08\xf9\x46\x22\x80\xe6\xf1\x02\x50\xa3\xbb\x05\x8d\xae\xab\x6c\xca\x39\x11\x0c\xa6\xe8\xa4\xab\x4b\x53\x34\x1c\x7a\x96\x83\x6d\x37\x84\x16\x33\xf5\x8b\xab\xf3\xe2\x4d\xf4\xef\x26\x20\x68\xa1\x49\xf6\x16\xc8\xe3\x56\x99\x14\x04\xbc\x83\x2f\xd1\xea\xe6\x76\xeb\x31\xcd\xf9\x20\x25\xeb\x47\xd4\xed\x4d\xfb\xaa\x2a\x68\x4e\xac\xa1\x81\x5d\x6c\xc5\x12\x61\x8f\x64\x65\x45\x8f\x27\x4f\x7a\xfc\x70\x06\x03\xad\xf1\x31\xaa\x30\x84\xfd\xf1\x60\xe0\x46\x30\x04\x21\x0c\x3c\x40\x9d\x5c\x33\xb5\xff\x01\xa2\xb2\x70\x29\x66\xde\x82\xb6\xe7\x09\xca\x96\xf0\x4a\x6d\x31\xc5\x29\xac\x49\x1a\xd3\x45\x26\x4e\xeb\xbe\x9d\x56\x42\x02\xff\x3d\xff\x3d\x07\x58\xa5\x8e\xed\x44\x6a\x7d\x5f\xb5\x72\x21\xbc\xda\x7a\xc0\xe8\x8a\xd8\x2e\x68\x17\x03\xb5\xc4\xd5\x47\x6b\x3a\x53\xe9\x22\xc9\x50\x3f\x00\x57\x5b\xb6\x13\x6a\xc6\x51\x9a\x86\x28\x7a\x39\x51\x11\x26\x89\x6b\x13\x91\x56\xb6\x42\x38\x5f\x80\x08\x8e\xa7\xd1\x09\x34\x24\x2c\x1a\x0e\xbd\x56\x1c\x22\x26\xbb\xeb\xba\x5c\xb9\x6a\xa3\x37\x45\xaa\x55\x0c\xa9\x12\xa5\xf9\x52\x63\xe3\x82\x64\x71\x7e\xc1\xbc\x79\x9e\xe2\xc1\x40\x7c\xf0\xd3\x7c\xa9\x7d\x71\xa9\x2e\x2c\xe5\x4e\x27\xa6\xd9\xec\x09\x6b\x38\x13\x7f\x27\xfd\xb1\x32\xa0\xb4\x08\x26\x97\xdc\xd2\x32\x27\xf5\x51\xad\x71\x6d\x49\x4a\x6d\xdd\x06\x28\xae\xd8\x4e\x9d\xc5\x38\x21\x19\x8e\xdb\x93\x42\xf3\x70\x31\x18\xe8\xa7\xe4\x86\x80\x36\x53\x0d\xd8\xba\x9f\xfc\x59\x8d\x8b\x4b\x0f\xdc\x2e\xfb\x54\x40\x31\xd4\x65\xb0\x11\x3f\x54\x53\x05\x44\x45\x35\xe9\x07\x20\x26\xe7\x24\xc6\x45\x39\x99\x2f\x40\xc2\xa5\x7a\xd2\x1f\x83\x94\x94\x15\x6b\xb3\x0a\x61\xb3\x67\x5f\x0e\xa7\xc7\xcf\x1a\xd8\xc0\xe6\x1b\x75\x62\x6e\x28\x44\x4f\x4c\x0f\xd1\x16\x30\x0d\x30\x4c\x69\xc3\x98\x44\x81\x7e\x66\xa1\x45\x81\x51\x85\x1f\xf3\xb5\x34\x5d\x00\xb1\x77\xe5\x84\x79\x9e\x62\x94\x69\x32\x1f\x6d\x36\x6e\x0c\x23\x10\xc1\xd0\x13\xf1\xe5\xfc\xfe\x49\xb9\x46\x59\x2f\x4a\x51\x59\x42\xe7\xfe\x50\xe5\xca\x67\xad\xb8\xf4\xe3\xbc\x1a\xde\xef\x55\xf9\xda\x39\x3d\x39\xa2\xf4\xa7\xf7\xc1\x1d\x06\x86\x79\x55\xe5\x67\xed\xd8\x85\xff\x49\x4e\x32\xd7\x71\xbc\x69\x3c\x18\xb8\x18\x3a\x8e\x07\x42\xa8\xcd\x90\xe6\x11\x4a\xc9\x5f\x53\xa4\xa6\xdc\x65\x1e\xc8\x2a\x07\x64\xe8\xf4\x9c\xa1\x1b\xcd\xa2\x89\xe3\x78\x7e\x95\x3f\xc9\x2f\x70\xf1\x88\x22\xef\x0d\xef\x3b\x87\xb2\x9f\xa2\x10\xa7\x8c\x7e\xe8\x86\xb3\x90\xce\x35\x74\xc4\x2e\x1c\x80\x41\xf3\x79\x01\x96\x10\xb9\x49\xb3\x2f\x3d\x6c\x91\x02\xc1\x15\x7c\xe9\x81\xa5\x3c\xbe\x27\xa4\xac\x3a\x22\xd4\x46\x2f\x52\x05\x98\x1d\x43\x00\xc1\xc0\x13\xe9\x88\x6e\x78\xe9\x24\xae\xca\x3f\xa0\x13\x69\x1c\x30\x09\xe4\xcb\x47\x1e\x88\xb6\xa0\xc0\x25\xae\x0c\x61\x93\xfb\xaf\xc8\x19\xee\x2c\xf2\x82\x9c\x61\x7d\x11\x6d\x44\x5e\x90\x25\xc9\x50\x3a\x7b\x8a\xaa\x95\x5f\xe4\x75\x16\xbb\x56\x02\x6f\x12\x80\xab\x33\x92\x91\xb3\xfa\xec\x31\x59\x92\xaa\x9c\x68\x74\x5a\x97\x34\xd0\x34\xb5\xb3\x4f\x07\xfa\x63\x6f\x0b\xd0\x7a\x8d\xb3\x98\x8d\x79\x91\x33\x7e\x35\x3b\x83\xfc\x7b\x38\xf5\x39\x91\xcb\x07\xc4\x9c\x5a\x23\xe3\x36\x9b\xe7\x46\xcd\xe9\xb8\x08\x5c\x09\x71\x98\x5c\xa1\xa8\x22\xe7\x78\x62\x95\x15\xde\x07\x42\x9c\xe4\xc5\x0e\x12\xde\x07\xe8\x6e\xec\x04\xb4\x87\xfa\x75\xae\xf9\xc6\xa6\xdc\xd0\xdb\x02\x66\x00\x26\x5a\x0e\x59\x56\xf9\x5a\x6f\xa1\x56\xe3\xc3\x2c\x2a\x8c\xb8\x47\x5f\x36\xaf\xb3\x2a\xce\x2f\xb2\x19\x17\x53\x2c\x06\xb8\x22\x7e\x21\x59\xd3\xb0\x05\xe4\xd6\xe9\xa8\xd4\xf8\x28\x8e\x9f\xe3\x28\xcf\x62\x3a\xa6\x99\x51\x1d\x13\x40\x68\x0e\x5b\xe2\x8a\x4a\x18\x1f\x5a\xba\xde\x4c\x23\xa1\x1b\x94\x4c\x69\xc3\xca\x3a\x6c\x57\x63\xb8\xea\x41\x01\x4f\x55\x58\x22\x8e\x7c\x8c\xa2\x95\x1b\x02\x5d\xdd\x64\x9a\xcd\x34\x84\x66\x4d\x78\xe6\x46\x9b\x4d\x08\x21\xf6\x63\x8a\x3d\x8d\x8e\xd7\x29\xba\x74\x3d\x80\xfd\x12\xa7\x38\x62\x59\xd0\x24\xf6\xa5\x10\xd1\x63\x79\x1b\x67\xc4\x0c\xff\x2e\x7f\x94\x91\x33\x44\xa7\xf8\x2e\x0d\x89\xc6\xf8\xab\x40\xf3\x50\x32\x20\x29\x27\x57\x31\x2e\xab\x22\x67\x0e\x8a\x4b\x2d\xfd\x44\xd3\x49\xfe\xb7\xc2\xc5\x39\x4a\xe9\x67\x2e\x38\xec\x43\xbe\xa6\x7f\xb9\x0d\xa0\x5e\xbe\x2b\xe0\x0e\x85\x72\xc4\x8a\x29\x23\xde\xe4\x48\xc1\x56\xbb\x78\x93\x23\xfd\xa5\xd6\x27\xda\x1c\x10\xe7\x95\xde\x91\x57\x0e\x60\x86\x55\x6b\x66\x2d\x0e\x3f\x43\xd6\xee\x00\x8a\xfa\xc4\xa1\xff\x3b\xe0\xa2\xa0\xea\xa0\x2f\x21\xda\x1c\xca\x7f\x1e\xbd\xa4\xee\x7d\xe2\x7c\x2b\xaf\x8b\xf4\xf2\x11\x95\x6c\xba\x7a\x23\xe2\x74\xc3\x31\x4e\x50\x9d\x56\x8f\x76\x52\x0b\x02\x1a\xaf\xd6\x68\x89\x27\x0e\xce\x96\x29\x29\x57\x0e\xb8\x87\x19\x86\x09\x1d\xd5\x0f\x40\x8a\xb2\x25\x0b\x09\x4c\x4a\xdd\xb6\x05\x40\x5a\x2a\x86\x78\x9d\x65\x84\x8f\xa4\xe2\x2b\xff\x16\xf4\xc3\x3d\xb9\x45\x7a\x90\x3b\xa2\x5c\xea\xd9\x63\xe6\xc4\x9b\xd8\x55\x8f\x2d\xc4\x02\x90\x4e\xdc\xc4\x1a\xb1\x20\xbd\x87\x53\x88\xdc\xd0\xa3\x82\xfb\x88\x9e\x38\x37\xa9\xd2\xf8\x88\xf5\x25\xb1\xf8\x0a\xe5\x50\xde\x2c\x77\x03\xa3\x6e\x72\x30\x8b\x26\x91\x6a\xfd\x23\x6a\xe9\x79\x11\x65\x8f\x37\xd1\x67\xb5\x7a\x06\xad\x09\xe8\x8a\xc1\x0d\xa5\xda\x24\x71\x61\xb8\x5a\x56\x2d\xf8\xb2\x12\x93\x34\x47\xb1\x3c\x6e\x0e\x87\x3c\x52\x85\xa0\x11\x18\x09\x98\xf8\xda\xcc\xb2\x33\x48\xd3\x46\x4f\x4c\xcb\x13\x81\x18\x3a\xb4\xcf\x01\x18\xf6\xc7\xd2\x67\x23\x88\xfc\x3a\x4a\x48\x51\x56\xae\x37\x8c\x1b\x57\x8b\x99\x25\x1c\x0c\x5c\xb1\x0a\xb5\x8a\x74\x60\xd0\x9e\xb0\xbf\xaa\xce\x52\x1a\x82\xb4\x30\xe8\x08\x42\x0b\xa8\x51\x6b\x86\x68\x02\xa0\xa1\x26\xb3\x5b\x10\x7a\x13\xbd\x43\x38\x0d\x5d\xa7\x86\x71\x43\x0e\x22\x9f\xc5\xb4\x6e\x1b\xeb\x62\x18\x01\xac\x83\xd4\xf6\x71\xb8\x1a\xe5\x33\x5c\xb3\x44\x27\x84\x7a\x6a\x39\x57\xc0\x5a\xcc\xf6\xf4\x4d\xba\x7d\x26\xb5\x41\xa3\x6e\x52\xb2\x26\xc3\x7a\x9a\xd6\x86\x94\x6b\x1e\xaa\xda\xce\x17\x36\x94\x53\x92\xb8\x7d\x24\x4a\xa0\xbd\xac\x4e\x53\x51\x2a\x46\x7a\x78\x3a\xed\x16\xb1\x94\x74\x96\x06\xee\x80\x06\xf2\xf3\x78\x31\xa3\xff\x4d\x90\x11\x10\x68\x81\xcc\x34\x94\x26\x61\xb3\x09\x5b\x7f\x3f\x18\xf4\x5d\xed\x6b\xc8\x05\x85\xfe\x77\x1a\x78\xb3\x90\x25\x8c\xce\x8b\xe2\x92\x64\xcb\x5e\x95\xf7\xd8\x1a\x3d\xa6\x53\xbd\x8b\x15\xce\x7a\xcd\xd8\x1e\x4a\x0b\x8c\xe2\xcb\x1e\xaa\x7a\x81\xe3\x4d\xdc\x50\x0a\x2a\x3d\x5d\x3e\xaf\x07\xf8\xdf\x42\xb4\x2a\xc1\x40\xc8\x63\x3b\x6b\x75\x14\x0d\x06\xc8\xf5\xb6\x9e\xd7\x8d\x70\x50\x1b\x78\x60\xa9\x05\xad\xd2\xf3\x16\xd4\x96\xba\x58\x7a\xd9\x1a\x4b\x4f\x89\x88\x57\xa8\xfc\xe8\x22\x7b\x56\xe4\x6b\x5c\x54\x97\x6e\xe8\x09\xd9\xe4\xd1\x40\xb8\x10\xb3\x1b\xf1\xb2\xe4\x80\x2f\xc8\x3a\x5d\xa4\x88\xb2\x68\xf2\x58\xea\x4b\xcd\x8e\x7d\x28\x37\x8d\x48\x09\x72\xfb\x01\xcf\xd2\xf5\x31\x5a\xc6\xcd\x06\xb1\x89\x1f\x35\xce\x6d\x47\x18\x2b\xce\x7e\x2a\xea\xfd\x82\x1a\xa2\x59\x3f\xa0\x6e\x26\xec\x58\x12\x55\x2d\x8d\x78\x4a\x03\x9d\xf3\x3a\x06\xe8\xce\x71\x0f\x8d\xa9\x77\x05\x3d\x2c\xde\x9a\x04\x7f\x6c\x38\xa2\xc4\x10\x5b\x35\x7f\x17\x3e\xfc\x5e\x1e\x7e\xc2\x62\xa3\x0a\x9f\xb1\x74\x3e\x45\x65\xf5\x58\x2e\xbd\x3f\xb5\x6f\xd2\x75\x20\x12\xbb\x25\xa9\x60\x28\xcd\x81\x98\x46\x36\x50\x7f\xdb\xcd\x26\xa4\xa1\xce\xc3\x4f\x5a\xcf\x1a\x9e\x06\x4d\x8d\x43\xc4\x97\x7a\x76\xa5\xa4\x2f\x72\x10\x13\x2d\x4a\x6c\xd4\xa4\xda\x2a\x4a\xab\x49\x33\x51\x90\xe7\x08\x2b\xbc\xab\x1b\xe9\x43\x7d\x1f\x9a\x28\x51\x16\x12\x92\xc5\xae\xe3\x3b\x43\x2d\x6e\xe0\xb0\x7b\x7e\x81\xcf\xf2\x73\x6c\x09\x2c\x04\xc1\xf4\x96\x79\xf8\xc9\xee\x99\x47\x12\xd8\xc3\x17\xb1\x0a\x50\x72\x27\x8a\xf8\x87\x15\x3e\xb3\x4d\xa3\xec\x9b\x9a\x27\xbe\xa8\x3c\x9e\xf6\x28\x5b\xaa\xed\x96\x07\xa4\xdd\xec\x87\x9d\x8d\xc6\x51\x37\xab\xa3\xc3\x0c\x1b\xa6\xde\x93\x4c\x85\x91\xc8\x6b\xcd\x36\xf2\xb4\x55\x85\x03\xed\x9a\xdb\xc8\x8c\x0a\x5f\x86\xfe\x9e\x5a\x6b\xa0\x68\x74\xe4\x9a\x5b\x96\xf9\xfd\x93\x94\xb4\xd5\x10\x17\xcd\x10\xab\x76\x88\x5a\x09\xea\xad\x0a\x9c\x40\xe7\x2b\xfc\x6b\x4c\xce\x25\x6d\xbd\x36\x9a\xca\x15\x8a\xf3\x0b\xe7\xf4\xe4\x28\x26\xe7\xdd\x4e\x92\x65\xdd\x72\x0a\xa5\x72\x40\xf3\x41\x23\xa7\xf6\xea\x8b\x5e\xc0\x39\x39\x42\xfc\x4f\x4a\x4e\x9d\xb6\x18\x25\xeb\x08\x3b\x80\x52\xfc\xb9\x0e\x28\xeb\x6e\xa5\x5b\x28\x6a\xe4\xd9\x8b\x3c\x9a\x56\x2d\x71\xf5\xac\xc0\xe7\x3c\x8d\x14\x49\xaf\xde\xa6\x28\x2b\x88\x20\x72\xe7\xf7\x4f\xea\xb4\x53\xb6\x52\xcb\x07\xbc\xf0\xa5\xc9\x83\xf0\x02\xb3\x9d\xc2\xd9\x9e\xb4\x6d\x77\x16\x2d\x03\xd2\x40\xed\x23\xb5\xe8\x99\x73\x72\x54\xa7\x2a\xe4\x4d\x95\x2a\x62\x5e\xef\x3b\xf8\x55\xd5\xa9\xcb\x34\xae\xef\xeb\x50\x51\xc6\x59\xa0\xe0\x32\xe4\x75\xea\x06\x34\xcb\xe0\x40\x1b\xfc\x75\x65\xf0\x68\x7c\x98\x07\xe3\x97\x30\xca\x95\x8e\x08\x2b\xe1\x0e\x27\x5d\xd6\x61\x59\x15\x6e\x00\xc6\x9e\x5f\xe5\xdf\xa3\xa9\x94\x28\x3f\xaa\xdd\xac\x16\xe5\x27\x99\x6f\xb9\xa1\x51\xef\x67\xb4\xc0\xdb\x45\x0c\x68\x4f\xdc\xb3\xb0\xf1\xc9\xfe\xf1\xfa\x1a\xac\x6b\x7b\x27\xb7\x4d\xcd\xd3\x2e\xb7\xcd\x52\xd9\x40\xf5\xb4\xdd\x1c\x78\x97\x57\xb5\xde\xdb\xb4\x97\x33\x91\x9e\xa6\x6c\x36\x6e\xa7\x05\x22\xfd\xbb\xa7\x64\xbd\x51\xb7\xc2\xde\x86\x3b\x2c\xea\x0a\xa9\x7e\xe6\xd9\x39\x2e\xb8\xc8\x95\x2f\xf2\x87\x45\xa1\x1a\xf8\x26\x96\x9a\x2f\xa6\x88\x85\xe9\xe2\x26\xae\x0d\x2d\x23\x18\x4c\xa3\x13\xa4\x5d\xea\xcc\xa3\x85\x7f\x86\xaa\x68\xe5\x1e\xfd\xd5\x0f\xe2\x77\xee\x1d\x2d\xbd\xf6\x72\x87\x5d\xea\xc8\x3c\x66\x0b\xe2\x7d\x45\xc8\x76\x3d\x10\xc1\x50\xae\x21\x07\xcf\xa3\x11\x5a\xcc\xf8\x1f\x56\xc4\x61\x73\x69\x69\x48\x63\xb4\xe6\xec\x31\xc0\x8e\x0a\x58\x08\x43\x75\xa9\x31\x6c\xd6\x62\x85\x67\x27\x70\x86\x61\xbb\xe3\x18\x06\xd3\xf8\xa4\xe1\x26\x1e\x0e\xbd\x88\xef\x2d\xf4\xa3\x15\x2a\x1e\x56\x6e\xec\x79\xec\xf0\xc4\xbb\x04\x33\xdf\x94\xe7\xa0\x9f\xa5\x1c\x20\xce\x4d\xeb\x3c\x6d\x3a\x5b\x3e\xa2\x86\x87\x13\x73\x80\x60\xac\xce\xca\x15\x49\x2a\xd7\x09\x9c\x8e\xa9\x79\x8c\x2a\xcc\x2f\x3a\xf7\x85\xd9\x46\x91\xa3\xe9\x61\xb9\x30\x6d\x72\x5d\xf9\xc9\x93\x55\x4c\xd7\x1b\x8e\xf1\x57\xdf\x91\xd6\x5c\x2b\x6c\x7a\x62\x75\x59\x74\xb2\x89\x9b\x1c\xf8\x18\x5d\x96\x32\x24\x59\xe2\xea\x5b\x79\x5d\x94\x6e\x93\xec\x2f\x71\xf5\x94\x64\x75\x85\x59\xdb\xa2\xa9\x1d\x34\xc2\x26\xa9\xe4\xda\xfd\xc0\xf3\x14\x73\xcc\x2f\x57\x24\x3b\xe5\x2e\x29\xec\x6e\xe0\xe8\xdd\x80\xfe\x3b\x7e\xa0\x2c\xe8\x86\x7f\x02\xdf\xf3\x00\xab\xfe\x24\x69\x9e\x17\x72\x62\xca\xb2\xb1\x51\xed\x89\x48\xc3\xca\x5c\xdf\xa6\x75\x93\xc0\xb2\xa5\x56\x9f\x50\xb3\x64\x7a\xb9\xeb\x50\x75\x9e\xdc\x96\xcb\x3b\xee\xbf\xb3\xf9\xe3\x07\xb6\xdd\x3f\x25\x29\xa9\x50\x71\xd9\x49\x00\xf9\x1b\x33\x33\x4a\xe7\x17\xd6\x74\x9b\x68\xb3\x71\x11\x6c\xc5\x40\x8a\xaa\xeb\xc9\x1b\xa0\x39\x52\xa1\x42\x2a\x4e\xad\x28\x84\x10\xc2\x7e\x30\x18\x08\x05\x45\x2a\x72\x86\x28\x44\x92\x67\x36\xcb\x1d\xd0\xe8\x40\xf1\x6e\x60\x87\x82\x4e\x7b\x57\x51\x68\xb6\x74\xc8\xc1\x2b\x7c\xd9\x6e\xb7\x39\xa6\x8d\xae\x82\x5b\x74\xdc\xbc\x23\x61\x9b\x3a\x43\xaf\x5a\x1f\xd2\xea\xfb\xd1\x18\x7f\x75\x84\xf4\xef\x20\xf0\x26\x9d\xa6\xd1\x8e\xa1\x13\x25\xaf\x37\x92\xff\x3f\x4e\x5e\x80\xfe\x04\x40\x88\x8f\x26\x3d\x31\x9c\x47\xa7\xe3\xe3\x59\x34\x1a\x1f\x4f\x02\x08\x61\x34\x1b\x1f\x4f\xa2\xdb\xa4\x2a\x3e\x44\xaa\x62\x7e\xfc\xb6\x73\xd9\x23\x55\xaa\x48\xbd\x1b\x40\x18\xce\x42\x18\x4c\x54\xe1\x8a\x30\x49\xa5\x6c\xfd\x39\xc6\x2f\xef\x6e\xbe\x8e\xde\xeb\x08\xee\xd7\x8e\x0d\xc1\xe5\xb9\xd7\x13\x8c\x62\x92\x2d\xff\x12\x17\x79\x69\xbd\x5f\x0a\x00\xa6\x2e\xb6\x89\xb5\x98\x9b\x8d\xb4\xb0\x2a\x3c\x45\xb3\x78\x08\xd7\xa8\x28\xf1\x87\x59\xe5\x46\x73\xb4\x00\xe3\xc0\x9b\x60\x71\x2b\x3b\x47\x0b\xea\x35\x29\xfc\xf1\x0c\x4f\x22\x76\x47\x69\x83\xed\x00\xdf\xe4\x97\xed\x89\x68\x12\xd7\xa0\x30\x44\x5e\x2b\x73\x43\x88\x94\xd5\xcc\x17\x08\x0d\x1f\x2c\x5c\x6d\xae\xdc\xbe\x04\xb6\x46\x2a\x5b\x23\xca\x56\xb3\x9a\xc1\x56\xcb\x07\x63\x6b\xcf\xfb\x29\xab\x1c\x28\x21\x8f\x3d\x01\xd8\x1b\x0a\x17\xbb\x62\xe1\x2f\xfa\x62\x8e\x5a\x1f\x3d\xb6\x6e\x46\xb2\x4b\x41\xe3\x9a\xf0\x96\x02\x96\xfd\xf1\x89\x48\xe6\x04\xef\xae\xf6\xad\x64\x6f\x53\xbd\xfd\x24\x7c\x83\x5c\x1d\xdf\x4f\xd1\xba\xc4\xf1\xae\x83\x60\x5b\xe2\xb1\x91\xdc\x8a\x3a\x6c\x47\xe1\xb3\x89\xb6\x5a\xa9\x95\xc7\x2a\x06\xb2\xd0\xca\x2c\xd3\x46\x29\x46\xc5\x87\x62\xa1\x76\xb4\xbc\xe6\x62\xdc\x40\x71\x51\xf5\x71\x89\xab\x86\x14\xdd\xb2\x65\xb6\x94\x79\x1f\x8f\x3a\x97\xe2\xcd\x8d\x9c\x5c\x83\x43\xc5\xaf\xa2\x6e\x5b\x83\xcd\xbd\xaf\xfc\xad\x2f\x31\x96\x4b\x68\x7b\xbe\x7d\x91\x7c\xdd\x25\x41\xde\x16\xe8\xd3\xa8\x67\xb2\x13\x53\x3a\xa8\xdd\x9e\xc9\x71\x07\xe0\x2d\xf8\x58\x28\xc9\x2e\x7a\x7d\x25\xed\x6a\x4f\xee\x56\xe1\x79\xff\x3e\xc5\x52\x74\x55\x62\x6c\x4a\xae\xb8\x5b\x07\x44\x35\x6d\xd7\xa2\x54\x8e\x86\xc3\x2d\x50\xb7\xb8\xfb\x52\xa6\xe1\x80\x4e\x20\x2e\x47\xa0\x0a\x8e\x76\x45\xa2\x52\x8b\x5a\x5f\x5b\xdd\xdb\x55\xbd\x50\x0c\xd6\x05\xce\xaa\xcb\x0f\xf2\xba\x60\xa1\xaf\xbc\xa9\x83\xfa\xcb\xb8\xfd\x0f\x7b\x3b\x16\x84\x7e\x35\x5e\xac\x69\x35\x32\x10\xeb\xaf\x3c\xda\xba\x6f\x9d\x3a\xa2\x5a\xa6\x3d\xe5\xa0\xff\x6d\x36\xf6\x57\x3e\x4d\xd8\x66\x16\x3a\xdf\xee\x99\x92\x88\xa3\x43\x1a\x59\x58\x9e\x95\x74\x62\x77\x57\xca\x5a\xb9\xca\x2f\x84\x07\xf1\xa6\x32\x07\x3e\x8d\x9b\x74\x79\x47\x86\x1d\xa9\x37\x00\xa1\xb7\xd5\x8a\x67\xe2\x05\x9f\xbb\xa3\x91\xd7\x7a\x9a\x27\x64\xf3\x60\xe1\xf9\x24\x2b\x71\x51\x7d\x83\x95\xe2\xc4\xa5\x35\xbb\xaf\x6a\x3f\x0a\x8e\x46\xc7\x0b\x76\x57\x60\xcc\x32\xbe\xd3\x2c\x0f\xc4\x2c\xed\xf9\x9b\xb7\x43\x8d\x7c\x68\x4f\x8f\xe8\x06\x20\x9a\xa1\xb7\xc4\x18\x74\x24\xee\x76\x59\x17\x19\xc6\x3e\xf9\x2e\x57\x79\x9d\xc6\x0f\xb5\x17\x52\xbb\xde\x5c\x68\x2f\xf2\x2c\x2f\x8b\x51\x7b\xf5\x0f\xc3\xf6\xb3\xbc\x5f\x53\x1f\x06\xb8\x72\x77\xdd\xc5\x79\x04\xdf\x3e\xb6\x82\xda\x15\x40\x7b\x6b\xd7\x1f\x03\xe4\x8b\x9b\x02\x57\x49\x2c\x8c\xd0\x66\xc8\xab\x8a\xcd\xe3\xab\x9d\x13\x06\x42\x09\xf6\x4d\x36\x65\x97\x53\xed\xc2\xe1\x88\xcf\x5e\xe2\xea\xfb\x28\xad\x95\x5f\x33\x85\x74\xf6\x86\x4c\x12\x89\x03\xd9\x43\x76\x8b\x55\x51\x2d\x27\x88\xf6\x19\x95\xae\xc9\x69\x37\xd5\xc9\x2c\x4d\x02\x2a\x91\x0b\x6f\x2a\x35\xb9\x29\x37\x35\x2a\x1d\xab\x2a\x1d\xb5\xbf\x81\x51\x14\x3b\xa2\xe3\xc5\x05\x5e\xe4\x31\xe1\x60\x43\x5b\x7d\xea\x16\xde\xc4\xbb\xb2\xed\x81\x9a\x65\x91\x1d\x71\x6d\x68\xe8\x5c\x9b\x0f\xbe\x35\x18\x86\xf2\xfd\x9f\xbe\x19\x55\x1f\x82\x1e\x62\x07\x1e\x23\xd2\xbc\xeb\xba\xc5\x18\x48\x6b\xc3\xde\xd5\xfd\x2f\xf8\x3e\x80\x61\x30\xdd\xe3\x79\xda\xda\xa0\x6b\x1a\xc4\x3f\xda\xe9\x28\xb3\xcd\x84\x57\xd0\x5d\x8e\x23\x7a\x1d\xef\x2d\x1c\xcd\x04\xc3\x63\x60\x9f\x56\x14\x12\xee\x36\xed\x83\x21\xd6\x3d\x58\x67\x52\x56\xc2\xb8\xdb\x94\xef\xde\x32\xe5\x63\x74\x59\x3a\xa0\x1f\xec\x99\x34\x38\xd0\x1f\x5a\xd4\xf0\x90\x43\x56\x1e\x95\xa9\x4a\x7d\x67\x5f\xa8\x3d\x6e\xfc\x72\x22\x3e\xfd\x8d\x2d\xc0\x7b\x63\xbe\x08\x46\xb3\xc8\x2e\xf3\x5a\x4d\xb6\xbd\x3a\xc0\x5d\x21\x8f\x3a\x42\x1e\x1b\x42\xfe\xc5\xcb\xf4\x17\x2a\xcf\x62\xca\x70\xb3\xd9\x27\xd0\x5f\xba\xf8\x69\x78\x1f\x26\x6f\xb7\x3d\xb2\x57\xc8\x9b\x22\x60\x13\xb3\x8a\xbd\x3c\x4c\x1a\x81\x57\x3b\x33\xfc\xaa\xe2\xbf\x68\xb3\xd6\x61\x5a\xa9\xd5\xaa\xc8\x1d\x89\x36\xa4\xbd\x15\x6f\x9a\x44\xbe\xff\x2a\xc2\x65\x29\x6e\x24\xf7\x86\x7b\xb7\xc6\x23\x2d\x81\x4c\xea\xcd\xc0\x56\xe1\xd3\xf6\x96\xea\x96\xe3\xe9\x8e\xbf\xbb\xee\xbf\xb8\xc0\xe9\x39\xde\x95\xef\xed\xc8\x08\x1b\xc4\xce\x70\x41\x62\x52\x9f\xb1\xbb\x5c\xf1\xf9\x05\x7e\x55\x4d\x9c\x87\x4f\x9d\x7d\x01\x9a\xf5\x51\xbf\x3d\x7d\x32\xcb\x51\xea\x4a\x4d\xf1\xf6\xa9\x68\x74\x3b\x44\xdd\xa7\x10\xca\x3b\x30\x49\xe2\x9c\xde\x07\xce\x49\x4a\xd8\x9b\x12\xf5\x05\xcb\xd0\x58\x6e\xa8\xbf\x09\xb1\x3c\x54\xd0\x57\x37\xe5\x79\x97\xc6\x8f\x17\xe2\xb5\xd6\xae\xa8\x4e\x65\xa3\x6f\xd9\x76\x73\x59\x7a\x17\x74\x84\xd9\x45\x8e\xc7\xdf\x03\x1b\x33\x68\x32\x75\xb7\x93\x03\xf2\x4a\x47\x0a\x89\x51\xad\xd3\x6f\x44\xc5\x35\xc3\x29\x1c\x1f\xcf\x9c\x67\x4f\x1d\x26\x44\x5b\x40\xca\x67\x4f\xcd\xa1\x94\x00\x5a\xb6\xc7\x33\x2a\x3a\xea\xa1\x6d\xd4\xc3\xbd\xa3\x6e\xb5\x2b\x4f\x50\xb6\xf4\x1f\x16\x28\x24\x11\xbc\xba\xc4\xa8\x28\x27\xce\xcd\xaf\xdf\xfc\xf4\xcd\xcf\x6f\x3e\xbd\xf9\x27\x07\x9c\xe5\x59\xb5\xa2\x6d\xbf\x79\xf3\xb3\x37\x3f\xbf\xf9\x57\x07\xc4\xe8\x92\x7e\xff\xfb\x37\xbf\xb8\xf9\xf4\xcd\x4f\x1c\xb0\x62\x57\x88\xce\xcd\xaf\x6f\x3e\xbd\xf9\xad\x18\x24\x6e\xd2\x9c\x9b\x7f\x7e\xf3\xa3\x9b\x4f\x6f\xfe\xf1\xcd\x8f\x1c\x50\x8a\x58\xd3\xb9\xf9\x15\x9d\xfc\xcd\x4f\xdf\xfc\xc2\xe9\xfe\x54\xd6\x47\x05\xb4\xf2\xd6\xa1\x9b\x3b\xa8\x18\xa1\xc2\x59\x1c\x44\xed\x23\xbe\x3f\x2b\xed\x21\xf8\x3c\x46\x19\x29\x57\x0d\x3e\x9f\xff\xa4\x68\x81\x79\xfa\xf9\x3f\x64\x98\xff\xdc\x83\x01\xf3\x18\x2d\x71\x03\x0a\x2b\xee\x29\x78\x30\xeb\xc6\x7e\x6f\xd1\xa0\xf1\x1c\xbf\xac\xb3\x98\xfd\x96\xa3\xb3\x6a\x8c\xba\x1c\x73\x3e\x0c\x2c\x62\x34\x8a\x5f\x9a\x58\x58\xa9\xfd\x98\xef\xc5\x4a\x7b\x08\x16\xdf\xc4\xc5\x19\xca\x1a\x2c\xbe\x8d\x56\x05\x56\xd0\xc8\x33\x54\x61\x89\xc5\x0b\x15\x8b\xe7\x15\xdd\x67\xd6\x45\x83\xb6\x74\xc1\xc8\x2c\x60\xe0\x2e\xcb\x9c\x11\x13\x0c\x3c\x8a\xb1\x09\x86\x95\xda\x5f\xf2\xcd\x58\x69\x0f\x01\xe3\x7d\xfe\xd3\x97\x06\x8d\xbf\xa0\x7f\x34\x34\xaa\x55\xd9\x4a\x06\x0d\xaf\x05\x1a\x22\xda\xe9\x60\x51\x6a\x58\xf0\xc0\xcd\x80\x02\x1b\x0c\x0b\x36\x0c\x2c\x70\x36\xaa\x4b\x13\x0b\x3b\xb9\x2f\x7e\xc7\xb3\x83\xfa\x10\x38\x9e\xaf\x75\x45\x79\x38\xf8\xca\xf1\x83\xf1\x34\x57\x21\xc1\x25\x6e\x11\x19\x7c\xe5\x38\xf8\xda\x14\xa9\xb0\x14\xc8\x80\x25\xd7\x61\x59\xd6\x59\x9c\x5b\x60\x29\xbb\x9c\x0b\x76\x4c\x58\xca\x11\xb6\xc0\x62\x27\xf7\xcb\xb5\x55\x63\x04\xf5\x21\xb0\x7c\x40\x32\x0d\x96\xef\xd7\x79\x55\xa1\x16\x93\x3f\xad\xeb\x97\xa8\x66\x4d\x1c\x97\x67\x9f\xff\x92\x9c\x7f\xfe\xcb\xcf\x7f\xd9\x1a\x92\x3a\xab\x08\xea\x00\x53\x57\xac\x4d\xd7\x1e\xda\x64\x60\x93\x90\x2e\xfb\x82\x27\x03\x9b\x84\x8c\x12\x62\x62\x63\x27\xa7\xbe\xd6\x86\x8d\xa0\x3e\x08\x9b\x02\x67\x91\x22\x31\x99\xa6\x3e\xa4\x11\x95\x6f\x73\x7d\x91\x62\x82\xeb\x02\x1f\xa2\x3e\xd8\xa2\x3f\x89\xe1\x63\x38\x13\x26\x16\xc5\x28\x42\x16\x2c\x6c\xd4\x7e\xc2\x37\x62\xa5\x3d\x04\x88\x0f\x2b\x94\x12\xc5\xb0\x3e\xcc\x32\xa2\xa9\x0d\x91\x50\x7c\x93\xe4\x05\xed\x13\x58\x7c\xc4\xec\xaf\x06\x04\x31\x70\x20\x26\x0c\xa4\xea\x32\x2b\x58\x30\x70\x20\xd5\x88\x54\x26\x0e\x76\x72\x9f\x88\x8d\xd8\xa9\x0f\x81\xe2\x09\xaa\xce\x55\x28\xbe\x89\x62\x15\x8a\xff\xfe\xdb\x0c\xff\xe1\xef\x1a\x38\x1e\x13\x9c\x29\x16\x84\xb9\x99\xae\x0d\xf9\xc3\xaf\xba\xc2\xc1\xfc\x8c\x45\x38\xd2\xf3\x2e\xe3\x82\x1b\x03\x95\xf4\x7c\x94\x9e\x9b\xa8\xd8\xc9\xfd\x54\xec\xc9\x4e\x7d\x50\x10\x52\x57\x8a\xa2\x7c\x1b\x15\xcc\x97\x4a\x50\x10\xe2\xce\xb5\x0d\x42\xb2\x06\x92\xef\x15\xb7\xba\x5d\xa6\x2b\x16\xb7\x9b\xa5\x46\xa4\x40\xd9\x30\xc0\xc8\xd2\x51\x68\xf1\xba\x36\x62\x3f\x66\x1b\xb1\x51\x1e\x02\xc3\x77\xf2\xe2\x02\x2f\x55\xf1\xb8\x35\x1c\x2b\xbe\x98\x78\x2c\xcb\xbb\x4c\x37\xbc\x18\xa4\xdd\x1f\xea\xed\x24\x9d\x3b\x59\x3e\xca\x42\x13\xba\xdd\x73\x37\x00\xec\x1a\x71\x08\x8c\xcf\xf2\xa2\xaa\x97\x35\x2e\xb1\x62\x71\x76\x3b\x6a\xf2\x36\x3e\xda\xa2\x5e\x6b\xc3\xe8\xb4\x8c\x18\xc8\xac\xab\x51\x68\x89\xf1\x77\x8e\xf0\xd7\xed\xa6\x76\x8e\x39\x04\x9c\xef\xd6\x65\xa9\x4a\xd8\xf5\x7f\x5e\x7f\xf6\xfa\x47\x2d\x36\xd7\xbf\xbb\xfe\xec\xf5\x0f\x5f\xff\xcd\xeb\x9f\x5e\x7f\x76\xfd\x6f\x12\xa4\xeb\xdf\x5c\xff\xfe\xfa\xb3\xeb\xdf\x36\x40\xbd\xfe\xd9\xf5\xbf\xbc\xfe\xe1\xf5\x7f\x51\x92\x06\xad\xeb\xdf\x5d\xff\xfb\xf5\xef\x5f\xff\xf8\xb5\x9a\x12\xbd\xfe\xe1\xf5\x67\xd7\xff\xf1\xfa\xc7\xd7\xbf\xbf\xfe\x8d\x09\x5a\x51\x77\xb7\x23\x18\x34\x10\x2b\xea\x51\x51\x9b\x88\xd9\xc9\xfd\x42\x6c\xd3\x4e\x7d\x50\xc0\x77\x81\xe3\xbd\x99\x11\xd2\x55\x11\x69\xaa\x78\x86\x0c\x5d\x3c\x50\x15\x4b\xc3\x4a\x0b\x4e\x0c\x44\xca\xf3\x51\x69\x31\x4c\x76\x72\xbf\x14\xfb\xb1\x53\x37\x88\x4c\xff\x27\x00\x00\xff\xff\x9d\x9d\x5c\x2f\x75\x51\x00\x00")

func webStaticFlipclockMinJsBytes() ([]byte, error) {
	return bindataRead(
		_webStaticFlipclockMinJs,
		"web/static/flipclock.min.js",
	)
}

func webStaticFlipclockMinJs() (*asset, error) {
	bytes, err := webStaticFlipclockMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/flipclock.min.js", size: 20853, mode: os.FileMode(438), modTime: time.Unix(1467867550, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStaticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xc1\x6e\x13\x31\x10\x86\xcf\x89\x94\x77\x18\x4c\xa5\x24\x52\xb1\x2b\x71\x00\x85\xdd\x5c\x52\x72\xa5\x0a\x2d\x12\xc7\xa9\x77\x92\x38\xf5\xda\x8b\x3d\x1b\xba\x42\xbc\x3b\xb2\xd7\x05\x45\x08\x4e\xb3\xf3\xcf\x3f\xb3\x9f\xed\xa9\x5e\xdd\x7e\xda\xdc\x7f\xbd\xfb\x08\x47\x6e\xed\x7a\x36\xad\x52\x04\x8b\xee\x50\x0b\x72\x22\x2b\x84\xcd\x7a\x36\x9d\x54\x2d\x31\x82\x3e\x62\x88\xc4\xb5\x78\xb8\xdf\xbe\x79\x2f\x72\x81\x0d\x5b\x5a\x1f\x76\x77\x1b\xd8\xf8\xde\x31\x85\x4a\x8d\x5a\xaa\x5a\xe3\x9e\x20\x90\xad\x45\xe4\xc1\x52\x3c\x12\xb1\x80\x96\x1a\x83\xb5\x40\x6b\x05\x1c\x03\xed\x6b\xb1\xb7\xa6\xd3\xd6\xeb\x27\xa9\x63\x14\xa0\x72\x73\xd4\xc1\x74\x0c\x31\xe8\x5a\xa8\x40\x16\x87\x20\xd6\x95\x1a\xe5\x44\xa7\x0a\x5e\xf5\xe8\x9b\x21\xb7\x34\xe6\x0c\xa6\xa9\x85\x1e\x51\x92\xbd\x31\xe7\xbf\xa6\x1d\x99\xbb\xb8\x52\x0a\x4f\xf8\x2c\x0f\xde\x1f\x2c\x61\x67\xa2\xd4\xbe\xcd\x9a\xb2\xe6\x31\xaa\xd3\xb7\x9e\xc2\xa0\xde\xca\x1b\x79\x53\x12\xd9\x1a\x27\x4f\xf1\x02\xe3\x72\xf4\x9f\x93\xfc\xcf\xca\x43\x47\xb5\x60\x7a\x66\x75\xc2\x33\x8e\x6a\xba\xd0\xd9\x74\x32\x39\x63\x80\x3c\xe2\x43\xca\xae\x16\x8d\xd7\x7d\x4b\x8e\x97\x32\x10\x36\xc3\x62\xdf\x3b\xcd\xc6\xbb\xc5\x12\x7e\x24\xc7\x64\x92\xdd\x50\x83\xa3\xef\xb0\xb5\xa6\xdb\xa4\x7c\x71\xb5\x98\xbf\x2e\x17\x31\x5f\x5e\xbf\x78\x47\xf3\x16\x35\xad\x60\x5e\x9e\x6c\x7e\x5d\x6a\xd8\xb3\xff\xcc\x18\x78\x05\x7b\xb4\x91\x5e\xf4\xd6\x38\xd3\xf6\xed\xad\x39\x18\x8e\x2b\x78\x37\xca\x3f\x97\x99\x70\x8c\xe9\x6b\x97\x1e\x69\xb7\xf1\xce\x51\x26\xfc\x17\xf0\xe8\x93\xf9\xef\xf9\x5b\x6a\x6b\xc8\xb1\x2c\xb8\x0f\x5d\x83\x4c\x0d\xd4\xf0\xbb\x35\x57\x2e\x0f\x2c\x23\xf1\x17\xb4\x3d\x95\x62\x81\xc9\x61\x44\xbb\x58\x96\xb2\x24\x95\x2a\xeb\xfe\x2b\x00\x00\xff\xff\x31\x1f\x4d\x21\x00\x03\x00\x00")

func webStaticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webStaticIndexHtml,
		"web/static/index.html",
	)
}

func webStaticIndexHtml() (*asset, error) {
	bytes, err := webStaticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/static/index.html", size: 768, mode: os.FileMode(438), modTime: time.Unix(1467900262, 0)}
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
	"web/static/flipclock.css": webStaticFlipclockCss,
	"web/static/flipclock.min.js": webStaticFlipclockMinJs,
	"web/static/index.html": webStaticIndexHtml,
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
	"web": &bintree{nil, map[string]*bintree{
		"static": &bintree{nil, map[string]*bintree{
			"flipclock.css": &bintree{webStaticFlipclockCss, map[string]*bintree{}},
			"flipclock.min.js": &bintree{webStaticFlipclockMinJs, map[string]*bintree{}},
			"index.html": &bintree{webStaticIndexHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
