package MD5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

//加密成MD5
func GetMD5(data string)string  {
	h := md5.New()
	_, _ = io.WriteString(h, data)
	//h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
