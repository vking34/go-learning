# Reflection

## Definition

- Reflection is the ability of a program to inspect its variables and values at runtime and find their types.

## What is the need to inspect a variable and find its type?

- The first question anyone gets when learning about reflection is why do we even need to inspect a variable and find its type at runtime when __each and every variable in our program is defined by us__ and we know its type at compile time itself. Well this is true most of the times, __but not always__.

## ```reflect``` package

- The ```reflect``` package implements run-time reflection in Go. The ```reflect``` package helps to __identify the underlying concrete type and the value of a interface{} variable__. 


- ```reflect.Type``` and ```reflect.Value```:

    - The concrete type of ```interface{}``` by ```reflect.Type``` and the underlying value is represented by ```reflect.Value```
    
    - There are two functions ```reflect.TypeOf()``` and ```reflect.ValueOf()``` which return the ```reflect.Type``` and ```reflect.Value``` respectively.

- ```reflect.Kind```: represents the specific kind of type, while ```Type``` represents the actual type of the interface{}
    
- ```NumberFiel()``` and ```Field()``` methods:
    - ```NumberFiel()``` method returns the number of fields in a struct
    - ```Field(i int)``` method returns the ```reflect.Value``` of the ```i```th field 

- ```Int()``` and ```String()``` methods extract the ```reflect.Value``` as an ```int64``` and ```string``` respectively
    