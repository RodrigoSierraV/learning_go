package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	buffReader := bufio.NewReader(os.Stdin)
	zoo := make([]Animal, 0)
	zoo = CreateAnimalList()

	fmt.Print("To request information about an animal, type a request in the following format:\n<animal_species> <request_type>\nAnimal Species is limited to Cow, Bird or Snake\nRequest type is limited to Eat, Move or Speak\n")
	fmt.Print("To Exit, type X\n\n")
	keepLooping := true
	for keepLooping {
		fmt.Print(">")
		inputRequest, err := buffReader.ReadString('\r')
		inputRequest = strings.Trim(inputRequest, "\r\n")

		if err != nil {
			fmt.Print("Something went wrong while reading request, exiting program\n")
			os.Exit(1)
		}
		keepLooping = HandleRequest(inputRequest, &zoo)
	}
}

type Animal struct {
	species    string
	food       string
	locomotion string
	sound      string
}

func newAnimal(spec string, eat string, move string, speak string) Animal {
	a := Animal{species: spec, food: eat, locomotion: move, sound: speak}
	return a
}
func (a Animal) Eat() {
	fmt.Print(a.species, " eats ", a.food, "\n")
}
func (a Animal) Move() {
	fmt.Print(a.species, " moves by ", a.locomotion, "ing\n")
}
func (a Animal) Speak() {
	fmt.Print(a.species, " goes ", a.sound, "\n")
}
func (a Animal) getSpecies() string {
	return a.species
}

func CreateAnimalList() []Animal {
	var a Animal
	list := make([]Animal, 0)

	a = newAnimal("cow", "grass", "walk", "moo")
	list = append(list, a)
	a = newAnimal("bird", "worms", "fly", "peep")
	list = append(list, a)
	a = newAnimal("snake", "mice", "slither", "hsss")
	list = append(list, a)

	return list
}

func HandleRequest(request string, zoo *[]Animal) bool {
	request = strings.ToLower(request)
	if request == "x" || request == "" {
		fmt.Print("Exiting program\n")
		return false
	}
	re := regexp.MustCompile("\\w+\\s\\w+")
	validRequest := re.MatchString(request)
	if !validRequest {
		fmt.Print("Invalid Request, try again\n")
		return true
	}

	requestTokens := strings.Split(request, " ")
	requestSpecies := requestTokens[0]
	requestType := requestTokens[1]
	re = regexp.MustCompile("eat|move|speak")
	if re.MatchString(requestType) {
		for _, animal := range *zoo {
			if strings.ToLower(animal.getSpecies()) == requestSpecies {
				switch requestType {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				}
			}
		}
	} else {
		fmt.Print("Invalid query type, please limit your input to Eat, Move or Speak")
	}

	return true
}
