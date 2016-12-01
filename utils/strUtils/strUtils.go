package strUtils

//
func Bytes2Str(p []byte) string{
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}


//
func Str2Bytes(s string) []byte{
	return []byte(s)
}