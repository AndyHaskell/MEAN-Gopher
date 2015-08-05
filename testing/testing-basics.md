# Basics of Go testing

## Why test?
 - If you have a structured set of automated tests for your code and run them as you keep developing your code, you'll know exactly when your feature breaks instead of having to backtracking.
 - You'll have the assumptions of how your code is supposed to work written down formally in the form of your tests.
 - Because of that, in many Go packages, tests serve as "the other documentation".
 - It has you break down how your code works into its elementary parts, so developing software with testing in mind is conducive to modular code that makes good use of Go's OOP idioms.

## Making test files

To make test files for your Go package, simply make a file in your package that ends with `_test.go`. For example, `webapp_test.go`.

Inside your test file, you import the testing package by importing `testing`.

To write a function, make a function whose name starts with `Test` and takes in a `*testing.T` (for example, if you're testing a function telling if a number is a Fibonacci number, you would do `func TestIsFib(*testing.T)`).

### `testing.T`

A `testing.T` is a struct that manages your test and allows you to log testing errors and other information. A `testing.T` logs an error with its `Errorf` method or logs an error and stops the test file altogether with its `Fatalf` method. If can also log information without making a test fail its `Logf` method.

In a Go test function, if a `testing.T` calls its `Fail` method (either directly or through `Errorf` or `Fatalf`), the test function fails.

### An example test file

Say we have a function in our package `fibonacci`

```go
package fibonacci

func IsFib(n int64) bool {
    var twoNumbersAgo int64 = 0
    var oneNumberAgo int64 = 0

    var i int64
    for i = 1; i < n; i = twoNumbersAgo + oneNumberAgo {
        twoNumbersAgo = oneNumberAgo
        oneNumberAgo = i
    }

    return i == n
}
```

that determines if a number is a Fibonacci number.

We could write this test file `test_fibonacci.go`:

```go
package fibonacci

import "testing"

func TestIsFib(t *testing.T) {
    if !IsFib(0) {
        t.Errorf("Testing IsFib failed for 0, expected true got false")
    }
    if !IsFib(1) {
        t.Errorf("Testing IsFib failed for 1, expected true got false")
    }
    if !IsFib(55) {
        t.Errorf("Testing IsFib failed for 55, expected true got false")
    }
    if IsFib(100) {
        t.Errorf("Testing IsFib failed for 100, expected false got true")
    }
}
```

### Running the test

We would then run this test by running `go test`, which would give us the output

```
--- FAIL: TestIsFib (0.00s)
        fibonacci_test.go:7: Testing IsFib failed for 0, expected true got false
FAIL
exit status 1
FAIL    fibonacci       1.283s
```

If we then add to the beginning of `IsFib` this if statement:

### Passing the test

```go
if n == 0 {
    return true
}
```

When we run `go test` we would get the test output:

```
PASS
ok      fibonacci       1.066s
```

## Your tests are Go code

A major point to note about Go tests is that **tests are go code**. Your tests are code as much as your main code is, so you can use Go code as you see fit to make your tests easy to work with and make them clear.

As an example, let's DRY (Don't Repeat Yourself) the test we had with the helper function:

```go
func expect(t *testing.T, cond string, expectation, result bool) {
    if result != expectation {
        t.Errorf("Expected %s to be %t, got %t", cond, expectation, result)
    }
}
```

Now we can change our testing code to the much more concise

```go
func TestIsFib(t *testing.T) {
    expect(t, "IsFib(0)",   true,  IsFib(0))
    expect(t, "IsFib(1)",   true,  IsFib(1))
    expect(t, "IsFib(55)",  true,  IsFib(55))
    expect(t, "IsFib(100)", false, IsFib(100))
}
```

The [test code for Negroni](https://github.com/codegangsta/negroni/blob/master/negroni_test.go) has a more generalized function for giving tests `expect` syntax:

```go
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
```

This also means you can write test helper functions to abstract away confusing or reused setup and teardown code for your Go tests, or, for data structures defined as `interface`s, you can make testing variations of their methods that allow you to add test-specific functionality for your data structures.
