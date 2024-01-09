// var idMap map[key_type]value_type

package main

import "fmt"

func main() {
//		var m map[int]string
//		m = make(map[int]string)
//	
//		m[100] = "Apple"
//		m[134] = "Grape"
//		m[777] = "Tomato"
//	
//		str := m[134]
//		println(str)
//	
//		noData := m[199]
//		println(noData)

//	tickers := map[string]string{
//		"GOOG": "Google Inc",
//		"MSFT": "Microsoft",
//	}
//
//	_, exists := tickers["MSFT"]
//	if !exists {
//		println("No MSFT ticker")
//	}

	myMap := map[string]string {
		"A": "Apple",
		"B": "Banana",
		"C": "Coffee",
	}

	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
