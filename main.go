package main

// * Example usage of the openwm package * //

import (
	"fmt"
	"main/openwm"
)

func main() {

	openWM := openwm.NewOpenWM("APP_KEY_HERE")

	wd, _ := openWM.GetWeatherByLatLon("63.83847", "23.13066")
	tempStr := fmt.Sprintf("%.2f", wd.Main.Temp-273.15)
	fmt.Println(tempStr)

	ld, _ := openWM.GetLocationsByCity("Kokkola")
	fmt.Println(ld)

	li, _ := openWM.GetLocationInfoByZip("67100,FI")
	fmt.Println(li)
}
