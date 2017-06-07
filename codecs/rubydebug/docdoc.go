// Code generated by "bitfanDoc -codec rubydebug"; DO NOT EDIT
package rubydebugcodec

import "github.com/vjeantet/bitfan/processors/doc"

func Doc() *doc.Codec {
	return &doc.Codec{
  Name:       "rubydebug",
  PkgName:    "rubydebugcodec",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/codecs/rubydebug",
  Doc:        "This codec pretty prints event",
  DocShort:   "",
  Decoder:    (*doc.Decoder)(nil),
  Encoder:    &doc.Encoder{
    Doc:     "Prettyprint event",
    Options: &doc.CodecOptions{
      Doc:     "Encode options",
      Options: []*doc.CodecOption{},
    },
  },
}
}