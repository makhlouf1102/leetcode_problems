package palindrome-number

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }
    
    half := 0
    for x > half {
        half = (half * 10) + (x % 10)
        x /= 10
    } 

    return (x == half) || (x == half/10)

}
