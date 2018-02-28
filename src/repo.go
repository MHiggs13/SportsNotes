package main

import (
)

var currentId int

var notes Notes

func RepoFindNote(id int) Note {
    for _, n := range notes {
        if n.Id == id {
            return n
        }
    }
    // no Note found so return empty Note
    return Note{}
}

func RepoCreateNote(n Note) Note {
    currentId += 1
    n.Id = currentId
    notes = append(notes, n)
    return n
}
