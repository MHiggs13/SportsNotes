package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"

)

func NotesIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
    var newNote Note
    // read up to a max of 500GB of json to protect against malicious attacks
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))
    if err != nil {
        panic(err)
    }
    err = r.Body.Close()
    if err!= nil {
        panic(err)
    }

    err = json.Unmarshal(body, &newNote)
    if err!= nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        // 422 = unprocessable entity, couldn't unmarshal body into Note
        w.WriteHeader(422)
        err = json.NewEncoder(w).Encode(err)
        if err != nil {
            panic(err)
        }
        panic(err)
    }

    t := RepoCreateNote(newNote)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)   // status code: 201
    err = json.NewEncoder(w).Encode(t);
    if err != nil {
        panic(err)
    }

}

