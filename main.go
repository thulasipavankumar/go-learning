package main
import "fmt"

func main(){

	// this is constant declaration and type will be assinged to float
	const pi = 3.14;
	// this is declaring a variable of int it can also be var age =30 and age will be predefined to 30
	var age int = 30
	// this is declaring a variable without datatype notice we use ":=" for dynamic declaring variable
	name := "Naruto"
	fmt.Println("Hello World!");
	fmt.Println(strings.Repeat("#",10));
	fmt.Println("Your name is ",name,"age is: ",age);
	fmt.Println(pi);

}
