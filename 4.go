package main

import "fmt"

func main() {
    elements := make(map[string]string)
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B": "Boron",
    "C": "Carbon",
    "N": "Nitrogen",
    "O": "Oxygen",
    "F": "Fluorine",
    "Ne": "Neon",

    fmt.Println(elements["Li"])
  
    if name, ok := elements["Un"]; ok {    
    fmt.Println(name, ok)
    }
  
  elements := map[string]map[string]string{
            "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        }
   if el, ok := elements["Li"]; ok {    
   fmt.Println(el["name"], el["state"])
   }
}
