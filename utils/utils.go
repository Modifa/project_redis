package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*FloatToString ...  */
func FloatToString(f float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func GenerateRandomNumbers(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

/*FloatToStringDecimal ...  */
func FloatToStringDecimal(f float64) string {
	// to convert a float number to a string
	return FloatToString(math.Round(f*100) / 100)
}

/*FloatToString ...  */
func StringToFloat64(f string) float64 {
	// to convert a float number to a string
	const bitSize = 64 // Don't think about it to much. It's just 64 bits.
	floatNum, _ := strconv.ParseFloat(f, bitSize)

	return floatNum
}

/*FloatToString ...  */
func StringToInt(f string) int {
	// to convert a float number to a string
	floatNum, _ := strconv.Atoi(f)

	return floatNum
}

/*Int64ToString ...  */
func Int64ToString(f int64) string {
	// to convert a float number to a string
	return strconv.FormatInt(f, 10)
}

/*IntToString ...  */
func IntToString(f int) string {
	// to convert a float number to a string
	return strconv.Itoa(f)
}

/*Int64ToString ...  */
func IntToInt64(f int) int64 {
	// to convert a float number to a string
	return int64(f)
}

/*StringToBase64 ...  */
func StringToBase64(f string) string {
	// to convert a float number to a string
	sEnc := base64.StdEncoding.EncodeToString([]byte(f))
	return sEnc
}

/*StringToBase64 ...  */
func Base64ToString(f string) (string, error) {
	// to convert a float number to a string
	sDec, err := base64.StdEncoding.DecodeString(f)
	return string(sDec), err
}

/*Int64ToString ...  */
// func ObjectToJSON(f []byte) string {
// 	// to convert a float number to a string

// 	if err := json.Unmarshal(f, &pwh); err != nil {
// 		c.JSON(http.StatusBadRequest, rb.New(false, "", err))
// 		return
// 	}

// 	return int64(f)
// }

func IsLaunchedByDebugger() bool {
	// gops executable must be in the path. See https://github.com/google/gops
	gopsOut, err := exec.Command("gops", strconv.Itoa(os.Getppid())).Output()
	if err == nil && strings.Contains(string(gopsOut), "\\dlv.exe") {
		// our parent process is (probably) the Delve debugger
		return true
	}
	return false
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func FormatMobileNumber(message string) (string, error) {

	countrycode := message[:len(message)-9]
	fmt.Println(countrycode)

	if len(message) == 11 && countrycode == "27" {
		return message, nil
	} else if len(message) == 10 && countrycode == "0" {

		message = "27" + message[1:10]
		return message, nil
	} else {
		err1 := errors.New("invalid mobile number")
		return "", err1
	}
}

func IsFormatMobileNumber(value interface{}) error {

	s, _ := value.(string)
	_, err := FormatMobileNumber(s)

	if err != nil {
		return errors.New("Invalid Mobile Number")
	}
	return nil
}

func HMAC256ENCRYPT(data string) string {

	secret := MustGetenv("VOUCHERMANAGERSECRET")
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func HashAuthCode(Authcode string) string {

	secret := MustGetenv("VOUCHERMANAGERSECRET")
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(Authcode))
	return StringToBase64(hex.EncodeToString(h.Sum(nil)) + ":" + Authcode)
}

func ValidateAuthCode(HashedAuthcode string) (string, bool) {

	unHashedAuthcode, err := Base64ToString(HashedAuthcode)

	sp := strings.Split(unHashedAuthcode, ":")

	if err == nil && len(sp) > 1 {

		HashAuth := strings.Split(unHashedAuthcode, ":")[0]
		Authcode := strings.Split(unHashedAuthcode, ":")[1]

		if HashAuthCode(Authcode) == HashAuth {
			return Authcode, true
		} else {
			return "", false
		}
	} else {
		return "", false
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncryptString(data string) string {

	// return hex string
	return hex.EncodeToString([]byte(data))
}

func DecryptString(ed string) string {

	dbytes, err := hex.DecodeString(ed)

	CheckError(err)

	return string(dbytes)
}

func Encrypt(data json.RawMessage) string {

	// return hex string
	return hex.EncodeToString(data)
}

func Decrypt(ed string) json.RawMessage {

	dbytes, err := hex.DecodeString(ed)

	CheckError(err)

	return json.RawMessage(dbytes)
}

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func PadWithZeros(length string, k string) string {

	i, err := strconv.Atoi("12345")

	if err != nil {
		// our parent process is (probably) the Delve debugger
		log.Printf("PadWithZeros: %v", err)
	}

	return fmt.Sprintf("%06d", i)
}

func StructToJSON1(in string) string {

	rawIn := json.RawMessage(in)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func StructToJSONRaw(in interface{}) json.RawMessage {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(in)

	rawIn := json.RawMessage(reqBodyBytes.String())

	return rawIn
}

func StructToJSON(response interface{}) string {

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(response)

	return StructToJSON1(reqBodyBytes.String())
}
