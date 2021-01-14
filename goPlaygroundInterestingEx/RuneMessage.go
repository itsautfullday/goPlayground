package main

import (
	"golang.org/x/tour/wc"
	"unicode"
)

func WordCount(s string) map[string]int {
	//Split the string into an array of string accumulating the entirety of the words
	r := []rune(s) // What the fuck is this type conversion
	var accumulator_slice = s[0:len(s)]
	var result = make(map[string]int)
	start := 0
	end:=0
	for i := 0 ; i < len(s) ; i++{
		if(unicode.IsSpace(r[i]) || i == len(s) - 1){
			if(!unicode.IsSpace(r[i])){
				end = i + 1
			} else {
				end = i
			}
			
			accumulator_slice = s[start:end]
			_,ok:=result[accumulator_slice] 
			if(ok){
				result[accumulator_slice] += 1
			} else{
				result[accumulator_slice] = 1
			}
			start = end + 1
		}
		
	}
	return result
}

func main() {
	wc.Test(WordCount)
}

