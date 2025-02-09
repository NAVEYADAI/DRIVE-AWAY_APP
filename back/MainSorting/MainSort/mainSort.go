package MainSort

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// מפה גלובלית של נהגים שבה המפתח הוא מזהה הנהג
var driverMap map[int]Driver

// פונקציה שממיינת רשימת נהגים ונסיעות ומחזירה רשימת נסיעות לנהגים
//
// קלט:
// - DriverList: רשימת נהגים
// - TransportationList: רשימת נסיעות
//
// פלט:
// - TravelListForDrivers: רשימת נסיעות לנהגים
//
// פעולת הפונקציה:
// הפונקציה ממיינת את הנסיעות לפי זמן ההתחלה שלהן, ואז משבצת נהגים לנסיעות תוך כדי שימוש בערימות עדיפות. היא עוקבת אחר השינויים שנעשו כדי לאפשר חזרה אחורה במקרה הצורך.
func Sort(DriverList []Driver, TransportationList []Transportation) (TravelListForDrivers, string) {
	finished := true
	trueSort := 0
	TransportationLen := len(TransportationList)
	indexTransportation := 0

	MinTeapsList := make([]MinHeap, TransportationLen)
	for i := 0; i < TransportationLen; i++ {
		MinTeapsList[i] = nil
	}
	ChangesToDownload := make([]ChangeToDriver, TransportationLen)
	var driveForDriver TravelListForDrivers
	driverMap = CreateMap(DriverList)

	TransportationList = SortFromTime(TransportationList)

	for finished {
		if trueSort == TransportationLen {
			finished = false
			break
		} else {
			if indexTransportation < 0 {
				return nil, "we didnt can sort"
			}
		}
		if MinTeapsList[indexTransportation] == nil {
			tmpHeap := &MinHeap{}
			heap.Init(tmpHeap)
			CreateHeap(tmpHeap, TransportationList[indexTransportation])
			if tmpHeap.Len() > 0 {
				LoadingDriver(tmpHeap, &ChangesToDownload, &driveForDriver,
					TransportationList[indexTransportation], &indexTransportation)
				MinTeapsList[indexTransportation] = *tmpHeap
				trueSort++
				indexTransportation++
			} else {
				trueSort--
				indexTransportation--
				continue
			}
		} else {
			RemoveDriveFromDriver(ChangesToDownload[indexTransportation], driveForDriver)
			LoadingDriver(&MinTeapsList[indexTransportation], &ChangesToDownload, &driveForDriver,
				TransportationList[indexTransportation], &indexTransportation)
		}
	}

	return driveForDriver, "good work"
}

// LoadingDriver מוסיפה נסיעה לנהג מתוך ערימת העדיפות
//
// קלט:
// - tmpHeap: מצביע לערימת העדיפות של הנהגים
// - ChangesToDownload: מצביע לרשימת השינויים שבוצעו
// - driveForDriver: מצביע לרשימת הנסיעות לנהגים
// - transportation: מבנה נסיעה שמכיל את פרטי הנסיעה
// - indexTransportation: מצביע למיקום הנסיעה הנוכחית ברשימת הנסיעות
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מוציאה נהג מערימת העדיפות, יוצרת שינוי חדש לנהג ומעדכנת את רשימת הנסיעות והשינויים בהתאם.
func LoadingDriver(tmpHeap *MinHeap, ChangesToDownload *[]ChangeToDriver, driveForDriver *TravelListForDrivers,
	transportation Transportation, indexTransportation *int) {
	if tmpHeap.Len() > 0 {
		getIdDriver := heap.Pop(tmpHeap).(Item).Driver.Id
		tmpChange := NewChange(driverMap[getIdDriver], transportation)
		driveForDriver.AddDrive(transportation.Id, getIdDriver)
		driver := driverMap[getIdDriver]
		driver.AddChangeFromDriver(tmpChange)
		driverMap[getIdDriver] = driver
		(*ChangesToDownload)[*indexTransportation] = tmpChange
	} else {
		panic(fmt.Sprintf("No available drivers for transportation: %v", transportation.Id))
	}
}

