package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Home  string `json:"home"`
	Shell string `json:"shell"`
}

func main() {
	path, format := parseFlags()
	users := collectUsers()

	var output io.Writer

	if path != "" {
		f, err := os.Create(path)
		handleError(err)
		defer f.Close()
		output = f
	} else {
		output = os.Stdout
	}

	if format == "json" {
		data, err := json.MarshalIndent(users, "", "  ")
		handleError(err)
		output.Write(data)
	} else if format == "csv" {
		output.Write([]byte("name,id,home,shell\n"))
		writer := csv.NewWriter(output)
		for _, user := range users {
			err := writer.Write([]string{user.Name, strconv.Itoa(user.Id), user.Home, user.Shell})
			handleError(err)
		}
		writer.Flush()
	}
}

func collectUsers() (users []User) {
	f, err := os.Open("/etc/passwd")
	handleError(err)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ':'

	lines, err := reader.ReadAll()
	handleError(err)

	// Example line:
	// cloud_user:x:1002:1003::/home/cloud_user:/bin/bash
	for _, line := range lines {
		id, err := strconv.ParseInt(line[2], 10, 64)
		handleError(err)

		if id < 1000 {
			continue
		}

		user := User{
			Name:  line[0],
			Id:    int(id),
			Home:  line[5],
			Shell: line[6],
		}

		users = append(users, user)

	}

	return
}

func parseFlags() (path, format string) {
	flag.StringVar(&path, "path", "", "the path to the export file.")
	flag.StringVar(&format, "format", "json", "the output format for the user information. Available options are 'csv' and 'json'. The default format is json.")

	flag.Parse()

	format = strings.ToLower(format)

	if format != "csv" && format != "json" {
		fmt.Println("Error: invalid format. Use 'json' or 'csv' instead.")
		flag.Usage()
		os.Exit(1)
	}

	return
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
