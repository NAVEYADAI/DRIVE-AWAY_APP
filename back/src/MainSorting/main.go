package MainSorting

//
//import (
//	"MainSorting/MainSort"
//	"fmt"
//)
//
//func main() {
//var tmp []MainSort.Driver
//tmp = append(tmp, MainSort.NewDriver(1, "ערבי נחל 24 בני עיש", true))
//tmp = append(tmp, MainSort.NewDriver(2, "ערבי נחל 14 בני עיש", true))
//tmp = append(tmp, MainSort.NewDriver(3, "ערבי נחל 24 בני עיש", false))
////tmp = append(tmp, MainSort.NewDriver(1, "31.787403213406456, 34.76082860860629", true))
////tmp = append(tmp, MainSort.NewDriver(2, "31.787403213406456, 34.76082860860629", true))
////tmp = append(tmp, MainSort.NewDriver(3, "31.79162991384637, 34.797584580644994", false))
////
//var tmp2 []MainSort.Transportation
//tmp2 = append(tmp2, MainSort.NewTransportation(5, 30, 7, "31.79162991384637, 34.797584580644994", "31.79344045283071, 34.82454407613016"))
//tmp2 = append(tmp2, MainSort.NewTransportation(6, 30, 6, "31.7932815086927, 34.76039908522838", "31.805385760544315, 34.776641435071454"))
//tmp2 = append(tmp2, MainSort.NewTransportation(3, 30, 15, "31.79344045283071, 34.82454407613016", "31.79162991384637, 34.797584580644994"))
//tmp2 = append(tmp2, MainSort.NewTransportation(4, 30, 15, "31.805385760544315, 34.776641435071454", "31.7932815086927, 34.76039908522838"))
//
//tn, _ := MainSort.Sort(tmp, tmp2)
//fmt.Println(tn)

// יצירת מפה לדוגמה
//	travelList := MainSort.TravelListForDrivers{
//		1: {101, 102, 103},
//		2: {201, 202},
//		3: {301, 302, 303, 304},
//	}
//
//	// קריאת הפונקציה להדפסת תוכן המפה
//	printTravelList(travelList)
//}
//
//// פונקציה להדפסת תוכן המפה
//func printTravelList(travelList MainSort.TravelListForDrivers) {
//	for driver, destinations := range travelList {
//		fmt.Printf("Driver %d:\n", driver)
//		for _, destination := range destinations {
//			fmt.Printf("  Destination: %d\n", destination)
//		}
//	}
//}
