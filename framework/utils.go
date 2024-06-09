package framework

import (
	"encoding/hex"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

func parseUUID(src string) (dst [16]byte, err error) {
	var uuidBuf [32]byte
	srcBuf := uuidBuf[:]

	switch len(src) {
	case 36:
		copy(srcBuf[0:8], src[:8])
		copy(srcBuf[8:12], src[9:13])
		copy(srcBuf[12:16], src[14:18])
		copy(srcBuf[16:20], src[19:23])
		copy(srcBuf[20:], src[24:])
	case 32:
		// dashes already stripped, assume valid
		copy(srcBuf, src)

	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	_, err = hex.Decode(dst[:], srcBuf)
	if err != nil {
		return dst, err
	}
	return dst, err
}

func StringToUUID(s string) pgtype.UUID {
	data, err := parseUUID(s)
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}
