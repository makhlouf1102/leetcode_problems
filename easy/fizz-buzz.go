package easy

import "strconv"

func fizzBuzz(n int) []string {
    result := make([]string, n);
    var str string;
    for i := 1; i <= n; i++ {
        str = "";
        if (i % 3) == 0 {
            str += "Fizz";
        } 
        if (i % 5) == 0 {
            str += "Buzz";
        } 
        
        if str == "" {
            str = strconv.Itoa(i);
        }

        result[i-1] = str;
    }

    return result;
}