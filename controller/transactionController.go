package controller

import (
	"fast-project-golang/model"
	"fast-project-golang/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vigneshuvi/GoDateFormat"
	"strconv"
	"time"
)

func FindTransaction(c *gin.Context) {
	db  := c.MustGet("db").(*gorm.DB)
	var FindTransaction []model.Transaction
	urlQuery := c.Request.URL.Query()
	totalRows := db.Find(&FindTransaction).RowsAffected
	pageSize := urlQuery.Get("page_size")
	fmt.Println(pageSize)
	paging := db.Scopes(tools.Paging(c.Request)).Find(&FindTransaction)
	result := model.PagingModel{
		Page: urlQuery.Get("page"),
		PageSize: urlQuery.Get("page_size"),
		TotalRows: strconv.Itoa(int(totalRows)),
		Data: paging.Value,
	}
	//c.JSON(http.StatusOK, gin.H{})
	tools.ResSuccess(c,result)
}

func CreateTransaction(c *gin.Context) {
	var input model.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tools.ResError(c, input,"Gagal Transaksi")
	}
	fmt.Println(input.Po_number)
	date      :=   time.Now()
	PO_number :=   tools.GetToday(GoDateFormat.ConvertFormat("yyyymm"))
	data  :=   model.Transaction{Po_number: PO_number,Po_date: date, Po_price_total: input.Po_price_total, Po_cost_total: input.Po_cost_total}
	db    :=   c.MustGet("db").(*gorm.DB)
	db.Create(&data)

	tools.ResSuccess(c,data)
	//c.JSON(http.StatusOK,gin.H{"code": "00","data": data})
}

func UpdateTransaction(c *gin.Context) {
	db   := c.MustGet("db").(*gorm.DB)
	var  findModel model.Transaction
	if err := db.Where("id = ?",c.Param("id")).First(&findModel).Error; err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "row tidak di temukan"})
		tools.ResError(c,"","Data tidak ditemukan")
		return
	}

	var input model.Transaction
	if err :=  c.ShouldBindJSON(&input); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tools.ResError(c,"",err.Error())
		return
	}

	var updateInput  model.Transaction
	updateInput.ID 	 	 	   =  findModel.ID
	updateInput.Po_number      =  input.Po_number
	updateInput.Po_cost_total  =  input.Po_cost_total
	updateInput.Po_price_total =  input.Po_price_total
	db.Model(&findModel).Updates(updateInput)
	//c.JSON(http.StatusOK,gin.H{"code": "00","data": findModel})
	tools.ResSuccess(c,updateInput)

}

func DeleteTransaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var findTransaction model.Transaction
	if err := db.Where("id = ?",c.Param("id")).First(&findTransaction).Error; err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Row tidak ditemukan"})
		tools.ResError(c,"","Data Tidak ditemukan")
		return
	}

	db.Delete(&findTransaction)
	//c.JSON(http.StatusOK,gin.H{"code": "00","data": "deleted"})
	tools.ResSuccess(c,findTransaction)
}