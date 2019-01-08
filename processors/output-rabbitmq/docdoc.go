// Code generated by "bitfanDoc "; DO NOT EDIT
package rabbitmqoutput

import "github.com/awillis/bitfan/processors/doc"

func (p *processor) Doc() *doc.Processor {
	return &doc.Processor{
		Name:       "rabbitmqoutput",
		ImportPath: "github.com/awillis/bitfan/processors/output-rabbitmq",
		Doc:        "",
		DocShort:   "",
		Options: &doc.ProcessorOptions{
			Doc: "",
			Options: []*doc.ProcessorOption{
				&doc.ProcessorOption{
					Name:           "AddField",
					Alias:          "add_field",
					Doc:            "Add a field to an event. Default value is {}",
					Required:       false,
					Type:           "hash",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Arguments",
					Alias:          "arguments",
					Doc:            "Extra exchange arguments. Default value is {}",
					Required:       false,
					Type:           "amqp.Table",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "ConnectRetryInterval",
					Alias:          "connect_retry_interval",
					Doc:            "Time in seconds to wait before retrying a connection. Default value is 1",
					Required:       false,
					Type:           "time.Duration",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "ConnectionTimeout",
					Alias:          "connection_timeout",
					Doc:            "Time in seconds to wait before timing-out. Default value is 0 (no timeout)",
					Required:       false,
					Type:           "time.Duration",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Durable",
					Alias:          "durable",
					Doc:            "Is this exchange durable - should it survive a broker restart? Default value is true",
					Required:       false,
					Type:           "bool",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Exchange",
					Alias:          "exchange",
					Doc:            "The name of the exchange to send message to. There is no default value for this setting.",
					Required:       true,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "ExchangeType",
					Alias:          "exchange_type",
					Doc:            "The exchange type (fanout, topic, direct). There is no default value for this setting.",
					Required:       true,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Heartbeat",
					Alias:          "heartbeat",
					Doc:            "Interval (in second) to send heartbeat to rabbitmq. Default value is 0\nIf value if lower than 1, server's interval setting will be used.",
					Required:       false,
					Type:           "time.Duration",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Host",
					Alias:          "host",
					Doc:            "RabbitMQ server address. There is no default value for this setting.",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Key",
					Alias:          "key",
					Doc:            "The routing key to use when binding a queue to the exchange. Default value is \"\"\nThis is only relevant for direct or topic exchanges (Routing keys are ignored on fanout exchanges).\nThis setting can be dynamic using the %{foo} syntax.",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Passive",
					Alias:          "passive",
					Doc:            "Use queue passively declared, meaning it must already exist on the server. Default value is false\nTo have BitFan to create the queue if necessary leave this option as false.\nIf actively declaring a queue that already exists, the queue options for this plugin (durable, etc) must match those of the existing queue.",
					Required:       false,
					Type:           "bool",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Password",
					Alias:          "password",
					Doc:            "RabbitMQ password. Default value is \"guest\"",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Persistent",
					Alias:          "persistent",
					Doc:            "Should RabbitMQ persist messages to disk? Default value is true",
					Required:       false,
					Type:           "bool",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Port",
					Alias:          "port",
					Doc:            "RabbitMQ port to connect on. Default value is 5672",
					Required:       false,
					Type:           "int",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "SSL",
					Alias:          "ssl",
					Doc:            "Enable or disable SSL. Default value is false",
					Required:       false,
					Type:           "bool",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Tags",
					Alias:          "tags",
					Doc:            "Add any number of arbitrary tags to your event. There is no default value for this setting.\nThis can help with processing later. Tags can be dynamic and include parts of the event using the %{field} syntax.",
					Required:       false,
					Type:           "array",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "User",
					Alias:          "user",
					Doc:            "RabbitMQ username. Default value is \"guest\"",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "VerifySSL",
					Alias:          "verify_ssl",
					Doc:            "Validate SSL certificate. Default value is false",
					Required:       false,
					Type:           "bool",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
				&doc.ProcessorOption{
					Name:           "Vhost",
					Alias:          "vhost",
					Doc:            "The vhost to use. Default value is \"/\"",
					Required:       false,
					Type:           "string",
					DefaultValue:   nil,
					PossibleValues: []string{},
					ExampleLS:      "",
				},
			},
		},
		Ports: []*doc.ProcessorPort{},
	}
}
