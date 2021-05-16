package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
    // Write your code here
     result := ""
     if s != ""{
         for i:= 0 ; i < len(s) ; i++{
             //For each letter need to determine what the start needs to be 
             charInQuestion := s[i]
             if (charInQuestion >= 'a' && charInQuestion <= 'z') || (charInQuestion >= 'A' && charInQuestion <= 'Z'){
            k = k %26
            startingLetter := 'a'
             endingLetter := 'z'
             if charInQuestion >= 'A' && charInQuestion <= 'Z'{
                 startingLetter = 'A'
                 endingLetter = 'Z'
             }
             
             if int32(charInQuestion) + k > endingLetter{
                 diff := int32(charInQuestion) + k - endingLetter
                 result += string(startingLetter + diff - 1)
             } else {
                 result += string(int32(charInQuestion) + k)
             }   
             } else {
                 result += string(int32(charInQuestion))
             }  
         }
         return result
     }
     return result

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    _, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
