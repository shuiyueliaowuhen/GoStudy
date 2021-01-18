package main

type Address struct {
	province string
	city     string
}

type Person struct {
	name    string
	gender  string
	address Address
}

func main() {
	p := Person{
		name:   "呵呵",
		gender: "男",
		address: Address{
			province: "江苏",
			city:     "南京",
		},
	}

}
