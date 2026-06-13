package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func New() string {
	const Retries = 3
	for i := 0; i < Retries; i++ {
		// Assumption no repeating of UUID
		uuid, err := generateUUID()
		if err != nil {
			// error while generating UUID retrying
			continue
		}

		// check if uuid exists
		return uuid
	}

	x, _ := generateUUIDFromNano()
	return x
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Set the 4 most significant bits of the 7th byte to 0100 (Version 4)
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// Set the 2 most significant bits of the 9th byte to 10 (Variant 1)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// Format matching the canonical 8-4-4-4-12 string layout
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), nil
}

func generateUUIDFromNano() (string, error) {
	// 1. Get current nanoseconds as an int64
	nanoseconds := time.Now().UnixNano()

	// 2. Create a 16-byte slice for the UUID
	uuidBytes := make([]byte, 16)

	// 3. Put the 8-byte nanosecond integer into the first 8 bytes
	binary.BigEndian.PutUint64(uuidBytes[0:8], uint64(nanoseconds))

	// 4. Fill the remaining 8 bytes with crypto random data to ensure uniqueness
	_, err := rand.Read(uuidBytes[8:16])
	if err != nil {
		return "", err
	}

	// 5. Apply RFC 9562 compliance bitmasks
	// Set Version 4 (0100) on the 7th byte
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40
	// Set Variant 1 (10) on the 9th byte
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80

	// 6. Format into the canonical 8-4-4-4-12 string layout
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuidBytes[0:4], uuidBytes[4:6], uuidBytes[6:8], uuidBytes[8:10], uuidBytes[10:16]), nil
}
