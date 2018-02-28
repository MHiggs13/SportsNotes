package main

import (
    "time"
)

type Note struct {
        Id int `json:"id"`
        Name string `json:"name"`
        Positives []string `json:"positives"`
        Negatives []string `json:"negatives"`
        Created time.Time `json:"created"`
        Edited time.Time `json:"edited"`

}

type Notes []Note
