package crypto

import (
	"encoding/json"
	"fmt"
)

func CreateJWT(header map[string]interface{}, payload map[string]interface{}) string {
	headerMapData, err := json.Marshal(header)
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}
	headerMapStr := string(headerMapData)
	fmt.Println(headerMapStr)

	payloadMapData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}
	payloadMapStr := string(payloadMapData)
	fmt.Println(payloadMapStr)

	b64Header := Base64Encode(headerMapStr)
	b64Payload := Base64Encode(payloadMapStr)

	verify := b64Header + "." + b64Payload
	sha := ComputeHMAC(verify, "256-bit-secret")

	return verify + "." + sha
}
