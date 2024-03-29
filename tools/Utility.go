package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fast-project-golang/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Luzifer/go-openssl"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

func GetDateFormat(format string, time time.Time) (todayString string) {
	today := time
	todayString = today.Format(format)
	return
}

func Dencrytion(key string, ciphertext string) string {
	openSSL := openssl.New()
	decryption, err := openSSL.DecryptBytes(key, []byte(ciphertext))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(decryption)
	return string(decryption)
}

func EncryptionSha256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	Encrypt := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return Encrypt
}

func SHA256(stringData []byte) string {
	secret := "ADIMARYA3014web-admin"
	data := stringData

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func GenerateToken(userId uint) (string, error) {
	tokenDuration, err := strconv.Atoi(os.Getenv("token_duration"))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["user_id"] = userId
	claim["exp"] = time.Now().Add(time.Minute * time.Duration(tokenDuration)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(os.Getenv("token_secret")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:  %v", token.Header["alg"])
		}
		return []byte(os.Getenv("token_secret")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func AES256Encrypt(key, src string) (string, error) {
	block, err := aes.NewCipher(getPaddedKey(key, 256))
	if err != nil {
		return "", err
	}
	if src == "" {
		return "", errors.New("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return hex.EncodeToString(crypted), nil
}

func AES256Decrypt(key, crypt string) (string, error) {
	block, err := aes.NewCipher(getPaddedKey(key, 256))
	if err != nil {
		return "", err
	}
	dcrypt, err := hex.DecodeString(crypt)
	if err != nil {
		return "", err
	}
	ecb := cipher.NewCBCDecrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	decrypted := make([]byte, len(dcrypt))
	ecb.CryptBlocks(decrypted, dcrypt)

	return string(PKCS5Trimming(decrypted)), nil
}
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func getPaddedKey(key string, bit int) []byte {

	paddedLen := bit/8 - len(key)
	if paddedLen < 0 {
		return []byte(key[0 : bit/8])
	}

	for i := 0; i < paddedLen; i++ {
		key = key + "f"
	}

	return []byte(key)
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ReadToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")

	return bearerToken
}

type MapKey struct {
	key string
	val []string
}

func DecryptJWT(tokenStr string, getStr string) interface{} {
	removeBearer := strings.Replace(tokenStr, "Bearer ", "", -1)
	claim := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(removeBearer, claim, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("token_secret")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return claim[getStr]
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GetToken(tokenStr string) string {
	removeBearer := strings.Replace(tokenStr, "Bearer ", "", -1)
	return removeBearer
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("token_secret")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}

func ResSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": "00", "data": data, "message": "success"})
}

func ResError(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{"code": "80", "data": data, "message": message})
}

func ResAll(c *gin.Context, data interface{}, code string, message string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "data": data, "message": message})
}

func Timer(tick time.Duration) {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	done := make(chan bool)
	sleep := 1 * time.Second
	go func() {
		time.Sleep(sleep)
		done <- true
	}()
	ticks := 0
	for {
		select {
		case <-done:
			fmt.Printf("%v × %v ticks in %v\n", ticks, tick, sleep)
			return
		case <-ticker.C:
			ticks++
		}
	}
}

func Paging(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page < 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		//switch  {
		//case pageSize > 100:
		//	pageSize = 100
		//case pageSize <= 0:
		//	pageSize = 10
		//}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CreateSignatureInquiry(payload *model.PayloadInquiry, key string) (string, error) {
	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	// Create the HMAC-SHA512 signature
	h := hmac.New(sha512.New, []byte(key))
	h.Write(payloadJSON)
	signature := fmt.Sprintf("%x", h.Sum(nil))
	// Return the signature as a hexadecimal string
	return signature, nil
}

func CreateSignatureFlagging(payload *model.PayloadFlagging, key string) (string, error) {
	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	// Create the HMAC-SHA512 signature
	h := hmac.New(sha512.New, []byte(key))
	h.Write(payloadJSON)
	signature := fmt.Sprintf("%x", h.Sum(nil))
	// Return the signature as a hexadecimal string
	return signature, nil
}

func GenerateTokenBRILINK() {
	signatureKey := ""
	transactionID := "66f2bd8d-8ab9-41c4-aeae-b3218ae33f05"
	token := "d051e189e088aae4c8b37b34864b85ec"
	amount := float64(200000)
	signatureINQ, _ := CreateSignatureInquiry(&model.PayloadInquiry{
		KodeAgen:      "test123",
		Token:         token,
		TransactionId: transactionID,
	}, signatureKey)
	fmt.Println("signature INQUIRY : " + signatureINQ)
	signatureFLG, _ := CreateSignatureFlagging(&model.PayloadFlagging{
		Amount:        amount,
		Status:        "success",
		Token:         token,
		TransactionId: transactionID,
	}, signatureKey)
	fmt.Println("signature FLAGGING : " + signatureFLG)
}
