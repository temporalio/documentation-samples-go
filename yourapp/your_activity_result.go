// /* @dac start
// id: how-to-define-activity-return-values-in-go
// title: How to define Activity return values in Go
// label: Activity return values
// description: A Go-based Activity Definition can return either just an `error` or a `customValue, error` combination.
// */

// /*
// A Go-based Activity Definition can return either just an `error` or a `customValue, error` combination (same as a Workflow Definition).
// You may wish to use a `struct` type to hold all custom values, just keep in mind they must all be serializable.
// */
// package yourapp

// // YourActivityResultObject is the struct returned from your Activity.
// // Use a struct so that you can return multiple values of different types.
// // Additionally, your function signature remains compatible if the fields change.
// type YourActivityResultObject struct {
// 	ResultFieldX string
// 	ResultFieldY int
// }

// // @dac end
