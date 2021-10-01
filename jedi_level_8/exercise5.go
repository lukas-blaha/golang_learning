package main

import (
	"fmt"
    "sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int            { return len(a) }
func (a ByAge) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool  { return a[i].Age < a[j].Age }

type ByLast []user

func (la ByLast) Len() int            { return len(la) }
func (la ByLast) Swap(i, j int)       { la[i], la[j] = la[j], la[i] }
func (la ByLast) Less(i, j int) bool  { return la[i].Last < la[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
    fmt.Println("--------------")
    sort.Sort(ByAge(users))
    printUsers(users)
    fmt.Println("--------------")
    sort.Sort(ByLast(users))
    printUsers(users)

}

func printUsers(users []user) {
    for _, u := range users {
        sort.Strings(u.Sayings)
        fmt.Println(u.First, u.Last, u.Age)
        for _, v := range u.Sayings {
            fmt.Println("\t", v)
        }
    }
}
