package main
import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"io"
	"flag"
	"time"
)

func main(){
	var fileName string;
	flag.StringVar(&fileName,"fileName", "../../trial.csv", "filename of csv for quiz")
	data,err := os.Open(fileName)
	if(err != nil){
		log.Fatalln("Error in opening file ",err)
	}
	r := csv.NewReader(data)
	correct := 0
	incorrect :=0
	for{
		record,err := r.Read()
		if(err == io.EOF){
			break
		}
		if(err != nil){
			log.Fatalln("Error in reading file ",err)
		}
		fmt.Println(record[0] + "?")
		var response string;
		fmt.Scanln(&response)
		if(response == record[0]){
			correct +=1
		} else {
			incorrect +=1
		}
	}
	fmt.Println("Correct = ",correct," incorrect =",incorrect);
	
}
