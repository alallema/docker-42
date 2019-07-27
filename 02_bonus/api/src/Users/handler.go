package Users


import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n- GET /pulse — heartbeat check if our API is online\n- GET /users — fetch all users from the database\n- GET /users/[email] — fetch user by email from the database\n")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var userList *[]User
	var errs []string
	
	enableCors(&w)

	userList = AllUsers()
	total := len(*userList)
	if len(*userList) != 0 {
		fmt.Println("Get User Request")
		json.NewEncoder(w).Encode(AllUserResponse{Total: total, User: *userList})
	} else {
		errs = append(errs,"No User Found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func GetByEmail(w http.ResponseWriter, r *http.Request) {
	var user *User
	var errs []string
	
	enableCors(&w)

	vars := mux.Vars(r)
	id := vars["email"]

	user = FindUserByEmail("email:" + id)

	if user != nil {
		fmt.Println("Get User Request")
		json.NewEncoder(w).Encode(GetUserResponse{User: *user})
	} else {
		errs = append(errs, "No User found")
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	fmt.Println("del")

	DelUserByEmail(email)
}