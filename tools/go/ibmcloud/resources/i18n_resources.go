// Code generated by go-bindata.
// sources:
// i18n/resources/de_DE.all.json
// i18n/resources/en_US.all.json
// i18n/resources/es_ES.all.json
// i18n/resources/fr_FR.all.json
// i18n/resources/it_IT.all.json
// i18n/resources/ja_JA.all.json
// i18n/resources/ko_KR.all.json
// i18n/resources/pt_BR.all.json
// i18n/resources/zh_Hans.all.json
// i18n/resources/zh_hant.all.json
// DO NOT EDIT!

package resources

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

var _i18nResourcesDe_deAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesDe_deAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesDe_deAllJson,
		"i18n/resources/de_DE.all.json",
	)
}

func i18nResourcesDe_deAllJson() (*asset, error) {
	bytes, err := i18nResourcesDe_deAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/de_DE.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesEn_usAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesEn_usAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesEn_usAllJson,
		"i18n/resources/en_US.all.json",
	)
}

func i18nResourcesEn_usAllJson() (*asset, error) {
	bytes, err := i18nResourcesEn_usAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/en_US.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesEs_esAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesEs_esAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesEs_esAllJson,
		"i18n/resources/es_ES.all.json",
	)
}

func i18nResourcesEs_esAllJson() (*asset, error) {
	bytes, err := i18nResourcesEs_esAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/es_ES.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesFr_frAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesFr_frAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesFr_frAllJson,
		"i18n/resources/fr_FR.all.json",
	)
}

func i18nResourcesFr_frAllJson() (*asset, error) {
	bytes, err := i18nResourcesFr_frAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/fr_FR.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesIt_itAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesIt_itAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesIt_itAllJson,
		"i18n/resources/it_IT.all.json",
	)
}

func i18nResourcesIt_itAllJson() (*asset, error) {
	bytes, err := i18nResourcesIt_itAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/it_IT.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesJa_jaAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesJa_jaAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesJa_jaAllJson,
		"i18n/resources/ja_JA.all.json",
	)
}

func i18nResourcesJa_jaAllJson() (*asset, error) {
	bytes, err := i18nResourcesJa_jaAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/ja_JA.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesKo_krAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesKo_krAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesKo_krAllJson,
		"i18n/resources/ko_KR.all.json",
	)
}

func i18nResourcesKo_krAllJson() (*asset, error) {
	bytes, err := i18nResourcesKo_krAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/ko_KR.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesPt_brAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesPt_brAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesPt_brAllJson,
		"i18n/resources/pt_BR.all.json",
	)
}

func i18nResourcesPt_brAllJson() (*asset, error) {
	bytes, err := i18nResourcesPt_brAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/pt_BR.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesZh_hansAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesZh_hansAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesZh_hansAllJson,
		"i18n/resources/zh_Hans.all.json",
	)
}

func i18nResourcesZh_hansAllJson() (*asset, error) {
	bytes, err := i18nResourcesZh_hansAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/zh_Hans.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _i18nResourcesZh_hantAllJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\x31\x0e\xc2\x20\x18\xc5\xf1\x9d\x53\xbc\x7c\x73\xc3\x01\xd8\x1c\x1c\x3c\x83\x75\x20\x2d\x69\x49\xc8\xc3\x7c\xd0\x89\x70\x77\xa3\x6e\xc6\x36\xdd\xff\xef\xf7\xee\x06\x68\x06\x00\x24\xce\xe2\x20\x37\x96\xea\x53\x8a\x5c\x50\xd7\x80\x67\xda\x96\x48\x6b\xad\x0c\xdf\xaa\xaa\x67\x49\xbe\xc6\x4c\x71\x38\xe8\x0d\xd0\x87\x5f\xfd\x42\x04\xd5\xac\x58\x7d\x41\x9e\xa6\x4d\x35\xcc\x3b\xb6\x5c\xff\x94\xe7\x55\x37\xb2\x35\xfb\x31\x7a\x1f\xb9\xf3\x71\x6a\xfa\x3e\x35\x8f\x57\x00\x00\x00\xff\xff\xd1\x02\xce\x3e\x2b\x01\x00\x00")

func i18nResourcesZh_hantAllJsonBytes() ([]byte, error) {
	return bindataRead(
		_i18nResourcesZh_hantAllJson,
		"i18n/resources/zh_hant.all.json",
	)
}

func i18nResourcesZh_hantAllJson() (*asset, error) {
	bytes, err := i18nResourcesZh_hantAllJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "i18n/resources/zh_hant.all.json", size: 299, mode: os.FileMode(420), modTime: time.Unix(1533113068, 0)}
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
	"i18n/resources/de_DE.all.json": i18nResourcesDe_deAllJson,
	"i18n/resources/en_US.all.json": i18nResourcesEn_usAllJson,
	"i18n/resources/es_ES.all.json": i18nResourcesEs_esAllJson,
	"i18n/resources/fr_FR.all.json": i18nResourcesFr_frAllJson,
	"i18n/resources/it_IT.all.json": i18nResourcesIt_itAllJson,
	"i18n/resources/ja_JA.all.json": i18nResourcesJa_jaAllJson,
	"i18n/resources/ko_KR.all.json": i18nResourcesKo_krAllJson,
	"i18n/resources/pt_BR.all.json": i18nResourcesPt_brAllJson,
	"i18n/resources/zh_Hans.all.json": i18nResourcesZh_hansAllJson,
	"i18n/resources/zh_hant.all.json": i18nResourcesZh_hantAllJson,
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
	"i18n": &bintree{nil, map[string]*bintree{
		"resources": &bintree{nil, map[string]*bintree{
			"de_DE.all.json": &bintree{i18nResourcesDe_deAllJson, map[string]*bintree{}},
			"en_US.all.json": &bintree{i18nResourcesEn_usAllJson, map[string]*bintree{}},
			"es_ES.all.json": &bintree{i18nResourcesEs_esAllJson, map[string]*bintree{}},
			"fr_FR.all.json": &bintree{i18nResourcesFr_frAllJson, map[string]*bintree{}},
			"it_IT.all.json": &bintree{i18nResourcesIt_itAllJson, map[string]*bintree{}},
			"ja_JA.all.json": &bintree{i18nResourcesJa_jaAllJson, map[string]*bintree{}},
			"ko_KR.all.json": &bintree{i18nResourcesKo_krAllJson, map[string]*bintree{}},
			"pt_BR.all.json": &bintree{i18nResourcesPt_brAllJson, map[string]*bintree{}},
			"zh_Hans.all.json": &bintree{i18nResourcesZh_hansAllJson, map[string]*bintree{}},
			"zh_hant.all.json": &bintree{i18nResourcesZh_hantAllJson, map[string]*bintree{}},
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
