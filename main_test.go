package main


func TestDivision(t *testing.T) {
    tests := []struct{
        Content      string
        Title      string
    }{
		{ Content: "Hello World", Title: "english-words" },
		{ Content: "Makeschool class of 20201", Title: "makeschool" },
		{ Content: "What's going on in the function", Title: "vibes-test" }
    }
    for _, test := range tests {
        result, err := divide(test.Content, test.Title)
        assert.IsType(t, test.err, err)
        assert.Equal(t, test.result, result)
    }
}
