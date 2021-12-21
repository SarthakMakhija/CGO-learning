# CGO-learning
Repository to understand invoking C functions from Golang

# C implementation
A Linked list where each node has an `int` key and an `int` value. Linked list supports the following behaviors -
- `put` which puts a key, and an associated value 
- `get` which returns the value by key 
- `getAllValues` which returns all the values
- `length` returns the number of nodes in the linked list
- `close` which frees the memory allocated for the linked list

# Building C Code
Run the following command in linked list package

`gcc -c linkedlist.c`

# GO implementation
- `Put` which invokes `C put`
- `GetBy` which invokes `C get`
- `GetAllValues` which invokes `C getAllValues`
- `Close` which invokes `C close`

# Building GO Code
Run the following command in linked list package

`go build`

# Running the tests
Run the following command in linked list package

`go test`

# References
- [Golang slices and C pointers](https://stackoverflow.com/questions/64852226/how-to-iterate-through-a-c-array)
- [Calling C code from go](https://karthikkaranth.me/blog/calling-c-code-from-go/)
- [ld: symbol(s) not found for architecture x86_64](https://github.com/golang/go/issues/31409)
