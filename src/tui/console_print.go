package tui

import (
	"fmt"
)

const DefaultTabSize = 4

func PrintMap(inputMap map[string]interface{}) {
	printMapWithPadding(inputMap, []int{DefaultTabSize})
}

func printMapWithPadding(inputMap map[string]interface{}, padding []int) {
	fmt.Println(" *** config *** ")
	fmt.Println("[")

	// mapLen := len(inputMap)
	cnt := 0
	for key, value := range inputMap {
		curPadding := padding[len(padding)-1]
		printSpace(curPadding)
		fmt.Printf("\"%s\": ", key)

		switch value.(type) {
		case map[string]interface{}, map[string]string, map[string]int, map[int]interface{}:
			fmt.Printf("{\n")

			padding = append(padding, curPadding+DefaultTabSize)
			printMapWithPadding(value.(map[string]interface{}), padding)
			padding = padding[:len(padding)-1]

			printSpace(curPadding)
			fmt.Printf("}")
		default:
			fmt.Printf("%v", value)
			// if we wanted comma
			// if cnt < mapLen-1 {
			// fmt.Print(",")
			// }
			cnt++
			fmt.Println()
		}
	}

	fmt.Println("]")
}

func printSpace(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" ")
	}
}
