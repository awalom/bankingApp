package helpers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, content interface{}) {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		panic(err)
	}

}

func GetVars(r *http.Request, str string) (string, error) {
	variables := mux.Vars(r)
	if variables == nil {
		return "", errors.New("variable was not found in the path")
	}
	return variables[str], nil
}
