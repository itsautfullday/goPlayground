package main

import "fmt"
import "net/http"
import "os"
import "io"
import "text/template"



// Compile templates on start of the application
var templates = template.Must(template.ParseFiles("public/fileUploadTemplate.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	err := templates.ExecuteTemplate(w, page+".html", data)
	if err != nil {
		fmt.Println("Error in executing template " + err.Error())
	}
}

func main(){
	http.HandleFunc("/uploadHome", displayHomePage);
	http.HandleFunc("/uploadFile",uploadFile)
	fmt.Print("Running a server on 8080");
	http.ListenAndServe(":8080",nil)
}

//Indicates a void type function 
//The intention of this fn is to ensure it is a GET request
//Print the content I want : This is a home page
func displayHomePage(res http.ResponseWriter, req * http.Request){
	fmt.Print("Calling displayHomePage");
	if req.Method == "GET" {
		display(res, "fileUploadTemplate", nil)
	} else {
		fmt.Fprintf(res, "Incorrect request type");
	}
}

func uploadFile(res http.ResponseWriter, req * http.Request){
	req.ParseMultipartForm(1e+7)
	file, handler, err := req.FormFile("data");
	if err != nil{
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close() // Defers the closing of the file to end of fn.
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	localFile, err := os.Create("localFileName");
	if err != nil{
		fmt.Println("Error Creating local file")
		fmt.Println(err)
		return
	}

	//Here  localfile is a pointer to file and file is multipart.File, which clearly does implement a reader interface but why does file pointer implement a writer interface?
	if _, err := io.Copy(localFile, file); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "Successful file upload");
}


