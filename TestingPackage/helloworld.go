package main

import (
	operations "TestingPackage/Test"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	result := operations.Sum(10, 20)
	fmt.Println("Sum:", result)
}

/*
go.mod file is important to maintain the dependencies of the project (Binds all the packages together)
for example, we added the operations package using operations "Golang/Operations"
If go.mod is not present, then we had to use full path of the that is "example.com/src/.../GoLang/Operations"
*/

/*
If we have to create a stand alone go file then we have to use package main and main function in that file.

This is due to Golangs file structure. As it work on packages, where each file must be part of a package.
Golang doesnt have a concept of classes, so we have to use packages to group the functions together.

Fact.. We cant files with differnt packages in the same directory. Each file of same directory must have the same package name.
*/
