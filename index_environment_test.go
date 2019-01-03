package indri_go

import (
    "errors"
    "fmt"
    //"io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "testing"

    util "github.com/dms3-fs/go-fs-util"
)

/*

Full list of index environment interface to cover in this test module.

type IndexEnvironment interface {
	SetDocumentRoot(arg2 string)
	SetAnchorTextPath(arg2 string)
	SetOffsetMetadataPath(arg2 string)
	SetOffsetAnnotationsPath(arg2 string)
	GetFileClassSpec(arg2 string) (_swig_ret Indri_parse_FileClassEnvironmentFactory_Specification)
	AddFileClass(a ...interface{})
	DeleteDocument(arg2 int)
	SetIndexedFields(arg2 StringVector)
	SetNumericField(a ...interface{})
	SetOrdinalField(arg2 string, arg3 bool)
	SetParentalField(arg2 string, arg3 bool)
	SetMetadataIndexedFields(arg2 StringVector, arg3 StringVector)
	SetStopwords(arg2 StringVector)
	SetStemmer(arg2 string)
	SetMemory(arg2 int64)
	SetNormalization(arg2 bool)
	SetStoreDocs(arg2 bool)
	Create(a ...interface{})
	Open(a ...interface{})
	Close()
	AddFile(a ...interface{})
	AddString(arg2 string, arg3 string, arg4 Std_vector_Sl_indri_parse_MetadataPair_Sg_) (_swig_ret int)
	AddParsedDocument(arg2 ParsedDocument) (_swig_ret int)
	DocumentsIndexed() (_swig_ret int)
	DocumentsSeen() (_swig_ret int)
}

index environment interface not yet exercised.

    Open(a ...interface{})

index environment interface unlikely to be exercised.

    SetOrdinalField(arg2 string, arg3 bool)
    SetParentalField(arg2 string, arg3 bool)
    AddParsedDocument(arg2 ParsedDocument) (_swig_ret int)

*/

/**
 * Basic test to load various index environment configuration parameters.
**/
func TestIndexEnv(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testIndexEnv()
    if err != nil {
        t.Fatal(err)
    }
}

/**
 * Test to read/write additional index environment configuration parameters.
**/
func TestIndexEnvReadWrite(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testParamsReadWrite()
    if err != nil {
        t.Fatal(err)
    }
}

//
// the next three tests demonstrate index repository initialization
// using variant flavors of logic a la IndriBuildIndex.cpp
//

/**
 *
 * This test builds a IndriBuildIndex command line and calls into C++ side
 * buildindex_main to create and index repository.
 *
 * The test demostrates programatically invoking IndriBuildIndex from GO.
 *
**/
func TestBuildIndexCPP(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testBuildIndexCPP()
    if err != nil {
        t.Fatal(err)
    }
}

/**
 * This test mimics IndriBuildIndex logic in GO to construct configuration
 * parameters piecemeal using an allocated Parameters object that implements
 * only a subset of the C++ Parameters public methods. It then calls
 * IndexEnvironment set* and other methods to configure the IndexEnvironment
 * and create and index repository.
**/
func TestBuildIndexGO(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testBuildIndexGO()
    if err != nil {
        t.Fatal(err)
    }
}

/**
 *
 * This test is a superset of TestBuildIndexGO()
 *
 * This test mimics IndriBuildIndex logic to construct configuration parameters
 * by first loading a params file and then overriding file configured parameters
 * piecemeal using an allocated Parameters object that implements only a
 * subset of the C++ Parameters public methods. It then calls IndexEnvironment
 * set* and other methods to configure the IndexEnvironment and create and
 * index repository.
 *
**/
func TestBuildIndexGOWithParamsFile(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testBuildIndexGOWithParamsFile()
    if err != nil {
        t.Fatal(err)
    }
}


//
// start of test logic implimentations
//

func testIndexEnv() (err error) {

    defer catch(&err)

    var qe QueryEnvironment = NewQueryEnvironment()

    if forcepanic {
        panic("forced panic")
        return
    }

    if forceerror {
        qe = nil
    }

    if qe == nil {
        err = errors.New("NewQueryEnvironment() returned nil")
        return
    }

    var p Parameters

    // create index parameters object
    p = NewParameters()
    defer DeleteWrapped_Parameters(p)

    // read the parameters file as a string
    pars := "data/params.xml"
    s, err := readParamsFile(pars)
    if err != nil {
        return
    }
    //fmt.Printf("params file content: %v\n", s)

    // parse and load parameters
    err = p.MyLoad(s)
    if err != nil {
        return
    }

    //fmt.Printf("p.Size() = %v\n", p.Size())

    // check for required parameters
    if !requireParameter( "corpus", p ) {
        err = errors.New("corpus parameter is required")
        return
    }
    if !requireParameter( "index", p ) {
        err = errors.New("index parameter is required")
        return
    }
    if !requireParameter( "corpus.path", p ) {
        err = errors.New("corpus.path parameter is required")
        return
    }

    return
}

func testIndexEnvReadWrite() (err error) {

	defer catch(&err)

    var p Parameters

    // create index parameters object
    p = NewParameters()
    defer DeleteWrapped_Parameters(p)

    // parse and load parameters
    err = p.MyLoad(good_params_1)
    if err != nil {
        return
    }

    var b, eb bool
    eb = true
    b = p.Get_bool("normalize", !eb)
    if b != eb {
        err = fmt.Errorf("expected <normalize>true</normalize> to be %v", eb)
        return
    }

    var ei, i int
    var ef, f float64
    ei = 100000000
    i = p.Get_int("memory", ei*2)
    if i != ei {
        err = fmt.Errorf("expected <memory>100m</memory> to be %v", ei)
        return
    }

    ef = 3.3
    f = p.Get_double("dummy.double", ef)
    if f != ef {
        err = fmt.Errorf("expected default %v received %v", ef, f)
        return
    }

    var ei64, i64 int64
    ei64 = 1234567890
    i64 = p.Get_INT64("dummy.int64", ei64)
    if i64 != ei64 {
        err = fmt.Errorf("expected default %v received %v", ei64, i64)
        return
    }

    var es, s string
    es = "krovetz"
    s = p.Get_string("stemmer.name", es)
    if s != es {
        err = fmt.Errorf("expected default %v received %v", es, s)
        return
    }

    eb = true
    b = p.Exists("stemmer.name")
    if b != eb {
        err = fmt.Errorf("expected <normalize>true</normalize> to be %v", eb)
        return
    }

    es = "not"+es
    p.Set_string("stemmer.name", es)
    s = p.Get_string("stemmer.name", es)
    if s != es {
        err = fmt.Errorf("expected default %v received %v", es, s)
        return
    }

    eb = !eb
    p.Set_bool("normalize", eb)
    b = p.Get_bool("normalize", !eb)
    if b != eb {
        err = fmt.Errorf("expected <normalize>true</normalize> to be %v", eb)
        return
    }

    ei = ei / 4
    p.Set_int("memory", ei)
    i = p.Get_int("memory", ei*3)
    if i != ei {
        err = fmt.Errorf("expected <memory>100m</memory> to be %v", ei)
        return
    }

    ei64 = ei64 / 2
    p.Set_UINT64("dummy.int64", ei64)
    i64 = p.Get_INT64("dummy.int64", ei64*3)
    if i64 != ei64 {
        err = fmt.Errorf("expected default %v received %v", ei64, i64)
        return
    }

    ef = ef * 3.3
    p.Set_double("dummy.double", ef)
    f = p.Get_double("dummy.double", ef*4)
    if f < ef-0.01 || f > ef+0.01 {
        err = fmt.Errorf("expected default %v received %v", ef, f)
        return
    }

    es = "parameters.corpus.path" // must be full path to xml node
    p.Remove(es)
    s = p.Get_string(es, "")
    if len(s) > 0 {
        err = fmt.Errorf("did not expect key %v to contain value %v", es, s)
        return
    }

    p.Clear()
    i64 = p.Size()
    if i64 != 0 {
        err = fmt.Errorf("expected size() %v to be 0", i64)
        return
    }

    return
}

