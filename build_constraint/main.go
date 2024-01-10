package main

func main() {
	// GOOS=linux
	// GOARCH=amd64
	// GOOS=linux GOARCH=amd64 go build -o build_constraint/linuxabc
	fooWithBuildTag()
	fooWithSuffix()

	// GOOS=windows
	// GOARCH=amd64
	// GOOS=windows GOARCH=amd64 go build -o build_constraint/windowsabc

}
