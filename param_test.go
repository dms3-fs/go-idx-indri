package indri_go

import (
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

var forcepanic, forceerror bool

var (
    ErrNotConfigured = errors.New("params file not configured.")
)

// missing all required params
var bad_params_1 string = `
<parameters>
</parameters>
`

// missing required corpus params
var bad_params_2 string = `
<parameters>
    <index>repo/base</index>
</parameters>
`

// missing required corpus path params
var bad_params_3 string = `
<parameters>
    <index>repo/base</index>
    <corpus>
    </corpus>
</parameters>
`

// malformed xml
var bad_params_4 string = `
<parameters>
    <index>repo/base</index>
    <corpus>
        <path>repo/base/corpus</path>
    </corpus>
`

// malformed xml
var bad_params_5 string = `
<parameters>
    <index>repo/base</index>
    <corpus>
        <path>repo/base/corpus
    </corpus>
</parameters>
`

const good_params_1 string = `
  <parameters>
      <index>repo/base</index>
      <corpus>
          <path>repo/base/corpus</path>
          <class>html</class>
          <metadata>repo/base/meta</metadata>
      </corpus>
      <metadata>

<!-- read-only life-cycle [system] metadata fields -->

          <forward>odmver</forward>
          <forward>schver</forward>
          <forward>kind</forward>
          <forward>basetime</forward>
          <forward>maxareas</forward>
          <forward>maxcats</forward>
          <forward>offset</forward>
          <forward>app</forward>

<!-- read-write life-cycle [document] metadata fields -->

          <forward>docno</forward>
          <forward>docver</forward>

<!-- start of [document] kind common metadata fields -->

          <backward>odmver</backward>
          <backward>schver</backward>
          <backward>kind</backward>
          <backward>basetime</backward>
          <backward>maxareas</backward>
          <backward>maxcats</backward>
          <backward>offset</backward>
          <backward>app</backward>
          <backward>docno</backward>
          <backward>docver</backward>
          <field>
              <name>odmver</name>
          </field>
          <field>
              <name>schver</name>
          </field>
          <field>
              <name>kind</name>
          </field>
          <field>
              <name>basetime</name>
          </field>
          <field>
              <name>maxareas</name>
          </field>
          <field>
              <name>maxcats</name>
          </field>
          <field>
              <name>offset</name>
          </field>
          <field>
              <name>app</name>
          </field>
          <field>
              <name>docno</name>
          </field>
          <field>
              <name>docver</name>
          </field>
      </metadata>

<!-- start of [document] kind specific metadata fields -->

      <field>
          <name>blog</name>
      </field>
      <field>
          <name>about</name>
      </field>
      <field>
          <name>address</name>
      </field>
      <field>
          <name>affiliation</name>
      </field>
      <field>
          <name>author</name>
      </field>
      <field>
          <name>brand</name>
      </field>
      <field>
          <name>citation</name>
      </field>
      <field>
          <name>description</name>
      </field>
      <field>
          <name>email</name>
      </field>
      <field>
          <name>headline</name>
      </field>
      <field>
          <name>keywords</name>
      </field>
      <field>
          <name>language</name>
      </field>
      <field>
          <name>name</name>
      </field>
      <field>
          <name>telephone</name>
      </field>
      <field>
          <name>version</name>
      </field>

<!-- optional index parameters -->

      <memory>100m</memory>
      <stemmer>
          <name>krovetz</name>
      </stemmer>
      <normalize>true</normalize>
      <stopper>
          <word>a</word>
          <word>an</word>
          <word>the</word>
          <word>as</word>
      </stopper>
  </parameters>
`

// readParamsFile returns the params file content as a string
func readParamsFile(paramsFile string) (string, error) {

    paramsFile = filepath.Clean(paramsFile)

    fileInfo, err := os.Stat(paramsFile)
    if err != nil {
        return "", err
    }

    if fileInfo.Size() > 8192 {
        return "", fmt.Errorf("params file too large, must be <=8192 bytes long: %s", paramsFile)
    }

    // if there is no file, assume there is no api addr.
    f, err := os.Open(paramsFile)
    if err != nil {
        if os.IsNotExist(err) {
            return "", ErrNotConfigured
        }
        return "", err
    }
    defer f.Close()

    // read up to 8192 bytes
    buf, err := ioutil.ReadAll(io.LimitReader(f, 8192))
    if err != nil {
        return "", err
    }

    s := string(buf)
    s = strings.TrimSpace(s)
    return s, nil

}

func requireParameter(key string, p Parameters) bool {
    exists := p.Exists(key)
    //fmt.Printf("parameter %s exists? %t\n", key, exists)

    //val := p.Get_string(key, "")
    //fmt.Printf("%s = %v\n", key, val)

    return exists
}


func TestParseParamsGoodString(t *testing.T) {
    err := testParseParamsFromString(good_params_1, true)
    if err != nil {
        t.Fatal(err)
    }
}

func TestParseParamsForcedPanic(t *testing.T) {
    forcepanic, forceerror = true, false
    err := testParseParams()
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsNullPointerOnNew(t *testing.T) {
    forcepanic, forceerror = false, true
    err := testParseParams()
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsStringNil(t *testing.T) {
    var str string = ""
    err := testParseParamsFromString(str, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsBadString1(t *testing.T) {
    err := testParseParamsFromString(bad_params_1, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsBadString2(t *testing.T) {
    err := testParseParamsFromString(bad_params_1, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsBadString3(t *testing.T) {
    err := testParseParamsFromString(bad_params_1, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsBadString4(t *testing.T) {
    err := testParseParamsFromString(bad_params_1, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsBadString5(t *testing.T) {
    err := testParseParamsFromString(bad_params_1, false)
    if err == nil {
        t.Fatal(err)
    }
}

func TestParseParamsGoodFile(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testParseParams()
    if err != nil {
        t.Fatal(err)
    }
}

func TestParseParamsReadWrite(t *testing.T) {
    forcepanic, forceerror = false, false
    err := testParamsReadWrite()
    if err != nil {
        t.Fatal(err)
    }
}

func testParseParamsFromString(s string, good bool) (err error) {

	defer catch(&err)

    var p Parameters

    // create index parameters object
    p = NewParameters()
    defer DeleteWrapped_Parameters(p)

    // parse and load parameters
    err = p.MyLoad(s)
    if err != nil {
        return
    }

    if !good {
        var k, v string
        k = "parameters"
        v = p.Get_string(k, "")
        if len(v) > 0 {
            err = fmt.Errorf("did not expect key %v to contain value %v", k, v)
            return
        }
    }

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

func testParseParams() (err error) {

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

func testParamsReadWrite() (err error) {

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
