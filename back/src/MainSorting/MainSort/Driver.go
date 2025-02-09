package MainSort

import (
	"time"
)

// מבנה נתונים לנהג
type Driver struct {
	Id           int
	TimeOfDriver time.Time
	Gps          string
	WageEarner   bool // האם הוא שכיר
}

// NewDriver יוצר מופע חדש של Driver
//
// קלט:
// - id: מזהה הנהג (מספר שלם)
// - gps: מיקום ה-GPS הנוכחי של הנהג (מחרוזת)
// - wageEarner: האם הנהג שכיר (בוליאני)
//
// פלט:
// - Driver: מבנה חדש של Driver עם הפרטים שסופקו והזמן הנוכחי מוגדר לשעה 6:00 בבוקר.
//
// פעולת הפונקציה:
// הפונקציה יוצרת מופע חדש של Driver עם הפרטים שסופקו ומגדירה את זמן הנהג לשעה 6:00 בבוקר ביום הנוכחי.
func NewDriver(id int, gps string, wageEarner bool) Driver {
	return Driver{Id: id, Gps: gps, WageEarner: wageEarner, TimeOfDriver: timeUntilSixAM()}
}

// GetId מחזירה את מזהה הנהג
//
// קלט:
// אין
//
// פלט:
// - int: מזהה הנהג
//
// פעולת הפונקציה:
// הפונקציה מחזירה את מזהה הנהג מתוך מבנה ה-Driver.
func (d *Driver) GetId() int {
	return d.Id
}

// timeUntilSixAM מחזירה את השעה 6:00 בבוקר ביום הנוכחי
//
// קלט:
// אין
//
// פלט:
// - time.Time: זמן המייצג את השעה 6:00 בבוקר ביום הנוכחי
//
// פעולת הפונקציה:
// הפונקציה מחזירה את הזמן המייצג את השעה 6:00 בבוקר ביום הנוכחי.
func timeUntilSixAM() time.Time {
	// קביעת השעה הנוכחית
	now := time.Now()
	// קביעת השעה 6:00 בבוקר
	now1 := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		6, 0, 0, 0,
		now.Location(),
	)

	return now1
}

// AddChangeFromDriver מעדכנת את נתוני הנהג לפי השינוי הנתון
//
// קלט:
// - tmpChange: מבנה מסוג ChangeToDriver המכיל את פרטי השינוי
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מעדכנת את מיקום ה-GPS של הנהג ואת הזמן שלו לפי הנתונים שבמבנה ChangeToDriver.
func (d *Driver) AddChangeFromDriver(tmpChange ChangeToDriver) {
	d.Gps = tmpChange.GpsEnd
	d.TimeOfDriver = tmpChange.TimeToSet
}

// MoveChangeFromDriver מחזירה את נתוני הנהג למצב קודם לפי השינוי הנתון
//
// קלט:
// - tmpChange: מבנה מסוג ChangeToDriver המכיל את פרטי השינוי
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מחזירה את מיקום ה-GPS של הנהג ואת הזמן שלו למצבם הקודם לפי הנתונים שבמבנה ChangeToDriver.
func (d *Driver) MoveChangeFromDriver(tmpChange ChangeToDriver) {
	d.Gps = tmpChange.GpsEnd
	d.TimeOfDriver = tmpChange.TimeReturnForDriver
}
