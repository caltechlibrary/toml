package toml

import "bytes"

//
// NOTE: This file provides a Marshal(), MarshalIndent()
// with the same function signature as the json package.
// I did this to app toml to be a drop in replacement (or addition)
// to where I currently support JSON configuration files.
// rsdoiel@library.caltech.edu, 2019-08-01
//

// MarshalIndent provides the TOML equivalent programming interface
// as the json package provides.
//
// if src, err := toml.MarshalIndent(myobject, "", "    "); err != nil {
//    log.Fatal(err)
// } else {
//    fmt.Fprintf(os.Stdout, "%s", src)
// }
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := NewEncoder(buf)
	encoder.Indent = prefix + indent
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Marshal provides the TOML equivalent programming interface
// as the json package provides.
//
// if src, err := toml.Marshal(myobject); err != nil {
//    log.Fatal(err)
// } else {
//    fmt.Fprintf(os.Stdout, "%s", src)
// }
func Marshal(v interface{}) ([]byte, error) {
	return MarshalIndent(v, "", "")
}
