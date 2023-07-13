package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePayment(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// imid, Imid := payload["imid"].(string)
	// merchantKey, MerchantKey := payload["merchantKey"].(string)
	// txId, TxId := payload["txId"].(string)
	// refNo, RefNo := payload["refNo"].(string)
	// amount, Amount := payload["amount"].(float64)
	// timestampTrx, TimestampTrx := payload["timestampTrx"].(string)
	// link := fmt.Sprintf("https://dev.nicepay.co.id/nicepay/direct/v2/payment?timeStamp=%s&tXid=%s&merchantToken=%s&amt=%d&callBackUrl=https://pluto.nicepay.co.id/result",
	// 	timestampTrx, txId, merchantKey, amount,
	// )

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success",
		"data":    payload,
	})
}
