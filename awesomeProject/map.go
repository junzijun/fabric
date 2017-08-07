package main
import "fmt"
func main(){
	var countryCapitalMap map[string] string
	countryCapitalMap = make(map[string] string
	countryCapitalMap[china]="bejing"
	countryCapitalMap["France"]="Paris"
    for country := range countryCapitalMap{
		fmt.Println("capital of"country,"is",countryCapitalMap[country])


}
