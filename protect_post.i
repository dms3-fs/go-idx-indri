
#ifdef SWIGGO

/*
 * protect_post.i - see protect_pre.i for comments
 */

%go_import("fmt")
%go_import("runtime")

%insert(go_wrapper) %{

//
// catch will recover from a panic and store the recover message to the error
// parameter. The error must be passed by reference in order to be returned to the
// calling function.
//
//
// see index_environment_test.go for more details on the following note.
// ...
// however, the "catch" panic handler does not trap all the exceptions and
// the cgo crosscall results in a segmentation violation [ most likely due
// the pair's value, both void * and value length, are in random state ]
//
// TODO: figure out how we may trap all C exceptions.
// there seems to be path cases where the execution thread context (thread
// stack and frame pointer) is switched on syscalls (to avoid GC and other
// GO thread from blocking. And on return, the completion may run on a
// different thread. This panic handler may be on a different stack???
// review comments in go/src/runtime/cgocall.go
// ...
//
func catch(err *error) {
    if r := recover(); r != nil {
        var sf []string = []string{}
        var pc []uintptr
        var skip, fc int = 0, 0
        fc = runtime.Callers(skip, pc)
        if fc > 0 {
            fp := runtime.CallersFrames(pc)
            if fp != nil {
                var more bool
                var f runtime.Frame
                for more = true; more; f, more = fp.Next() {
                    sf = append(sf, fmt.Sprintf("%v %v %v %v\n", f.PC, f.Function, f.File, f.Line))
                }
            }
        }
        *err = fmt.Errorf("Panic stack %v, error %v", sf, r)
    }
}

%}

%include "Parameters_post.i"
%include "IndexEnvironment_post.i"
%include "IndriBuildIndex_post.i"
%include "MetadataPair_post.i"
%include "MetadataPairVector_post.i"


#endif
