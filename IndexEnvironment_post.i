
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend IndexEnvironment.i
//
type IndexEnvironment interface {
    Wrapped_IndexEnvironment

	SetDocumentRoot(arg2 string) (err error)
	SetAnchorTextPath(arg2 string) (err error)
	SetOffsetMetadataPath(arg2 string) (err error)
	SetOffsetAnnotationsPath(arg2 string) (err error)
	GetFileClassSpec(arg2 string) (_swig_ret Indri_parse_FileClassEnvironmentFactory_Specification, err error)
	AddFileClass(a ...interface{}) (err error)
	DeleteDocument(arg2 int) (err error)
	SetIndexedFields(arg2 StringVector) (err error)
	SetNumericField(a ...interface{}) (err error)
	SetOrdinalField(arg2 string, arg3 bool) (err error)
	SetParentalField(arg2 string, arg3 bool) (err error)
	SetMetadataIndexedFields(arg2 StringVector, arg3 StringVector) (err error)
	SetStopwords(arg2 StringVector) (err error)
	SetStemmer(arg2 string) (err error)
	SetMemory(arg2 int64) (err error)
	SetNormalization(arg2 bool) (err error)
	SetStoreDocs(arg2 bool) (err error)
	Create(a ...interface{}) (err error)
	Open(a ...interface{}) (err error)
	Close() (err error)
	AddFile(a ...interface{}) (err error)
	AddString(arg2 string, arg3 string, arg4 MetadataPairVector) (_swig_ret int, err error)
	AddParsedDocument(arg2 ParsedDocument) (_swig_ret int, err error)
	DocumentsIndexed() (_swig_ret int, err error)
	DocumentsSeen() (_swig_ret int, err error)
}

func NewIndexEnvironment() IndexEnvironment {
    return (IndexEnvironment)(SwigcptrWrapped_IndexEnvironment(C._wrap_new_Wrapped_IndexEnvironment_indri_go_add17ee78870902e()))
}

