package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Number string
	Name   string
	Age    int
}

func main() {
	var students []Student
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var student Student

		fmt.Print("Enter student number (or type 'exit' to finish): ")
		scanner.Scan()
		student.Number = scanner.Text()
		if student.Number == "exit" {
			break
		}

		fmt.Print("Enter student name: ")
		scanner.Scan()
		student.Name = scanner.Text()

		fmt.Print("Enter student age: ")
		scanner.Scan()
		age, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid age. Please enter a number.")
			continue
		}
		student.Age = age

		students = append(students, student)
	}

	file, err := os.Create("students.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, student := range students {
		record := []string{student.Number, student.Name, strconv.Itoa(student.Age)}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to file:", err)
			return
		}
	}

	fmt.Println("Student records have been written to students.csv")
}
