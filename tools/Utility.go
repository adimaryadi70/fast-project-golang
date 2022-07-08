package tools

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetToday(format string) (todayString string) {
	today := time.Now()
	todayString = today.Format(format)
	return
}

func GetDateFormat(format string,time time.Time) (todayString string) {
	today := time
	todayString = today.Format(format)
	return
}



func EncryptionSha256(data []byte) string {
	hash   	 :=  	sha256.New()
	hash.Write(data)
	Encrypt  := 	base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return Encrypt
}

func GenerateToken(userId uint) (string, error) {
	tokenDuration, err := strconv.Atoi(os.Getenv("token_duration"))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["authorized"]   =   true
	claim["user_id"]  	  =   userId
	claim["exp"] 		  =   time.Now().Add(time.Minute * time.Duration(tokenDuration)).Unix()
	token  :=  jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
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
		return  err
	}
	return nil
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

func DecryptJWT(tokenStr string,getStr string) interface{} {
	removeBearer := strings.Replace(tokenStr,"Bearer ","",-1)
	claim   :=  jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(removeBearer,claim, func(token *jwt.Token) (i interface{}, err error) {
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
	removeBearer := strings.Replace(tokenStr,"Bearer ","",-1)
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


func ResSuccess(c *gin.Context,data ...interface{})  {
	c.JSON(http.StatusOK, gin.H{"code": "00","data": data,"message": "success"})
}

func ResError(c *gin.Context,data interface{},message string)  {
	c.JSON(http.StatusOK, gin.H{"code": "80","data": data,"message": message})
}

func ResAll(c *gin.Context,data interface{},code string,message string)  {
	c.JSON(http.StatusOK, gin.H{"code": code,"data": data,"message": message})
}

func Timer(tick time.Duration) {
	ticker   := time.NewTicker(tick)
	defer ticker.Stop()
	done     := make(chan bool)
	sleep  	 := 1 * time.Second
	go func() {
		time.Sleep(sleep)
		done <- true
	}()
	ticks := 0
	for  {
		select {
		case <- done:
			fmt.Printf("%v × %v ticks in %v\n", ticks, tick, sleep)
			return
		case <-ticker.C:
			ticks++
		}
	}
}