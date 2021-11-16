package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return "", err
	}
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", fmt.Errorf("buffer length is not enough")
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
