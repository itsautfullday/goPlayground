package main
import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"io"
	"flag"
	"time"
	"strings"
)

type Question struct {
	question string
	answer string
}

func main(){
	//Read data from file name given from flag
	var fileName string
	var timerSeconds int
	flag.StringVar(&fileName,"fileName", "../../trial.csv", "filename of csv for quiz")
	flag.IntVar(&timerSeconds,"time",30,"timer for quiz")
	flag.Parse()
	
	data,err := readCSV(fileName)
	if(err != nil){
		log.Fatalln("Error in opening file ",err)
		return
	}

	questionsList,err :=generateQuestionsList(data)
	if(err != nil){
		log.Fatalln("Error in generateQuestionsList ",err)
		return
	}
	correct := 0
	incorrect :=0

	//Start timer 
	timer := time.NewTimer(time.Duration(timerSeconds) * time.Second)
	fmt.Println("Starting timer for 10!");
	
	done := make(chan string)
	for i := 0 ; i < len(questionsList) ; i++{
		go getParsedInput(done)
		res, err := askQuestion(timer.C, done,questionsList[i].question,questionsList[i].answer)
		if(err != nil){
			log.Fatalln("Error in askQuestion ",err)
			return
		}
		if(res == -1){
			break
		}
		if(res == 1){
			correct += 1
		}
		if(res == 0){
			incorrect +=1
		}	
	}
	
	
	fmt.Println("\nCorrect = ",correct," incorrect = ",incorrect);
}

func readCSV(fileName string) (io.Reader, error){
	return os.Open(fileName)
}

func generateQuestionsList (data io.Reader) ([]Question, error){
	//What is type of variable being returned here?
	allQuestions,err := csv.NewReader(data).ReadAll()
	if(err != nil){
		log.Fatalln("Error in reading questions from file ",err)
		return nil, fmt.Errorf("No Question in file")
	}
	numOfQues := len(allQuestions)
	if(numOfQues == 0){
		log.Fatalln("Error in reading numOfQues == 0 ",err)
		return nil, fmt.Errorf("No Question in file")
	}
	var questionsData []Question
	for _,line := range allQuestions {
		ques := Question{}
		ques.question = line[0]
		ques.answer = line[1]
		questionsData = append(questionsData,ques)
	}
	return questionsData, nil
}


//Funciton gets a channel and provides the correct input
func getParsedInput(done chan<- string){
	var response string;
	fmt.Scanln(&response)
	done <- response
}

//Function gets 2 channels : 1 for timer process and 1 for input read process!
//Returns -1 : if timer over; 0 if incorrect and 1 if correct!
func askQuestion( timer <-chan time.Time, done <-chan string, question string, answer string ) (int, error) {
	fmt.Printf("%s ?", question)
	select {
	case <-timer:
		return -1, nil
	case ans := <-done:
		if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), answer) == 0 {
			return 1, nil
		} else {
			return 0, nil
		}
	}	
}
