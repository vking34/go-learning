# Structs

## Definition
- A struct is a user-defined type that represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than having each of them as separate values.

## Declaration
```
type Employee struct {  
    firstName string
    lastName  string
    age       int
}
```

## Exported structs and fields
- If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages. Similarly, if the fields of a struct start with caps, they can be accessed from other packages.

## Notes:
- Anonymous struct only creates a new struct variable and does not define any new struct type like named struct.
- Zero value of struct: When a struct is defined and it is not explicitly initialized with any value, the fields of the struct are assigned their zero values by default.
- The Go language gives us the option to use ```emp4.firstName``` instead of the explicit dereference ```(*emp4).firstName``` to access the firstName field
- Anonymous field is fields without name in a struct. By default, the name of anonymous fields is the name of type.
- Promoted fields: fields of a anonymous struct field in the (nested) struct.
- Export: If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages. Similarly, if the fields of a struct start with caps, they can be accessed from other packages.
- Struct Equality:
    - Structs are value types and are comparable if each of their fields are comparable. Two struct variables are considered equal if their corresponding fields are equal.
    - Struct variables are not comparable if they contain fields that are not comparable
    - ```map``` is not comparable.