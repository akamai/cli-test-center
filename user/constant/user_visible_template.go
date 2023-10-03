package constant

/**
This class will be used for all output print templates. Do not modify spaces/newlines for existing templates unless
there is a requirement.
Blogs foe templates:
https://blog.logrocket.com/using-golang-templates/
https://pkg.go.dev/text/template
*/

var GroupByTestRequest = "{{ $count := 1 }}" +
	"{{ $trMapLength := len . }}{{ $trMapLength = dec $trMapLength }}" +
	"{{ range $index, $mapValue := . }}" +
	"{{ range $key, $value := $mapValue.Value }}" +
	"{{if eq $key 0}}|-- {{ bold \"TR\" }}{{ bold $count }} {{ bold $value.TestRequest.RequestMethod }} {{ printRequestURL $value.TestRequest }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}\n|       Keywords: {{ join $value.TestRequest.Tags }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|       Customized headers: {{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}{{ printHeader $header }}{{if ne $headerIndex $length }}\n|			    {{ end }}{{ end }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|       Request body: {{ printRequestBody $value.TestRequest }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|       URL encode: {{ $value.TestRequest.EncodeRequestBody }}{{ end }}{{ end }}" +
	"{{ if ne $value.ParentOrder 0 }}\n|       |-- {{ bold $value.ParentOrder }}.{{ bold $value.Order }} {{ else }}\n|       |-- {{ bold $value.Order }} {{ end }}" +
	"{{ printCondition $value.Condition }} | {{$ipVersion := printClientProfile $value.ClientProfile }}{{ bold $ipVersion }}" +
	"{{ if ne $value.ParentOrder 0 }}{{ $length := len $value.SetVariables }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|               Set variables: {{ range $variableIndex, $variable := $value.SetVariables }}{{printSetVariables $variable }}{{if ne $variableIndex $length }}\n|                              {{end}}{{end}}{{end}}{{end}}" +
	"{{ if eq $value.ParentOrder 0 }}{{ $length := len $value.SetVariables }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|             Set variables: {{ range $variableIndex, $variable := $value.SetVariables }}{{printSetVariables $variable }}{{if ne $variableIndex $length }}\n|                           {{end}}{{end}}{{end}}{{end}}" +
	"{{ end }}{{if ne $index $trMapLength }}\n|\n{{ end }}" +
	"{{ $count = inc $count }}{{end}}"

var GroupByCondition = "{{ $condLength := len . }}{{ $condLength = dec $condLength }}" +
	"{{ $count := 1 }}" +
	"{{ range $index, $mapValue := . }}" +
	"|-- {{ bold \"C\" }}{{ bold $count }} {{$mapValue.Key}}" +
	"{{ $tcLength := len $mapValue.Value }}{{ $tcLength = dec $tcLength }}" +
	"{{ range $tc_index, $value := $mapValue.Value }}" +
	"{{ if ne $value.ParentOrder 0 }}\n|      |-- {{ bold $value.ParentOrder }}.{{ bold $value.Order }} {{ else }}\n|      |-- {{ bold $value.Order }} {{ end }}" +
	" {{ bold $value.TestRequest.RequestMethod }} {{ printRequestURL $value.TestRequest }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}\n|          Keywords: {{ join $value.TestRequest.Tags }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|          Customized headers: {{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}{{ printHeader $header }}{{if ne $headerIndex $length }}\n|			       {{ end }}{{ end }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|          Request body: {{ printRequestBody $value.TestRequest }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|          URL encode: {{ $value.TestRequest.EncodeRequestBody }}{{ end }}" +
	"{{ $length := len $value.SetVariables }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|          Set variables: {{ range $variableIndex, $variable := $value.SetVariables }}{{printSetVariables $variable }}{{if ne $variableIndex $length }}\n|                         {{end}}{{end}}{{end}}" +
	"\n|          Client profile: {{$ipVersion := printClientProfile $value.ClientProfile }}{{ bold $ipVersion }} {{if ne $tc_index $tcLength }}\n|{{ end }}" +
	"{{ end }}{{if ne $index $condLength }}\n|\n{{ end }}" +
	"{{ $count = inc $count }}{{ end }}"

var GroupByClientProfile = "{{ $cpLength := len . }}{{ $cpLength = dec $cpLength }}" +
	"{{ range $index, $mapValue := . }}" +
	"|-- {{ bold $mapValue.Key }}" +
	"{{ $tcLength := len $mapValue.Value }}{{ $tcLength = dec $tcLength }}" +
	"{{ range $tc_index, $value := $mapValue.Value }}" +
	"{{ if ne $value.ParentOrder 0 }}\n|      |-- {{ bold $value.ParentOrder }}.{{ bold $value.Order }} {{ else }}\n|      |-- {{ bold $value.Order }} {{ end }}" +
	"{{ bold $value.TestRequest.RequestMethod }} {{ printRequestURL $value.TestRequest }}" +
	"{{ $length := len $value.TestRequest.Tags }}{{ if ne $length 0 }}\n|          Keywords: {{ join $value.TestRequest.Tags }}{{ end }}" +
	"{{ $length := len $value.TestRequest.RequestHeaders }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|          Customized headers: {{ range $headerIndex, $header := $value.TestRequest.RequestHeaders }}{{ printHeader $header }}{{if ne $headerIndex $length }}\n|			       {{ end }}{{ end }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|          Request body: {{ printRequestBody $value.TestRequest }}{{ end }}" +
	"{{ if eq $value.TestRequest.RequestMethod \"POST\" }}\n|          URL encode: {{ $value.TestRequest.EncodeRequestBody }}{{ end }}" +
	"{{ $length := len $value.SetVariables }}{{ if ne $length 0 }}{{ $length = dec $length }}\n|          Set variables: {{ range $variableIndex, $variable := $value.SetVariables }}{{printSetVariables $variable }}{{if ne $variableIndex $length }}\n|                         {{end}}{{end}}{{end}}" +
	"\n|          Condition: {{ printCondition $value.Condition }}{{if ne $tc_index $tcLength }}\n|{{ end }}" +
	"{{ end }}{{if ne $index $cpLength }}\n|\n{{ end }}{{end}}"
