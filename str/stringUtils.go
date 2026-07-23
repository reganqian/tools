package str

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/xid"
	// . "yunapi/log"
)

const chars = "AD1JSTXFGPHUB2CVW6ILK9QRE3M7N845YZ"

func UserIDToCode(userID uint, length int) string {
	if userID <= 0 {
		panic("userID must be positive")
	}
	userID += 100000

	var result []byte
	for userID > 0 {
		result = append(result, chars[userID%36])
		userID /= 36
	}

	// 反转
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	s := string(result)

	// 截断（理论上不会触发）
	if len(s) > length {
		s = s[len(s)-length:]
	}

	return s
}

func GenerateShortRandomString(str string, length int) string {
	if length <= 0 {
		length = 8 // 默认 8 位
	}

	// 1. 用 SHA-256 处理 userID，得到稳定哈希
	hash := sha256.Sum256([]byte(str))

	// 2. 用哈希前 8 字节做随机种子
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	rng := rand.New(rand.NewSource(seed))

	// 3. 62 字符集：a-z A-Z 0-9
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 4. 生成短随机串
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rng.Intn(len(charset))]
	}
	return string(result)
}

// GetImageMD5FromURL 从网络URL获取图片并生成MD5
func GetImageMD5FromURL(imageURL string) (string, error) {

	resp, err := http.Get(imageURL)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, resp.Body); err != nil {
		return "", fmt.Errorf("读取响应体失败: %w", err)
	}

	md5Bytes := hash.Sum(nil)
	data := hex.EncodeToString(md5Bytes)

	return data, nil
}

func UniqAppend(dataList []string, data string) []string {
	for _, data1 := range dataList {
		if data1 == data {
			return dataList
		}
	}
	dataList = append(dataList, data)
	return dataList
}

func GetRandomStringWithStr(inStr string, l int) string {
	oldStr := Md5String(inStr)
	str := "0123456789abcdefghijklmnopqrstuvwxyz" + oldStr
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func GetRandomNum(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		br := r.Intn(10)
		result = append(result, bytes[br])
	}
	return string(result)
}

func CreateMd5String(addStr, salt string) string {

	m5 := md5.New()

	m5.Write([]byte(addStr))
	m5.Write([]byte(salt))

	st := m5.Sum(nil)
	token := hex.EncodeToString(st)

	return string(token)
}

func Md5String(addStr string) string {
	m5 := md5.New()
	m5.Write([]byte(addStr))
	st := m5.Sum(nil)
	token := hex.EncodeToString(st)
	return string(token)
}

func ProdPwd(oldPwd, salt string) string {
	oldKey := Md5String(salt)
	key := []byte(strings.ToTitle(oldKey[0:8]))
	result, err := DesEncrypt([]byte(oldPwd), key)
	if err != nil {
		panic(err)
	}
	hexstr := fmt.Sprintf("%X", result)
	fmt.Println(hexstr)
	return hexstr
}

func PwdDes(pwd, salt string) string {
	oldKey := Md5String(salt)
	key := []byte(strings.ToTitle(oldKey[0:8]))
	oldPwd := pwd
	desStr := HexToString(pwd)

	tryErr := func() (_e error) {

		defer func() {
			if r := recover(); r != nil {
				// fmt.Println("Recovered from panic:", r)
				// _e = errors.New(r.(string))
				_e = fmt.Errorf("Decrypt error: %v", r)
				return
			}
		}()

		result, err := DesDecrypt([]byte(desStr), key)
		if err != nil {
			//如果异常, 则返回空字符串
			return err
		}
		oldPwd = fmt.Sprintf("%s", result)

		return nil
	}()

	fmt.Println(oldPwd)
	if tryErr != nil {
		return oldPwd
	}

	return oldPwd
}

func HexToString(hexStr string) string {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		fmt.Println("Error decoding hex string:", err)
		return ""
	}
	return string(bytes)
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {

		return nil, err
	}

	return m, nil
}

// Convert map json string
func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {

		return "", nil
	}

	return string(jsonByte), nil
}

func MakeTimestamp() int64 {

	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}

// 生成订单号：YYYYMMDDHHmmssSSS
func CreateOrderNo(orderPre string) string {
	date := GetTodyDay()
	data := GetTimeTick64()
	code := fmt.Sprintf("%s%s%s%s", orderPre, date, data, GetRandomNum(3))
	return code
}

