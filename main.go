package main

import (
	"fmt"
)
func main() {
	bangun()
}
func bangun (){

	// belah ketupat
	fmt.Println("belah ketupat")
	for f := 0; f < 5; f++ {
		for g := 5; g > f; g-- {
			fmt.Print("  ")
		}
		for h := 0; h <= f; h++ {
			fmt.Print("* ")
		}
		for i := 0; i <= f; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for j := 0; j < 5; j++ {
		for l := 0; l <= j; l++ {
			fmt.Print("  ")
		}
		for k := 5; k > j; k-- {
			fmt.Print("* ")
		}
		for l := 5; l > j; l-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// persegi
	fmt.Println("persegi")
	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n ++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}


	// persegi panjang
	fmt.Println("persegi panjang")
	for o := 0; o < 5; o++ {
		for p := 0; p < 10; p++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// segitiga siku
	fmt.Println("segitiga siku")
	for q := 0; q < 5; q++ {
		for r := 0; r <= q; r++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// segitga sama sisi lancip diatas
	fmt.Println("segitiga sama sisi")
	for s := 0; s < 5; s++ {
		for t := 5; t > s; t-- {
			fmt.Print("  ")
		}
		for u := 0; u <= s; u++ {
			fmt.Print("* ")
		}
		for v := 0; v < s; v++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// trapesium 1
	fmt.Println("trapesuim 1")
	for w := 0; w < 5; w++ {
		for x := 0; x < 5; x++ {
			fmt.Print("* ")
		}
		for y := 0; y < w; y++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// trapesium 2
	fmt.Println("trapesium")
	for b1 := 0; b1 < 5; b1++ {
		for b2 := 0; b2 < 5; b2++ {
			fmt.Print("* ")
		}
		for b3 := 5; b3 > b1; b3-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// trapesium 3
	fmt.Println("trapesium 3")
	for a1 := 0; a1 < 5; a1++ {
		for a2 := 5; a2 > a1; a2-- {
			fmt.Print("  ")
		}
		for a3 := 0; a3 <= a1; a3++ {
			fmt.Print("* ")
		}
		for a4 := 0; a4 < 5; a4++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// trapesuim 4
	fmt.Println("trapesium 4")
	for c1 := 0; c1 < 5; c1++ {
		for c3 := 0; c3 <= c1; c3++ {
			fmt.Print("  ")
		}
		for c2 := 5; c2 > c1; c2-- {
			fmt.Print("* ")
		}
		for c4 := 0; c4 < 5; c4++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}