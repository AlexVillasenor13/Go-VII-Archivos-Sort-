package main

import (
	"fmt"
	"sort"
	"os"
	"strings"
)

type ByString []string

func (a ByString) Len() int           { return len(a) }
func (a ByString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByString) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var n int
	var aux_string string
	fmt.Println("Capture numero de cadenas: ")
	fmt.Scan(&n)
	string_slide := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		fmt.Println("Capture cadena numero ", i)
		fmt.Scan(&aux_string)
		string_slide = append(string_slide, aux_string)
	}

	sort.Sort(ByString(string_slide))
	writeFile(strings.Join(string_slide, "\n"), "asecendente")
	sort.Sort(sort.Reverse(ByString(string_slide)))
	writeFile(strings.Join(string_slide, "\n"), "descendente")

}

func writeFile(content, file_name string){
	file, err := os.Create(file_name + ".txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}
