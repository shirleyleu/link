package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func Test_extractLink(t *testing.T) {
	tests := []struct {
		name string
		file string
		want []link
	}{
		{
			name: "Single link ex1",
			file: "ex1.html",
			want: []link{{href:"/other-page", text:"A link to another page"}},
		},
		{
			name: "Multiple links ex2",
			file: "ex2.html",
			want: []link{{href:"https://www.twitter.com/joncalhoun", text:"Check me out on twitter"},{href:"https://github.com/gophercises", text: "Gophercises is on Github !"}},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := os.Open(tt.file)
			require.NoError(t, err)
			doc, err := html.Parse(d)
			require.NoError(t, err)
			var links []link
			extractLink(doc, &links)
			assert.Equal(t, tt.want, links)
		})
	}
}

func Test_full_document_ex3(t *testing.T) {
	want:= []link{
		{href:"#", text:"Login"},
		{href:"/lost", text:"Lost? Need help?"},
		{href:"https://twitter.com/marcusolsson", text:"@marcusolsson"},
	}
	d, err := os.Open("ex3.html")
	require.NoError(t, err)
	doc, err := html.Parse(d)
	require.NoError(t, err)
	var links []link
	extractLink(doc, &links)
	assert.Equal(t, want, links)
}


func Test_commented_text_should_not_be_included_ex4(t *testing.T) {
	want:= []link{{href:"/dog-cat", text:"dog cat"}}
	d, err := os.Open("ex4.html")
	require.NoError(t, err)
	doc, err := html.Parse(d)
	require.NoError(t, err)
	var links []link
	extractLink(doc, &links)
	assert.Equal(t, want, links)
}

func Test_multiple_text_tags_ex5(t *testing.T) {
	want:= []link{{href:"/dog", text:"Something in a span Text not in a span Bold text!"}}
	d, err := os.Open("ex5.html")
	require.NoError(t, err)
	doc, err := html.Parse(d)
	require.NoError(t, err)
	var links []link
	extractLink(doc, &links)
	assert.Equal(t, want, links)
}

func Test_Deeper_Link_ex6(t *testing.T) {
	want:= []link{{href:"/dog", text:"Something in a span Text not in a span Bold Stuff more Mountain text!"}}
	d, err := os.Open("ex6.html")
	require.NoError(t, err)
	doc, err := html.Parse(d)
	require.NoError(t, err)
	var links []link
	extractLink(doc, &links)
	assert.Equal(t, want, links)
}