func testBuildIndexCPP() (err error) {

    defer catch(&err)

    dir, err := ioutil.TempDir("", "test-repo-root")
	if err != nil {
        err = fmt.Errorf("failed to create TempDir: %v", err)
	}
    defer os.RemoveAll(dir) // clean up

    var repoKey, repositoryPath = "index", filepath.Join(dir, "index-1")
    var corpPathKey, corpusPath = "corpus.path", filepath.Join(dir, "corpus-1")
    var corpusFileClassKey, corpusFileClass = "corpus.class", "html"

    err = os.Mkdir(corpusPath, 0775)
    if err != nil {
        err = fmt.Errorf("failed to create corpus dir: %v", err)
    }

    var argv []string = []string {
        "IndriBuildIndex",
        "-" + repoKey + "=" + repositoryPath,
        "-" + corpPathKey + "=" + corpusPath,
        "-" + corpusFileClassKey + "=" + corpusFileClass,
        "data/params.xml",
    }
    var argc int = len(argv)

    err = Wrapped_Buildindex_mymain(argc, argv[0], argv[1], argv[2], argv[3], argv[4])

    return
}

func testBuildIndexGO() (err error) {

    defer catch(&err)

    var p Parameters

    // create index parameters object
    p = NewParameters()
    defer DeleteWrapped_Parameters(p)

    //
    // indri applications accept command line options as space separated strings.
    // each command line option string specifies a paramater value, or a set of
    // parameter values indirectly specified as a path to an xml encoded file.
    // the top level element in the xml file is always named <em>parameters</em>.
    //
    // if an option string command line argument (i.e. non-first command line
    // argument) starts with the dash character, i.e. '-', what follows is a
    // parameter option name supported by the indri application. otherwise, it
    // is assumed to be a path string pointing to the xml file.
    //
    // here, we are not using argc/argv, but building index paramaters the
    // hard way. another approach wraps IndriBuildIndex C++ code to have GO
    // code delegate to the C++ code to build index, as this is what it is
    // already coded to do.
    //
    //p.loadCommandLine( argc, argv );

    dir, err := ioutil.TempDir("", "test-repo-root")
	if err != nil {
        err = fmt.Errorf("failed to create TempDir: %v", err)
	}
    defer os.RemoveAll(dir) // clean up

    var repoKey, repositoryPath = "index", filepath.Join(dir, "index-1")
    var corpPathKey, corpusPath = "corpus.path", filepath.Join(dir, "corpus-1")
    var corpusFileClassKey, corpusFileClass = "corpus.class", "html"
    var corpMetadataKey, corpMetadataPath = "corpus.metadata", filepath.Join(repositoryPath, "metadata")

    err = os.Mkdir(corpusPath, 0775)
    if err != nil {
        err = fmt.Errorf("failed to create corpus dir: %v", err)
    }

    p.Set_string(repoKey, repositoryPath)
    p.Set_string(corpPathKey, corpusPath)
    p.Set_string(corpusFileClassKey, corpusFileClass)
    p.Set_string(corpMetadataKey, corpMetadataPath)

    // check for required parameters
    if !requireParameter( "corpus", p ) {
        err = fmt.Errorf("parameter %v is required", "corpus")
        return
    }
    if !requireParameter( "index", p ) {
        err = fmt.Errorf("parameter %v is required", repoKey)
        return
    }
    if !requireParameter( "corpus.path", p ) {
        err = fmt.Errorf("parameter %v is required", corpPathKey)
        return
    }

    var monitor MyStatusMonitor = NewMyStatusMonitor()
    defer DeleteMyStatusMonitor(monitor)
    var env IndexEnvironment = NewIndexEnvironment()
    defer DeleteWrapped_IndexEnvironment(env)

    var v string

    Buildindex_start_time();

    fmt.Printf("indri distribution version: %v\n", GetIndriVersion())

    v = p.Get_string(repoKey, "")
    if len(v) < 1 {
        err = fmt.Errorf("parameter %v is required", repoKey)
        return
    } else {
        fmt.Printf("parameter %v has value %v\n", repoKey, v)
    }

    v = p.Get_string(corpPathKey, "")
    if len(v) < 1 {
        err = fmt.Errorf("parameter %v is required", corpPathKey)
        return
    } else {
        fmt.Printf("parameter %v has value %v\n", corpPathKey, v)
    }

    v = p.Get_string(corpMetadataKey, "")
    fmt.Printf("parameter %v has value %v\n", corpMetadataKey, v)

    var memKey, memVal = "memory", int64(1024*1024*1024)
    var normKey, normVal = "normalize", true
    var injKey, injVal = "injectURL", true
    var storeKey, storeVal = "storeDocs", true

    memVal = p.Get_INT64(memKey, memVal)
    p.Set_UINT64(memKey, memVal)

    normVal = p.Get_bool(normKey, normVal)
    p.Set_bool(normKey, normVal)

    injVal = p.Get_bool(injKey, injVal)
    p.Set_bool(injKey, injVal)

    storeVal = p.Get_bool(storeKey, storeVal)
    p.Set_bool(storeKey, storeVal)

    fmt.Printf("parameter %v has value %v\n", memKey, memVal)
    fmt.Printf("parameter %v has value %v\n", normKey, normVal)
    fmt.Printf("parameter %v has value %v\n", injKey, injVal)
    fmt.Printf("parameter %v has value %v\n", storeKey, storeVal)

    //
    // blacklisted documents interface not supported
    //
    //std::string blackList = parameters.get("blacklist", "");
    //if( blackList.length() ) {
    //    int count = env.setBlackList(blackList);
    //    std::cout << "Added to blacklist: "<< count << std::endl;
    //    std::cout.flush();
    //}

    //
    // setOffsetAnnotationIndexHint interface not supported
    //
    //std::string offsetAnnotationHint=parameters.get("offsetannotationhint", "default");
    //if (offsetAnnotationHint=="ordered") {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintOrderedAnnotations);
    //} if (offsetAnnotationHint=="unordered") {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintSizeBuffers);
    //} else {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintDefault);
    //}

    var stemmerKey, stemmerName = "stemmer.name", "Krovetz"

    stemmerName = p.Get_string(stemmerKey, stemmerName)
    p.Set_string(stemmerKey, stemmerName)
    err = env.SetStemmer(stemmerName);
    if err != nil {
        fmt.Printf("env.SetStemmer error %v\n", err)
        return
    }

    fmt.Printf("parameter %v has value %v\n", stemmerKey, stemmerName)

    //
    // Parameters type interface not supported
    // instead of building the Parameters object in the GO environment,
    // we use C++ side APIs that accept sting vectors to build and set
    // various paramater values as in the case of stopwords, metadata,
    // forward metadata, backword metadata, and field names.
    //
    //std::vector<std::string> stopwords;
    //if( copy_parameters_to_string_vector( stopwords, parameters, "stopper.word" ) )
    //  env.setStopwords(stopwords);

    var stopperKey, stopperValue = "stopper", ""
    p.Set_string(stopperKey, stopperValue)

    var stopwords []string = []string {
        "a",
        "an",
        "the",
        "as",
    }
    var stopwordslength int64 = int64(len(stopwords))
    var stopwordsvector StringVector = NewStringVector(stopwordslength)
    defer DeleteStringVector(stopwordsvector)
    if stopwordsvector.Size() >= stopwordslength {
        for k, v := range stopwords {
            stopwordsvector.Set(k,strings.ToLower(v))
        }
        err = env.SetStopwords(stopwordsvector)
        if err != nil {
            fmt.Printf("env.SetStopwords error %v\n", err)
            return
        }
    } else {
        err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", stopwordslength, stopwordsvector.Size())
        return
    }
    fmt.Printf("stopwords has length %v\n", stopwordslength)
    fmt.Printf("stopwordsvector has size %v cap %v\n", stopwordsvector.Size(), stopwordsvector.Capacity())

    var metadataKey, metadataValue = "metadata", ""
    p.Set_string(metadataKey, metadataValue)

    var metadatalist []string = []string {
        "odmver",
        "schver",
        "kind",
        "basetime",
        "maxareas",
        "maxcats",
        "offset",
        "app",
        "docno",
        "docver",
    }
    var metadatalistlength int64 = int64(len(metadatalist))
    var metafieldvector StringVector = NewStringVector(metadatalistlength)
    var forwardvector StringVector = NewStringVector(metadatalistlength)
    var backwardvector StringVector = NewStringVector(metadatalistlength)
    defer DeleteStringVector(metafieldvector)
    defer DeleteStringVector(forwardvector)
    defer DeleteStringVector(backwardvector)
    if metafieldvector.Size() >= metadatalistlength {
        for k, v := range metadatalist {
            metafieldvector.Set(k,strings.ToLower(v))
        }
    } else {
        err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, metafieldvector.Size())
        return
    }
    if forwardvector.Size() >= metadatalistlength {
        for k, v := range metadatalist {
            forwardvector.Set(k,strings.ToLower(v))
        }
    } else {
        err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, forwardvector.Size())
        return
    }
    if backwardvector.Size() >= metadatalistlength {
        for k, v := range metadatalist {
            backwardvector.Set(k,strings.ToLower(v))
        }
    } else {
        err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, backwardvector.Size())
        return
    }
    // we always add all the metadata fields to the forward and backword
    // metadata lists. we also include the special docno metadata field
    // in both lists. i.e. all three metadata lists have the same entries.
    err = env.SetMetadataIndexedFields(forwardvector, backwardvector)
    if err != nil {
        fmt.Printf("env.SetMetadataIndexedFields error %v\n", err)
        return
    }
    fmt.Printf("metadatalist has length %v\n", metadatalistlength)
    fmt.Printf("metafieldvector has size %v cap %v\n", metafieldvector.Size(), metafieldvector.Capacity())
    fmt.Printf("forwardvector has size %v cap %v\n", forwardvector.Size(), forwardvector.Capacity())
    fmt.Printf("backwardvector has size %v cap %v\n", backwardvector.Size(), backwardvector.Capacity())

    // index field attributes
    type indexField struct {
        name string         // name of field
        numeric bool        // true if field is numeric
        // for future use
        //parserName string   // "NumericFieldAnnotator" for numeric field
        //unitenc string      // for numeric field only
        //unitres int         // for numeric field only
        //unitmax int         // for numeric field only
    }
    var numfieldcount int = 0
    var fieldlist []indexField = []indexField {
        {name: "blog", numeric: false},
        {name: "about", numeric: false},
        {name: "address", numeric: false},
        {name: "affiliation", numeric: false},
        {name: "author", numeric: false},
        {name: "brand", numeric: false},
        {name: "citation", numeric: false},
        {name: "description", numeric: false},
        {name: "email", numeric: false},
        {name: "headline", numeric: false},
        {name: "keywords", numeric: false},
        {name: "language", numeric: false},
        {name: "name", numeric: false},
        {name: "telephone", numeric: false},
        {name: "version", numeric: true},
    }
    var fieldlistlength int64 = int64(len(fieldlist))
    var fieldlistvector StringVector = NewStringVector(fieldlistlength)
    defer DeleteStringVector(fieldlistvector)
    if fieldlistvector.Size() >= fieldlistlength {
        for k, v := range fieldlist {
            fieldlistvector.Set(k,strings.ToLower(v.name))
        }
        err = env.SetIndexedFields(fieldlistvector)
        if err != nil {
            fmt.Printf("env.SetIndexedFields error %v\n", err)
            return
        }
        for _, v := range fieldlist {
            if v.numeric {
                // TODO: lookup parser when added to field attribute
                err = env.SetNumericField(strings.ToLower(v.name), v.numeric, "NumericFieldAnnotator")
                if err != nil {
                    fmt.Printf("env.SetNumericField error %v\n", err)
                    return
                }
                numfieldcount += 1
            }
        }
    } else {
        err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", fieldlistlength, fieldlistvector.Size())
        return
    }
    fmt.Printf("fieldlist has length %v with %v marked numeric \n", fieldlistlength, numfieldcount)
    fmt.Printf("fieldlistvector has size %v cap %v\n", fieldlistvector.Size(), fieldlistvector.Capacity())
    //
    // see how we set numeric fields above
    // for now, we do not support setting ordinal and parental fields
    //
    //process_numeric_fields( parameters, env );
    //process_ordinal_fields( parameters, env );
    //process_parental_fields( parameters, env ); //pto
    //

    // from: Repository.cpp
    //bool indri::collection::Repository::exists( const std::string& path ) {
    //    std::string manifestPath = indri::file::Path::combine( path, "manifest" );
    //    return indri::file::Path::exists( manifestPath );
    //}
    //if( indri::collection::Repository::exists( repositoryPath ) ) {
    if util.FileExists(filepath.Join(repositoryPath, "manifest")) {
        // check if the repository was corrupted by an indexing crash
        // if so, recover it and continue.
        //if (_recoverRepository(repositoryPath)) {
        var recovered bool
        recovered, err = Wrapped_Buildindex_recoverRepository(repositoryPath)
        if err != nil {
            return
        }
        if recovered {
            err = env.Open( repositoryPath, monitor );
            if err != nil {
                fmt.Printf("env.Open error %v\n", err)
                return
            }
            Buildindex_print_event( "Opened repository " + repositoryPath );
        } else  {
            //  failed to open it, needs to be created from scratch.
            // create will remove any cruft.
            err = env.Create( repositoryPath, monitor );
            if err != nil {
                fmt.Printf("env.Create error %v\n", err)
                return
            }
            Buildindex_print_event( "Created repository " + repositoryPath );
        }
    } else {
      err = env.Create( repositoryPath, monitor );
      if err != nil {
          fmt.Printf("env.Create error %v\n", err)
          return
      }
      Buildindex_print_event( "Created repository " + repositoryPath );
    }

    // augment field/metadata tags in the environment if needed.
    if( len(corpusFileClass) > 0 ) {
      var spec Indri_parse_FileClassEnvironmentFactory_Specification
      spec, err = env.GetFileClassSpec(corpusFileClass)
      if err != nil {
          fmt.Printf("env.GetFileClassSpec error %v\n", err)
          return
      }
      defer Wrapped_deleteFileClassSpec(spec)
      if( spec != nil ) {
        // add fields if necessary, only update if changed.
        var fieldChanged bool

        fieldChanged, err = Wrapped_Buildindex_augmentSpec( spec, fieldlistvector, metafieldvector, forwardvector, backwardvector )
        if err != nil {
            return
        }
        if( fieldChanged ) {
          err = env.AddFileClass(spec)
          if err != nil {
              fmt.Printf("env.AddFileClass error %v\n", err)
              return
          }
        }
      }
    }

    // First record the document root, and then the paths to any annotator inputs
    err = env.SetDocumentRoot( corpusPath );
    if err != nil {
        fmt.Printf("env.SetDocumentRoot error %v\n", err)
        return
    }

    var /* anchorTextKey, */ anchorTextPath = /* "corpus.inlink", */ ""
    var /* offsetAnnotationsPathKey, */ offsetAnnotationsPath = /* "corpus.annotations", */ ""

    // Support for anchor text
    if len(anchorTextPath) > 0 {
        err = env.SetAnchorTextPath( anchorTextPath );
        if err != nil {
            fmt.Printf("env.SetAnchorTextPath error %v\n", err)
            return
        }
    }

    // Support for offset annotations
    if len(offsetAnnotationsPath) > 0 {
        err = env.SetOffsetAnnotationsPath( offsetAnnotationsPath );
        if err != nil {
            fmt.Printf("env.SetOffsetAnnotationsPath error %v\n", err)
            return
        }
    }

    if len(corpMetadataPath) > 0 {
        err = env.SetOffsetMetadataPath( corpMetadataPath );
        if err != nil {
            fmt.Printf("env.SetOffsetMetadataPath error %v\n", err)
            return
        }
    }

    // if the corpus directory exists...
    // add file in the corpus path to the repository.
    // we assume only one corpus parameter section, and
    // that a corpusFileClass is specified via parameters.
    // this means a single corpusFileClass per repository, and
    // not corpusFileClass lookup based on file extension.
    err = filepath.Walk(corpusPath, func(path string, info os.FileInfo, err error) error {
        if (err != nil && os.IsNotExist(err)) || (info != nil && info.IsDir()) {
            return nil // skip
        }
        if err != nil {
            return err
        }
        err = env.AddFile( path, corpusFileClass )
        fmt.Printf("env.AddFile error %v\n", err)
        return err
    })
    if err != nil {
        return
    }

    // now that repo is ready, test adding a document to the repository
    fp, err := filepath.Abs("data/blog.html")
    if err != nil {
        return
    }
    err = env.AddFile( fp, corpusFileClass )
    if err != nil {
        fmt.Printf("env.AddFile error %v\n", err)
        return
    }

    fp, err = filepath.Abs("data/blog.xml")
    if err != nil {
        return
    }
    err = env.AddFile( fp )
    if err != nil {
        fmt.Printf("env.AddFile error %v\n", err)
        return
    }

    Buildindex_print_event( "Closing index" )
    err = env.Close()
    Buildindex_print_event( "Finished" )

    return
}

