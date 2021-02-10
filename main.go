package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Name`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/task", taskHandler).Methods("GET")
	r.HandleFunc("/task", createTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTask).Methods("GET")
	r.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")

}

// CompileDaemon --command="./ApiRest.exe"

func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Print("Insert a Valid Task!")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Print("ID invalid")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-type", "application/json")

			json.NewEncoder(w).Encode(task)
		}
	}

}
