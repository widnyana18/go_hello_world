package main

type author struct {
	name string
	from string
	age  int
}

type media struct {
	size    int
	release int
	author
}

type music struct {
	title    string
	genre    string
	duration string
	media
}

type game struct {
	name  string
	genre string
	rate  float64
	media
}
