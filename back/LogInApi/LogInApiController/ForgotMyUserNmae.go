package LogInApiController

import (
	"back/LogInApi/LogInApiGormModels"
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
	err := c.BindJSON(&tmp)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
		return
	}
	IdFromMail := LogInApiGormModels.GetUserIdFromMail(tmp.Mail)
	IdFromPhone := LogInApiGormModels.GetUserIdFromPhone(tmp.Phone)
	IdFromID := LogInApiGormModels.GetUserIdFromID(tmp.IdentityCard)
	if IdFromMail == IdFromPhone && IdFromPhone == IdFromID && IdFromMail != -1 && IdFromPhone != -1 && IdFromID != -1 {

	}
}
