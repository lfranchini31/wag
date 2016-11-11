// Code generated by go-bindata.
// sources:
// _hardcoded/doer.go
// _hardcoded/middleware.go
// DO NOT EDIT!

package hardcoded

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

var __hardcodedDoerGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x39\x7f\x73\xdb\x36\xb2\x7f\x93\x9f\x62\xa3\x99\x36\xa4\xa3\x90\x76\x5a\xf7\xbd\x51\x9e\xfb\xa6\xb5\x9d\x4b\x66\xd2\x24\x53\xb9\xd7\xcc\x64\x32\x0d\x44\xae\x28\xc4\x24\xc0\x00\xa0\x64\x8d\xeb\xef\x7e\xb3\x00\x48\x91\x92\xac\xa4\x3f\xe6\xe6\xce\x7f\x58\x12\xb0\xbb\xd8\x5d\xec\x6f\xd4\x2c\xbb\x66\x05\x42\x56\x72\x14\x26\x0c\x79\x55\x4b\x65\x20\x0a\x83\x51\x26\x85\xc1\x1b\x33\x0a\x83\x11\x8a\x4c\xe6\x5c\x14\xe9\x47\x2d\x85\x5d\x50\x4a\x2a\x4d\xdf\xe6\x95\x85\x28\x65\x41\x1f\x15\x33\x8b\x54\x31\x91\xd3\x0f\x81\xc6\x7f\xa4\x0b\x63\x6a\xfa\xae\xd7\x22\xa3\x4f\xc3\x2b\x1c\x85\x61\x30\x2a\xb8\x59\x34\xb3\x24\x93\x55\xca\xe6\x78\x93\x2e\xd6\xda\x28\x7e\xf3\xb8\x90\xed\xd7\xd1\x10\x2a\x97\x42\x2e\x99\x58\xf0\x1c\x53\x5c\xa2\x30\x5a\x36\x2a\xc3\x51\x18\xc8\x1a\x85\x51\x2c\xe3\xa2\x80\x3e\x46\x6f\xbd\xff\xfd\x71\x21\x2d\x69\x59\x32\x51\x24\x52\x15\xe9\x4d\x4a\xac\x7a\xb1\xd3\xcc\xdc\x78\xae\x4b\x59\x14\xa8\x60\x54\xc8\xfa\xba\x48\xb8\x48\xcf\x4b\x5c\xa2\x4a\xaf\xd9\x7a\x89\xf8\xb8\x90\xc9\xf2\x34\x75\x40\xa3\x30\x0e\xc3\x34\x85\x5c\xa2\x02\xae\x81\x09\xe0\xc2\xa0\x9a\xb3\x0c\x61\x2e\x15\x8c\x72\xc9\x45\x31\x02\x22\x0d\x0a\x3f\x35\xa8\x8d\x86\x5a\x6a\xcd\x67\xe5\x1a\x56\xdc\x2c\x60\xa5\x58\x5d\x73\x51\x84\x66\x5d\xa3\x27\xd5\x11\xb9\x0d\x83\x0b\x19\x65\x70\x44\x14\x92\x73\x7b\x6d\x63\x50\xfe\xf7\xcf\x8e\x62\x0c\x51\xfb\x5b\xd7\x52\x68\x1c\x83\xbd\xb1\x38\xbc\x0b\x1d\x55\x59\xbf\x62\x15\x9e\x9b\x1b\xd0\x46\x35\x99\xb9\xbd\xb3\x7c\xcf\x98\x46\x4f\xe3\x39\x13\x79\x89\x0a\x6a\x54\x73\xa9\x2a\x0d\x66\x81\x76\x7f\xc0\xbb\xa3\x46\xcb\x17\xc4\xe7\x86\xd8\xbc\x11\x19\x44\x79\xb7\x15\xc3\x5f\x60\x9b\xa4\x56\x68\x1a\x25\xc0\x5f\x4b\x72\x21\x23\x95\x9c\xbb\xbb\x8a\xe2\x31\x64\x63\x70\xe2\xa5\x29\xf8\x1b\xb6\x1c\xb1\x3c\xd7\xed\x02\x18\x39\x54\xbc\xe3\xbe\x0f\xee\x04\xa0\xf3\x72\xab\xf9\xb0\x27\x4a\x0f\xee\xaf\x4a\x13\x06\x99\xb9\x81\xc9\x19\xf4\x64\x20\x0b\xa6\x4b\xa1\xe5\xcc\xdc\x24\xff\x64\x65\x83\x51\x77\x51\xb7\x77\x71\x12\x91\x47\x88\x22\x0e\x83\x25\x53\xa0\x6b\xe8\xd9\x73\x32\xad\x99\x08\x83\x34\x85\xab\xd7\x17\xaf\x27\x24\x38\x18\x56\x68\x50\x58\x32\xe3\xa5\xe7\xa2\x6e\x0c\xe4\xcc\xb0\xff\x0f\x03\x3e\x87\x9a\x29\x14\x86\x30\xe9\xd4\x6d\x6a\xcf\x94\xac\x5a\xf6\x32\x73\x13\x3f\xed\xc3\x3f\x38\x03\xc1\x4b\xd2\x54\xa0\x6b\xd8\x42\x36\x4c\x59\x28\xcf\xfe\x78\xb0\x7b\xbe\xe0\x65\xfe\x7a\x1e\x6d\x88\x6d\x94\x10\xc7\x61\x70\x07\x58\x6a\xfc\x32\xca\x04\x6e\x25\x41\xa5\x48\x04\x5d\x27\x57\x8a\x65\xa8\xa2\x38\x79\x21\x3e\x62\x66\x22\x5d\xf7\xec\x24\x0c\xfa\x71\x22\x79\x7e\x75\xf5\xe6\x39\xb2\x1c\x95\x3e\xb0\x75\xce\x94\xe2\xa8\x22\x95\xb8\x85\x38\x7e\x6a\xcf\xeb\xa9\xc0\x5b\xa7\xe0\xe5\x18\xe6\x95\x49\x2e\xe9\xa6\xe7\xd1\x28\x93\x4d\x99\x8b\x87\x06\xb8\x65\xa6\xb3\xc4\x85\xa3\x0c\xd1\x57\xcb\x78\x64\x0d\xc3\x49\xe2\xe9\xe4\x49\x4e\x16\xde\x37\x6a\x85\x46\xad\x5b\xa7\xa4\x1f\x1c\x35\x9c\x1e\xbf\xdd\x67\xd1\x16\x76\xc7\x9e\x37\x7f\xd6\xb2\x03\x0b\xf5\x46\x96\x3c\x5b\xc3\xcf\x9b\xef\xfe\xbc\xde\x0a\xe4\x38\xe7\x02\x35\x30\x47\x19\x6a\xbb\x9c\xb8\xc3\xfa\x80\x83\x38\x95\xa6\xf0\x23\xcb\xae\xe5\x7c\x4e\x46\x48\x72\xb9\x20\x22\x9a\x6a\x46\x9e\x29\x72\x30\xbc\x22\x6d\xc8\xb9\x27\xcc\x8c\xc1\xaa\x36\x3a\x09\x83\x16\x35\x8a\xe1\xdd\x7b\xca\x17\xc9\x45\xa3\x98\xe1\xd2\x19\xb9\x3d\x15\x14\x66\xc8\x97\xe8\xe8\xf6\x15\x31\x06\xa6\x61\x85\x65\x49\x9f\xb4\xa9\x50\x37\xa5\x01\x39\xb7\xd8\x6d\x4e\xf2\xbe\xfb\x50\xc3\x87\x0b\xf9\x01\x2a\x34\x0b\x99\x27\x61\x60\xa9\x47\x03\x77\x1e\xc3\x3d\xde\x3c\x93\xb2\xf4\x3a\x9b\x72\x51\x94\x78\x48\x73\x66\xc1\x4c\x77\x7b\xac\xe5\x16\xa4\xc8\xd0\x69\x73\x97\xc4\x20\x40\xef\x51\x28\x33\xb0\x96\x0d\xe8\x05\x99\x5a\x77\x0c\x76\xb4\x4f\x34\x66\x52\xe4\xc0\xe6\x86\x32\x89\x81\x39\xe3\xa5\x4e\x7c\x58\xdb\x39\x2f\x86\xfb\x35\xdf\x8b\xc3\x5b\x3b\xb7\x27\x70\x04\x76\x65\x6a\x4f\xbb\xeb\x5b\x11\xac\x78\x59\x7a\xce\x84\x14\x8f\xdf\xbc\x9e\x5e\x8d\xdd\xb7\x1f\xae\xce\x9f\x6f\xf2\xa0\x15\xe6\xf4\xed\xdb\x24\xec\xc2\xd8\x0b\x43\xd6\xaa\x41\x48\x03\x59\xa3\x28\x62\x94\xeb\xd6\x5a\xc4\xda\x5d\x42\xab\x0e\xcc\x61\xb6\xbe\xe7\x72\x0f\x08\xec\x6e\x5b\xe1\x27\xd8\xba\x71\x85\xba\xde\x77\xed\xfd\xab\x27\x9d\xf8\xf8\xe3\xe3\xc1\xef\xbf\x93\x44\xc9\x4f\xd6\x98\xe0\xec\x0c\x46\x24\xf0\x68\xdf\x3a\x89\x4f\x1b\x36\x82\xe8\x9a\xc2\x9b\x69\xf4\xb9\xcc\x11\xfe\x0f\x4e\x8f\x8f\xfb\xb1\x65\xce\x4a\x8d\xfd\x20\x61\x54\x83\x5e\xcd\x97\x37\xb5\x14\x28\x0c\x67\xe5\x5e\xeb\x13\x80\x1b\x88\x81\x13\x3b\xab\xbb\x07\xff\xb0\xe9\xcd\xf9\x92\x52\xbf\x5f\xb5\xa5\x4b\xef\x94\x92\xa2\x41\xa6\x90\x69\xf2\xf0\x15\xe3\xc6\xda\x87\xb6\x65\x06\x9a\x15\xa2\xe8\x2e\x7e\x02\x27\xc7\xc7\x63\x78\x42\xff\xbe\xa5\x7f\xff\x4b\xff\x28\x40\x9c\x7c\x77\x7c\x0c\x15\x2f\x4b\xee\xac\x58\xc3\xa3\xf4\x31\x34\x35\x25\xb3\xd3\xaf\xe0\x23\x37\x06\x55\x7b\xb5\xfb\xa5\xf8\x02\x83\xa6\xcc\x51\xb1\x6b\x8c\xb6\xb6\xc7\x70\x1a\x87\x81\xc0\x1b\x0b\x71\x72\x7c\xdc\x1a\xf9\x4f\x1b\x8e\xc2\x40\x89\xdc\xa6\x72\x26\xf2\xe4\x15\xae\xa2\xf6\xcb\xd4\x56\xa6\x91\x45\x78\x25\x57\x51\x9c\xfc\x22\xf8\xcd\x2b\x26\xa4\xcb\x73\x36\xd3\x1f\x27\xc7\xa7\x90\xa6\x56\xac\x53\x2a\xb6\x32\x14\xc6\xcb\x15\x06\x54\x32\x72\x4f\xbc\xb0\x71\xbd\xb5\x88\x77\xfc\x3d\x9c\x81\x65\xed\x11\x0c\x98\x8e\xa2\x48\x89\x3c\x79\x56\x4a\x66\xbe\xfb\x36\x8a\x8f\x9e\xc4\x8f\x4f\xe2\x23\x3c\x9a\xfb\x15\x42\xa2\xf3\x9d\x60\x47\x67\xf0\xa4\x6f\x55\x0a\xcd\x7f\xac\xef\xde\x77\xc1\xff\xfd\x0e\xfc\x4a\xee\xcf\x1a\xce\x49\x29\x41\x0a\xea\x39\x9c\xf2\xc8\xa1\xba\xdc\xe1\xb3\xf0\x90\xc0\x61\xc7\xa5\x70\x50\xd5\x66\x0d\xba\xe4\x19\xb6\xca\x1d\x50\xf8\x93\x59\x60\x18\xf6\x59\xb9\x62\xeb\x5e\xbc\x20\xf9\xef\x39\xed\x8f\x26\xdc\x1e\x0f\x4e\xad\x6d\x63\x63\x15\xe4\xeb\xbd\xa1\x1a\x7e\xe5\x66\xd1\xd7\x51\xa7\x0d\x10\xb8\x02\xdf\xf6\x39\x33\x96\x4b\x54\x8a\xe7\xbe\xae\x70\xbd\x31\xc8\x19\x55\x70\x0f\x75\x57\x8c\x75\x65\x90\x95\x68\x8b\x3c\x15\xcc\x2d\xd1\xb6\xfe\x1c\xc3\x3d\x25\x57\xbc\x0d\xd9\x6f\x78\xfc\x0e\xd1\x77\x2d\x41\x66\x6e\xc6\x03\x39\x6f\xef\x06\x94\xe3\x7e\xd3\x72\xd4\xd5\x82\x7f\x47\x07\xd6\x1e\x31\x06\x79\xbd\xd5\xbe\xf8\x7e\x65\xc8\x57\x9c\x44\x7d\x31\xad\xa3\x3d\x90\xd7\xad\x57\x74\xba\x38\x83\x3c\xe9\xfd\xb6\x4e\xd2\x65\x16\x3a\x67\xb3\x97\x6c\x0c\xd3\xb5\x41\x7b\x7c\xdc\x6d\x74\x7e\xee\xe2\x68\x5b\x75\x51\xd0\x7d\x6a\x7d\xef\x69\xbb\xf6\xe8\x91\xe7\x48\xd7\x2e\x3c\x9c\x0d\x2b\xf0\x80\xf8\x6e\xf1\xcf\xce\xa0\x44\x11\xb5\xec\xc5\x14\x16\x1e\xf4\x19\xf4\xd1\xc8\x45\x1f\x57\xdb\x5b\xf2\xc1\x4c\x21\xbb\x0e\x03\x12\x8e\x8a\xd0\xf3\x52\x6a\x6c\x8b\x53\xcb\x37\xcc\x64\xbe\x76\x85\x31\xb9\x4f\xc1\xb8\x68\x03\xcb\x8f\x32\x5f\x27\x16\x83\xe4\x0e\x5c\xa9\x55\x22\xd6\x1d\x23\xef\x3c\x83\xef\xe3\x61\x38\xf7\x3c\x78\xcf\xcc\xb8\xca\x1a\x6e\x7e\x24\x56\x50\xd9\x1e\x81\x57\x75\x89\x15\x0a\xe3\xcd\xdd\x41\xc0\xcc\x81\x40\x4d\x45\xb9\x12\x36\xae\x4f\xd1\x40\x8e\xb3\xc6\x36\x93\xa4\x42\xfa\x94\x35\x2a\x66\x70\x80\xcc\x05\x30\xa8\x28\x0c\xae\x16\xa8\x90\xca\x4e\x9b\x48\xa4\x28\xd7\x50\xca\xa2\xc7\x0a\x68\xc3\x0c\xba\x74\x6f\x61\x28\x5d\xcc\x4a\x99\x5d\xdb\x24\xd1\x25\x98\xb9\x92\x15\x14\xd2\x76\xb2\x0b\x25\x9b\x62\xe1\x43\xdf\x1e\x91\x0e\xb5\x3d\x8e\x7f\xf7\x67\x6b\xf7\xc0\x13\xb0\xcd\xb7\xeb\xb0\x49\x59\x64\x42\x7e\x67\x3a\xbd\x7c\x2d\x32\x04\xbd\x16\x59\x42\xdf\xac\x2a\x9f\xbb\x01\xd5\x74\x7a\x79\xb9\xa4\x00\xc1\x35\x60\x45\x89\xdb\x66\xb2\xcd\x24\x0b\x96\x9c\x81\x46\xb5\x44\xf5\x58\x13\xa0\x9b\x5b\x25\x36\x3f\xa2\xce\x14\x9f\xb9\xb2\x88\x14\x68\x95\x41\xf1\x9e\xb5\x87\x7b\x29\xb7\x4f\xdb\x88\x78\x45\xdb\x07\xfe\x9c\x48\xf0\xe1\xa3\x96\x62\x32\x22\x62\xa3\x0f\x61\x60\xa5\xfd\x62\x2c\xc1\x2a\x8b\xe5\xe3\xc4\xb9\x6c\x84\xd9\x87\xc5\xdd\xb2\xc7\x52\x3d\x68\xc2\xb6\x5d\xf1\xbd\xb8\xdb\xd8\xd8\x41\x77\xb8\x6f\x5c\x7d\xc4\x8a\x5d\xd6\xf7\xe0\x6e\xa0\x89\xc0\x0b\x7d\x3e\x30\x94\xd7\x35\x8a\x3e\x01\x9b\x56\x3a\x02\x7c\x0f\xb4\x55\x80\x2c\x4b\x2e\x0a\xcb\xd6\x33\xc6\xcb\x46\xe1\x21\x05\xec\x42\xef\x12\x29\x4b\x72\xe1\x3e\xb1\x43\x44\x06\xd0\xf7\x11\x9b\x36\x59\x86\x5a\x7f\x21\x31\x0f\xbd\x4d\x6c\xba\x90\xca\x78\x2d\x60\xfe\x39\xf1\x86\xd0\x3b\xb4\x7a\x0c\x7d\x5e\x55\xf7\x30\x74\xb5\x50\xc8\xf2\x37\x52\x96\x3f\x23\xe5\x62\x62\xea\x7e\x22\xbb\xd0\x3b\xf4\x78\x85\xb2\x31\x5f\xc8\x94\x87\x26\x22\xe7\xae\x98\x3d\x97\xc2\x97\xb5\x97\x37\x98\x35\x54\x01\x39\xe3\x1e\x12\xc9\x0e\x43\x13\xc1\x97\xcc\xa0\xc8\xd6\x57\xd2\xb0\xf2\x27\x64\xe2\xb0\x71\x97\x3d\xe8\xdf\x2a\x64\x64\x98\x6d\xca\x2f\x65\x61\xe3\x43\x54\xc2\x91\x9b\x49\x27\x2f\xed\xc7\x18\x76\x62\x88\x4d\x4b\x65\xf2\x42\xcc\xe5\x45\x84\x89\x9b\xd4\x55\xac\x7e\xe7\x7c\xff\x7d\x37\xcd\xb9\xbd\xa3\xfc\x35\x74\xe7\xc9\x3e\x07\xc6\xa4\x1f\x20\xc6\x84\xd4\xf3\xe2\xbd\x28\x84\xb4\x89\x0b\x1b\x94\x9e\xf3\xee\xe2\x79\x94\x0d\x8c\xc5\xdb\xeb\xb3\x93\x21\xde\xbe\x28\x60\x91\xf7\xb9\xea\x64\x5b\xb8\x5d\x98\x3d\xb8\x43\x0f\x9d\xec\xc5\x1d\xc0\xdc\x4b\xa3\xf5\x83\x43\x34\x3c\xcc\x0e\x8d\x2d\x7f\x9c\xec\x93\x61\x08\xb3\x4b\x62\x70\xfc\x7e\x35\xdc\x77\xfc\x1e\xef\x9b\x6c\xe3\xee\xc2\xec\x92\xf1\x4e\x77\x88\x05\x0f\x63\x71\x3f\xe7\x6b\x13\xc0\xe4\x33\xde\x6b\xe9\x94\x5b\x0e\xb9\x6b\x83\x98\x6c\x3b\xed\x38\x0c\xee\xe2\x70\x50\x7c\xef\x96\x26\x31\x70\xc1\x4d\x14\xfb\x01\xe9\x1b\x54\x5c\xe6\x3c\xb3\xf3\x92\x52\x16\x6e\x82\x2a\x58\x39\xac\x8e\xa8\xcc\x62\x5a\x73\x6d\x2b\x2b\x8d\xc6\x50\x91\x12\xf4\x8a\x28\xb3\x50\xa8\x17\xb2\xcc\xb5\x2d\xa4\x1a\x91\xa3\xd2\x86\x89\xdc\x36\x8b\x75\x5d\xf2\xcc\x35\x6f\x33\x5c\xb0\x25\x97\x2a\xb1\xe8\xbf\x88\xb9\x54\xa6\x11\xcc\x60\xb9\x1e\xf7\x8b\x16\xea\xdc\xc5\x43\x03\x0b\xb6\x44\x60\x50\x28\x64\x06\x56\x6c\x4d\xac\xe0\x4d\xed\x4a\x57\xae\x2d\x95\x9c\x19\x36\x06\x2d\x61\x65\x6b\x59\xa9\x8c\xad\x0f\x1b\x25\xec\x74\x57\x00\x37\x1a\x74\x53\xdb\xb7\x45\xaa\xc5\x67\x4a\xb2\x3c\x63\x7a\x47\x8c\x8a\x6a\xd8\x4c\xdb\x92\xc9\xce\x73\x77\xeb\x26\x88\xa6\xd3\xcb\xd8\xb1\x3f\x45\x37\xf6\xd5\x93\x34\x3d\xf0\x94\xc8\xb5\x6e\x50\xa7\xa7\xdf\x25\x5d\xb1\xe7\x4b\x3a\xaa\xf1\xe9\xb6\xdc\x85\x04\x1e\x67\x6a\x14\xb2\xaa\x1d\xb6\x4f\xce\x5a\xbd\xd8\xa9\x4e\x7f\xcf\xd6\xe2\xfb\x90\xdc\x43\x85\xdd\x2e\xb9\x36\x28\x6c\x00\x76\x6f\x14\x02\x4d\xf2\xd2\x2e\x46\x23\x93\xd5\xa3\x31\x8c\x4e\x9e\xfc\x4f\x72\x9c\x1c\x27\x27\x93\xe3\x91\x6f\x37\xb6\x1e\x18\x82\x9a\x09\x9e\x45\xee\xa1\xc0\x76\x10\x85\x7b\xcc\x4a\xa6\xa4\xa1\x68\x73\xca\x3e\x76\x62\x07\xdf\x93\xb4\x7d\xd5\x9c\x9c\x81\x4f\x13\xaf\x70\x15\x8d\x56\xac\xb0\x0c\xd0\xf6\x33\x1b\xef\x45\xb6\x26\xa0\x6f\x8e\x87\x03\x5e\x0b\xc3\xb4\xb1\x89\x64\x8a\x28\xdc\x0c\xad\xcb\x1d\x5b\xa9\xe6\xf6\x6e\x80\xf0\x52\x16\xe4\xb5\x5b\x38\x96\x3c\xad\x53\xc7\x1e\x04\xb6\x6b\xfb\x0d\xda\xe1\x97\xdf\xcd\xae\xa3\xad\x59\x73\x0c\xb7\xd0\xf5\xe4\x5c\xd8\xe1\x09\x27\x7b\x26\x1a\x41\xa3\x4a\x3a\x66\x44\xba\x9a\xa4\xe9\x08\x1e\x41\xab\xab\xe4\x87\x3c\x57\x51\x9c\x4c\xed\xf1\xf6\xb2\xa8\xdb\xfa\xd4\xdd\x94\xd5\xef\x2b\x5c\xf9\xa4\x16\x8d\xfe\x71\x79\x35\x1a\x43\xa3\xca\x31\x5d\x8c\x43\xd8\x73\x57\x56\x7d\xc9\x1b\xc5\x85\x99\x47\x2e\x9d\x51\x1f\x2f\x30\x6b\x1f\xe2\x5a\x7b\x9f\x4e\x2f\xa9\xd2\x46\x56\x4d\xe0\x2b\xdd\xbd\x04\xd1\x1f\x35\xfe\x5c\x34\x68\x7f\x59\xf5\x05\x0e\xb2\x63\xaf\xf7\xfe\x9d\x4c\x9b\x99\xeb\x24\x7e\xe5\x66\x11\x8d\x46\x63\xf8\xba\xd7\xe8\xbb\x09\xc1\xa7\x7f\x3b\xc3\x74\x83\xb8\xdc\x0c\x30\x1d\x66\x72\xe9\x1c\xd9\x1f\x6d\x7b\xf5\xed\xda\xc4\x6d\x6d\x1e\xf6\xa8\xf2\x49\x7e\x11\x15\x53\x7a\xc1\xca\xe8\xdd\xfb\xd9\xda\x60\x84\xcb\xe4\x82\x19\x16\xc5\xf1\x18\xbe\xc6\xdd\x67\xb9\x1d\xb6\x3c\x5f\x96\x70\x62\x1b\xa7\x07\x67\x30\xf2\x67\x9f\xcb\xaa\x62\x22\x1f\x1d\x46\x25\x23\x26\x83\x6f\x67\x20\x03\x2f\x78\xe7\xca\xa8\xf7\x1b\xd0\x9d\x2d\x38\x03\x0c\xdd\x7e\x9a\xda\x98\xdf\x6f\x9c\x5d\xb8\xe7\x73\xe0\xe6\xa1\x06\xe4\x66\x41\x3d\xab\x07\x8e\x4e\x62\x0b\x3c\xe7\x4a\xfb\x70\x08\x2b\x7c\xb8\x44\xd0\xe4\x81\xa4\xed\x3e\x2d\xd9\xc3\x7c\x12\x0f\xb7\xa8\xf9\x71\x67\x2d\x98\x86\x6c\x41\xd7\x93\x0f\x30\xbe\x89\xc9\xe9\xdb\xb1\xbb\x4d\x03\x35\xd3\x1a\x73\xd0\x9c\x9a\xe1\x15\xba\xc8\x91\x83\x96\x15\x9a\x05\x19\xca\x16\x07\x9d\xae\xdb\x99\x4f\xb0\x27\x08\xf4\xf4\xb2\x19\x95\xb7\xb0\x5d\x35\xdb\xd6\xb0\xf1\x67\x6e\xb5\xbd\x9d\xbd\x85\x1e\x5d\xf6\xfe\x12\xf0\xef\xe5\xce\xbf\x60\xf3\x79\x0f\x87\x3c\x34\xba\x8f\x7a\x0c\xdf\xc3\x20\xe2\xfe\xbd\xec\x6c\x3c\x92\xfe\xdd\x11\xc6\x5d\xfc\xd9\x42\xe5\x2f\x8e\x0b\xf9\x1c\xf2\xc4\x0d\x5b\x86\xc3\xb5\xc9\xce\x74\xad\x4d\xaf\x17\x32\xca\x93\xde\x34\x66\xdc\x66\x2c\x17\x91\xac\x52\xf6\x45\x2f\x3f\xeb\x42\xa5\x5a\x19\xdd\xc0\x6e\x38\x90\xff\xfe\xac\x9b\xc8\x07\xee\x11\x18\x73\x0a\x6d\x1e\x9b\x09\x7f\x8c\x1d\x6d\x71\x9b\x1d\xfb\xee\x62\x4b\x95\xc1\x61\x52\x69\x97\x35\x4f\xdf\xbe\x75\x59\xd3\x1e\x3d\xe0\xe6\xae\xcb\x14\x3b\x03\xb9\x80\xb2\xdc\x97\x4e\x30\xdd\x54\xf2\x0f\x28\xea\xf0\x30\x73\x5b\x85\x43\x9e\xc3\xcf\xea\xef\x4f\xab\xef\x80\xf6\xee\xc2\x21\x1f\x4e\x75\x90\x1e\x81\x90\x30\xf7\xed\x8e\x15\xd4\xbd\xcd\xad\xd1\xc0\x51\x1a\xef\x1d\x75\xfe\x2b\x00\x00\xff\xff\xa8\x58\xf2\xc5\xd8\x26\x00\x00")

