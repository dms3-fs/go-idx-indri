
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend MetadataPair.i
//

type MetadataPair interface {
    Wrapped_MetadataPair

    WSetKey(arg2 string) (err error)
	WGetKey() (_swig_ret string, err error)
	WSetValue(arg2 []byte) (err error)
	WGetValue() (_swig_ret []byte, err error)

    WReset()
    WClear() (err error)
}

func NewMetadataPair() (_swig_ret MetadataPair, err error) {
    _swig_ret = (MetadataPair)(SwigcptrWrapped_MetadataPair(C._wrap_new_Wrapped_MetadataPair_indri_go_add17ee78870902e()))
    if _swig_ret == nil {
        return
    }
    _swig_ret.WReset()
	return
}

func DeleteMetadataPair(arg1 MetadataPair) {

    // don't leak resource
    arg1.WClear()

    DeleteWrapped_MetadataPair(arg1)
    return
}

func (m SwigcptrWrapped_MetadataPair) WReset() {
    var zvs string
    var zvv uintptr = 0

    // initialize key/value to zero values
    m.SetKey(zvs)
    m.SetValue(zvv)
    m.SetValueLength(0)

    return
}

func (m SwigcptrWrapped_MetadataPair) WClear() (err error) {
    defer catch(&err)

    // overrwrite reference to GO resource
    var zvs string
    m.SetKey(zvs)

    // don't leak C++ resource
    p := m.GetValue()
    l := m.GetValueLength()
    if l > 0 && p > 0 {
        Swig_free(p)
    }
    m.WReset()

    return
}

func (m SwigcptrWrapped_MetadataPair) WSetKey(arg2 string) (err error) {
    defer catch(&err)
    m.SetKey(arg2)
    return
}

func (m SwigcptrWrapped_MetadataPair) WGetKey() (_swig_ret string, err error) {
    defer catch(&err)
    _swig_ret = m.GetKey()
    return
}

func (m SwigcptrWrapped_MetadataPair) WSetValue(arg2 []byte)  (err error) {
    defer catch(&err)

    // don't leak C++ resource that may have already been set
    p := m.GetValue()
    l := m.GetValueLength()
    if l > 0 && p > 0 {
        Swig_free(p)
    }

    p = Swig_malloc(len(arg2))
    s := (*[1<<30]byte)(unsafe.Pointer(p))[:len(arg2)]

    i := copy(s, arg2)
    if i != len(arg2) {
        err = fmt.Errorf("Wrapped_MetadataPair.WSetValue: expected to copy %v but copied %v", len(arg2), i)
        return
    }

    m.SetValue(uintptr(unsafe.Pointer(p)))
    m.SetValueLength(len(arg2))

    //fmt.Printf("WSetValue: len %v cap %v arg2 %v i %v\n", len(arg2), cap(arg2), arg2, i)
    return
}

func (m SwigcptrWrapped_MetadataPair) WGetValue() (_swig_ret []byte, err error) {
    defer catch(&err)

    p := m.GetValue()
    len := m.GetValueLength()
    s := (*[1<<30]byte)(unsafe.Pointer(p))[:len]
    _swig_ret = make([]byte, len, len)

    i := copy(_swig_ret, s)
    if i != len {
        err = fmt.Errorf("Wrapped_MetadataPair.WGetValue: expected to copy %v but copied %v", len, i)
    }

	return
}

%}

#endif
