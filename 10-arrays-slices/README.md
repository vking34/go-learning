# Arrays and Slices


## Notes

### Arrays: [n]T
- We can ignore the length of the array in the declaration by replacing it with ```...```.

- The size of the array is a part of the type. Hence ```[5]int``` and ```[7]int``` are distinct types. Because of this, arrays cannot be resized. Don't worry about this restriction since slices exist to overcome this.

- Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array. __Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.__

- Length: ```len(a)```

- 
### Slices: []T

- Two ways to create a slice:
    - create reference from existing array:
        ```
        a := [5]int{76, 77, 78, 79, 80}
        var b []int = a[1:4]
        ```

    - create reference from array declaration with length specific
        ```
        c := []int{6, 7, 8}
        ```

- A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.
- It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.

- ```capacity```: the number of elements in the underlying array starting from the index from which the slice is created.

- A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.

- methods:
    - ```make```: func make([]T, len, cap) []T

- When new elements are appended to the slice, a new array is created. The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. The capacity of the new slice is now twice that of the old slice.

- the zero value of a slice type is ```nil```.

- the structure type of slice:
    ```
    type slice struct {  
        Length        int
        Capacity      int
        ZerothElement *byte
    }
    ```

- A slice contains the length, capacity and a pointer to the zeroth element of the array. When a slice is passed to a function, even though it's passed by value, the pointer variable will refer to the same underlying array. Hence when a slice is passed to a function as parameter, __changes made inside the function are visible outside the function too__.


- Memory optimisation:
    - Slices hold a reference to the underlying array. As long as the slice is in memory, the array cannot be garbage collected.
    - Lets assume that we have a very large array and we are interested in processing only a small part of it. The important thing to be noted here is that the array will still be in memory since the slice references it.
    - Ways to solve:
        - the ```copy``` function: 

