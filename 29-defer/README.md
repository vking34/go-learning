# Defer

- ```defer``` statement is used to execute a function call just before the surrouding function returning.

- Defer is allowed to methods

- The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.

- Stack of defers:
    - When a function has multiple defer calls, they are pushed on to a stack and executed in Last In First Out (LIFO) order.

- Advantages:
    - Stack of defers
    - No need to call multiple times of a function before return.
    