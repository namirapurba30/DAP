/* SID: 1301184423
Name: Diya Namira Purba
This program is doing Bit
*/

package main

import (
	"fmt"
)

func main() {
	var (
		w, h, wh, nt int
		p            string
	)

	fmt.Scan(&p)
	fmt.Scan(&w)
	fmt.Scan(&h)
	fmt.Scan(&nt)

	for i := 0; i < h; i++ {
		for k := 0; k < w; k++ {
			fmt.Scan(&wh)
			if wh < 100 {
				fmt.Print("#")

			} else if wh >= 100 {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}

}
