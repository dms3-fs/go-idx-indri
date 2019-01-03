
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend Parameters.i
//
type Parameters interface {
    Wrapped_Parameters

    // can't seem to name function Load, compiler declares it is a duplicate.
    // the wrapped method varies in signature only with the return type,
    // which is empty/void. perhaps this is a signature mapping 'feature'
    // when the c++ method return is void?
    MyLoad(string) error
}

func NewParameters() Parameters {
    return (Parameters)(SwigcptrWrapped_Parameters(C._wrap_new_Wrapped_Parameters_indri_go_add17ee78870902e()))
}

// Load catching panics generated in the %exception block and turn them into
// a go error type
func (e SwigcptrWrapped_Parameters) MyLoad(s string) (err error) {
	defer catch(&err)
    e.Load(s)
	return
}

%}

#endif
