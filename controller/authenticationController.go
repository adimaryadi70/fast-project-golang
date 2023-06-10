package controller

import (
	"errors"
	"fast-project-golang/model"
	"fast-project-golang/tools"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var authfind []model.Authentification
	db.Find(&authfind)
}

func MiddlewareAuth(c *gin.Context) {
	var input model.Authentification

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "99", "message": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	hashPassword := tools.EncryptionSha256([]byte(input.Password))
	if err := db.Where("username = ? AND password = ? ", input.Username, hashPassword).First(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "78", "message": "username password salah"})
		return
	}
	token, err := tools.GenerateToken(input.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "08", "message": err.Error()})
		return
	}
	var session model.SessionToken
	if err := db.Where("user_id = ?", input.ID).First(&session).Error; err != nil {
		startTime := time.Now()
		tokenDuration, err := strconv.Atoi(os.Getenv("token_duration"))
		if err != nil {
			log.Fatal(err)
		}
		endTime := startTime.Add(time.Minute * time.Duration(tokenDuration))
		createSession := model.SessionToken{UserId: input.ID, Username: input.Username, StartSession: startTime, EndSession: endTime}
		db.Create(&createSession)
	}
	startTime := time.Now()
	tokenDuration, err := strconv.Atoi(os.Getenv("token_duration"))
	endTime := startTime.Add(time.Minute * time.Duration(tokenDuration))
	var updateSession model.SessionToken
	updateSession.StartSession = startTime
	updateSession.EndSession = endTime
	db.Model(&session).Updates(updateSession)
	userInfo := &model.UsersInfo{Token: token, Data: input}
	tools.ResSuccess(c, userInfo)
	//c.JSON(http.StatusOK, gin.H{"code": "00","message": token})
}

func QueryCheckSession(c *gin.Context) bool {
	token, _ := tools.ExtractTokenID(c)
	var session model.SessionToken
	session.UserId = token
	db := c.MustGet("db").(*gorm.DB)
	result := db.Find(&session)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func RegisterAuth(c *gin.Context) {
	var input model.Authentification

	if err := c.ShouldBindJSON(&input); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"code": "99", "message": err.Error()})
		tools.ResError(c, input, "Gagal Input user")
		return
	}
	hashPassword := tools.EncryptionSha256([]byte(input.Password))
	save := model.Authentification{Username: input.Username, Password: hashPassword}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&save)
	tools.ResSuccess(c, save)
	//c.JSON(http.StatusOK, gin.H{"code": "00","message": "Create users "+input.Username})
}
