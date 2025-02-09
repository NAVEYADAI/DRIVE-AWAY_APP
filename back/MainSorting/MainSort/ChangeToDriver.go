package MainSort

import "time"

// מבנה לצורך שמירה של שינויים שאנו מבצעים בעת הוספת נסיעה לנהג על מנת שנוכל לחזור אחורה
type ChangeToDriver struct {
	DriverId            int
	TransportationId    int
	TimeToSet           time.Time
	TimeReturnForDriver time.Time
	GpsStart            string
	GpsEnd              string
}

// NewChange יוצר מופע חדש של ChangeToDriver
//
// קלט:
// - driver: מבנה מסוג Driver המכיל את פרטי הנהג
// - transportation: מבנה מסוג Transportation המכיל את פרטי הנסיעה
//
// פלט:
// - ChangeToDriver: מבנה חדש המכיל את פרטי השינוי שבוצע
//
// פעולת הפונקציה:
// הפונקציה מחשבת את הזמן שבו הנהג יסיים את הנסיעה (כולל תוספת של 5 דקות), ויוצרת מבנה חדש של ChangeToDriver המכיל את כל הפרטים הנדרשים.
func NewChange(driver Driver, transportation Transportation) ChangeToDriver {
	// חישוב הזמן שצריך להוריד לנהג - החישוב הוא הזמן שלוקחת הנסיעה והזמן שלוקח לנהג בין המיקום האחרון שהיה לתחילת הנסיעה
	tmpTime := transportation.TimeStart.Add(transportation.TimeForAllTransportation).Add(time.Duration(5) * time.Minute)
	return ChangeToDriver{
		DriverId:            driver.Id,
		TransportationId:    transportation.Id,
		TimeToSet:           tmpTime,
		TimeReturnForDriver: driver.TimeOfDriver,
		GpsStart:            driver.Gps,
		GpsEnd:              transportation.GpsEnd,
	}
}

// GetIdDriverAndTransportation מחזירה את מזהה הנהג ומזהה הנסיעה
//
// קלט:
// אין
//
// פלט:
// - int: מזהה הנהג
// - int: מזהה הנסיעה
//
// פעולת הפונקציה:
// הפונקציה מחזירה את מזהי הנהג והנסיעה מתוך מבנה ChangeToDriver.
func (t *ChangeToDriver) GetIdDriverAndTransportation() (int, int) {
	return t.DriverId, t.TransportationId
}

// GetTimeTimeInTimeDuration מחזירה זמן בפורמט time.Duration מתוך מבנה time.Time
//
// קלט:
// - tmp: מבנה time.Time שממנו יש לחשב את הזמן
//
// פלט:
// - time.Duration: זמן מחושב הכולל את השעות, הדקות והשניות
//
// פעולת הפונקציה:
// הפונקציה מחשבת ומחזירה את הזמן בפורמט של time.Duration מתוך מבנה time.Time. היא מפרקת את הזמן לשעות, דקות ושניות ומחברת אותם לתוצאה הסופית.
func GetTimeTimeInTimeDuration(tmp time.Time) time.Duration {
	second := time.Duration(tmp.Second()) * time.Second
	minute := time.Duration(tmp.Minute()) * time.Minute
	hour := time.Duration(tmp.Hour()) * time.Hour
	return hour + minute + second
}
