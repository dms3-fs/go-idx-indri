
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend MetadataPairVector.i
//

type Wrapped_MetadataPairVector interface {
    MetadataPairVector

    WReserve(arg2 int64) (err error)
	WClear() (err error)
	WAdd(arg2 MetadataPair) (err error)
	WGet(arg2 int) (_swig_ret MetadataPair, err error)
	WSet(arg2 int, arg3 MetadataPair) (err error)
}

func NewWrapped_MetadataPairVector(a ...interface{}) Wrapped_MetadataPairVector {
	argc := len(a)
	if argc == 0 {
		return (Wrapped_MetadataPairVector)(SwigcptrMetadataPairVector(C._wrap_new_MetadataPairVector__SWIG_0_indri_go_add17ee78870902e()))
	}
	if argc == 1 {
        _swig_i_0 := a[0].(int)
    	return (Wrapped_MetadataPairVector)(SwigcptrMetadataPairVector(C._wrap_new_MetadataPairVector__SWIG_1_indri_go_add17ee78870902e(C.swig_type_128(_swig_i_0))))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrMetadataPairVector) WReserve(arg2 int64) (err error) {
    defer catch(&err)
    arg1.Reserve(arg2)
    return
}

func (arg1 SwigcptrMetadataPairVector) WClear() (err error) {
    defer catch(&err)
    arg1.Clear()
    return
}

func (arg1 SwigcptrMetadataPairVector) WAdd(arg2 MetadataPair) (err error) {
    defer catch(&err)
    arg1.Add(arg2)
    return
}

func (arg1 SwigcptrMetadataPairVector) WGet(arg2 int) (_swig_ret MetadataPair, err error) {
    defer catch(&err)

    //
    // we want to return the extended wrapped pair interface instead
    // (named MetadataPair). The C++ MetadataPair class is first renamed to
    // Wrapped_MetadataPair so we can subclass it in GO using the director
    // feature. When we wrap the (GO) MetadataPairVector specialization, it
    // references the renamed C++ class type and not the GO subclass. So the
    // generated code returns the renamed interface type Wrapped_MetadataPair
    // instead of the GO MetadataPairVector vector interface type.
    //_swig_ret = arg1.Get(arg2)
    //return

	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_ret = (MetadataPair)(SwigcptrWrapped_MetadataPair(C._wrap_MetadataPairVector_get_indri_go_add17ee78870902e(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return
}

func (arg1 SwigcptrMetadataPairVector) WSet(arg2 int, arg3 MetadataPair) (err error) {
    defer catch(&err)
    arg1.Set(arg2, arg3)
    return
}

func DeleteWrapped_MetadataPairVector(arg1 Wrapped_MetadataPairVector) {
    DeleteMetadataPairVector(arg1)
}

%}

#endif
