package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("请选择使用场景 ==> ")
		for _, sc := range Scenarios {
			fmt.Printf("场景: %s ,", sc.Name)
			printDescription(sc.Description)
		}
		return
	}
	for _, arg := range os.Args[1:] {
		sc := matchScenario(arg)
		if sc != nil {
			printDescription(sc.Description)
			printExamples(sc.Examples)
			sc.RunExample()
		}
	}
}

func printDescription(str []string) {
	fmt.Printf("场景描述: %s \n", str)
}

func printExamples(str []string) {
	fmt.Printf("场景举例: %s \n", str)
}

func matchScenario(name string) *Scenario {
	for _, sc := range Scenarios {
		if sc.Name == name {
			return sc
		}
	}
	return nil
}

var doSomething = func(i int) string {
	time.Sleep(time.Millisecond * time.Duration(10))
	fmt.Printf("Goroutine %d do things .... \n", i)
	return fmt.Sprintf("Goroutine %d", i)
}

var takeSomthing = func(res string) string {
	time.Sleep(time.Millisecond * time.Duration(10))
	tmp := fmt.Sprintf("Take result from %s.... \n", res)
	fmt.Println(tmp)
	return tmp
}
