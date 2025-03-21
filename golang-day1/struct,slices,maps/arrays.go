	package main
	import "fmt"

	func main(){
		//diff ways to initialize an array

		//1st way
		var a [5]int
		fmt.Println("Array 1:", a) //since uve not assigned any value, each element is 0 for int array.

		// set specific index values.
		a [4] = 100
		fmt.Println("Set to:", a)
		//get specific index
		fmt.Println("Get index:", a[4])

		//get length of arr
		fmt.Println("Length of arr:", len(a))

		//2nd way to initialize an array, but in 1 line only
		b := [5]int{1, 2, 3, 4, 5}
    	fmt.Println("declared one line:", b)

		//Let coompiler count no of elements using "..."
		b = [...]int{1, 2, 3, 4, 5}
		fmt.Println("declared one line, but u dont have to enter the no of elements:", b)

		//3rd way, but you're initializing index, and the between elements = 0
		b = [...]int{100, 3: 400, 500}
    	fmt.Println("elements btween index = 0:", b)

		// 2D array initializing
		var twoD [2][3] int
		for i :=0; i<2; i++{
			for j:=0; j<3; j++{
				twoD[i][j] = i+j
			}
		}
		fmt.Println("2d: ", twoD)

		// 2D array in one line only like 1D
		twoD = [2][3] int{
			{1, 2, 3},
			{3, 2, 1},
		}
		fmt.Println("2D in one line:", twoD)

	}