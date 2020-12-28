# Arrays and Slices


## Notes

### Arrays
- We can ignore the length of the array in the declaration by replacing it with ```...```.

- The size of the array is a part of the type. Hence ```[5]int``` and ```[7]int``` are distinct types. Because of this, arrays cannot be resized. Don't worry about this restriction since slices exist to overcome this.

- Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected in the original array. __Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.__

- Length: ```len(a)```

- 
### Slices

