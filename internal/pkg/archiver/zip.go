package archiver

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func unzip(src string, dst string, strip bool) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	var prefixToStrip string
	if strip {
		var prefix []string
		for _, f := range r.File {
			var dir string
			if !f.Mode().IsDir() {
				dir = filepath.Dir(f.Name)
			} else {
				continue
			}
			if prefix != nil {
				dirSplit := strings.Split(dir, string(filepath.Separator))
				i, e, dse := 0, len(prefix), len(dirSplit)
				if dse < e {
					e = dse
				}
				for i < e {
					if prefix[i] != dirSplit[i] {
						prefix = prefix[0:i]
						break
					}
					i++
				}
			} else {
				prefix = strings.Split(dir, string(filepath.Separator))
			}
		}
		prefixToStrip = strings.Join(prefix, string(filepath.Separator))
	}
	dirCache := make(map[string]bool) // todo: radix tree would perform better here
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	for _, f := range r.File {
		var dir string
		if !f.Mode().IsDir() {
			dir = filepath.Dir(f.Name)
		} else {
			dir = filepath.Clean(f.Name)
			if !strings.HasPrefix(dir, prefixToStrip) {
				continue
			}
		}
		dir = strings.TrimPrefix(dir, prefixToStrip)
		if dir != "" && dir != "." {
			cached := dirCache[dir]
			if !cached {
				if err := os.MkdirAll(filepath.Join(dst, dir), 0755); err != nil {
					return err
				}
				dirCache[dir] = true
			}
		}
		if !f.Mode().IsDir() {
			name := filepath.Base(f.Name)
			fr, err := f.Open()
			if err != nil {
				return err
			}
			d, err := os.OpenFile(filepath.Join(dst, dir, name),
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC, (f.Mode()|0600)&0777)
			if err != nil {
				return err
			}
			_, err = io.Copy(d, fr)
			d.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
