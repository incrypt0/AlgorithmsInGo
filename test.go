// package main

// import (
// 	"bufio"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	inpVim := scanner.Text()
// 	arrVim := strings.Fields(inpVim)
// 	intArrVim := stringToIntSlice(arrVim)
// 	inpSush := scanner.Text()
// 	arrSush := strings.Fields(inpSush)
// 	intArrSush := stringToIntSlice(arrSush)

// }
// func stringToIntSlice(sli []string) []int {
// 	intSli := []int{}
// 	for _, item := range sli {
// 		i, err := strconv.Atoi(item)
// 		if err != nil {
// 			log.Fatal(err)
// 			return nil
// 		}
// 		intSli = append(intSli, i)
// 	}
// 	return intSli
// }
