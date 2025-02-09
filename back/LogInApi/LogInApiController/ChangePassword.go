package LogInApiController

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"back/LogInApi/LogInApiGormModels"

	"github.com/gin-gonic/gin"
)

type changePassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
	IP          string `json:"ip"`
	MacAddress  string `json:"macAddress"`
}

func ChangePassword(c *gin.Context) {
	var changePassword changePassword
	if err := c.BindJSON(&changePassword); err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
		return
	}

	// יצירת מופע של הנתונים המבוקשים
	logInData := LogIn{
		UserName:   changePassword.Username,
		Password:   changePassword.OldPassword,
		IP:         changePassword.IP,
		MacAddress: changePassword.MacAddress,
	}

	// המרת הנתונים לתבנית JSON
	jsonData, err := json.Marshal(logInData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to marshal JSON data")
		return
	}

	// שליחת הבקשה POST עם הנתונים המרובים
	resp, err := http.Post("http://localhost:8080/logIn/LogIn", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Failed to send POST request")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %s", err)
		}
	}(resp.Body)

	// טיפול בתגובה מהשרת, אם נדרש
	// בדיקת קוד המצב
	if resp.StatusCode == http.StatusOK {
		// קריאת תוכן התגובה
		var responseData Response
		err = json.NewDecoder(resp.Body).Decode(&responseData)
		if err != nil {
			log.Fatal(err)
		}

		// כעת אתה יכול לגשת ל־UserId ולעשות איתו מה שאתה רוצה
		userId := responseData.UserId
		num, err := strconv.Atoi(userId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Failed to convert UserId to integer")
			return
		}
		LogInApiGormModels.AddPassword(changePassword.NewPassword, num)
		c.IndentedJSON(http.StatusOK, "Password changed successfully")
		return
	} else {
		// החיבור לא הצליח - טיפול במצב זה
		log.Printf("Connection failed with status code: %d", resp.StatusCode)
		c.IndentedJSON(resp.StatusCode, resp.Body)
	}
}
