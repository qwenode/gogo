package main

import (
    "log"
    
    "github.com/qwenode/gogo/mmap"
)

func main() {
    checker := mmap.NewDuplicateChecker[string](true)
    log.Println(checker.Exist("xx"))
    checker.Add("xx")
    log.Println(checker.Exist("xx"))
}
