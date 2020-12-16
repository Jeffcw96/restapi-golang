package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	//package to serve our HTTP request and server
	"github.com/gorilla/mux"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Age        int64  `json:"age"`
	Occupation string `json:"occupation"`
}

var allUsers []User

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/createUser", CreateUser).Methods("POST")
	router.HandleFunc("/getAllUsers", GetUsers).Methods("GET")
	router.HandleFunc("/getUser/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/updateUser/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", DeleteUser).Methods("DELETE")
	fmt.Println("Server is at Port 3000")
	http.ListenAndServe(":3000", router)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User                              //initialize with struct
	json.NewDecoder(r.Body).Decode(&user)      //passing JSON object to struct
	user.Id = strconv.Itoa(rand.Intn(1000000)) //Generate random ID
	allUsers = append(allUsers, user)          //save user into all users in memory
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("user Created")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(allUsers)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userFound := false
	var userDetails User

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	for _, user := range allUsers {
		if params["id"] == user.Id {
			userDetails = user
			userFound = true
			break
		}
	}

	if userFound {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userDetails)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not Found")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	for index, user := range allUsers {
		if params["id"] == user.Id {
			var user User
			allUsers = append(allUsers[:index], allUsers[index+1:]...)
			fmt.Println("latest all user", allUsers)
			json.NewDecoder(r.Body).Decode(&user)
			user.Id = params["id"]
			allUsers = append(allUsers, user)
			json.NewEncoder(w).Encode("User Profile Updated !!!")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	for index, user := range allUsers {
		if params["id"] == user.Id {
			allUsers = append(allUsers[:index], allUsers[index+1:]...)
			fmt.Println("latest all user", allUsers)
			json.NewEncoder(w).Encode("Successfully deleted !!!")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User Not Found")
}
