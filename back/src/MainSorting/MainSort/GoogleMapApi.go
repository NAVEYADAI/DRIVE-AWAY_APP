// package MainSort
//
// import (
//
//	"io"
//	"time"
//
// )
//
// import (
//
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"strconv"
//	"strings"
//
// )
//
//	func myApiKey() string {
//			apiKey := os.Getenv("MY_GOOGLE_API_KEY")
	// if apiKey == "" {
	// 	fmt.Println("Warning: GOOGLE_MAPS_API_KEY is not set")
	// }
	// return apiKey
//	}
//
//	func FindWay(origin, destination string) time.Duration {
//		googleMapsAPIKey := myApiKey()
//		// יצירת URL לקריאה ל-Google Maps Distance Matrix API
//		apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%s&destinations=%s&key=%s",
//			url.QueryEscape(origin), url.QueryEscape(destination), googleMapsAPIKey)
//
//		// ביצוע קריאה ל-API
//		response, err := http.Get(apiURL)
//		if err != nil {
//			fmt.Println("שגיאה בביצוע קריאה ל-Google Maps Distance Matrix API:", err)
//			return 0
//		}
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//				fmt.Println("שגיאה בסגירת גוף התשובה:", err)
//			}
//		}(response.Body)
//
//		// קריאת תוכן התשובה
//		body, err := ioutil.ReadAll(response.Body)
//		if err != nil {
//			fmt.Println("שגיאה בקריאת תוכן התשובה:", err)
//			return 0
//		}
//
//		// פענוח והצגת התוצאה
//		var distanceMatrixResponse map[string]interface{}
//		err = json.Unmarshal(body, &distanceMatrixResponse)
//		if err != nil {
//			fmt.Println("שגיאה בפענוח תשובת JSON:", err)
//			return 0
//		}
//
//		// הצגת זמן הנסיעה
//		duration := distanceMatrixResponse["rows"].([]interface{})[0].(map[string]interface{})["elements"].([]interface{})[0].(map[string]interface{})["duration"].(map[string]interface{})["text"].(string)
//		fmt.Println(duration)
//		tmp := GetInTime(duration)
//		fmt.Println(tmp)
//		return tmp
//	}
//
//	func GetInTime(durationText string) time.Duration {
//		// פרק את הטקסט ליחידות וכמות
//		parts := strings.Fields(durationText)
//
//		var days, hours, minutes int
//
//		for i := 0; i < len(parts); i += 2 {
//			value, err := strconv.Atoi(parts[i])
//			if err != nil {
//				fmt.Println("שגיאה בהמרת ערך:", err)
//				return time.Duration(0)
//			}
//			unit := parts[i+1]
//			switch unit {
//			case "day", "days":
//				days = value
//			case "hour", "hours":
//				hours = value
//			case "min", "mins", "minute", "minutes":
//				minutes = value
//			default:
//				fmt.Println("יחידות לא מזוהות:", unit)
//				return time.Duration(0)
//			}
//		}
//
//		// המרת היחידות לתקופת זמן
//		duration := time.Hour*24*time.Duration(days) + time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes)
//
//		// הצגת התוצאה
//		fmt.Printf("הזמן שהוזן: %v\n", duration)
//		return duration
//	}
package MainSort

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func myApiKey() string {
	apiKey := os.Getenv("MY_GOOGLE_API_KEY")
	if apiKey == "" {
		fmt.Println("Warning: GOOGLE_MAPS_API_KEY is not set")
	}
	return apiKey
}

func FindWay(origin, destination string) time.Duration {
	googleMapsAPIKey := myApiKey()
	apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%s&destinations=%s&key=%s",
		url.QueryEscape(origin), url.QueryEscape(destination), googleMapsAPIKey)

	response, err := http.Get(apiURL)
	if err != nil {
		return 0
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0
	}

	var distanceMatrixResponse map[string]interface{}
	if err := json.Unmarshal(body, &distanceMatrixResponse); err != nil {
		return 0
	}

	// Check if "rows" field exists and is not empty
	rows, ok := distanceMatrixResponse["rows"].([]interface{})
	if !ok || len(rows) == 0 {
		return 0
	}

	// Check if "duration" field exists
	elements, ok := rows[0].(map[string]interface{})["elements"].([]interface{})
	if !ok || len(elements) == 0 {
		return 0
	}

	// Check if "duration" field exists
	durationMap, ok := elements[0].(map[string]interface{})["duration"].(map[string]interface{})
	if !ok {
		return 0
	}

	durationText, ok := durationMap["text"].(string)
	if !ok {
		return 0
	}

	return GetInTime(durationText)
}

func GetInTime(durationText string) time.Duration {
	parts := strings.Fields(durationText)

	var days, hours, minutes int

	for i := 0; i < len(parts); i += 2 {
		value, err := strconv.Atoi(parts[i])
		if err != nil {
			fmt.Println("error converting value:", err)
			return time.Duration(0)
		}
		unit := parts[i+1]
		switch unit {
		case "day", "days":
			days = value
		case "hour", "hours":
			hours = value
		case "min", "mins", "minute", "minutes":
			minutes = value
		default:
			fmt.Println("unrecognized units:", unit)
			return time.Duration(0)
		}
	}

	return time.Hour*24*time.Duration(days) + time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes)
}
