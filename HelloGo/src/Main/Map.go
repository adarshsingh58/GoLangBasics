package main

import "fmt"

func main() {

	fmt.Println(getAge("Adarsh"))
	fmt.Println(getAge_ShortHandMap("Adarsh"))
	deleteFromMap()
	checkExistenceOfKeyInMap()

}

/*Basic way of declaring map and adding values
we declare map as:
	var ageMap map[string]int;
Here, [string] is the key type and int is the value type
	=make(map[string]int)
This is used to create a map, like doing new HashMap<>() in java. W/O this map will be null
make is like new keyword
*/
func getAge(name string) (int) {
	var ageMap map[string]int = make(map[string]int)
	ageMap["Adarsh"] = 25
	ageMap["John"] = 32
	ageMap["William"] = 21
	return ageMap[name]
}
func getAge_ShortHandMap(name string) (int) {
	ageMap := map[string]int{
		"Adarsh":  22,
		"John":    32,
		"William": 21,// final comma mandatory
	}

	return ageMap[name]
}
func deleteFromMap(){
	ageMap := map[string]int{
		"Adarsh":  22,
		"John":    32,
		"William": 21,// final comma mandatory
	}
	fmt.Println(ageMap)
	delete(ageMap,"Adarsh")
	fmt.Println(ageMap)
}
func checkExistenceOfKeyInMap(){
	ageMap := map[string]int{
		"Adarsh":  22,
		"John":    32,
		"William": 21,// final comma mandatory
	}
//value,keyExist:=ageMap["Adarsh"] inside an if clause will return me a boolean value, we save
//that in key keyExist. value var will contain the actual value of the key in map
	if value,keyExist:=ageMap["Adarsh"]; keyExist{
		fmt.Println(keyExist)
		fmt.Println(value)
	}
}