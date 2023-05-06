package link

import (
    "crypto/rand"
    "encoding/base64"
)

func RandomString(length int) string {
    bytes := make([]byte, length)
    rand.Read(bytes)
    return base64.StdEncoding.EncodeToString(bytes)[:length]
}
