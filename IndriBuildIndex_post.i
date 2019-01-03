
#ifdef SWIGGO

%insert(go_wrapper) %{

//
//  extend wrapped IndriBuildIndex.i
//
func Wrapped_Buildindex_recoverRepository(arg1 string) (_swig_ret bool, err error) {
    defer catch(&err)
    _swig_ret = Buildindex_recoverRepository(arg1)
    return
}

func Wrapped_Buildindex_augmentSpec(spec Indri_parse_FileClassEnvironmentFactory_Specification,
                        fields StringVector,
                        metadata StringVector,
                        metadataForward StringVector,
                        metadataBackward StringVector ) (_swig_ret bool, err error) {
    defer catch(&err)
    _swig_ret = Buildindex_augmentSpec(spec,
                            fields,
                            metadata,
                            metadataForward,
                            metadataBackward )
    return
}

func Wrapped_Buildindex_mymain(argc int, arg0, arg1, arg2, arg3, arg4 string) (err error) {
    defer catch(&err)
    Buildindex_mymain(argc, arg0, arg1, arg2, arg3, arg4)
    return
}

%}

#endif