func _hardcodedDoerGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedDoerGo,
		"_hardcoded/doer.go",
	)
}

func _hardcodedDoerGo() (*asset, error) {
	bytes, err := _hardcodedDoerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_hardcoded/doer.go", size: 9944, mode: os.FileMode(436), modTime: time.Unix(1478885327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __hardcodedMiddlewareGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x56\xcd\x6e\xdc\x36\x10\x3e\x4b\x4f\x31\x55\x11\x54\x2c\x64\xee\xa9\x17\x07\x3e\x04\x8e\xdd\x04\x70\xd2\x20\x76\x90\x43\x10\x14\xb4\x34\x92\x98\x95\x48\x95\x1c\xad\x6c\x04\x7e\xf7\x62\x48\x69\x57\x1b\xdb\x6d\x4f\x3d\x78\xcd\xe5\xfc\x7f\x33\xf3\x71\x07\x55\x6e\x55\x83\xe0\xd1\xed\xd0\xa5\xa9\xee\x07\xeb\x08\xf2\x34\xc9\xea\x9e\xb2\x34\xc9\x0c\xd2\xa6\x25\x1a\xf8\xec\x46\x43\xba\xc7\x4d\x85\xb7\x63\x93\xa5\x69\x92\x35\x76\xd8\x36\x52\x9b\xcd\x79\x87\x3b\x74\x9b\xad\xba\xdf\x21\x9e\x34\x56\xee\x7e\xdb\x74\xb6\x69\xd0\x65\x69\xb2\xdd\xbd\xd3\x55\xd5\xe1\xa4\x1c\xc2\x3f\xdb\xf4\x7b\x45\xf6\x6f\x07\x34\xe4\x54\xa9\x4d\x03\x59\xa3\xa9\x1d\x6f\x65\x69\xfb\xcd\xea\x7e\x7d\x3e\x69\x6c\x96\x8a\x34\xad\x47\x53\xc2\xa4\xa9\x3d\x84\xcd\xb9\x42\x5d\xe2\x7b\xd5\x23\x78\x72\xda\x34\x05\x38\x3b\x12\x3a\xe0\xf2\xe4\x1b\x65\xaa\x0e\x9d\x38\xfa\x06\xdf\xd3\xa4\x9d\x8f\xa7\x67\x30\x87\x59\x79\x8d\x1e\xc4\x41\xeb\x0c\x06\x65\x74\xb9\x52\x99\x25\x22\x4d\x36\x1b\xb8\xb2\x4d\xc3\xc5\x1c\xca\x84\xd2\xf6\xe8\xa1\x53\x9e\x0a\xd0\x12\x25\x4c\xba\xeb\xe0\x16\xc1\x8d\x06\x6a\xed\x3c\xc9\x60\x7a\xd3\x6a\x0f\xbd\xda\xa2\x07\x4d\xe0\x2d\x50\xab\x08\x2c\xb5\xe8\xd6\xee\x5a\xe5\x41\x95\x25\x7a\x0f\xc4\x3a\x08\xb1\x0d\xc1\x47\x30\x39\xea\x86\x36\xdf\xb0\x24\x0f\xda\xcc\xda\x0e\xff\x1a\xd1\x13\x94\xd6\x10\xde\x71\xec\x43\x69\x6b\x4b\xf9\x1e\xa7\xa5\xb6\x02\x56\xe8\x8a\x34\x71\x48\xa3\x33\x30\x4b\xd3\x87\x34\xdd\x6c\x7e\xc4\x85\xd3\xf2\xa0\xcc\x7d\x14\x78\x09\x97\xd6\x81\xb1\x53\x01\x13\xfe\x12\x70\x31\xa4\xcd\x88\x40\xad\xb3\x13\x83\xc6\xe9\x05\x65\x18\x87\x34\x54\xc3\xad\x54\xe5\x36\xa2\x11\xe0\xb9\x87\xd2\x29\xdf\x46\x5d\x67\x19\x07\x19\xe7\xe1\x51\x5f\xfe\xa5\xef\x4b\x11\xab\xeb\xcb\xd1\x94\x39\x3b\xcb\xa7\x78\xff\x11\xfd\x60\x8d\xc7\xcf\x4e\x13\xc3\xe0\xe0\xd7\xf9\x3e\x80\x28\xd8\x4f\x52\x61\x8d\x0e\x82\x59\xbc\x48\x42\x2a\x17\x2e\xcc\x94\xc3\xd2\xee\xd0\xe5\x82\x05\xba\x86\xbd\xec\xec\x0c\x8c\xee\xa2\xc1\x9c\x0c\x1f\x1f\xf8\x63\xa7\x1c\xa0\x0b\x7f\xd6\xa5\x7c\xe3\x27\x4d\x65\x0b\x6b\xcf\xcb\x59\xe6\x74\x3f\xe0\x1c\xba\x54\x7e\x99\xff\xd3\xe0\x99\xdd\x9c\x41\xdd\x93\xbc\x60\x67\x75\xbe\x98\x89\xbd\x7a\x88\xb2\xd6\x5e\x54\xf8\xaa\xc2\x5a\x8d\x1d\x3d\xe3\x2c\x1b\xcd\xd6\xd8\xc9\xcc\x7d\x7b\xf1\xf3\x0e\x6c\x0d\x9c\x0f\xbc\xb8\xc9\x8a\xbd\xa7\xc3\x49\xc4\x22\xf9\x33\x0e\xae\xbc\x74\xb6\x3f\x8f\xd3\x98\x3b\xb9\x9c\x84\x88\x31\x5e\xe7\x59\x30\xcd\x8a\x90\xc1\x6c\xf3\xee\x7b\x86\xce\x65\xa7\x9c\x7b\x01\x59\x98\x12\xde\x5e\xcc\x4e\xe7\xea\xf3\x40\x62\xf2\x9a\x25\xb9\x10\x0f\x62\xdf\x98\x23\x04\x1e\x42\x63\x5a\x79\xcd\x04\xf9\xe6\xe6\xe6\x43\x3e\x15\xc0\x92\x07\x31\xcf\xf5\x23\x52\x80\xd2\xa1\x22\xf4\xa0\xc0\xe0\x04\x7e\x50\x06\x8c\xea\xb1\x02\x55\x33\xdd\xf0\x68\x7e\xfa\x78\x05\x83\xa2\x36\xa0\x71\xd8\x3a\xc9\x0e\xdf\x12\x0c\x9d\x2a\xd1\xc7\x99\x0e\xf6\xda\x3c\xb5\x9c\x05\xd4\xd6\xc1\xe8\x11\x6e\xef\x67\x26\x98\x77\xce\xc3\x4e\x2b\x58\x51\xa3\xbc\x1e\x94\x59\x23\x29\x42\xa8\x1a\x54\x0c\x80\x77\xda\x07\x1a\xd8\xc7\x68\x51\x55\xe8\x7c\x11\xd7\x8c\x75\x62\x5d\x15\x07\x8b\xdb\x76\x28\x79\xe1\x2c\x05\x65\xab\xbb\x2a\x96\xa5\x28\xd8\xcd\xfb\xf7\x98\x3c\xff\xb7\x0d\xdc\x6c\xe0\x15\x11\xf6\x03\x31\x27\x7e\xb3\xda\x2c\x65\xdf\xde\x43\x83\x44\x81\x5d\x78\x3c\x40\x9b\xda\x42\xed\x6c\x1f\xca\x9e\x31\x90\x69\x92\xd8\x21\x3c\x1d\xbc\xb1\xf2\xd3\xc7\x2b\xf9\x41\x51\x9b\xc6\x4d\xf4\xc3\x23\xa4\xd3\xb0\xcb\xbe\x2c\xc2\x9a\x9e\x9e\x1d\x29\xfc\xde\xd9\x5b\xd5\xdd\x70\x3c\x97\x0b\x76\x9e\x5c\xdc\xb1\x90\xf2\xb5\x1a\x4f\xdb\x9b\xb9\x09\x61\xb6\x9f\x11\x9e\x2b\xe7\x34\xba\xdc\xc9\x78\x21\xc4\xcb\x10\xf5\xa7\x15\x81\xf8\x01\x8e\x73\xb8\x26\xe5\x88\x33\xcd\x63\x61\x61\xd6\x01\x3b\x8f\xff\xd1\xa0\x38\x92\x9e\x73\xd7\xff\xa8\x73\x5f\x8a\xe0\x69\x4f\x7b\x7e\x90\x97\xda\x68\xdf\xe6\x22\x8d\x9d\x88\x2f\x4e\x84\xff\xed\xeb\xf8\xf0\x84\xb7\x80\x2c\x28\x5d\xad\x47\x30\xac\x28\xbf\x97\x69\x92\x10\xa3\xc8\xcf\x5f\xde\xab\xe1\x4b\xdc\xe1\xaf\xf1\x9f\x88\x68\xcf\x48\xfb\x41\xee\xa1\x7d\x1b\x62\xe5\x7e\x38\xd0\xc6\x71\xde\x37\x78\x47\xef\xd4\x10\xf0\x7d\xe2\x7e\x81\x96\x16\x4c\xd7\xa4\xcc\x0d\x66\xaa\xa8\x0a\xb0\xdb\xf0\xf3\xe0\x4b\x66\xe9\x24\xcc\x91\x3b\x89\xa2\xec\xeb\x4b\x16\x7e\x5f\x73\xd3\xb3\x7c\xf6\xaa\xaa\x96\x6f\x8f\x1d\x15\x73\x30\xb1\x3c\x02\x81\x23\xfd\x20\xaf\x6c\x73\xb1\x43\x43\x79\x36\xc3\xf6\xa7\xc3\x12\xf5\x0e\xab\x4c\x3c\xf9\xfa\x3c\x69\x53\x87\x26\xcd\x36\x91\xf4\x0c\x4e\xe7\x74\xf7\xe3\xf0\xce\x19\x7e\xd6\xd4\x86\x69\x70\x6b\x68\xfd\x10\xba\xfc\x23\x5d\x4a\xd6\x5e\xd4\xa2\x5b\xb1\x50\xe8\xdf\x01\x00\x00\xff\xff\x29\x98\x17\x0d\x82\x0a\x00\x00")

func _hardcodedMiddlewareGoBytes() ([]byte, error) {
	return bindataRead(
		__hardcodedMiddlewareGo,
		"_hardcoded/middleware.go",
	)
}

func _hardcodedMiddlewareGo() (*asset, error) {
	bytes, err := _hardcodedMiddlewareGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_hardcoded/middleware.go", size: 2690, mode: os.FileMode(436), modTime: time.Unix(1478573361, 0)}
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
	"_hardcoded/doer.go":       _hardcodedDoerGo,
	"_hardcoded/middleware.go": _hardcodedMiddlewareGo,
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
	"_hardcoded": &bintree{nil, map[string]*bintree{
		"doer.go":       &bintree{_hardcodedDoerGo, map[string]*bintree{}},
		"middleware.go": &bintree{_hardcodedMiddlewareGo, map[string]*bintree{}},
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
