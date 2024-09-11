package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"strconv"
	"strings"
)

const base64Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Ceil(a int, b int) int {
	return (a + b - 1) / b
}

func CharToBin(c byte) string {
	ascii := int(c)

	bin := strconv.FormatInt(int64(ascii), 2)

	for len(bin) < 8 {
		bin = "0" + bin
	}
	return bin
}

func BinToB64(bin string) byte {
	if i, err := strconv.ParseInt(bin, 2, 64); err != nil {
		panic("aaaaaah!")
	} else {
		return base64Charset[i]
	}
}

func Base64Encode(input string) string {
	inputChars := []byte(input)

	var binary string
	for _, c := range inputChars {
		binary = binary + CharToBin(c)
	}

	chunksLength := Ceil(len(binary), 6)
	// chunksLength := len(binary) / 6
	var chunks []string

	for i := 0; i < chunksLength; i++ {
		startIndex := i * 6
		endIndex := startIndex + 6

		if endIndex > len(binary) {
			endIndex = len(binary)
		}

		chunk := binary[startIndex:endIndex]

		if len(chunk) < 6 {
			chunk += strings.Repeat("0", 6-len(chunk))
		}
		chunks = append(chunks, chunk)
	}

	var b64Encoded string

	for _, chunk := range chunks {
		dec := BinToB64(chunk)
		b64Encoded = b64Encoded + string(dec)
	}

	return b64Encoded
}

func ComputeHMAC(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))

	h.Write([]byte(message))

	sha := Base64Encode(string(h.Sum(nil)))

	return sha
}