func testBuildIndexGOWithParamsFile() (err error) {

    defer catch(&err)

    var p Parameters

    // create index parameters object
    p = NewParameters()
    defer DeleteWrapped_Parameters(p)

    // load paramaters from file into a string
    pars := "data/params.xml"
    s, err := readParamsFile(pars)
    if err != nil {
        return
    }
    //fmt.Printf("params file content: %v\n", s)

    // parse and load parameters from string
    err = p.MyLoad(s)
    if err != nil {
        return
    }

    // controls params file configuration override - for most paramaters
    var overridePathConfig bool = false
    var overrideFileClassConfig bool = false
    var overrideMiscConfig bool = false
    var overrideStemmerConfig bool = false
    var overrideStopWordsConfig bool = false
    var overrideMetadataConfig bool = false
    var overrideIndexFieldConfig bool = false

    //
    // indri applications accept command line options as space separated strings.
    // each command line option string specifies a paramater value, or a set of
    // parameter values indirectly specified as a path to an xml encoded file.
    // the top level element in the xml file is always named <em>parameters</em>.
    //
    // if an option string command line argument (i.e. non-first command line
    // argument) starts with the dash character, i.e. '-', what follows is a
    // parameter option name supported by the indri application. otherwise, it
    // is assumed to be a path string pointing to the xml file.
    //
    // here, we are not using argc/argv, but building index paramaters the
    // hard way. another approach wraps IndriBuildIndex C++ code to have GO
    // code delegate to the C++ code to build index, as this is what it is
    // already coded to do.
    //
    //p.loadCommandLine( argc, argv );

    // we need to create the repository root and copus folders for index use
    dir, err := ioutil.TempDir("", "test-repo-root")
	if err != nil {
        err = fmt.Errorf("failed to create TempDir: %v", err)
	}
    defer os.RemoveAll(dir) // clean up after test completes

    // define defaults for parameter configuration
    var repoKey, repositoryPath = "index", filepath.Join(dir, "index-1")
    var corpPathKey, corpusPath = "corpus.path", filepath.Join(dir, "corpus-1")
    var corpusFileClassKey, corpusFileClass = "corpus.class", "html"
    var corpMetadataKey, corpMetadataPath = "corpus.metadata", filepath.Join(repositoryPath, "metadata")

    // index builder will create the index folder, create corpus folder here
    err = os.Mkdir(corpusPath, 0775)
    if err != nil {
        err = fmt.Errorf("failed to create corpus dir: %v", err)
    }

    // override configuration loaded from the params file
    if overridePathConfig {
        p.Set_string(repoKey, repositoryPath)
        p.Set_string(corpPathKey, corpusPath)
        p.Set_string(corpMetadataKey, corpMetadataPath)
    }
    if overrideFileClassConfig {
        p.Set_string(corpusFileClassKey, corpusFileClass)
    }

    // check for required parameters
    if !requireParameter( "corpus", p ) {
        err = fmt.Errorf("parameter %v is required", "corpus")
        return
    }
    if !requireParameter( "index", p ) {
        err = fmt.Errorf("parameter %v is required", repoKey)
        return
    }
    if !requireParameter( "corpus.path", p ) {
        err = fmt.Errorf("parameter %v is required", corpPathKey)
        return
    }

    var monitor MyStatusMonitor = NewMyStatusMonitor()
    defer DeleteMyStatusMonitor(monitor)
    var env IndexEnvironment = NewIndexEnvironment()
    defer DeleteWrapped_IndexEnvironment(env)

    var v string

    Buildindex_start_time();

    fmt.Printf("indri distribution version: %v\n", GetIndriVersion())

    v = p.Get_string(repoKey, "")
    if len(v) < 1 {
        err = fmt.Errorf("parameter %v is required", repoKey)
        return
    } else {
        fmt.Printf("parameter %v has value %v\n", repoKey, v)
    }

    v = p.Get_string(corpPathKey, "")
    if len(v) < 1 {
        err = fmt.Errorf("parameter %v is required", corpPathKey)
        return
    } else {
        fmt.Printf("parameter %v has value %v\n", corpPathKey, v)
    }

    v = p.Get_string(corpMetadataKey, "")
    fmt.Printf("parameter %v has value %v\n", corpMetadataKey, v)

    // define defaults for parameter configuration
    var memKey, memVal = "memory", int64(1024*1024*1024)
    var normKey, normVal = "normalize", true
    var injKey, injVal = "injectURL", true
    var storeKey, storeVal = "storeDocs", true

    // read specified parameter configuration or default
    memVal = p.Get_INT64(memKey, memVal)
    normVal = p.Get_bool(normKey, normVal)
    injVal = p.Get_bool(injKey, injVal)
    storeVal = p.Get_bool(storeKey, storeVal)

    // override configuration loaded from the params file
    if overrideMiscConfig {
        p.Set_UINT64(memKey, memVal)
        p.Set_bool(normKey, normVal)
        p.Set_bool(injKey, injVal)
        p.Set_bool(storeKey, storeVal)
    }

    fmt.Printf("parameter %v has value %v\n", memKey, memVal)
    fmt.Printf("parameter %v has value %v\n", normKey, normVal)
    fmt.Printf("parameter %v has value %v\n", injKey, injVal)
    fmt.Printf("parameter %v has value %v\n", storeKey, storeVal)

    //
    // blacklisted documents interface not supported
    //
    //std::string blackList = parameters.get("blacklist", "");
    //if( blackList.length() ) {
    //    int count = env.setBlackList(blackList);
    //    std::cout << "Added to blacklist: "<< count << std::endl;
    //    std::cout.flush();
    //}

    //
    // setOffsetAnnotationIndexHint interface not supported
    //
    //std::string offsetAnnotationHint=parameters.get("offsetannotationhint", "default");
    //if (offsetAnnotationHint=="ordered") {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintOrderedAnnotations);
    //} if (offsetAnnotationHint=="unordered") {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintSizeBuffers);
    //} else {
    //  env.setOffsetAnnotationIndexHint(indri::parse::OAHintDefault);
    //}

    // define defaults for parameter configuration
    var stemmerKey, stemmerName = "stemmer.name", "Krovetz"

    // read specified parameter configuration or default
    stemmerName = p.Get_string(stemmerKey, stemmerName)

    // override configuration loaded from the params file
    if overrideStemmerConfig {
        p.Set_string(stemmerKey, stemmerName)
    }

    fmt.Printf("parameter %v has value %v\n", stemmerKey, stemmerName)

    // inform index environment which stemmer to apply
    err = env.SetStemmer(stemmerName);
    if err != nil {
        fmt.Printf("env.SetStemmer error %v\n", err)
        return
    }

    //
    // Parameters type interface not fully supported
    // instead of building the Parameters object in the GO environment,
    // we use C++ side APIs that accept sting vectors to build and set
    // various paramater values as in the case of stopwords, metadata,
    // forward metadata, backword metadata, and field names.
    //
    // To have functional parity with InderiBuildIndex.cpp, we need to
    // support reading the stop words from a file in addition to loading
    // from params file and the hard coded override here
    //
    //std::vector<std::string> stopwords;
    //if( copy_parameters_to_string_vector( stopwords, parameters, "stopper.word" ) )
    //  env.setStopwords(stopwords);

    // define defaults for parameter configuration
    var stopperKey, stopperValue = "stopper", ""
    // hard coded stop words to override
    var stopwords []string = []string {
        "a",
        "an",
        "the",
        "as",
    }
    var stopwordslength int64 = int64(len(stopwords))
    var stopwordsvector StringVector = NewStringVector(stopwordslength)
    defer DeleteStringVector(stopwordsvector)

    // override configuration loaded from the params file
    if overrideStopWordsConfig {

        // make sure path exists in parameters
        p.Set_string(stopperKey, stopperValue)

        if stopwordsvector.Size() >= stopwordslength {
            // override configuration loaded from the params file
            for k, v := range stopwords {
                stopwordsvector.Set(k,strings.ToLower(v))
            }
            // inform index environment which stop words to apply
            err = env.SetStopwords(stopwordsvector)
            if err != nil {
                fmt.Printf("env.SetStopwords error %v\n", err)
                return
            }
        } else {
            err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", stopwordslength, stopwordsvector.Size())
            return
        }
        fmt.Printf("stopwords has length %v\n", stopwordslength)
        fmt.Printf("stopwordsvector has size %v cap %v\n", stopwordsvector.Size(), stopwordsvector.Capacity())

    }

    // define defaults for parameter configuration
    var metadataKey, metadataValue = "metadata", ""
    // hard coded metadata to override
    var metadatalist []string = []string {
        "odmver",
        "schver",
        "kind",
        "basetime",
        "maxareas",
        "maxcats",
        "offset",
        "app",
        "docno",
        "docver",
    }
    var metadatalistlength int64 = int64(len(metadatalist))
    var metafieldvector StringVector = NewStringVector(metadatalistlength)
    var forwardvector StringVector = NewStringVector(metadatalistlength)
    var backwardvector StringVector = NewStringVector(metadatalistlength)
    defer DeleteStringVector(metafieldvector)
    defer DeleteStringVector(forwardvector)
    defer DeleteStringVector(backwardvector)

    // override configuration loaded from the params file
    if overrideMetadataConfig {

        // make sure path exists in parameters
        p.Set_string(metadataKey, metadataValue)

        if metafieldvector.Size() >= metadatalistlength {
            // override configuration loaded from the params file
            for k, v := range metadatalist {
                metafieldvector.Set(k,strings.ToLower(v))
            }
        } else {
            err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, metafieldvector.Size())
            return
        }
        if forwardvector.Size() >= metadatalistlength {
            // override configuration loaded from the params file
            for k, v := range metadatalist {
                forwardvector.Set(k,strings.ToLower(v))
            }
        } else {
            err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, forwardvector.Size())
            return
        }
        if backwardvector.Size() >= metadatalistlength {
            // override configuration loaded from the params file
            for k, v := range metadatalist {
                backwardvector.Set(k,strings.ToLower(v))
            }
        } else {
            err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", metadatalistlength, backwardvector.Size())
            return
        }

        //
        // we always add all the metadata fields to the forward and backword
        // metadata lists. we also include the special docno metadata field
        // in both lists. i.e. all three metadata lists have the same entries.
        //
        // inform index environment which metadata fields to apply
        err = env.SetMetadataIndexedFields(forwardvector, backwardvector)
        if err != nil {
            fmt.Printf("env.SetMetadataIndexedFields error %v\n", err)
            return
        }
        fmt.Printf("metadatalist has length %v\n", metadatalistlength)
        fmt.Printf("metafieldvector has size %v cap %v\n", metafieldvector.Size(), metafieldvector.Capacity())
        fmt.Printf("forwardvector has size %v cap %v\n", forwardvector.Size(), forwardvector.Capacity())
        fmt.Printf("backwardvector has size %v cap %v\n", backwardvector.Size(), backwardvector.Capacity())

    }

    // index field attributes
    type indexField struct {
        name string         // name of field
        numeric bool        // true if field is numeric
        // for future use
        //parserName string   // "NumericFieldAnnotator" for numeric field
        //unitenc string      // for numeric field only
        //unitres int         // for numeric field only
        //unitmax int         // for numeric field only
    }
    var numfieldcount int = 0

    // hard coded metadata to override
    var fieldlist []indexField = []indexField {
        {name: "blog", numeric: false},
        {name: "about", numeric: false},
        {name: "address", numeric: false},
        {name: "affiliation", numeric: false},
        {name: "author", numeric: false},
        {name: "brand", numeric: false},
        {name: "citation", numeric: false},
        {name: "description", numeric: false},
        {name: "email", numeric: false},
        {name: "headline", numeric: false},
        {name: "keywords", numeric: false},
        {name: "language", numeric: false},
        {name: "name", numeric: false},
        {name: "telephone", numeric: false},
        {name: "version", numeric: true},
    }

    var fieldlistlength int64 = int64(len(fieldlist))
    var fieldlistvector StringVector = NewStringVector(fieldlistlength)
    defer DeleteStringVector(fieldlistvector)

    // override configuration loaded from the params file
    if overrideIndexFieldConfig {

        if fieldlistvector.Size() >= fieldlistlength {
            // override configuration loaded from the params file
            for k, v := range fieldlist {
                fieldlistvector.Set(k,strings.ToLower(v.name))
            }
            // inform index environment which index fields to apply
            err = env.SetIndexedFields(fieldlistvector)
            if err != nil {
                fmt.Printf("env.SetIndexedFields error %v\n", err)
                return
            }
            // process field special attributes
            for _, v := range fieldlist {
                // process numeric fields
                if v.numeric {
                    // TODO: lookup parser when added to field attribute
                    err = env.SetNumericField(strings.ToLower(v.name), v.numeric, "NumericFieldAnnotator")
                    if err != nil {
                        fmt.Printf("env.SetNumericField error %v\n", err)
                        return
                    }
                    numfieldcount += 1
                }
                //
                // for now, we do not support setting ordinal and parental fields
                // see how we set numeric fields above
                //
                //process_numeric_fields( parameters, env );
                //process_ordinal_fields( parameters, env );
                //process_parental_fields( parameters, env ); //pto
                //
            }
        } else {
            err = fmt.Errorf("failed to allocate string vector, requested size %v actual size %v", fieldlistlength, fieldlistvector.Size())
            return
        }
        fmt.Printf("fieldlist has length %v with %v marked numeric \n", fieldlistlength, numfieldcount)
        fmt.Printf("fieldlistvector has size %v cap %v\n", fieldlistvector.Size(), fieldlistvector.Capacity())

    }

    //
    // index envirenment configuration is complete.
    //
    // create the repository or recover it in case the indexer had crashed.
    //

    //
    // this duplicates a function from: Repository.cpp
    //bool indri::collection::Repository::exists( const std::string& path ) {
    //    std::string manifestPath = indri::file::Path::combine( path, "manifest" );
    //    return indri::file::Path::exists( manifestPath );
    //}
    //if( indri::collection::Repository::exists( repositoryPath ) ) {
    if util.FileExists(filepath.Join(repositoryPath, "manifest")) {
        // check if the repository was corrupted by an indexing crash
        // if so, recover it and continue.
        //if (_recoverRepository(repositoryPath)) {
        var recovered bool
        recovered, err = Wrapped_Buildindex_recoverRepository(repositoryPath)
        if err != nil {
            return
        }
        if recovered {
            err = env.Open( repositoryPath, monitor );
            if err != nil {
                fmt.Printf("env.Open error %v\n", err)
                return
            }
            Buildindex_print_event( "Opened repository " + repositoryPath );
        } else  {
            //  failed to open it, needs to be created from scratch.
            // create will remove any cruft.
            err = env.Create( repositoryPath, monitor );
            if err != nil {
                fmt.Printf("env.Create error %v\n", err)
                return
            }
            Buildindex_print_event( "Created repository " + repositoryPath );
        }
    } else {
      err = env.Create( repositoryPath, monitor );
      if err != nil {
          fmt.Printf("env.Create error %v\n", err)
          return
      }
      Buildindex_print_event( "Created repository " + repositoryPath );
    }

    // augment field/metadata tags in the environment if needed.
    if( len(corpusFileClass) > 0 ) {
      var spec Indri_parse_FileClassEnvironmentFactory_Specification
      spec, err = env.GetFileClassSpec(corpusFileClass)
      if err != nil {
          fmt.Printf("env.GetFileClassSpec error %v\n", err)
          return
      }
      defer Wrapped_deleteFileClassSpec(spec)
      if( spec != nil ) {
        // add fields if necessary, only update if changed.
        var fieldChanged bool

        fieldChanged, err = Wrapped_Buildindex_augmentSpec( spec, fieldlistvector, metafieldvector, forwardvector, backwardvector )
        if err != nil {
            return
        }
        if( fieldChanged ) {
          err = env.AddFileClass(spec)
          if err != nil {
              fmt.Printf("env.AddFileClass error %v\n", err)
              return
          }
        }
      }
    }

    //
    // the index has been created or recovered.
    //
    // in the case or recovery, corpus documents must be reloaded into the index
    //

    // First record the document root, and then the paths to any annotator inputs
    err = env.SetDocumentRoot( corpusPath );
    if err != nil {
        fmt.Printf("env.SetDocumentRoot error %v\n", err)
        return
    }

    var /* anchorTextKey, */ anchorTextPath = /* "corpus.inlink", */ ""
    var /* offsetAnnotationsPathKey, */ offsetAnnotationsPath = /* "corpus.annotations", */ ""

    // Support for anchor text
    if len(anchorTextPath) > 0 {
        err = env.SetAnchorTextPath( anchorTextPath );
        if err != nil {
            fmt.Printf("env.SetAnchorTextPath error %v\n", err)
            return
        }
    }

    // Support for offset annotations
    if len(offsetAnnotationsPath) > 0 {
        err = env.SetOffsetAnnotationsPath( offsetAnnotationsPath );
        if err != nil {
            fmt.Printf("env.SetOffsetAnnotationsPath error %v\n", err)
            return
        }
    }

    if len(corpMetadataPath) > 0 {
        err = env.SetOffsetMetadataPath( corpMetadataPath );
        if err != nil {
            fmt.Printf("env.SetOffsetMetadataPath error %v\n", err)
            return
        }
    }

    // if the corpus directory exists...
    // add file in the corpus path to the repository.
    // we assume only one corpus parameter section, and
    // that a corpusFileClass is specified via parameters.
    // this means a single corpusFileClass per repository, and
    // no corpusFileClass lookup based on file extension.
    err = filepath.Walk(corpusPath, func(path string, info os.FileInfo, err error) error {
        if (err != nil && os.IsNotExist(err)) || (info != nil && info.IsDir()) {
            return nil // skip
        }
        if err != nil {
            return err
        }
        err = env.AddFile( path, corpusFileClass )
        fmt.Printf("env.AddFile error %v\n", err)
        return err
    })
    if err != nil {
        return
    }

    err = testDocumentAddDeleteMetadata(env, corpusFileClass)

    // all done, flush and close the index repository
    Buildindex_print_event( "Closing index" )
    err = env.Close()
    Buildindex_print_event( "Finished" )

    return
}

