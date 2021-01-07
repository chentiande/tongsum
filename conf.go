package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type sqls struct {
	Sqlname string `xml:"sqlname"`
	Sql     string `xml:"sql"`
}
type sqlconf struct {
	Sqls []sqls `xml:"sqls"`
}

type Configuration struct {
	Dbip   string `xml:"dbip"`
	Dbport string `xml:"dbport"`
	Dbname string `xml:"dbname"`
	Dbuser string `xml:"dbuser"`
	Dbpwd  string `xml:"dbpwd"`
}

func getconf(filename string) (Configuration, error) {
	xmlFile, err := os.Open(filename)
	var conf Configuration
	if err != nil {
		fmt.Println("Error opening file:", err)
		return conf, err
	}
	defer xmlFile.Close()

	if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
		fmt.Println("Error Decode file:", err)
		return conf, err
	}

	return conf, nil

}

func getsql(filename string) (sqlconf, error) {
	xmlFile, err := os.Open(filename)
	var conf sqlconf
	if err != nil {
		fmt.Println("Error opening file:", err)
		return conf, err
	}
	defer xmlFile.Close()

	if err := xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
		fmt.Println("Error Decode file:", err)
		return conf, err
	}

	return conf, nil

}
