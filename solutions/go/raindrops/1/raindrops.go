package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
    var result strings.Builder
    
	if number % 3 == 0 {
        result.WriteString("Pling")
    } 
    if number % 5 == 0 {
        result.WriteString("Plang")
    }
    if number % 7 == 0 {
        result.WriteString("Plong")
    }

    if result.Len() == 0 {
        return strconv.Itoa(number)
    }

    return result.String()
}
