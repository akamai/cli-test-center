package internal

/**
This class will be used for all output print templates. Do not modify spaces/newlines for existing templates unless
there is a requirement.
*/

var groupByTestRequest = "{{ $count := 1 }}" +
	"{{ range $mapKey, $mapValue := . }}" +
	"{{ range $key, $value := $mapValue.Value }}" +
	"{{if eq $key 0}}|-- {{ bold \"TR\" }}{{ bold $count }} {{$value.TestRequest.TestRequestUrl}}{{ end }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}{{if eq $key 0}}\n|   Keywords: {{ join $value.TestRequest.Tags }}{{ end }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{if eq $key 0}}{{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}\n|   {{ printHeader $header.HeaderName $header.HeaderAction $header.HeaderValue }}{{ end }}{{ end }}{{ end }}" +
	"\n|   |-- {{ bold $value.Order }} {{ $value.Condition.ConditionExpression }} {{ $ipVersion := printClientProfile $value.ClientProfile.IpVersion }}{{ bold $ipVersion }}" +
	"{{ end }}\n" +
	"{{ $count = inc $count }}{{end}}"

var groupByCondition = "{{ $count := 1 }}" +
	"{{ range $mapKey, $mapValue := . }}" +
	"|-- {{ bold \"C\" }}{{ bold $count }} {{$mapValue.Key}}" +
	"{{ range $key, $value := $mapValue.Value }}" +
	"\n|   |-- {{ bold $value.Order }} {{ $value.TestRequest.TestRequestUrl }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}\n|         Keywords: {{ join $value.TestRequest.Tags }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}\n|         {{ printHeader $header.HeaderName $header.HeaderAction $header.HeaderValue }}{{ end }}{{ end }}" +
	"\n|         {{ $ipVersion := printClientProfile $value.ClientProfile.IpVersion }}{{ bold $ipVersion }}" +
	"{{ end }}\n" +
	"{{ $count = inc $count }}{{ end }}"

var groupByIpVersion = "{{ range $mapKey, $mapValue := . }}" +
	"|-- {{ $ipVersion := printClientProfile $mapValue.Key }}{{ bold $ipVersion }}" +
	"{{ range $key, $value := $mapValue.Value }}" +
	"\n|   |-- {{ bold $value.Order }} {{ $value.TestRequest.TestRequestUrl }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}\n|         Keywords: {{ join $value.TestRequest.Tags }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}\n|         {{ printHeader $header.HeaderName $header.HeaderAction $header.HeaderValue }}{{ end }}{{ end }}" +
	"\n|         {{ $value.Condition.ConditionExpression }}" +
	"{{ end }}\n{{end}}"
