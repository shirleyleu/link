package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_extractLink(t *testing.T) {
	tests := []struct {
		name string
		html string
		want []link
	}{
		{
			name: "Single link",
			html: `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`,
			want: []link{{href:"/other-page", text:"A link to another page"}},
		},
		{
			name: "Multiple links",
			html: `<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`,
			want: []link{{href:"https://www.twitter.com/joncalhoun", text:"Check me out on twitter"},{href:"https://github.com/gophercises", text: "Gophercises is on Github!"}},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := strings.NewReader(tt.html)
			doc, _ := html.Parse(d)
			var links []link
			extractLink(doc, &links)
			assert.Equal(t, tt.want, links)
		})
	}
}

func Test_Deeper_Link(t *testing.T) {
	h := `<a href="/dog">
    <span>Something in a span</span>
    Text not in a span
    <b>Bold<div>Stuff<i>more</i>Mountain</div>text!</b>
</a>`
	want:= []link{{href:"/dog", text:"Something in a span Text not in a span BoldStuffmoreMountaintext!"}}

	d := strings.NewReader(h)
	doc, _ := html.Parse(d)
	var links []link
	extractLink(doc, &links)
	assert.Equal(t, want, links)
}
