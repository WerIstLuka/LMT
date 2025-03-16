package main

import (
	"fmt"
	"os"
	"io"
	"strings"
	"slices"
	"math/big"
)

var Version string = "0.0.1"


func Help(){
	fmt.Println(`Math Stuff
Usage:
	lmt [Options] [Integers]
Options:
	add
	sub
	mul
	div
	pow
	sqrt
	abs`)
	os.Exit(0)
}

func HasPipeInput()bool{
	FileInfo, _ := os.Stdin.Stat()
	return FileInfo.Mode() & os.ModeCharDevice == 0
}

func GetArguments() ([]string, []string, string){
	Options := []string{}
	Numbers := []string{}
	Arguments := []string{}
	//read input from pipe if it exists
	if HasPipeInput(){
 		bytes, _ := io.ReadAll(os.Stdin)
		lines := strings.Split((string(bytes)), "\n")
		for i:=0; i<len(lines)-1; i++{
			line := strings.Split(lines[i], " ")
			Arguments = slices.Concat(Arguments, line)
		}
	}
	Arguments = slices.Concat(Arguments, os.Args[1:])
	//get arguments
	var IsNum bool
	for i := 0; i < len(Arguments); i++ {
		IsNum = true
		for j:=0;j<len(Arguments[i]);j++{
			if !strings.Contains("0123456789", string(Arguments[i][j])){
				IsNum = false
			}
		}
		if !IsNum{
			Options = append(Options, Arguments[i])
		}else{
			Numbers = append(Numbers, Arguments[i])
		}
	}
	if len(Arguments) == 0{
		Help()
	}
	return Options, Numbers, os.Args[0]
}

func add(Options []string, Numbers []*big.Int){
	Result := new(big.Int)
	for i:=0;i<len(Numbers);i++{
		Result.Add(Result, Numbers[i])
	}
	fmt.Println(Result)
}

func Parser(Options []string, Numbers []*big.Int, ExecName string){
	var GotOperation bool
	var Operation string
	var Remove int = -1
	if ExecName == "./LMT"{
		GotOperation = false
	}else{
		GotOperation = true
	}
	for i:=0;i<len(Options);i++{
		if slices.Contains([]string{"add", "sub", "mul", "div", "sqrt", "abs", "pow"}, Options[i]){
			if GotOperation{
				fmt.Println("Error: Operation given twice")
				os.Exit(1)
			}
			GotOperation = true
			Operation = Options[i]
			Remove = i
		}
	}
	if Remove != -1{
		Options = append(Options[:Remove], Options[Remove+1:]...)
	}
	if Operation == "add" || ExecName == "addl"{
		add(Options, Numbers)
	}
}

func ConvertNumbers(Numbers []string) ([]*big.Int){
	var ConvertedNumbers []*big.Int
	for i:=0;i<len(Numbers);i++{
		var BigNum big.Int
		Num, _ := BigNum.SetString(Numbers[i], 10)
		ConvertedNumbers = append(ConvertedNumbers, Num)
	}
	return ConvertedNumbers
}

func main(){
	Options, Numbers, ExecName := GetArguments()
	ConvertedNumbers := ConvertNumbers(Numbers)
	Parser(Options, ConvertedNumbers, ExecName)
	//Options has been changed in Parser
	//if Options is used after return it from Parser
}
