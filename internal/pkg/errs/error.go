package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrVersionNotFound 版本不存在
	ErrVersionNotFound = errors.New("version not found")
	// ErrPackageNotFound 版本包不存在
	ErrPackageNotFound = errors.New("installation package not found")
)

var (
	// ErrUnsupportedChecksumAlgorithm 不支持的校验和算法
	ErrUnsupportedChecksumAlgorithm = errors.New("unsupported checksum algorithm")
	// ErrChecksumNotMatched 校验和不匹配
	ErrChecksumNotMatched = errors.New("file checksum does not match the computed checksum")
	// ErrChecksumFileNotFound 校验和文件不存在
	ErrChecksumFileNotFound = errors.New("checksum file not found")
)

// URLUnreachableError URL不可达错误
type URLUnreachableError struct {
	err error
	url string
}

// NewURLUnreachableError 返回URL不可达错误实例
func NewURLUnreachableError(url string, err error) error {
	return &URLUnreachableError{
		err: err,
		url: url,
	}
}

func (e URLUnreachableError) Error() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("URL %q is unreachable", e.url))
	if e.err != nil {
		buf.WriteString(" ==> " + e.err.Error())
	}
	return buf.String()
}

func (e URLUnreachableError) Err() error {
	return e.err
}

func (e URLUnreachableError) URL() string {
	return e.url
}

// DownloadError 下载失败错误
type DownloadError struct {
	url string
	err error
}

// NewDownloadError 返回下载失败错误实例
func NewDownloadError(url string, err error) error {
	return &DownloadError{
		url: url,
		err: err,
	}
}

// Error 返回错误字符串
func (e DownloadError) Error() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("Resource(%s) download failed", e.url))
	if e.err != nil {
		buf.WriteString(" ==> " + e.err.Error())
	}
	return buf.String()
}

// Err 返回错误对象
func (e DownloadError) Err() error {
	return e.err
}

// URL 返回资源URL
func (e DownloadError) URL() string {
	return e.url
}
