
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend Specification.i
//
func DeleteFileClassSpec(spec Indri_parse_FileClassEnvironmentFactory_Specification) (err error) {
    defer catch(&err)
    Wrapped_deleteFileClassSpec(spec)
    return
}

%}

#endif
