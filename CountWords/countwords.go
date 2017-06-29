/*
* Saúl Ibáñez Cerro
* Grado en Telemática
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	f := func(c rune) bool {
		return c == ',' || c == ' ' || c == '\n' || c == '\t' || c == '\r'
	}

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error en la lectura del fichero: %v\n", err)
			continue
		}
		for _, line := range strings.FieldsFunc(string(data), f) {
			counts[line]++
		}
	}

	keys := []string{}
	for key, _ := range counts {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for i := range keys {
		key := keys[i]
		value := counts[key]
		fmt.Printf("%v\t%v", key, value)
		fmt.Println()
	}
}
