package MainSort

import "time"

// Transportation מייצג נסיעה עם פרטי זיהוי, זמני התחלה וסיום, ומיקומי GPS
type Transportation struct {
	Id                       int           // מזהה הנסיעה
	TimeStart                time.Time     // זמן התחלת הנסיעה
	GpsStart                 string        // מיקום ה-GPS של תחילת הנסיעה
	GpsEnd                   string        // מיקום ה-GPS של סוף הנסיעה
	TimeForAllTransportation time.Duration // זמן כולל לנסיעה
}

// TimeToDisperseThem מייצג את זמן הפיזור של הנסיעה, מוגדר כאן כ-10 דקות
var TimeToDisperseThem = GetTimeFromHourAndMin(0, 10)

// GetTimeFromHourAndMin מחזירה משך זמן המבוסס על שעות ודקות
//
// קלט:
// - hour: מספר שעות
// - minute: מספר דקות
//
// פלט:
// - time.Duration: משך הזמן המחושב
//
// פעולת הפונקציה:
// הפונקציה יוצרת משך זמן מבוסס על השעות והדקות שניתנו ומחזירה אותו.
func GetTimeFromHourAndMin(hour, minute int) time.Duration {
	tmp := time.Minute*time.Duration(minute) + time.Hour*time.Duration(hour)
	return tmp
}

// createTime יוצר אובייקט זמן לפי שעה ודקה שניתנים כקלט
//
// קלט:
// - hour: מספר שעות
// - minute: מספר דקות
//
// פלט:
// - time.Time: אובייקט זמן לפי השעה והדקה הנתונים
//
// פעולת הפונקציה:
// הפונקציה יוצרת ומחזירה אובייקט זמן עם השעה והדקה הנתונים תוך שימוש בתאריך הנוכחי.
func createTime(hour, minute int) time.Time {
	// קבלת התאריך הנוכחי לשימוש בשנה, חודש ויום הנוכחיים
	now := time.Now()
	// יצירת אובייקט זמן חדש עם השעה והדקה הנתונים והתאריך הנוכחי
	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
}

// NewTransportation יוצר ומחזיר מבנה נסיעה חדש
//
// קלט:
// - tmpId: מזהה הנסיעה
// - min: מספר הדקות של תחילת הנסיעה
// - hour: מספר השעות של תחילת הנסיעה
// - GpsStart: מיקום ה-GPS של תחילת הנסיעה
// - GpsEnd: מיקום ה-GPS של סוף הנסיעה
//
// פלט:
// - Transportation: מבנה נסיעה חדש עם הפרטים הנתונים
//
// פעולת הפונקציה:
// הפונקציה יוצרת ומחזירה מבנה נסיעה חדש עם מזהה, זמן התחלה, מיקומי GPS, וזמן כולל לנסיעה המבוסס על החישוב של זמן הנסיעה.
func NewTransportation(tmpId int, min int, hour int, GpsStart string, GpsEnd string) Transportation {
	return Transportation{
		Id:                       tmpId,
		TimeStart:                createTime(hour, min),
		GpsStart:                 GpsStart,
		GpsEnd:                   GpsEnd,
		TimeForAllTransportation: FindWay(GpsStart, GpsEnd),
	}
}

// GetFinishTransportation מחזירה את זמן סיום הנסיעה
//
// קלט:
// אין (הפונקציה פועלת על אובייקט Transportation)
//
// פלט:
// - time.Time: זמן סיום הנסיעה
//
// פעולת הפונקציה:
// הפונקציה מחשבת ומחזירה את זמן סיום הנסיעה על בסיס זמן התחלה, זמן הפיזור וזמן הנסיעה הכולל.
func (t *Transportation) GetFinishTransportation() time.Time {
	return t.TimeStart.Add(TimeToDisperseThem).Add(t.TimeForAllTransportation)
}
