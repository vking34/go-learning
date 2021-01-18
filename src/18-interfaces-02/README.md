# Interface 02

## Notes:

- Interface with pointer receiver vs value receiver
    - __Value receiver__: it is legal to call the interface method with __both value receiver or pointer receiver__.
    - __Pointer receiver__: it is legal to call the interface method with __pointer receiver only__. If we try to assign the interface to value receiver while the interface value accepts pointer, the compiler can not references the value receiver.

- A type can implement multiple interfaces.

- Although Go does not support inheritance, it is possible to create a new interface by embedding other interfaces.

- Zero value of interface: ```nil```
    - If we try to call a method on the ```nil``` interface, the program will panic since the ```nil``` interface has no underlying value or a concrete type.