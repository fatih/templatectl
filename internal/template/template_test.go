package template

import (
	"os"
	"testing"
)

func TestTemplate_Execute(t *testing.T) {
	test := []struct {
		name string
		in   string
		out  string
		envs map[string]string
	}{
		{
			name: "template with environment variable",
			envs: map[string]string{
				"SOME_ENV_VARIABLE_FOO": "foo",
			},
			in:  `This is: {{ env "SOME_ENV_VARIABLE_FOO" }}`,
			out: `This is: foo`,
		},
	}

	for _, ts := range test {
		t.Run(ts.name, func(t *testing.T) {
			for key, val := range ts.envs {
				if err := os.Setenv(key, val); err != nil {
					t.Fatal(err)
				}
			}

			tmpl := NewTemplate()
			out, err := tmpl.Execute(ts.in)
			if err != nil {
				t.Fatal(err)
			}

			if ts.out != out {
				t.Errorf("test case: %+v \n===== WANT =====\n\n%+v\n\t==== GOT ====\n\n%+v",
					ts.name, ts.out, out)
			}
		})
	}

}
