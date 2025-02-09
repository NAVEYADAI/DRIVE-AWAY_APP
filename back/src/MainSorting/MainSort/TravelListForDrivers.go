package MainSort

// TravelListForDrivers מייצג מפה שבה המפתח הוא מזהה הנהג והערך הוא רשימת מזהי הנסיעות שהוקצו לנהג זה.
type TravelListForDrivers map[int][]int

// AddDrive מוסיף נסיעה לרשימת הנסיעות של נהג
//
// קלט:
// - idTransportation: מזהה הנסיעה
// - idDriver: מזהה הנהג
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מוסיפה את מזהה הנסיעה לרשימת הנסיעות של הנהג במפה. אם המפה עדיין לא מאותחלת, היא מאתחלת אותה.
func (t *TravelListForDrivers) AddDrive(idTransportation int, idDriver int) {
	if *t == nil {
		*t = make(map[int][]int)
	}
	(*t)[idDriver] = append((*t)[idDriver], idTransportation)
}

// RemoveLastTransportation מסירה את הנסיעה האחרונה שהוקצתה לנהג מרשימת הנסיעות שלו
//
// קלט:
// - idDriver: מזהה הנהג
//
// פלט:
// אין
//
// פעולת הפונקציה:
// הפונקציה מסירה את הנסיעה האחרונה שנוספה לרשימת הנסיעות של הנהג במפה, אם קיימת רשימה כזו והיא אינה ריקה.
func (t *TravelListForDrivers) RemoveLastTransportation(idDriver int) {
	if *t != nil && len((*t)[idDriver]) > 0 {
		(*t)[idDriver] = (*t)[idDriver][:len((*t)[idDriver])-1]
	}
}

// LenOfTransportation מחזירה את מספר הנסיעות שהוקצו לנהג
//
// קלט:
// - idDriver: מזהה הנהג
//
// פלט:
// - int: מספר הנסיעות שהוקצו לנהג
//
// פעולת הפונקציה:
// הפונקציה מחזירה את מספר הנסיעות שהוקצו לנהג במפה. אם המפה לא מאותחלת או אין נסיעות לנהג זה, הפונקציה מחזירה 0.
func (t *TravelListForDrivers) LenOfTransportation(idDriver int) int {
	if *t != nil {
		return len((*t)[idDriver])
	}
	return 0
}