// 生成订单号：YYYYMMDDHHmmssSSS
func CreateOrderNoWithUserId(orderPre string, userId uint) string {
	date := GetTodyDay()
	data := GetTimeTick64()
	randStr := GetRandomStringWithStr(strconv.Itoa(int(userId)), 8)
	code := fmt.Sprintf("%s%s%s%s", orderPre, randStr, date, data)
	return code
}

func GetTodyDay() string {
	timeTemplate := "060102"
	return time.Now().UTC().Format(timeTemplate)
}

func GetTimeTick64() string {
	intData := time.Now().UnixNano() / 1e6
	s := strconv.FormatInt(intData, 10)
	content := s[4 : len(s)-1]
	return content
}

func GetIdStr() string {
	return xid.New().String()
}

func FormatTimeToStr(t time.Time) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	return t.Format(timeTemplate1)
}

func FormatTimeToStrForScript(t time.Time) string {
	timeTemplate1 := "2006-01-02T15:04:05"
	return t.Format(timeTemplate1)
}

// func FormatTimeToStrWithZone(t time.Time, zone int) string {
// 	timeTemplate1 := "2006-01-02 15:04:05"
// 	loc := time.FixedZone("UTC-8", zone * 60 * 60)
// 	t = t.In(loc)
// 	return  t.Format(timeTemplate1)
// }

func GetTimeUnix(t time.Time) int64 {
	return t.UTC().Unix()
}

func GetTodayStr() string {
	timeTemplate := "20060102"
	return time.Now().UTC().Format(timeTemplate)
}

func GetTodayHourStr() string {
	timeTemplate := "2006010215"
	return time.Now().UTC().Format(timeTemplate)
}

func GetYesterdayNumStr() string {
	inTime := time.Now().Unix()
	timeTemplate := "20060102"
	yesterday := time.Unix(inTime, 0).AddDate(0, 0, -1)
	return yesterday.Format(timeTemplate)
}

// 1明天, -1昨天
func GetDayNumStr(addNum int) string {
	if addNum == 0 {
		return GetTodayStr()
	}

	inTime := time.Now().Unix()
	timeTemplate := "20060102"
	yesterday := time.Unix(inTime, 0).AddDate(0, 0, addNum)
	return yesterday.Format(timeTemplate)
}

func GetTimeStr(inTime time.Time) string {
	timeTemplate := "2006-01-02"
	return inTime.Format(timeTemplate)
}

func GetTimeMonthAndDay(inTime time.Time) (timeMonth string, timeDay int) {
	timeTemplate := "2006-01"
	timeMonth = inTime.Format(timeTemplate)
	timeDay = inTime.Day()
	return timeMonth, timeDay
}

// 查询昨天
func GetPointTimeStr(inTime time.Time, addNum int) string {
	timeTemplate := "2006-01-02"
	yesterday := inTime.AddDate(0, 0, addNum)
	return yesterday.Format(timeTemplate)
}

// 查询昨天
func GetYesterdayStr(inTime time.Time) string {
	timeTemplate := "2006-01-02"
	yesterday := inTime.AddDate(0, 0, -1)
	return yesterday.Format(timeTemplate)
}

func GetTimeHourStr(inTime time.Time) string {
	timeTemplate := "2006-01-02 15"
	return inTime.Format(timeTemplate)
}

func GetTimeMonthStr(inTime time.Time) string {
	timeTemplate := "2006-01"
	return inTime.Format(timeTemplate)
}

func GetNowStr() string {
	timeTemplate := "20060102150405"
	return time.Now().UTC().Format(timeTemplate)
}

type ProxyTag struct {
	Proxy string `json:"proxy"`
}

func (s *ProxyTag) AppendTag(tag string) {
	if s.Proxy == "" {
		s.Proxy = tag
	} else {
		s.Proxy = s.Proxy + " && " + tag
	}
}

func (s *ProxyTag) GetBase64Str() string {
	return GetBase64Data(s.Proxy)
}

func GetProxyTag(ipId, region, ipUrl string) ProxyTag {
	proxy := ProxyTag{}
	if ipId != "" {
		tag := "uid==\"" + ipId + "\""
		proxy.AppendTag(tag)
	} else {
		if ipUrl != "" {
			tag := "url==\"" + ipUrl + "\""
			// tag := "ext.at(\"url\")==\"" + region + "\""
			proxy.AppendTag(tag)
		}
		if region != "" {
			tag := "ext.has(\"region\") && ext.at(\"region\")==\"" + region + "\""
			proxy.AppendTag(tag)
		}
	}
	return proxy
}

