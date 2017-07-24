// Code generated by "bitfanDoc "; DO NOT EDIT
package templateprocessor

import "github.com/vjeantet/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
  Name:       "templateprocessor",
  ImportPath: "/Users/sodadi/go/src/github.com/vjeantet/bitfan/processors/template",
  Doc:        "",
  DocShort:   "",
  Options:    &doc.ProcessorOptions{
    Doc:     "",
    Options: []*doc.ProcessorOption{
      &doc.ProcessorOption{
        Name:           "Add_field",
        Alias:          "",
        Doc:            "If this filter is successful, add any arbitrary fields to this event.",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Tags",
        Alias:          "",
        Doc:            "If this filter is successful, add arbitrary tags to the event. Tags can be dynamic\nand include parts of the event using the %{field} syntax.",
        Required:       false,
        Type:           "array",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Type",
        Alias:          "",
        Doc:            "Add a type field to all events handled by this input",
        Required:       false,
        Type:           "string",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "",
      },
      &doc.ProcessorOption{
        Name:           "Location",
        Alias:          "location",
        Doc:            "Go Template content\n\nset inline content, a path or an url to the template content\n\nGo template : https://golang.org/pkg/html/template/",
        Required:       true,
        Type:           "location",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "location => \"test.tpl\"",
      },
      &doc.ProcessorOption{
        Name:           "Var",
        Alias:          "var",
        Doc:            "You can set variable to be used in template by using ${var}.\neach reference will be replaced by the value of the variable found in Template's content\nThe replacement is case-sensitive.",
        Required:       false,
        Type:           "hash",
        DefaultValue:   nil,
        PossibleValues: []string{},
        ExampleLS:      "var => {\"hostname\"=>\"myhost\",\"varname\"=>\"varvalue\"}",
      },
      &doc.ProcessorOption{
        Name:           "Target",
        Alias:          "target",
        Doc:            "Define the target field for placing the template execution result. If this setting is omitted,\nthe data will be stored in the \"output\" field",
        Required:       false,
        Type:           "string",
        DefaultValue:   "\"output\"",
        PossibleValues: []string{},
        ExampleLS:      "target => \"mydata\"",
      },
    },
  },
  Ports: []*doc.ProcessorPort{},
}
}