package main

import "fmt"

func main61() {

	//1. General format
	// %v    value default format
	// %+v   struct with field names
	// %#v   Go-syntax representation
	// %T    type of the value
	// %%    literal percent sign

	// i := 15.5
	i := 111_505.5
	string := "Happy code!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", string)
	fmt.Printf("%#v\n", string)
	fmt.Printf("%T\n", string)

	//2. INTEGER format
	// 	%d    decimal (base 10)
	// %b    binary
	// %c    character (rune)
	// %o    octal
	// %O    octal with 0o prefix
	// %x    hex (a-f)
	// %X    hex (A-F)
	// %U    Unicode format (U+1234)
	// %q    quoted character literal

	int := 255

	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	fmt.Printf("%c\n", int)

	//3. String format
	// 	%s    string
	// %q    quoted string ("hello")
	// %x    hex dump (lowercase)
	// %X    hex dump (uppercase)

	txt := "Coding!"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// 4. Boolean format
	// %t    true / false

	t := true
	f := false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t", f)

	//5. Float format
	// 	%f    decimal notation (no exponent)
	// %e    scientific notation (e.g. 1.23e+06)
	// %E    scientific (uppercase E)
	// %g    shortest of %f or %e
	// %G    shortest of %F or %E

	flt := 94735.189723723

	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%4.2f\n", flt)

	fmt.Printf("%e\n", flt)
	fmt.Printf("%E\n", flt)
	fmt.Printf("%g\n", flt)
	fmt.Printf("%G\n", flt)

}
