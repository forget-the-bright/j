package archiver

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/xi2/xz"
)

func untgx(src string, dst string, strip bool) error {
	xzFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer xzFile.Close()
	var prefixToStrip string
	if strip {
		xzr, err := xz.NewReader(xzFile, 0)
		if err != nil {
			return err
		}
		r := tar.NewReader(xzr)
		var prefix []string
		for {
			header, err := r.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			var dir string
			if header.Typeflag != tar.TypeDir {
				dir = filepath.Dir(header.Name)
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
	xzFile.Seek(0, 0)
	xzr, err := xz.NewReader(xzFile, 0)
	if err != nil {
		return err
	}
	r := tar.NewReader(xzr)
	dirCache := make(map[string]bool) // todo: radix tree would perform better here
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var dir string
		if header.Typeflag != tar.TypeDir {
			dir = filepath.Dir(header.Name)
		} else {
			dir = filepath.Clean(header.Name)
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
		target := filepath.Join(dst, dir, filepath.Base(header.Name))
		switch header.Typeflag {
		case tar.TypeReg:
			d, err := os.OpenFile(target,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode|0600)&0777)
			if err != nil {
				return err
			}
			_, err = io.Copy(d, r)
			d.Close()
			if err != nil {
				return err
			}
		case tar.TypeSymlink:
			if err = os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		}
	}
	return nil
}
