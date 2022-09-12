package test

// govulncheck fails on the next line with "undeclared name: bar"
func doSomething() string { return bar }
