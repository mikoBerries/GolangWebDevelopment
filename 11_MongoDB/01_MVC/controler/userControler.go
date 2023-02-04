package controler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"m.test/11_MongoDB/01_MVC/models"
)

type UserControler struct{}

func NewUserControler() *UserControler {
	return &UserControler{}
}

func (uc UserControler) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "agus",
		Gender: "male",
		Age:    45,
		Id:     p.ByName("id"),
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserControler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	//decode json to var u
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = "10"
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserControler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.Header().Set("content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "some magic code to delete user")
}
