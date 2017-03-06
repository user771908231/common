package pushService

import (
	"crypto/md5"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
)

var  sECRET_KEY = []byte{0x93, 0x46, 0x78, 0x20 }

func md5d(data []byte) []byte{
	//log.T("需要校验的data%v",data)
	md5data := append(data, sECRET_KEY[0], sECRET_KEY[1], sECRET_KEY[2], sECRET_KEY[3])
	h := md5.New()
	h.Write(md5data)
	resultByte := h.Sum(nil)
	//log.T("校验计算出来的完整md5:%v",resultByte)

	var resultByte4 []byte
	resultByte4 = append(resultByte4,resultByte[4],resultByte[6],resultByte[8],resultByte[10])
	//log.T("校验的结果sign:%v",resultByte4)
	return resultByte4
}

func md5IdAndData(id,data []byte) []byte{
	//fmt.Println("id:",id)
	//fmt.Println("data:",data)
	//fmt.Println("SECRET_KEY",SECRET_KEY)
	m2 := make([]byte, 2+len(data))
	copy(m2[0:2],id)
	copy(m2[2:],data)
	return md5d(m2)
}

func AssembleData(idv ddproto.HallEnumProtoId, data proto.Message) []byte {
	//fmt.Println("需要转化的id",idv)
	//1,把id转化成 []byte
	id := make([]byte, 2)
	idv2 := uint16(idv)
	binary.BigEndian.PutUint16(id, idv2)
	//	//2,data 转化成[]byte
	data2, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}

	//3计算md5
	md5byte := md5IdAndData(id, data2)
	leng := len(data2)
	m2 := make([]byte, 4 + leng + 4)
	binary.BigEndian.PutUint16(m2, uint16(2 + leng + 4))                //// 默认使用大端序
	copy(m2[2:4], id)
	copy(m2[4:leng + 4], data2)
	copy(m2[leng + 4:], md5byte)

	//4,返回值
	//fmt.Println("发送的m2:", m2)
	return m2
}
