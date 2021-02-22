package main

//Start by not creating a different package
//We can list that later

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	storyMap := openAndReadJSONStory("jsonStory.json")
	mapHandler := mapHandler(storyMap)
	if storyMap == nil || len(storyMap) == 0 {
		fmt.Println("Error in Reading file")
		return
	}
	//Create a htttp response handler for all of the stories sent here
	//The response handler will take an arc type
	// Will print the stories Joined
	// Currently use redirect through simple redirection
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

//To read a file:
//Use OS package to open it  --> returns a file handler
//Use a package specific reading techniques
//For JSON the pattern given below
//For CSV use ReadQuiz.go for reference :quiz/readfileModule/readQuiz.go
func openAndReadJSONStory(location string) map[string]arc {
	mapOfArcs := make(map[string]arc)
	var result map[string]interface{}

	jsonFile, er := os.Open(location)
	if er != nil {
		fmt.Println("Error in opening file")
		return mapOfArcs
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//Json parse
	json.Unmarshal([]byte(byteValue), &result)

	for key, element := range result {
		arcElement := createArcFromMap(key, element)
		mapOfArcs[key] = arcElement
	}
	return mapOfArcs
}

func createArcFromMap(key string, givenMap interface{}) arc {
	var storyArc arc
	dbByte, err := json.Marshal(givenMap)
	if err != nil {
		println("Some error ")
	}
	json.Unmarshal(dbByte, &storyArc)
	storyArc.Key = key
	return storyArc

}

func mapHandler(mapOfArcs map[string]arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		path = path[1:len(path)]
		el := mapOfArcs["intro"]
		storyToBeShown := el.Story
		titleToBeShown := el.Title
		_, ok := mapOfArcs[path]
		if ok {
			el = mapOfArcs[path]
			storyToBeShown = el.Story
			titleToBeShown = el.Title
		}
		fmt.Fprintln(w, "Hello, world!\n title: ", titleToBeShown, "\n Story :", strings.Join(storyToBeShown, "."))
		return
	}
}

//Following a json serialisation and deserialiation approch

type arc struct {
	Key     string
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text    string `json:"text"`
	NextArc string `json:"arc"`
}