// Grade מחשבת את ציון ההתאמה של נהג לנסיעה
//
// קלט:
// - driver: מבנה נהג
// - transportation: מבנה נסיעה
//
// פלט:
// - float64: ציון ההתאמה של הנהג לנסיעה
//
// פעולת הפונקציה:
// הפונקציה מחשבת את ציון ההתאמה של נהג לנסיעה בהתבסס על זמן ההגעה המשוער וסטטוס הנהג כשכיר או עצמאי.
func Grade(driver Driver, transportation Transportation) float64 {
	timeToAriveForTransportation := FindWay(driver.Gps, transportation.GpsStart)
	ExcessTimeBetweenTrips := transportation.TimeStart.Add(-timeToAriveForTransportation)
	if driver.TimeOfDriver.Before(ExcessTimeBetweenTrips) {
		grade := timeToAriveForTransportation.Minutes()
		if driver.WageEarner {
			return 0.75 * float64(grade)
		}
		return grade
	} else {
		return -1
	}
}

// CreateMap יוצר מפה של נהגים
//
// קלט:
// - driver: רשימת נהגים
//
// פלט:
// - map[int]Driver: מפה של נהגים שבה המפתח הוא מזהה הנהג
//
// פעולת הפונקציה:
// הפונקציה יוצרת ומחזירה מפה של נהגים שבה המפתח הוא מזהה הנהג והערך הוא מבנה הנהג עצמו.
func CreateMap(driver []Driver) map[int]Driver {
	dictionary := make(map[int]Driver)
	for i := 0; i < len(driver); i++ {
		dictionary[driver[i].GetId()] = driver[i]
	}
	return dictionary
}

// CreateHeap יוצר ערימת עדיפות של נהגים לפי ציון ההתאמה שלהם לנסיעה
//
// קלט:
// - tmpHeap: מצביע לערימת העדיפות של הנהגים
// - transportation: מבנה נסיעה שמכיל את פרטי הנסיעה
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מחשבת את ציון ההתאמה של כל נהג לנסיעה ומכניסה את הנהגים עם ציון חיובי לערימת העדיפות.
func CreateHeap(tmpHeap *MinHeap, transportation Transportation) {
	for _, driver := range driverMap {
		tmpGrade := Grade(driver, transportation)
		if tmpGrade != -1 {
			heap.Push(tmpHeap, Item{Grate: tmpGrade, Driver: driverMap[driver.GetId()]})
		}
	}
}

// RemoveDriveFromDriver מחזירה את נתוני הנהג למצבם הקודם לפי השינוי הנתון
//
// קלט:
// - tmpChange: מבנה ChangeToDriver המכיל את פרטי השינוי
// - listDriveInDriver: רשימת הנסיעות לנהגים
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מחזירה את מיקום ה-GPS של הנהג ואת זמן הנהג למצבם הקודם לפי הנתונים שבמבנה ChangeToDriver, ומסירה את הנסיעה האחרונה מהנהג.
func RemoveDriveFromDriver(tmpChange ChangeToDriver, listDriveInDriver TravelListForDrivers) {
	var driver Driver
	driver = driverMap[tmpChange.DriverId]
	driver.MoveChangeFromDriver(tmpChange)
	listDriveInDriver.RemoveLastTransportation(tmpChange.DriverId)
}

// Item מייצג פריט בערימת העדיפות של הנהגים
type Item struct {
	Grate  float64
	Driver Driver
}

// SortFromTime ממיינת רשימת נסיעות לפי זמן התחלת הנסיעה
//
// קלט:
// - list: רשימת נסיעות
//
// פלט:
// - []Transportation: רשימת הנסיעות הממוינת לפי זמן ההתחלה
//
// פעולת הפונקציה:
// הפונקציה ממיינת את רשימת הנסיעות לפי זמן ההתחלה שלהן באמצעות מיון מהיר.
func SortFromTime(list []Transportation) []Transportation {
	if len(list) < 2 {
		return list
	}
	left, right, center := 0, len(list)-1, rand.Int()%len(list)
	list[center], list[right] = list[right], list[center]
	for i := range list {
		if list[i].TimeStart.Before(list[right].TimeStart) {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}
	list[left], list[right] = list[right], list[left]
	SortFromTime(list[:left])
	SortFromTime(list[left+1:])
	return list
}
