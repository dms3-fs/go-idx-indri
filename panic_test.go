package indri_go

import (
    //"errors"
    "fmt"
    //"io"
    //"io/ioutil"
    //"os"
    //"path/filepath"
    //"strings"
    "testing"
)

func dontPanic() (err string) {
    return
}

func doPanic() (err string) {
    panic("intentional panic")
    err = ""
    return
}

// protect against exception when [Indri] call does not return value
func protect(g func() (err string)) (err string) {
    err = "" // initialize return here for case when recover call returns nil.
             // otherwise (if set inside the deferred function when x is nill),
             // the return from g() in the normal flow would get wiped out.
    defer func() {
        if x := recover(); x != nil {
            err = fmt.Sprintf("run time panic: %v\n", x)
        }
        return
    }()
    err = g()
    return
}

// Test Go no panic
func TestGoNoPanic(t *testing.T) {
    expectedErr := ""
    err := protect(dontPanic)
    if err != "" {
        t.Errorf("Expected no error message %v but got %v", expectedErr, err)
    }
}

// Test Go panic
func TestGoPanic(t *testing.T) {
    err := protect(doPanic)

	if err == "" {
		t.Error("Expected an error with intentional panic.")
    } else {
        //fmt.Printf("library returned %v\n", err)
    }
}

// Test Indri static panic
func TestStaticMethodPanic(t *testing.T) {
    err := protect(StaticMethodPanic)
    if err == "" {
        t.Fatal(err)
    } else {
        //fmt.Printf("library returned %v\n", err)
    }
}

// Test Indri object panic
func TestObjectMethodPanic(t *testing.T) {
    tester := NewTester()
    if tester == nil {
        t.Fatal("failed to create Tester object")
    }
    defer DeleteTester(tester)

    sname := "testname"
    stype := "testtype"

    tester.SetName(sname)
    tester.SetType(stype)

    sn := tester.GetName()
    if sn != sname {
        t.Fatal(fmt.Sprintf("failed name test, expected: %v, received: %v", sname, sn))
    }
    st := tester.GetType()
    if st != stype {
        t.Fatal(fmt.Sprintf("failed type test, expected: %v, received: %v", stype, st))
    }

    err := protect(tester.InjectFault)
    if err == "" {
        t.Fatal(err)
    } else {
        //fmt.Printf("library returned %v\n", err)
    }

}

var (
	demo = NewDemoLib()
)

func TestThrowsNegativeThrows(t *testing.T) {
	expectedErr := "NegativeThrows threw exception"
	_, err := demo.NegativeThrows(-1)

	if err == nil {
		t.Fatal("Expected an error.")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %v but got %v", expectedErr, err.Error())
	}
}

func TestDivideByZero(t *testing.T) {
	expectedErr := "Cannot divide by zero"
	_, err := demo.DivideBy(0)

	if err == nil {
		t.Fatal("Expected an error when dividing by zero.")
	}
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %v but got %v", expectedErr, err.Error())
	}
}

func TestNeverThrowsReturnsInput(t *testing.T) {
	n := demo.NeverThrows(-1)

	if n != -1 {
		t.Errorf("Expected -1 but got %v", n)
	}
}
