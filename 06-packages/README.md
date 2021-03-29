# Packages & Module !!!


## Modules
- A Go Module is nothing but a collection of Go packages.
- The import path for the custom package we create is derived from the name of the go module.

- For custom package in other workspace, we need to create to module.
    ```
    go mod init packagename
    ```


## Packages
- Packages are used to __organize Go source code for better reusability and readability__. Packages are a collection of Go sources files that reside in the same directory. Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

- Every executable Go application must contain the main function. This function is the entry point for execution. The main function should reside in the main package.

- ```package packagename``` specifies that a particular source file belongs to package ```packagename```. This should be the first line of every go source file.


### init() function

- init() is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.

- If a package has one or more init() functions they are automatically executed before the main package's main() function is called.

    ```
    import --> const --> var --> init()
    ```

- A package can have multiple init functions (either in a single file or distributed across multiple files) and they are called in the order in which they are presented to the compiler.

- A package will be initialized only once event if it is imported from multiple packages
