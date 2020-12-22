# Packages & Module !!!

## Packages
- Packages are used to __organize Go source code for better reusability and readability__. Packages are a collection of Go sources files that reside in the same directory. Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

- Every executable Go application must contain the main function. This function is the entry point for execution. The main function should reside in the main package.

- ```package packagename``` specifies that a particular source file belongs to package ```packagename```. This should be the first line of every go source file.

## Modules
- A Go Module is nothing but a collection of Go packages.
- The import path for the custom package we create is derived from the name of the go module.

- For custom package in other workspace, we need to create to module.
    ```
    go mod init packagename
    ```
