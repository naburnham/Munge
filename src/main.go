package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func deduplicate(list []string) []string {
	keys := make(map[string]bool)
	dedupList := []string{}
	for _, word := range list {
		if _, value := keys[word]; !value {
			keys[word] = true
			dedupList = append(dedupList, word)
		}
	}
	return dedupList
}

func simpleLeet(word string) (list []string) {
	/* Apply basic LeetCode replacements to passwords on single characters. */
	list = append(list, strings.ReplaceAll(word, "a", "4"))
	list = append(list, strings.ReplaceAll(word, "a", "@"))
	list = append(list, strings.ReplaceAll(word, "b", "8"))
	list = append(list, strings.ReplaceAll(word, "b", "13"))
	list = append(list, strings.ReplaceAll(word, "b", "!3"))
	list = append(list, strings.ReplaceAll(word, "e", "3"))
	list = append(list, strings.ReplaceAll(word, "i", "1"))
	list = append(list, strings.ReplaceAll(word, "i", "!"))
	list = append(list, strings.ReplaceAll(word, "l", "1"))
	list = append(list, strings.ReplaceAll(word, "o", "0"))
	list = append(list, strings.ReplaceAll(word, "s", "5"))
	list = append(list, strings.ReplaceAll(word, "s", "$"))
	list = append(list, strings.ReplaceAll(word, "t", "7"))
	list = append(list, strings.ReplaceAll(word, "t", "+"))
	list = append(list, strings.ReplaceAll(word, "1", "i"))
	list = append(list, strings.ReplaceAll(word, "1", "l"))
	list = append(list, strings.ReplaceAll(word, "1", "!"))
	list = append(list, strings.ReplaceAll(word, "4", "a"))
	list = append(list, strings.ReplaceAll(word, "7", "t"))
	return deduplicate(list)
}

func comboLeet(word string) (list []string) {
	/* Apply 2 different, but common, leetspeak styles for chracter replacement */
	temp := word
	temp = strings.ReplaceAll(word, "a", "4")
	temp = strings.ReplaceAll(word, "A", "4")
	temp = strings.ReplaceAll(word, "b", "8")
	temp = strings.ReplaceAll(word, "B", "8")
	temp = strings.ReplaceAll(word, "e", "3")
	temp = strings.ReplaceAll(word, "E", "3")
	temp = strings.ReplaceAll(word, "i", "!")
	temp = strings.ReplaceAll(word, "l", "1")
	temp = strings.ReplaceAll(word, "o", "0")
	temp = strings.ReplaceAll(word, "O", "0")
	temp = strings.ReplaceAll(word, "s", "$")
	temp = strings.ReplaceAll(word, "S", "$")
	temp = strings.ReplaceAll(word, "t", "7")
	list = append(list, temp)

	temp = word
	temp = strings.ReplaceAll(word, "a", "@")
	temp = strings.ReplaceAll(word, "A", "@")
	temp = strings.ReplaceAll(word, "b", "13")
	temp = strings.ReplaceAll(word, "B", "13")
	temp = strings.ReplaceAll(word, "e", "3")
	temp = strings.ReplaceAll(word, "E", "3")
	temp = strings.ReplaceAll(word, "i", "1")
	temp = strings.ReplaceAll(word, "l", "1")
	temp = strings.ReplaceAll(word, "o", "0")
	temp = strings.ReplaceAll(word, "O", "0")
	temp = strings.ReplaceAll(word, "s", "5")
	temp = strings.ReplaceAll(word, "S", "5")
	temp = strings.ReplaceAll(word, "t", "+")
	list = append(list, temp)

	return deduplicate(list)
}

func appendMunge(word string) (list []string) {
	/* Appends common characters to the end of the word */
	characters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "."}
	other_endings := []string{"@123", "123", "@1234", "1234", "12345", "234", "007"}
	// Common character endings
	for _, each := range characters {
		list = append(list, word+each)
	}
	// Other common password endings
	for _, each := range other_endings {
		list = append(list, word+each)
	}
	// Append all numbers between 0 and 99
	for i := 0; i <= 99; i++ {
		list = append(list, word+strconv.Itoa(i))
	}
	// Append birth years
	for i := 1980; i <= 2023; i++ {
		list = append(list, word+strconv.Itoa(i))
	}
	return deduplicate(list)
}

func simpleMunge(word string) (list []string) {
	// Append original word and TitleCased word
	list = append(list, word)
	caser := cases.Title(language.Und)
	list = append(list, caser.String(strings.ToLower(word)))

	return deduplicate(list)
}

func generator(inputFile string) <-chan string {
	out := make(chan string)
	go func() {
		file, err := os.Open(inputFile)
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		close(out)
	}()
	return out
}

func writeFile(outputFile string, information []string) {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	check(err)

	defer file.Close()

	for _, element := range information {
		_, err := file.WriteString(element + "\n")
		check(err)
	}
}

func main() {
	var helpFlag = flag.Bool("help", false, "Display help.")
	var inputFile = flag.String("i", "dictionary.txt", "Define desired input file.")
	var outputFile = flag.String("o", "munged.txt", "Define desired output file.")
	var level = flag.Int("level", 4, "Define level of munges [0-4]. Defaults to 4.")
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *level < 0 {
		*level = 0
	}
	if *level > 4 {
		*level = 4
	}

	information := generator(*inputFile)

	for word := range information {
		var wordlist []string
		fmt.Println(word)

		if *level >= 0 {
			for _, i := range simpleMunge(word) {
				wordlist = append(wordlist, i)
			}
		}
		if *level >= 1 {
			for _, i := range simpleLeet(word) {
				wordlist = append(wordlist, i)
			}
		}
		if *level >= 2 {
			for _, i := range comboLeet(word) {
				wordlist = append(wordlist, i)
			}
		}
		if *level >= 4 {
			for _, word := range wordlist {
				for _, i := range simpleMunge(word) {
					wordlist = append(wordlist, i)
				}
				for _, i := range simpleLeet(word) {
					wordlist = append(wordlist, i)
				}
				for _, i := range comboLeet(word) {
					wordlist = append(wordlist, i)
				}
			}
		}

		if *level >= 3 {
			for _, word := range wordlist {
				for _, i := range appendMunge(word) {
					wordlist = append(wordlist, i)
				}
			}
		}
		writeFile(*outputFile, deduplicate(wordlist))
		fmt.Println("Finished with", word)
	}
}
