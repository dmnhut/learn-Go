package main
import "fmt"

func main(){
	var countryCapitalMap map[string] string
	/* create a map */
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys */
	for country := range countryCapitalMap{
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* test if entry is present in the map or not */
	// captial, ok := countryCapitalMap["United States"]
	/* if ok is true, entry is present otherwise emtry is absent */

	if(!(countryCapitalMap["United States"] == "")){
		fmt.Println("Capital of United States is", countryCapitalMap["United States"])
	}else{
		fmt.Println("Capital of United States is not present")
	}

}