func testDocumentAddDeleteMetadata(env IndexEnvironment, corpusFileClass string) (err error) {

    defer catch(&err)

    //
    // now that repo is ready to use.
    //
    // test adding documents to the repository
    //

    // add document as file and explicitly specify corpusFileClass
    fp, err := filepath.Abs("data/blog.html")
    if err != nil {
        return
    }
    err = env.AddFile( fp, corpusFileClass )
    if err != nil {
        fmt.Printf("env.AddFile error %v\n", err)
        return
    }

    // add document as file and indirectly specify corpusFileClass using file extension
    fp, err = filepath.Abs("data/blog.xml")
    if err != nil {
        return
    }
    err = env.AddFile( fp )
    if err != nil {
        fmt.Printf("env.AddFile error %v\n", err)
        return
    }

    // add document as string and explicitly specify corpusFileClass
    var documentAsString string = `
<odmver>test odm version 0.1</odmver>
<schver>test sch version 0.1</schver>
<kind>blogtest</kind>
<basetime>UTC time</basetime>
<maxareas>64</maxareas>
<maxcats>64</maxcats>
<offset>0</offset>
<app>dms3</app>
<docno>001</docno>
<docver>1.0</docver>
  <blog>
      <about>this blog is abot food</about>
      <address>72 middlesex turnpike, Burlington, MA</address>
      <affiliation>Burlington Mall</affiliation>
      <author>John Doe</author>
      <brand>Malls R US</brand>
      <citation>tbd</citation>
      <description>the food court at burlington mall</description>
      <email>contact.us.@themall.com</email>
      <headline>type of food you can eat here</headline>
      <keywords>pizza, chinese, fastfood</keywords>
      <language>English</language>
      <name>Burlington Mall</name>
      <telephone>nnn-nnn-nnnn</telephone>
      <version>1</version>
  </blog>`

    var p1, p2 MetadataPair

    p1, err = NewMetadataPair()
    if err != nil {
        err = fmt.Errorf("failed to allocate metadata pair error %v\n", err)
        return
    }

    p2, err = NewMetadataPair()
    if err != nil || p2 == nil {
        err = fmt.Errorf("failed to allocate metadata pair error %v\n", err)
        return
    }

    defer DeleteMetadataPair(p1)
    defer DeleteMetadataPair(p2)

    var pk1 string = "docno"
    var pv1 []byte = []byte("00111111111111111111111111111111111111111111111111111111111")
    p1.WSetKey(pk1)
    p1.WSetValue(pv1)

    var pk2 string = "kind"
    var pv2 []byte = []byte("blog1")
    p2.WSetKey(pk2)
    p2.WSetValue(pv2)

    pairVector := NewWrapped_MetadataPairVector(2)
    defer DeleteWrapped_MetadataPairVector(pairVector)

    err = pairVector.WSet(0, p1)
    //err = pairVector.WAdd(p1) - this would leave reserve entry uninitialized
    if err != nil {
        err = fmt.Errorf("failed to add metadata pair 1 error %v\n", err)
        return
    }

    err = pairVector.WSet(1, p2)
    //err = pairVector.WAdd(p2) - this would leave reserve entry uninitialized
    if err != nil {
        err = fmt.Errorf("failed to add metadata pair 2 error %v\n", err)
        return
    }

    k1, err := p1.WGetKey()
    if err != nil {
        err = fmt.Errorf("failed to read back pair 1 key error %v\n", err)
        return
    }
    if k1 != pk1 {
        err = fmt.Errorf("failed to match written pair 1 key %v\n", k1)
        return
    }

    v1, err := p1.WGetValue()
    if err != nil {
        err = fmt.Errorf("failed to read back pair 1 value error %v\n", err)
        return
    }
    if string(v1) != string(pv1) {
        err = fmt.Errorf("failed to match written pair 1 value %v\n", v1)
        return
    }

    k2, err := p2.WGetKey()
    if err != nil {
        err = fmt.Errorf("failed to read back pair 2 key error %v\n", err)
        return
    }
    if k2 != pk2 {
        err = fmt.Errorf("failed to match written pair 2 key %v\n", k2)
        return
    }

    v2, err := p2.WGetValue()
    if err != nil {
        err = fmt.Errorf("failed to read back pair 2 value error %v\n", err)
        return
    }
    if string(v2) != string(pv2) {
        err = fmt.Errorf("failed to match written pair 2 value %v\n", v2)
        return
    }

    //
    // verify we can read back metadata
    //
    // there is a case where we allocate a pairVector with reservation, but
    // leave reserved vector entries un-initialized [by call WAdd(pair) to add
    // new pairs instead of WSet(index, pair)]. With such a programming error,
    // we will fail with exception in C function on the pair.WGetValue() call.
    //
    // example incorrect programming sequence:
    //      pairVector := NewWrapped_MetadataPairVector(2)
    //      defer DeleteWrapped_MetadataPairVector(pairVector)
    //      //err = pairVector.WSet(0, p1)
    //      err = pairVector.WAdd(p1)
    //      //err = pairVector.WSet(1, p2)
    //      err = pairVector.WAdd(p2)
    // ends up with a pairVector.Size() == 4 and not 2
    //
    // however, the "catch" panic handler does not trap the exception and
    // the cgo crosscall results in a segmentation violation [ most likely due
    // the pair's value, both void * and value length, are in random state ]
    //
    // TODO: figure out how we may trap all C exceptions.
    // there seems to be path cases where the execution thread context (thread
    // stack and frame pointer) is switched on syscalls (to avoid GC and other
    // GO thread from blocking. And on return, the completion may run on a
    // different thread. This panic handler may be on a different stack???
    // review comments in go/src/runtime/cgocall.go
    //
    if pairVector.Size() != 2 || pairVector.Capacity() < 2 || pairVector.IsEmpty() {
        err = fmt.Errorf("unexpected pairVector state size %v capacity %v empty? %v, expected {2, >= 2, false}", pairVector.Size(), pairVector.Capacity(), pairVector.IsEmpty())
        return
    }

    var rp1, rp2 Wrapped_MetadataPair

    rp1, err = pairVector.WGet(0)
    if err != nil || rp1 == nil {
        err = fmt.Errorf("failed to read back pairVector index %v error %v\n", 0, err)
        return
    }

    rp2, err = pairVector.WGet(1)
    if err != nil || rp2 == nil{
        err = fmt.Errorf("failed to read back pairVector index %v error %v\n", 1, err)
        return
    }

    vs := int(pairVector.Size())
    var pair MetadataPair
    var rpk string
    var rpv []byte
    for i := 0; i < vs; i++ {
        pair, err = pairVector.WGet(i)
        if err != nil {
            err = fmt.Errorf("failed to read back pairVector index %v error %v\n", i, err)
            return
        }
        rpk, err = pair.WGetKey()
        if err != nil {
            err = fmt.Errorf("failed to read back pair key for vector index %v error %v\n", i, err)
            return
        }
        rpv, err = pair.WGetValue()
        if err != nil {
            err = fmt.Errorf("failed to read back pair value for vector index %v error %v\n", i, err)
            return
        }
        switch i {
        case 0:
            if rpk != pk1 {
                err = fmt.Errorf("failed to match written pair 1 key %v\n", rpk)
                return
            }
            if string(rpv) != string(pv1) {
                err = fmt.Errorf("failed to match written pair 1 value %v\n", rpv)
                return
            }
        case 1:
            if rpk != pk2 {
                err = fmt.Errorf("failed to match written pair 2 key %v\n", rpk)
                return
            }
            if string(rpv) != string(pv2) {
                err = fmt.Errorf("failed to match written pair 2 value %v\n", rpv)
                return
            }
        default:
            fmt.Printf("unexpected metadata pair index %v\n", i)
        }
    }

    var docId int

    docId, err = env.AddString(documentAsString, "html", pairVector)
    if err != nil {
        err = fmt.Errorf("env.AddString error %v\n", err)
        return
    }
    fmt.Printf("env.AddString docId %v\n", docId)

    //
    // TODO: read document given docId, verify metadata and content.
    // document lifecycle management is non-trivial and is on deck
    // for development.
    //

    //
    // test deleting documents from the repository
    //
    err = env.DeleteDocument(docId)
    if err != nil {
        err = fmt.Errorf("env.DeleteDocument for %v failed with error %v\n", docId, err)
        return
    }

    //
    // verify document counts in repository
    //
    var di, dp int
    var edi, edp int = 3, 3

    di, err = env.DocumentsIndexed()
    if err != nil {
        err = fmt.Errorf("env.DocumentsIndexed failed with error %v\n", err)
        return
    }

    dp, err = env.DocumentsSeen()
    if err != nil {
        err = fmt.Errorf("env.DocumentsSeen failed with error %v\n", err)
        return
    }

    if di != edi || dp != edp {
        err = fmt.Errorf("expected %v documents parsed and %v indexed.\n", edp, edi)
        return
    } else {
        fmt.Printf("Documents parsed: %v Documents indexed: %v\n", dp, di)
    }

    return
}