func (e SwigcptrWrapped_IndexEnvironment) SetDocumentRoot(arg2 string) (err error) {
    defer catch(&err)
    e.Wrapped_setDocumentRoot(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetAnchorTextPath(arg2 string) (err error) {
    defer catch(&err)
    e.Wrapped_setAnchorTextPath(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetOffsetMetadataPath(arg2 string) (err error) {
    defer catch(&err)
    e.Wrapped_setOffsetMetadataPath(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetOffsetAnnotationsPath(arg2 string) (err error) {
    defer catch(&err)
    e.Wrapped_setOffsetAnnotationsPath(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) GetFileClassSpec(arg2 string) (_swig_ret Indri_parse_FileClassEnvironmentFactory_Specification, err error) {
    defer catch(&err)
    _swig_ret = e.Wrapped_getFileClassSpec(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) AddFileClass(a ...interface{}) (err error) {
    defer catch(&err)
    argc := len(a)
	if argc == 1 {
		e.Wrapped_addFileClass(a[0].(Indri_parse_FileClassEnvironmentFactory_Specification))
		return
	}
	if argc == 12 {
		e.Wrapped_addFileClass(a[0].(string), a[1].(string), a[2].(string), a[3].(string), a[4].(string), a[5].(string), a[6].(string), a[7].(StringVector), a[8].(StringVector), a[9].(StringVector), a[10].(StringVector), a[11].(Std_map_Sl_indri_parse_ConflationPattern_Sm__Sc_std_string_Sg_))
		return
	}
	panic("No match for overloaded function call")
    return
}

func (e SwigcptrWrapped_IndexEnvironment) DeleteDocument(arg2 int) (err error) {
    defer catch(&err)
    e.Wrapped_deleteDocument(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetIndexedFields(arg2 StringVector) (err error) {
    defer catch(&err)
    e.Wrapped_setIndexedFields(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetNumericField(a ...interface{}) (err error) {
    defer catch(&err)
    argc := len(a)
	if argc == 2 {
		e.Wrapped_setNumericField(a[0].(string), a[1].(bool))
		return
	}
	if argc == 3 {
		e.Wrapped_setNumericField(a[0].(string), a[1].(bool), a[2].(string))
		return
	}
	panic("No match for overloaded function call")
    // can't nest variadit function
    //e.Wrapped_setNumericField(a)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetOrdinalField(arg2 string, arg3 bool) (err error) {
    defer catch(&err)
    e.Wrapped_setOrdinalField(arg2, arg3)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetParentalField(arg2 string, arg3 bool) (err error) {
    defer catch(&err)
    e.Wrapped_setParentalField(arg2, arg3)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetMetadataIndexedFields(arg2 StringVector, arg3 StringVector) (err error) {
    defer catch(&err)
    e.Wrapped_setMetadataIndexedFields(arg2, arg3)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetStopwords(arg2 StringVector) (err error) {
    defer catch(&err)
    e.Wrapped_setStopwords(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetStemmer(arg2 string) (err error) {
    defer catch(&err)
    e.Wrapped_setStemmer(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetMemory(arg2 int64) (err error) {
    defer catch(&err)
    e.Wrapped_setMemory(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetNormalization(arg2 bool) (err error) {
    defer catch(&err)
    e.Wrapped_setNormalization(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) SetStoreDocs(arg2 bool) (err error) {
    defer catch(&err)
    e.Wrapped_setStoreDocs(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) Create(a ...interface{}) (err error) {
    defer catch(&err)
    argc := len(a)
	if argc == 1 {
		e.Wrapped_create(a[0].(string))
		return
	}
	if argc == 2 {
		e.Wrapped_create(a[0].(string), a[1].(IndexStatus))
		return
	}
	panic("No match for overloaded function call")
    return
}

func (e SwigcptrWrapped_IndexEnvironment) Open(a ...interface{}) (err error) {
    defer catch(&err)
    argc := len(a)
	if argc == 1 {
		e.Wrapped_open(a[0].(string))
		return
	}
	if argc == 2 {
		e.Wrapped_open(a[0].(string), a[1].(IndexStatus))
		return
	}
	panic("No match for overloaded function call")
    return
}

func (e SwigcptrWrapped_IndexEnvironment) Close() (err error) {
    defer catch(&err)
    e.Wrapped_close()
    return
}

func (e SwigcptrWrapped_IndexEnvironment) AddFile(a ...interface{}) (err error) {
    defer catch(&err)
    argc := len(a)
	if argc == 1 {
		e.Wrapped_addFile(a[0].(string))
		return
	}
	if argc == 2 {
		e.Wrapped_addFile(a[0].(string), a[1].(string))
		return
	}
	panic("No match for overloaded function call")
    return
}

func (e SwigcptrWrapped_IndexEnvironment) AddString(arg2 string, arg3 string, arg4 MetadataPairVector) (_swig_ret int, err error) {
    defer catch(&err)
    _swig_ret = e.Wrapped_addString(arg2, arg3, arg4)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) AddParsedDocument(arg2 ParsedDocument) (_swig_ret int, err error) {
    defer catch(&err)
    _swig_ret = e.Wrapped_addParsedDocument(arg2)
    return
}

func (e SwigcptrWrapped_IndexEnvironment) DocumentsIndexed() (_swig_ret int, err error) {
    defer catch(&err)
    _swig_ret = e.Wrapped_documentsIndexed()
    return
}

func (e SwigcptrWrapped_IndexEnvironment) DocumentsSeen() (_swig_ret int, err error) {
    defer catch(&err)
    _swig_ret = e.Wrapped_documentsSeen()
    return
}


//
// the GO type MyStatusMonitor is subclassed from C++ type indri::api::IndexStatus
// this requires that director feature be enabled, look for
// '%module(directors="1") modulename' under ifdef SWIGGO in indri.i
//
// MyStatusMonitor is a superset of IndexStatus and hence can be used as a
// drop in replacement for IndexStatus but the reverse causes a compile time
// error. refer to swig docs for details
//
type MyStatusMonitor interface {
  IndexStatus
  DeleteIndexStatus()
  IsIndexStatus()
}

// Via embedding myStatusMonitor "inherits" all methods of IndexStatus.
type myStatusMonitor struct {
  IndexStatus
}

func (isgs *myStatusMonitor) DeleteIndexStatus() {
  DeleteDirectorIndexStatus(isgs.IndexStatus)
}

func (isgs *myStatusMonitor) IsIndexStatus() {}

type overwrittenMethodsOnIndexStatus struct {
  is IndexStatus
}

func (om *overwrittenMethodsOnIndexStatus) Status( code int, documentFile string, error string, documentsIndexed int, documentsSeen int ) {

    // C++ code defines class local enums, method signature uses int for code
    // swiggo generates constants we cannot use because of signature mismatch
    // note: must match C++ definition
    const (
        FileOpen = iota
        FileSkip
        FileError
        FileClose
        DocumentCount
    )

    switch(code) {
    case FileOpen:
        Buildindex_print_event( fmt.Sprintf("Opened %v", documentFile) )

    case FileClose:
        Buildindex_print_event( fmt.Sprintf("Closed %v", documentFile) )

    case FileSkip:
        Buildindex_print_event( fmt.Sprintf("Skipped %v", documentFile) )

    case FileError:
        Buildindex_print_event( fmt.Sprintf("Error in %v : %v", documentFile, error) )

    default:
        fallthrough
    case DocumentCount:
        if( documentsSeen < 10 || (documentsSeen % 500) == 0 ) {
            // note: C++ signature wants second count as INT64, but it is int here in GO
            //Buildindex_print_status( "Documents parsed: ", documentsSeen, " Documents indexed: ", documentsIndexed )
            // so use this instead.
            Buildindex_print_event( fmt.Sprintf("Documents parsed: %v Documents indexed: %v", documentsSeen, documentsIndexed) )
            Buildindex_flush_status()
        }
    }
}

func NewMyStatusMonitor() MyStatusMonitor {
  om := &overwrittenMethodsOnIndexStatus{}
  ins := NewDirectorIndexStatus(om)
  om.is = ins

  isgs := &myStatusMonitor{IndexStatus: ins}
  return isgs
}

func DeleteMyStatusMonitor(isgs MyStatusMonitor) {
    isgs.DeleteIndexStatus()
}

%}

#endif
