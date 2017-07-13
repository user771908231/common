package encodingUtils

import (
	"encoding/gob"
	"bytes"
)

//序列化
func GobEncode(obj interface{}) (bin []byte, err error) {
	writer := bytes.NewBuffer(bin)
	err = gob.NewEncoder(writer).Encode(obj)
	return writer.Bytes(), err
}

//反序列化
func GobDecode(bin []byte, obj interface{}) (err error) {
	reader := bytes.NewBuffer(bin)
	err = gob.NewDecoder(reader).Decode(obj)
	return
}
