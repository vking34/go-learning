package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title:", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Author:", p.fullName())
	fmt.Println("Bio:", p.bio)
}

type website struct {
	posts []post
}

func (w website) contents() {
	fmt.Println("Content of website:")
	for _, post := range w.posts {
		post.details()
		fmt.Println()
	}
}

func main() {
	// Embedding a struct
	a1 := author{
		firstName: "Vuong",
		lastName:  "Le",
		bio:       "Software Engineer",
	}

	p1 := post{
		title:   "Go Learning",
		content: "Composition - Inheritance",
		author:  a1,
	}

	p1.details()
	fmt.Println()

	// Embedding slice of structs
	p2 := post{
		"Spring Bean",
		"Definition, Scope, Lifecycle",
		a1,
	}

	p3 := post{
		"Javascript Fundamentals",
		"Event Loop",
		a1,
	}

	w := website{
		posts: []post{p1, p2, p3},
	}

	w.contents()
}
