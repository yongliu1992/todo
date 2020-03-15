package routehandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	exec "github.com/yongliu1992/todo/dbexecutor"
	todo "github.com/yongliu1992/todo/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

// AddTodoHandler : Handler to add a todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	/**
			https://golang.org/pkg/encoding/json/#NewDecoder
	**/
	var td todo.Todo

	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&td)
	exec.AddTodo(td)
	json.NewEncoder(w).Encode(td)
	log.Println(td)
}

// UpdateTodoHandler : Handler to update a given todo
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var td todo.Todo
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	json.NewDecoder(r.Body).Decode(&td)
	err := exec.UpdateTodo(tid, td)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
	}
	json.NewEncoder(w).Encode(td)
	log.Println(td)
}

// DeleteTodoHandler : Handler to delete a given todo
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	d, err := exec.DeleteTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
	}
	json.NewEncoder(w).Encode(d)
	log.Println(d)
}

// GetTodoHandler : Handler to read a given todo
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	t, err := exec.GetTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
	}
	json.NewEncoder(w).Encode(t)
	log.Println(t)
}

// GetTodosHandler : Handler to read all todos
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	t, err := exec.GetTodos()
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err)
	}
	json.NewEncoder(w).Encode(t)
	log.Println(t)
}

func GetHandler2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Write([]byte(params["TID"]))
	json.NewEncoder(w)
}
