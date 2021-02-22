# Select

## Definition
- The ```select``` statement is used to choose from multiple send/receive channel operations. The select statement __blocks until one of the send/receive operation is ready__. If __multiple operations are ready__, one of them is __chosen at random__. The syntax is similar to ```switch``` except that each of the case statement will be a channel operation.

## Notes

- The __default case__ in a select statement is executed when none of the other case is ready. This is generally used to __prevent the select statement from blocking.__
- __Do not let select be empty__, because the select statement will block until one of its cases is executed. In this case the select statement doesn't have any cases and hence it will block forever resulting in a deadlock. This program will panic