package main

import (
	"fmt"
	"tools/str"
)

// func ProdPwd(oldPwd, salt string) string {
// 	oldKey := Md5String(salt)
// 	key := []byte(strings.ToTitle(oldKey[0:8]))
// 	result, err := DesEncrypt([]byte(oldPwd), key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	hexstr := fmt.Sprintf("%X", result)
// 	fmt.Println(hexstr)
// 	return hexstr
// }

func main() {
	// key := []byte{0x01, 0x21, 0x25, 0x17, 0x19, 0xFB, 0xFD, 0xDF} // 8字节密钥
	pwd := "Qaz123!@#"
	salt := "wisfqd"

	enPwd := str.ProdPwd(pwd, salt)

	fmt.Println(">>>", enPwd)

	dePwd := str.PwdDes(enPwd, salt)

	fmt.Println(">>>", dePwd)
}
