package toml

//
// NOTE: This file provides a Marshal(), MarshalIndent(), Unmarshal()
// funcs using the same implementation pattern os the json package.
// this makes it easy to swap out toml and json without changing
// allow of code.
//

// MarshalIndent provides the TOML equivalent programming interface
// as the json package provides.
//
// if src, err := toml.MarshalIndent(myobject, "", "    "); err != nil {
//    log.Fatal(err)
// } else {
//    fmt.Fprintf(os.Stdout, "%s", src)
// }
func MarshalIdent(v interface{}, prefix, indent string) ([]byte, error) {
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

// Unmarshal provides toml with the same programming interface
// as the json package provides.
//
// obj := new(MyObject)
// if err := toml.Marshal(src, &obj); err != nil {
//     log.Fatal(err)
// } else {
//    .... do something with obj.
// }
func Unmarshal(src []byte, v interface{}) error {
	_, err := Decoder(src, v)
	return err
}