func GetProxyGroup(userId uint32) string {
	userIdStr := strconv.Itoa(int(userId))
	str := "ext.at(\"group\") == \"" + userIdStr + "\""
	return GetBase64Data(str)
}

func GetProxyGroupNot64(userId uint32) string {
	userIdStr := strconv.Itoa(int(userId))
	str := "ext.at(\"group\") == \"" + userIdStr + "\""
	return str
}

func GetBase64Data(data string) string {
	input := []byte(data)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return string(encodeString)
}

func GetTodayZeroTime() time.Time {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := time.Now().UTC().Format("2006-01-02")
	t2 := timeStr + " 00:00:00"
	stamp, _ := time.ParseInLocation(timeTemplate1, t2, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp
}

func IntsToUints(dataList []int) (resList []uint) {
	for _, data := range dataList {
		res := uint(data)
		resList = append(resList, res)
	}

	return resList
}

func GetPointTime(inTime time.Time, pointStr string) time.Time {
	timeTemplate1 := "2006-01-02 15:04:05"
	timeStr := inTime.Format("2006-01-02")
	t2 := timeStr + " " + pointStr
	stamp, _ := time.ParseInLocation(timeTemplate1, t2, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp
}

func GetTodayZeroStr() string {
	timeStr := time.Now().UTC().Format("2006-01-02")
	t2 := timeStr + " 00:00:00"
	return t2
}

func FormatStringToInt(data string) (int, error) {
	j, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(j), nil
}

func FormatStringToInt32(data string) (int32, error) {
	j, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(j), nil
}

// 获取一天前的字符串
func GetTimeBefore(inTime time.Time) time.Time {
	next := inTime.Add(-time.Hour * 24)
	return next
}

func GetZeroTime(inTime time.Time) int64 {
	return time.Date(inTime.Year(), inTime.Month(), inTime.Day(), 0, 0, 0, 0, inTime.Location()).Unix()
}

func GetLastTime(inTime time.Time) int64 {
	return time.Date(inTime.Year(), inTime.Month(), inTime.Day(), 23, 59, 59, 0, inTime.Location()).Unix()
}

func GetTimeBeforeMonth(inTime time.Time) time.Time {
	next := inTime.AddDate(0, -1, 0)
	return next
}

// strconv.Itoa(int(i))
func FormatInt32ToStr(data int32) string {
	return strconv.Itoa(int(data))
}

// 是否是不重复的数组
func CheckSameFriend(dataList []string) bool {
	fmap := make(map[string]string)
	for _, fid := range dataList {
		dbD := fmap[fid]
		if dbD == "" {
			fmap[fid] = fid
		} else {
			return false
		}
	}
	return true
}

func CheckInStrings(data string, datas []string) bool {
	for _, v := range datas {
		if v == data {
			return true
		}
	}

	return false
}

func StringsToUint32s(accidList []string) []uint32 {
	var intList []uint32
	for _, accStr := range accidList {
		data, _ := strconv.Atoi(accStr)
		intList = append(intList, uint32(data))
	}
	return intList
}

func StringsToUint64s(strList []string) []uint64 {
	var intList []uint64
	for _, str := range strList {
		su64, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			intList = append(intList, su64)
		}

	}
	return intList
}

func StringToUint64(str string) (uint64, error) {
	su64, err := strconv.ParseUint(str, 10, 64)
	return su64, err
}

func StringsToUints(accidList []string) []uint {
	var intList []uint
	for _, accStr := range accidList {
		data, err := strconv.Atoi(accStr)
		if err == nil {
			intList = append(intList, uint(data))
		}
	}
	return intList
}

func StrToTime(str string) time.Time {

	var timeLayoutStr = "2006-01-02 15:04:05"
	st, err := time.Parse(timeLayoutStr, str) //string转StrToTime
	if err != nil {
		var timeLayoutStr = "2006-1-2 15:04:05"
		st, err := time.Parse(timeLayoutStr, str) //string转StrToTime
		if err != nil {
			return time.Now()
		}
		return st
	}

	return st
}

// CalculateAge 根据生日计算年龄
func CalculateAge(birthday string) (int, error) {
	// 解析生日字符串
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, err
	}
	// 获取当前时间
	now := time.Now()
	// 计算年份差
	age := now.Year() - birthDate.Year()
	// 判断生日是否已经过了
	if now.YearDay() < birthDate.YearDay() {
		age--
	}
	return age, nil
}

