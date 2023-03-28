package check

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func PrintSha256(path string) string {
	strContent := readFileToStr(path)
	sum := sha256.Sum256([]byte(strContent))
	//fmt.Printf("sha256: %x\n", sum)
	return hex.EncodeToString(sum[:])
}

func readFileToStr(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return ""
	}
	return string(content)
}
