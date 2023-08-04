package repository

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

func decodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
	bytes, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	array := strings.Split(string(bytes), ",")
	if len(array) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	res, err = time.Parse(time.RFC3339Nano, array[0])
	if err != nil {
		return
	}
	
	uuid = array[1]
	return
}

func encodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}