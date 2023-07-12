package registration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func HandleRegistration(c *gin.Context) {
	// Konfigurasi logger
	file, err := os.OpenFile("logger/log.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Tidak dapat membuka file log:", err)
	} else {
		log.SetOutput(file)
		log.SetFormatter(&log.JSONFormatter{})
	}
	defer file.Close()

	// Memproses payload body
	payload := make(map[string]interface{}) // Membentuk seperti array
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.WithFields(log.Fields{
		"parameter": "Request",
	}).Info(payload)

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://dev.nicepay.co.id/nicepay/direct/v2/registration", strings.NewReader(string(payloadBytes)))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	// Melakukan Request
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "false",
			"message": "Terjadi kesalahan saat melakukan request",
		})
		return
	}
	defer res.Body.Close()

	// Mendapatkan Result Dari Request
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "false",
			"message": "Terjadi kesalahan saat melakukan request",
			"error":   err.Error(),
		})
		return
	}

	// Merubah Result
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Merubah Json Struct
	log.WithFields(log.Fields{
		"parameter": "Response",
	}).Info(data)

	// Respon ke webhook dengan OK
	c.JSON(http.StatusOK, gin.H{"data": data})
}
