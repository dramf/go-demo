package templates

import "testing"

const scanTemplate = `
{{ .Name }}

{{ .Vu
`

const scan2 = `hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}`

func TestBuilder(t *testing.T) {
	tests :=[]struct{
		userTemplate string
		data *Data
	}{
		{
			"{{ .Title }}\n{{ .Description }}",
//			scan2,
			&Data{
				Title:       "Title 1",
				Description: "Description 1",
			},
		},
	}

	for _, test := range tests {
		r, err := BuildTemplate(test.userTemplate, test.data)
		if err != nil {
			t.Errorf("BuildTemplate error: %v", err)
			continue
		}
		t.Logf("Result: %q", r)
	}
}