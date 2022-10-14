package common

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

// UUIDString 生成32位UUID
func UUIDString() string {
	u := uuid.New()
	var buf [32]byte
	hex.Encode(buf[:], u[:])
	return string(buf[:])
}

// GetNanoTimeLabel 获取唯一的时间标识(21位)
func GetNanoTimeLabel() (cur string) {
	now := time.Now()
	cur = fmt.Sprintf("%s%09d", now.Format("060102150405"), now.UnixNano()%1e9)
	return
}

// base64x
func base64x() string {
	ff, _ := os.Open("d:/車牌.jpg")
	defer ff.Close()
	buf := make([]byte, 600000)
	n, _ := ff.Read(buf)
	res := base64.StdEncoding.EncodeToString(buf[:n])
	return res
	/*
		//写入临时文件
		ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)
		//读取临时文件
		cc, _ := ioutil.ReadFile("a.png.txt")

		//解压
		dist, _ := base64.StdEncoding.DecodeString(string(cc))
		//写入新文件
		f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(dist)
	*/
}