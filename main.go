package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numTask int64
	var err error

	fmt.Println("Sprint")

	readerNum := bufio.NewReader(os.Stdin)
	fmt.Print("Enter amount tasks: ")
	nt, _, _ := readerNum.ReadLine()

	numTask, err = strconv.ParseInt(string(nt), 10, 64)
	if err != nil {
		fmt.Printf("F: %v\n", err)
		panic("Worng first parameter")
	}

	if numTask%2 != 0 {
		panic("Not even number of tasks")
	}

	readerCost := bufio.NewReader(os.Stdin)
	fmt.Print("Enter cost of every task: ")
	costLine, _, _ := readerCost.ReadLine()
	cl := strings.Split(string(costLine), " ")

	if len(cl) != int(numTask) {
		panic("!!!!")
	}

	sliceCosts := make([]int64, numTask)
	for _, i := range cl {
		c, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic("Worng cost parameter")
		}
		sliceCosts = append(sliceCosts, c)
	}

}
