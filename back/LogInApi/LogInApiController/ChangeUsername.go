package LogInApiController

import (
	"back/LogInApi/LogInApiGormModels"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChangeUserName struct {
	OldUsername string `json:"oldUsername"`
	Password    string `json:"password"`
	IP          string `json:"ip"`
	MacAddress  string `json:"macAddress"`
	NewUsername string `json:"newUsername"`
}

func ChangeUsername(c *gin.Context) {
	var changeUserName ChangeUserName
	err := c.BindJSON(&changeUserName)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "error in BindJSON"})
		return
	}

	// יצירת מופע של הנתונים המבוקשים
	logInData := LogIn{
		UserName:   changeUserName.OldUsername,
		Password:   changeUserName.Password,
		IP:         changeUserName.IP,
		MacAddress: changeUserName.MacAddress,
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
	// קריאת תוכן התגובה וטיפול בו
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
		listUser := LogInApiGormModels.GetUserListById(num)
		for i := 0; i < len(listUser); i++ {
			if listUser[i].Name == changeUserName.NewUsername {
				c.IndentedJSON(http.StatusConflict, "Username already exists")
				return
			}
		}
		LogInApiGormModels.AddName(changeUserName.OldUsername, num)
		LogInApiGormModels.UpdateUser(changeUserName.NewUsername, num)
		c.IndentedJSON(http.StatusOK, "Username changed successfully")
		return
	} else {
		// החיבור לא הצליח - טיפול במצב זה
		log.Printf("Connection failed with status code: %d", resp.StatusCode)
		c.IndentedJSON(resp.StatusCode, resp.Body)
	}
}
