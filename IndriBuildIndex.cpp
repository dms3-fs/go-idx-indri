//
// contents of IndriBuildIndex.cpp moved to IndriBuildIndex.i
//
// swig compiles all the cpp file in the folder, but when linking the package
// code wrapped via some swig interface file such as IndriBuildIndex.i cannot
// resolve to an id defined in a cpp file outside of the Module_wrap.cxx.
// attempts to do so results in:
//      error: ‘_recoverRepository’ was not declared in this scope
//
