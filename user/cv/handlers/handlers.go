package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user/cv/models"
	"github.com/user/cv/repository"
)


// GET /
func Index(w http.ResponseWriter, r *http.Request) {
    candidates := repository.GetCandidates() // list of all candidates
    if err := json.NewEncoder(w).Encode(candidates); err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    return
}

// GET /candidate/{id}
func GetCandidate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

    vars := mux.Vars(r)
    id := vars["id"] // param id

   candidate := repository.GetCandidate(id) // get the candidate by id
    if candidate.Email != "" {
    		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    		w.WriteHeader(http.StatusOK)
    		if err := json.NewEncoder(w).Encode(candidate); err != nil {
    			panic(err)
    		}
  		  return
  	}

  	// If we didn't find it, 404
  	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  	w.WriteHeader(http.StatusNotFound)
}

// POST /candidate
func CreateCandidate(w http.ResponseWriter, r *http.Request) {
    var candidate models.Candidate
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &candidate); err != nil { // unmarshall body contents as a type Candidate
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := repository.CreateCandidate(candidate) // create candidate
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

// DELETE /candidate
func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"] // param id

    if err := repository.DeleteCandidate(id); err != "" { // delete a candidate by id
      panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    return
}
