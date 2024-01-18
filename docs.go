//go:build docs

package simplewg

//go:generate rm -r docs
//go:generate go run "github.com/johnstarich/go/gopages" -internal -out docs -source-link "https://github.com/point-c/$GOPACKAGE/blob/main/{{.Path}}{{if .Line}}#L{{.Line}}{{end}}"

import _ "github.com/johnstarich/go/gopages"
