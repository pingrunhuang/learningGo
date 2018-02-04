package main
import(
	"os"
	"fmt"
	"strconv"
)

func main() {
	for index, value := range(os.Args) {
		fmt.Println(strconv.Itoa(index) + value)
	}
}
