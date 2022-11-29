package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
)

func main() {

	list := readLineFromFile("links.txt")
	for _, v := range list {

		dnsCheck(v)
		fmt.Println()

	}

}

func readLineFromFile(fileName string) []string {

	filePath := "./" + fileName
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines

}

func dnsCheck(d string) {
	colorRed := "\033[33m"
	ips, err := net.LookupIP(d)

	checkNilErr(err)
	for _, IP := range ips {
		cfmt.Println(string(colorRed), d+" : {{"+IP.String()+"}}::red")

	}

}
func checkNilErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
