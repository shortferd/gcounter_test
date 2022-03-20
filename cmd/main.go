package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "io/ioutil"
    "os/signal"
	"syscall"
    "encoding/json"
    "github.com/gorilla/mux"
)


//  Define the Counter structure

type Counter struct {
    Number int 
}

// Create Counter structure
var Count Counter

// Initilize counter
func(ctr *Counter) fill_defaults(filename string){
  
    // setting default values
    // if no values present
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        ctr.Number = 0
        }
    json.Unmarshal(file, &ctr)
}
// Dump counter if application is closed
func(ctr *Counter) dump_couter(filename string){
        jsonString, err := json.Marshal(&ctr)
        if err != nil {
        fmt.Println(err)
        return
        }
        ioutil.WriteFile(filename, jsonString, os.ModePerm)
    }


// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the

// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler(filename string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Count.dump_couter(filename)
        os.Exit(0)
	}()
}

// RESTful API 

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnNumbers(w http.ResponseWriter, r *http.Request){
    ctr := Count.Number
    Count.Number = ctr +1
    fmt.Println("Count number " + strconv.Itoa(Count.Number))
    json.NewEncoder(w).Encode("Number of gets: "+strconv.Itoa(Count.Number))
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/num", returnNumbers).Methods("GET")
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// Application entrypoint

func main() {
    Count = Counter{}
    Count.fill_defaults("json/counter.json")
    SetupCloseHandler("json/counter.json")
    handleRequests()
}