// Go
package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "strconv"

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)
	line, _ = reader.ReadString('\n')
	splitted := strings.Split(line, " ")
	result := make([]int64, len(splitted))
	for i, elem := range splitted {
		x, _ := strconv.ParseInt(strings.TrimSpace(elem), 10, 64)
		result[i] = x
	}
	fmt.Fprintf(writer, "%d\n", n)
	for _, elem := range result {
		fmt.Fprintf(writer, "%d ", elem)
	}
	writer.Flush()
}
