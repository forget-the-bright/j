package archiver

import (
	"errors"
	"strings"
)

func Run_unzip(src string, dst string, strip bool) error {
	return Unarchive(src, dst, strip)
}

func Unarchive(src string, dst string, strip bool) (err error) {
	fileTypes := strings.Split(src, ".")
	if len(fileTypes) == 0 {
		return errors.New("fileType is not supported")
	}
	//fmt.Println(fileTypes[len(fileTypes)-1])
	switch fileTypes[len(fileTypes)-1] {
	case "gz":
		err = untgz(src, dst, strip)
	case "gx":
		err = untgx(src, dst, strip)
	case "zip":
		err = unzip(src, dst, strip)
	default:
		err = errors.New(fileTypes[len(fileTypes)-1] + " is not supported")
	}
	return err
}
