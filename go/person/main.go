package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
)

type Person struct {
    Name string
    Age  int
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func (person *Person) Save() {
    fileName := fmt.Sprintf("%s.persondata", person.Name)
    personFile, err := os.Create(fileName)
    check(err)

    defer personFile.Close()

    _, err = personFile.WriteString(fmt.Sprintf("%d", person.Age))
    check(err)
}

func loadPerson(name string) Person {
    fileName := strings.Trim(name, " ") + ".persondata"
    data, err := os.ReadFile(fileName)
    check(err)

    personAge, err := strconv.Atoi(string(data))
    check(err)

    return Person{name, personAge}
}

func parseAdd() {
    var name string
    var ageStr string

    fmt.Println("What is the name?")
    fmt.Scanln(&name)

    fmt.Println("What is the age?")
    fmt.Scanln(&ageStr)

    age, err := strconv.Atoi(ageStr)
    check(err)

    person := Person{name, age}
    person.Save()

    fmt.Println(fmt.Sprintf("Saved %s.", person.Name))
}

func parseView() {
    var name string

    fmt.Println("What is the name?")
    fmt.Scanln(&name)

    person := loadPerson(name)
    fmt.Println(fmt.Sprintf("%s is %d years old.", person.Name, person.Age))
}

func main() {
    for {
        var userInput string
        fmt.Println("What do you want to do? (add/view):")
        fmt.Scanln(&userInput)
        switch userInput {
        case "add":
            parseAdd()
        case "view":
            parseView()
        default:
            fmt.Println("Invalid input. Please try again.")
        }
    }
}
