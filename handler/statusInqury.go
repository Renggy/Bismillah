package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandlerStatusInquiry(c *gin.Context) {
	// Konfigurasi logger
	file, err := os.OpenFile("logger/log.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Tidak dapat membuka file log:", err)
	} else {
		log.SetOutput(file)
		log.SetFormatter(&log.JSONFormatter{})
	}
	defer file.Close()

	payload := make(map[string]interface{})
	if err := c.BindJSON(&payload); err != nil {
		log.WithFields(log.Fields{
			"parameter": "Request",
			"message":   "Payload tidak merupakan json",
			"endpoint":  c.Request.URL,
		}).Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payloadByte, err := json.Marshal(&payload)
	response, err := http.Post("https://dev.nicepay.co.id/nicepay/direct/v2/inquiry", "application/json", bytes.NewReader(payloadByte))
	if err != nil {
		log.WithFields(log.Fields{
			"parameter": "Response",
			"message":   "API tidak memberikan response",
			"endpoint":  c.Request.URL,
		}).Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	// Membaca response payload body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"parameter": "Response",
			"message":   "Terjadi kesalhan saat membaca data",
			"endpoint":  response.Request.URL,
		}).Info(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Merubah Result
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.WithFields(log.Fields{
			"parameter": "Response",
			"message":   "Gagal melakukan unmarshal json",
			"endpoint":  response.Request.URL,
		}).Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menyimpan log respon
	log.WithFields(log.Fields{
		"parameter": "Response",
		"message":   "Berhasil menyimpan log data response",
		"endpoint":  response.Request.URL,
	}).Info(data)

	// Respon ke webhook dengan OK
	c.JSON(http.StatusOK, gin.H{"data": data})
}
