/*

	The code was written and tested on a Windows 8.1 machine running Go 1.7
	The idea is to use a recursion (goroutine) instead of FOR loops. I did not use channels because I do not think I need it here.
	
	Visit royleekiat.com to look at my other mini projects 
	(Software development (Chrome extensions, Splunk, Nodejs), InfoSecurity (Buffer overflow attacks- C program, Vulnerable FTP server) using Ubuntu and Kali Linux and Windows)

*/

package main

//import scanners, printers and string manipulators
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

	//Initialize Scanner to take in stdin
	Scanner := bufio.NewScanner(os.Stdin)

	//checkNum is used to store a counter so that only the lines with the desired input are taken in.
	//sumArray stores all the sum of squares.
	//numTestCases is to check against checkNum so that scanning ends when all test cases are entered and calculated.
	checkNum := 0
	sumArray := []int{}
	numTestCases := 0

	//Scan stdin, process and print stdout
	scanInput(checkNum, numTestCases, sumArray, Scanner)
	
}


//Recursive function
func squareSum (testcase []string)(int){

	//Convert from string to int for further processing
	i, err := strconv.Atoi(testcase[0])

	//Error checking 
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
	}

	//Initialize new array(slice), removing first element from previous array(slice)
	rest :=  testcase[1:] 
	//Calculate square of first element in array
	square := i*i  

	//Switch-case for deciding to call recursive function or return sum on its own
	switch {
		//int element is negative and this is the last int element in array(slice), return 0 because negative elements are not considered.
	    case i <0 && len(rest) == 0:
	        return 0
	    //int element is negative and there are more int elements in array(slice), return 0 and call function because negative elements are not considered.
	    case i < 0 && len(rest) > 0:
	        return 0 + squareSum(rest) 
	    //return square of int element in array, and do not call function because this is the last element in the array(slice)
	    case i >=0  && len(rest) == 0:
	        return square
	    //return square of int element in array, and call function because this is not the last element in the array(slice)
	    case i >=0  && len(rest) > 0:
	        return square + squareSum(rest) 
    }

    //Redundant, but Go needs it
    return 0

}

//Print array(slice) without using for loops
func printArray (sumArray []int){
	rest := sumArray[1:]
	fmt.Println(sumArray[0])

	if len(rest) == 0 {
		return
	}else{
		printArray(rest)
	}
	return
}

//Use recursion to scan instead of the classical "for Scanner.scan(){}"", returns nothing. Prints using printArray within this function itself
func scanInput(checkNum int, numTestCases int, sumArray []int, Scanner *bufio.Scanner){

	//Read user input from cmd line by line 
	Scanner.Scan() 

	//Takes in a line of input
	input := Scanner.Text()

	//Store Number of test cases in numTestCases
	if checkNum==0 {

		//Convert from string to int (Number of test cases)
		firstInput, err := strconv.Atoi(input)

		//Error checking 
		if err != nil {
		    fmt.Println(err)
		    os.Exit(0)
		//If 0 test cases added
		}else if err == nil && firstInput == 0{
			//No test cases, hence EXIT.
			os.Exit(0)
		}

		//Store parsed input into numTestCases
		numTestCases = firstInput
	}
	
	

	//Slice input by spaces
	result := strings.Split(input, " ")

	//ignore first line of input (which stands for number of test cases), and lines of input that stands for number of INT arguments in each test case
	if checkNum >0 && checkNum%2 ==0{
		//Call recursive function squareSum
		sum := squareSum(result)
		//Add sum into array(slice)
		sumArray = append(sumArray,sum)
	}

	//Break when all test cases are processed and entered
	if(numTestCases == checkNum/2){
		//Print all sum of squares without using for loop
		printArray(sumArray)
		return 
	}else{
		//Increment counter
		checkNum += 1
		scanInput(checkNum, numTestCases, sumArray, Scanner)
	}
	
	return
}

//Thanks
//Visit royleekiat.com to look at my other out-of school mini projects
//Roy Lee, Singapore