func StrTimeAddYear(str string, addType string, addVal int) (targetTime time.Time) {
	d := StrToTime(str)
	if addType == "add" {
		targetTime = d.AddDate(addVal, 0, 0)
	} else {
		targetTime = d.AddDate(0-addVal, 0, 0)
	}
	return targetTime
}

func TimeAddYear(d time.Time, addType string, addVal int) (targetTime time.Time) {
	if addType == "add" {
		targetTime = d.AddDate(addVal, 0, 0)
	} else {
		targetTime = d.AddDate(0-addVal, 0, 0)
	}
	return targetTime
}

func Int64ToTime(timeInt int64) time.Time {

	return time.Unix(timeInt, 0)
}

func RecerseList(x []interface{}) []interface{} {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return x
}

func TimeIntToTimeStr(timeInt int64) string {
	loc := time.FixedZone("UTC-8", 8*60*60)
	d := time.Unix(timeInt, 0).In(loc)
	timeTemplate := "2006-01-02 15:04:05"
	return d.Format(timeTemplate)

}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	// pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	// 定义邮箱正则表达式
	emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// // 判断是否符合邮箱格式
	// if regexp.MatchString(email) {
	// 	fmt.Println("是邮箱")
	// } else {
	// 	fmt.Println("不是邮箱")
	// }
	// // reg := regexp.MustCompile(pattern)

	return emailRegexp.MatchString(email)
}

// 获取一个月最后一天的年月日
func GetLastTimeOfMonth(timeStr string) string {
	inTime := StrToTime(timeStr)
	lastTime := inTime.AddDate(0, 1, -1)
	lastTimeStr := GetTimeStr(lastTime)
	return lastTimeStr
}

type KeyValData struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func DealKeyValData(queryStr string) (dataList []KeyValData) {
	strList := strings.Split(queryStr, ",")
	for _, str := range strList {
		kvs := strings.Split(str, ":")
		if len(kvs) == 2 {
			k := kvs[0]
			v := kvs[1]
			data := KeyValData{}
			data.Key = k
			data.Val = v
			dataList = append(dataList, data)
		}
	}

	return dataList
}

func JsonStringToMap(body string) (map[string]interface{}, error) {
	if body == "" {
		return nil, errors.New("data is null")
	}
	var result map[string]interface{}
	s := fmt.Sprintf(body)
	bb := []byte(s)
	err := json.Unmarshal(bb, &result)
	return result, err
}

func StrInStrs(target string, str_array []string) bool {

	for _, element := range str_array {

		if target == element {

			return true

		}

	}

	return false

}

// 获取上个月的开始和结束的时间戳
func GetLastMonthStartAndEndSec() (startTime, endTime int64) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	startTime = time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
	endTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix() - 1

	return startTime, endTime
}

func GetLastMonthStr() string {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	timeStr := GetTimeMonthStr(lastMonth)
	return timeStr

}

// getRandomNumber 生成 m 到 n 之间的随机数
func GetRandomNumber(m, n int) int {
	if m > n {
		m, n = n, m
	}

	return m + r.Intn(n-m+1)
}

func CreateKey() string {
	keyLength := 32
	key := make([]byte, keyLength)
	// 生成随机字节
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}
	// 将字节转换为十六进制字符串
	keyHex := hex.EncodeToString(key)
	fmt.Println("Generated JWT Key:", keyHex)
	return keyHex
}

func HiddenStr(s string) string {
	rs := []rune(s)
	n := len(rs)

	switch {
	case n == 0:
		return ""
	case n == 1:
		return "*"
	case n <= 4:
		// 除首字符外都替换，保留第一个字符
		return string(rs[0]) + strings.Repeat("*", n-1)
	case n <= 7:
		// 保留首尾各一个字符
		middleLen := n - 2
		return string(rs[0]) + strings.Repeat("*", middleLen) + string(rs[n-1])
	default:
		// 保留前2个和后2个字符
		middleLen := n - 4
		return string(rs[:2]) + strings.Repeat("*", middleLen) + string(rs[n-2:])
	}
}
