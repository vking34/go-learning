# Array & Slices

- Array Declaration:
    ```
    [n]T
    ```



## Slices
- A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.
- It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.

- ```capacity```: the number of elements in the underlying array starting from the index from which the slice is created.

- A slice can be re-sliced upto its capacity. Anything beyond that will cause the program to throw a run time error.

- methods:
    - ```make```: func make([]T, len, cap) []T
    - 