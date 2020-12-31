# Maps

## How to create a map?

```
make(map[string]int)
```


## Notes
- The zero value of a map is ```nil```. If you try to add elements to a ```nil``` map, a run-time __panic__ will occur. Hence the map has to be initialized before adding elements

- When accessing an element not present in a map, the map will return the zero value of the type of that element.

- Maps are reference types. When a map is assigned to a new variable, they both point to the same internal data structure. Hence changes made in one will reflect in the other.

- Maps can't be compared using the ```==``` operator. The ```==``` can be only used to check if a map is ```nil```.
    ```
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

    map2 := map1

    if map1 == map2 {
    }
    ```

    - this will fail

- The ways to compare 2 maps:
    - Compare each individual elements one by one
    - Use ```reflection```

