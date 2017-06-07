// Code generated by "bitfanDoc -codec json_lines"; DO NOT EDIT
package jsonlinescodec

import "github.com/vjeantet/bitfan/processors/doc"

func Doc() *doc.Codec {
	return &doc.Codec{
  Name:       "json_lines",
  PkgName:    "jsonlinescodec",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/codecs/jsonlines",
  Doc:        "",
  DocShort:   "",
  Decoder:    &doc.Decoder{
    Doc:     "",
    Options: &doc.CodecOptions{
      Doc:     "",
      Options: []*doc.CodecOption{
        &doc.CodecOption{
          Name:           "Delimiter",
          Alias:          "",
          Doc:            "Change the delimiter that separates lines",
          Required:       false,
          Type:           "string",
          DefaultValue:   "\"\\n\"",
          PossibleValues: []string{},
          ExampleLS:      "",
        },
      },
    },
  },
  Encoder: (*doc.Encoder)(nil),
}
}