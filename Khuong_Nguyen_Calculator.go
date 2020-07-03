package main

import (
 "bufio"
 "fmt"
 "os"
 "stack"
 "math"
 "strings"
)
// Global operator and operand stacks
var operandStack stack.Stack
var operatorStack stack.Stack

/* Power does Returns x ^ y. This is a brute force integer power routine using successive
 multiplication. (There are more efficient ways to do this.) */
func Power(x float64, y float64) (pow float64) {
	pow = 1.0
	for i := 0.0 ; i < y ; i++ {
		pow *= x    
	}
	return
}

// Returns true if the character is a digit.
func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// Returns the precedence of the operator.
func precedence(op byte) (prec int) {
	switch op {
	case '+', '-': prec = 0
	case '*', '/': prec = 1
	case '^': prec = 2
	default: panic("unknown operator")
	}
	return
}

func apply(){
	// Pop the operator off the operator stack
	op, err := operatorStack.Pop()
     
	if err != nil { 
		panic("operator stack underflow")
	}
	// Pop the right operand off the operand stack
	right, err := operandStack.Pop()
    
	if err != nil {
		panic("operand stack underflow")
	}
	// Pop the left operand off the operand stack
	left, err := operandStack.Pop()
    
	if err != nil {
		panic("operand stack underflow")
	}
	// Apply the operator to the left and right operands and push the result
	// onto the operand stack
    //If left and right are integer
    //Have extra 
   
	switch op.(byte) {
	case '+': operandStack.Push(left.(float64) + right.(float64))
	case '-': operandStack.Push(left.(float64) - right.(float64))
	case '*': operandStack.Push(left.(float64) * right.(float64))
	case '/': operandStack.Push(left.(float64) / right.(float64))
	case '^': operandStack.Push(Power(left.(float64), right.(float64)))
	default: panic("unknown operator")
	}
	return

} 
//Evaluate expression and print the result

func evaluate(expr string) (float64,int){
    // Initialize the operator and operand stacks
	operandStack = stack.New()
	operatorStack = stack.New()
       
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("illegal expression:", r)   
		}
	}()
    //
	closeBracket := 0 
    //
    breakBracket := false
	// Process the expression character by character left to right
	operandExpected := true
	i := 0
     
	for i < len(expr) {
     
		switch expr[i] {
		// Digit: Extract the operand and push it on the operand stack
        
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if !operandExpected {
				panic("operator expected but operand found")
			}
			integralPart := 0.0
            //Before . 
			for i < len(expr) && isDigit(expr[i]) {
				integralPart = float64(10.0*integralPart)+float64(expr[i]-'0')
				i++
			}
           
            fractionalPart := 0.0
            if i < len(expr) && expr[i] == '.' {
                i++
                //After .
                c := 0
                for i < len(expr) && isDigit(expr[i]) {
                    c++
                    fractionalPart = (math.Pow(10,float64(-c))*float64(expr[i]-'0') + fractionalPart)
                    i++
                }
            }
               
            v := float64(integralPart + fractionalPart)
            //Take floating point here
            //Need to push floating point 
            
			operandStack.Push(v)
			operandExpected = false
        //Do floating point
            
		// Operator: Apply pending operators of greater or equal precedence
		// then push the operator on the operator stack
		case '+', '-', '*', '/', '^':
			if operandExpected {
				panic("operand expected but operator found")
			}
			for !operatorStack.IsEmpty() {  
				op,_:= operatorStack.Top()
                //Need to fix precedence for parenthese
                //Right to left ^ ^ ^ ^ 
				if precedence(op.(byte)) >= precedence(expr[i]) {
					apply()
				} else {
					break
				}
			}
			operatorStack.Push(expr[i])
			i++
			operandExpected = true
		case '(':
			i++
            a :=  operatorStack
            b :=  operandStack
                
            operand, closeBracket := evaluate(expr[i:len(expr)])
            
            operandExpected = false
            operatorStack = a
            operandStack = b
            
            operandStack.Push(operand)
            
            j := i + closeBracket
            for i < j {
             i++
            }
		case ')':
            
            i++
            operandExpected = false
            closeBracket = i
              //Break the loop
            breakBracket = true
		case ' ':
			i++
		default:
			panic(fmt.Sprintf("%q is an illegal character", expr[i]))
		}
        
        if breakBracket {
             break
        }
	}
	
	for !operatorStack.IsEmpty() {
		apply()
	}
	// The result is the one operator remaining on the stack.
	result, _ := operandStack.Pop()
   
	if !operandStack.IsEmpty() {
		panic("too many operands")
	}
	
    return result.(float64) ,closeBracket
 }

//Make string to create a new string to do right to left for ^
func makeString(expr string) string{
    i := 0
    newline := ""
    bracket := 0
    for i < len(expr){
        if expr[i] == '+'{
             bracket--
             newline += string(')')
             newline += string(expr[i])
        }else{
            newline += string(expr[i])
        }
        if expr[i] == '^'{
            newline += string('(')
            bracket++
            newline += makeString(expr[i+1:len(expr)])
            break
         }
        i++
    }
    for bracket > 0{
        newline += string(')')
        bracket--
    }
   
    return newline
}
// Main routine to read expressions from standard input, calculate their values,
// and print the result. (Use an end of file, control-Z, to exit.)
func main() {
	
	// Make a scanner to read lines from standard input
	scanner := bufio.NewScanner(os.Stdin)
	
	// Process each of the lines from standard input
	for scanner.Scan() {
		// Get the current line of text
		line := scanner.Text()
		// fmt.Println(line)
	
        if strings.Count(line, "^") >= 2 {
            newline := makeString(line)
            result , _ := evaluate(newline)
            fmt.Println(result)
        }else{
	
         result , _ := evaluate(line)
         fmt.Println(result)
         }
            
       
        }
}
