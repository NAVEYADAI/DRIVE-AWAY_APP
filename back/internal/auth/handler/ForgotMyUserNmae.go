package handler

import (
	"back/internal/auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type forgetMyUserName struct {
	Phone        string `json:"Phone"`
	Mail         string `json:"Mail"`
	IdentityCard int    `json:"IdentityCard"`
}

func ForgetMyUserName(c *gin.Context) {
	var tmp forgetMyUserName
	err := c.ShouldBindJSON(&tmp)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
		return
	}
	IdFromMail := model.GetUserIdFromMail(tmp.Mail)
	IdFromPhone := model.GetUserIdFromPhone(tmp.Phone)
	IdFromID := model.GetUserIdFromID(tmp.IdentityCard)
	if IdFromMail == IdFromPhone && IdFromPhone == IdFromID && IdFromMail != -1 && IdFromPhone != -1 && IdFromID != -1 {

	}
}
