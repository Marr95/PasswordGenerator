package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

type instellingen struct {
	Char           bool `json:"char"`
	Nummers        bool `json:"nummers"`
	SpecialeTekens bool `json:"specialeTekens"`
	HoofdLetters   bool `json:"hoofdLetters"`
}

// dit is een lege array. 
var myslice []string
var nieuwewachtwoord = ""
var CharactersLijst = [26]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var HoofdLetterLijst = [26]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var NummersLijst = [10]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var SpecialeTekensLijst = [24]string{
	"!", "?", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "<", ">", ",", ".", "]", "[", "{", "}", "|"}

// met deze flag laat ik zien hoe lang een wachtwoord word, met deze flag laat ik de gebruiker de lengte van een wachtwoord bepalen.
var amount int

func init() {
	flag.IntVar(&amount, "m", 10, "met deze flag laat ik zien hoe lang een wachtwoord word")
	flag.Parse()
}

func main() {
	config := readjson() 
	// hier onder voeg ik de array samen toe in een array ligt aan config file.
	if config.Char == true {
		for index := range CharactersLijst { // de index is nu de grote van de Array.
			myslice = append(myslice, CharactersLijst[index])
		}
	}
	if config.HoofdLetters == true {
		for index := range HoofdLetterLijst { // de index is nu de grote van de Array.
			myslice = append(myslice, HoofdLetterLijst[index])
		}
	}
	if config.Nummers == true {
		for index := range NummersLijst { // de index is nu de grote van de Array.
			myslice = append(myslice, NummersLijst[index])
		}
	}
	if config.SpecialeTekens == true {
		for index := range SpecialeTekensLijst { // de index is nu de grote van de Array.
			myslice = append(myslice, SpecialeTekensLijst[index])
		}
	}
	generatePassword()
	fmt.Println(nieuwewachtwoord)
}

func readjson() *instellingen {
	// open de file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("File kan niet gemaakt worden: ", err)
	}
	defer file.Close()
	// lees file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("file kan niet gelezen worden: ", err)
	}
	// object word gemaakt
	var newInstellingen instellingen
	if err := json.Unmarshal(data, &newInstellingen); err != nil {
		log.Fatal("kan niet unmarshallen: ", err)
	}
	return &newInstellingen

}
func generatePassword() {
	index := len(myslice)
	for i := 0; i <= amount; i++ {
		nieuwewachtwoord += myslice[rand.Intn(index)]
	
		func randInt(min int, max int) int {
			rand.Seed(time.Now().UTC().UnixNano())
			return min + rand.Intn(max-min)
	}
}
