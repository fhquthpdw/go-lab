package main

import (
	"errors"
	"fmt"
)

// errors.Is
// errors.As

func main() {
	var err error

	err = fmt.Errorf("one: %w", errors.New("yes, one"))
	fmt.Printf("%s\n", errors.Unwrap(err))
	fmt.Println(errors.Unwrap(err))
	fmt.Println()

	err = fmt.Errorf("two: %w", errors.New("yes, two"))
	fmt.Printf("%s\n", errors.Unwrap(err))
	fmt.Printf("%s\n", err)
	fmt.Println()

	err = fmt.Errorf("three: %w", errors.New("yes, three"))
	fmt.Printf("%s\n", errors.Unwrap(err))
	fmt.Printf("%s\n", err)
	fmt.Println()

	/*
		// errors with stack
		err := f3()
		fmt.Printf("%+v", err)

		fmt.Println("==============")

		// wrap errors
		errWrap := ff3()
		fmt.Printf("%+v", errWrap)
	*/
}

/*
// stack
func f1() error {
	err := errors.New("this is f1")
	return err
}
func f2() error {
	err := f1()
	err = errors.WithMessage(err, "this is f2")
	return err
}
func f3() error {
	err := f2()
	err = errors.WithMessage(err, "this is f3")
	return err
}

// wrap
func ff1() error {
	err := errors.New("this is f1")
	return err
}
func ff2() error {
	err := ff1()
	err = fmt.Errorf("%w, this is wrap in ff2", err)
	return err
}
func ff3() error {
	err := ff2()
	err = fmt.Errorf("%w, this is wrap in ff3", err)
	return err
}
*/
