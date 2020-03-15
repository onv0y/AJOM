package main

import(
	"fmt"
	"os"
	"bufio"
)

var(
	work_xp = bufio.NewScanner(os.Stdin)
	person = bufio.NewScanner(os.Stdin)
	time = bufio.NewScanner(os.Stdin)
	skill1 = bufio.NewScanner(os.Stdin)
	skill2 = bufio.NewScanner(os.Stdin)
	skill3 = bufio.NewScanner(os.Stdin)
	position = bufio.NewScanner(os.Stdin)
	init1 = bufio.NewScanner(os.Stdin)
)

func nl() {
	fmt.Print("\n")
}

func initial() {
	fmt.Println("Enter name to initialize AJOM.")
	nl()
	fmt.Print("---> ")
	init1.Scan()
	nl()
}

func greeter() {
	fmt.Println("Welcome to AJOM " + init1.Text() + ". This program was built to help you come up with a snazzy job objective that is sure to wow. Lest get started shall we.")
}

func start() {
	fmt.Println("First things first, to get this resume built im going to need a couple of things.")
}

func work() {
	fmt.Println("Enter your most recent work experience.")
	fmt.Println("eg. Macdonalds, Home Depot, SYKES. ")
	nl()
	fmt.Print("---> ")
	work_xp.Scan()
	nl()
	fmt.Println("How long have you worked there? ")
	fmt.Println("eg. 1 year, 2 years, 3 months.")
	nl()
	fmt.Print("---> ")
	time.Scan()

}

func persontype() {
	nl()
	fmt.Println("If you were to describe yourself with one word what would it be? ")
	fmt.Println("eg. Adaptable, Honest, Commited, Dynamic, Motivated, Persistent.")
	nl()
	fmt.Print("---> ")
	person.Scan()
}

func skillz() {
	nl()
	fmt.Println("Please enter your top 3 skills: ")
	fmt.Println("eg. Conflict Resolution, Customer Communication, Contract Compliance.")
	nl()
	fmt.Println("Skill no 1: ")
	fmt.Print("---> ")
	skill1.Scan()
	nl()
	fmt.Println("Skill no 2: ")
	fmt.Print("---> ")
	skill2.Scan()
	nl()
	fmt.Println("Skill no 3: ")
	fmt.Print("---> ")
	skill3.Scan()
}

func pos() {
	nl()
	fmt.Println("Enter the position you are applying for. ")
	fmt.Print("---> ")
	position.Scan()
}

func write() {
    f, err := os.Create("AJOM.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(person.Text() + " professional with " + time.Text() + " of experience and a proven knowledge of " + skill1.Text() + ", " + skill2.Text() + ", and " + skill3.Text() + ". Aiming to leverage my skill to successfully fill the " + position.Text() + " position at your company.\n")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "Job Objective Successfully Created.")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func main(){
	nl()
	initial()
	nl()
	greeter()
	nl()
	start()
	nl()
	work()
	nl()
	persontype()
	nl()
	skillz()
	nl()
	pos()
	nl()
	write()
	nl()
}
