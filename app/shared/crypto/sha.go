package sha

import (
	"crypto/sha256"
	"fmt"
)

// ToSHA256String converts string to SHA256 string
func ToSHA256String(s string) string {
	bs := []byte(s + "key:4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d")
	return fmt.Sprintf("%x", (sha256.Sum256(bs)))
}
