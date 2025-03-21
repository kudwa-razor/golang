// Testing code only for the output 3

/*
package main
import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer){
	fmt.Fprint(out, "3")
}
func main(){
	Countdown(os.Stdout)

}

*/

// Till here printing 3 is done. Now we can make it print 2,1,Go
// The backtick syntax used in below code is another way of creating a string but lets you include things like newlines, which is perfect for our test.
//Since it works for 3, we are going ahead with printing 2,1,Go


/*
package main
import (
	"fmt"
	"io"
	"os"
)
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}
*/


// The above is a code with for loop to count from 3 to 1 and then print GO, BUT  It does not have the 1 second stop.
// This is achieved thru time.sleep 
// The same above code is copied with one added line.

package main
import (
	"fmt"
	"io"
	"os"
)
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1*time.Second)
	}
	fmt.Fprint(out, "Go!")
}

