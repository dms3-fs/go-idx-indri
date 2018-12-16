/* -----------------------------------------------------------------------------
 * Tester.i
 * used to validate go panic protection schemes for C/C++ triggered panics.
 *
 * notes:
 * 1) exception handling order controls behavior. Meaning, if the wrong
 * handler is reached first, the resulting behavior may not be as expected.
 * use exception.i when throwing std_exception, see below for usage example.
 * use setEx(method) when throwing LemurException, see below for usage example.
 *
 * 2) TODO:
 * need to figure out what controls wrapper generation of id suffixes, as in
 * _libindri_go_c5c8d3b43a2c09c9 . does not seem to change on each compile.
 * Would like to use %module name expansion formatting and append to fixed
 * id name. The suffixes below were manually edited into the MODULE.go file
 * after first running swig to generate the go wrapper module.
 *
 *
 * This interface file generates code to demonstrate several C++ panic patterns.
 * The panic_test.go file demonstrate corresponding go protection patterns:
 *
 * a) An example that shows how to protect against go panic pattern can be found
 * in panic_test TestGoPanic. This protection pattern converts a go
 * panic message into a go string, an empty string is returned when there is no
 * panic. This pattern is suitable when calling go functions that do not
 * return a value. The go panic message returned is of type "string".
 *
 * b) a static method (StaticMethodPanic) defined in this file that vectors to
 * SWIG_exception defined in swig's go/exception.i, that dispatches to
 * _swig_gopanic defined in go/goruntime.swg, that in turn crosses into the
 * cgo compiler environment (_cgo_panic). The C/C++ panic message returned is
 * of type "const char *".
 *
 * An example that shows how to protect against this panic pattern can be found
 * in panic_test TestStaticMethodPanic. This protection pattern converts a
 * panic message into a go string, an empty string is returned when there is no
 * panic. This pattern is suitable when calling static C/C++ methods that do not
 * return a value. The go panic message returned is of type "string".
 *
 * c) a C++ object class (Tester) method defined in this file that vectors to
 * the LemurExcepion.i handler, a common exception pattern used in Indri. This
 * pattern requires the C++ methods that throw an exception be wrapped with
 * "setEx" to dispatch exceptions to SWIG_exception as discussed above.
 *
 * An example that shows how to protect against this panic pattern can be found
 * in panic_test TestObjectMethodPanic.
 *
 * d) Preferred pattern - wrap arbitrary C/C++ class and add error return type
 * to class method signatures. from the reference pattern:
 *      https://github.com/jsolmon/go-swig-exceptions
 * In this pattern, a C++ class is renamed to wrap original methods that throw
 * exceptions to vector into corresponding wrapped methods that set error the
 * return using the same SWIG_exception dispatch path as discussed above.
 *
 * Examples that show how to protect against this panic pattern can be found
 * in panic_test TestThrowsNegativeThrows and TestDivideByZero.
 *
 * ----------------------------------------------------------------------------- */

#ifdef SWIGGO

%{

#ifdef INDRI_STANDALONE
#include "lemur/Exception.hpp"
#else
#include "Exception.hpp"
#endif

static char *StaticMethodPanic() {
    SWIG_exception(SWIG_RuntimeError,"forced exception");
    char *s = (char *)"";
    return s;
}

namespace indri{

    namespace api{

        class Tester {
            std::string name;
            std::string type;

        public:

            std::string GetName() {
                return name;
            }
            std::string GetType() {
                return type;
            }

            void SetName(std::string v) {
                name = v;
            }
            void SetType(std::string v) {
                type = v;
            }

            std::string InjectFault() {
                LEMUR_THROW( LEMUR_GENERIC_ERROR, "forced exception");
                std::string s = "";
                return s;
            }
        };
    }
}

class DemoLib {

public:
    DemoLib() {}

    double DivideBy(int n) {
        if (n == 0) {
            // use exception.i when throwing std_exception
            throw std::invalid_argument("Cannot divide by zero");
        }
        return 1.0 / n;
    }

    int NegativeThrows(int in) {
        if (in < 0) {
            // use exception.i when throwing std_exception
            throw std::range_error("NegativeThrows threw exception");
        }
        return in;
    }

    int NeverThrows(int in) {
        return in;
    }
};

%}

static char *StaticMethodPanic();

namespace indri{

    namespace api{

        setEx(Tester::InjectFault());

        class Tester {
            std::string name;
            std::string type;

        public:

            std::string GetName();
            std::string GetType();
            void SetName(std::string v);
            void SetType(std::string v);
            std::string InjectFault();
        };

    }
}


%include "exception.i"

// The %exception directive will catch any exception thrown by the C++ library and
// panic() with the same message.
%exception {
    try {
        $action;
    } catch (std::exception &e) {
        _swig_gopanic(e.what());
    }
}

// Rename the DemoLib class to Wrapped_DemoLib so that it can be wrapped in an
// DemoLib interface in the go_wrapper. Same for all methods that throw exceptions
// which need to be turned into errors.
%rename(Wrapped_DemoLib) DemoLib;
%rename(Wrapped_NegativeThrows) DemoLib::NegativeThrows;
%rename(Wrapped_DivideBy) DemoLib::DivideBy;

class DemoLib {

public:
    DemoLib();

    // throws invalid_argument exception if arg = 0
    double DivideBy(int);
    // throws range_error exception if arg < 0
    int NegativeThrows(int);

    // does not throw any exceptions
    int NeverThrows(int);
};

%go_import("fmt")

%insert(go_wrapper) %{

type DemoLib interface {
    Wrapped_DemoLib
    NegativeThrows(int) (int, error)
    DivideBy(int) (float64, error)
}

func NewDemoLib() DemoLib {
    //// TODO: need to figure out what controls id suffixes generation
    //// they don't seem to change each time we recompile/.
    //// the suffix used here and in functions below was manually edited into
    //// the MODULE.go file after swig generates it.
    return (DemoLib)(SwigcptrWrapped_DemoLib(C._wrap_new_Wrapped_DemoLib_indri_go_8e24520e0b567faa()))
}

// catch will recover from a panic and store the recover message to the error
// parameter. The error must be passed by reference in order to be returned to the
// calling function.
func catch(err *error) {
    if r := recover(); r != nil {
        *err = fmt.Errorf("%v", r)
    }
}

// NegativeThrows catched panics generated in the %exception block and turns them into
// a go error type
func (e SwigcptrWrapped_DemoLib) NegativeThrows(n int) (i int, err error) {
	defer catch(&err)
	i = e.Wrapped_NegativeThrows(n)
	return
}

func (e SwigcptrWrapped_DemoLib) DivideBy(n int) (f float64, err error) {
    defer catch(&err)
	f = e.Wrapped_DivideBy(n)
	return
}

%}

#endif
