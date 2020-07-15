package utils

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
)

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func SplitInts(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}

	sArr := strings.Split(s, ",")
	res := make([]int64, 0)

	for _, sc := range sArr {
		num, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, num)
	}
	return res, nil
}
