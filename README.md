### Go foreign export implementations for the standard library

Note that many values are currently missing. Only a minimum set has been implemented – just beyond enough for the standard compiler tests to pass. Please feel free to contribute any that you have implemented yourself (via a pull request).

#### Some basic conventions used in this code (not absolutely required, but please try to follow them)
* Run your code through `gofmt`
* Use provided type alias `Any` instead of `interface{}`
* Use provided type alias `Dict` instead of `map[string]Any` (intended for when you're dealing with ps records)
* Use provided type `[]Any` for arrays
* For input parameters used as-is (i.e. when no type assertion from `Any` is necessary), simply use a name without a trailing underscore.
  * For example:
    ```go
    exports["foo"] = func(bar Any) Any {
      return bar
    }
    ```
* For input parameters needing a type assertion from `Any` to a concrete type, use a name with a trailing underscore, then declare a type-inferred variable, with underscore removed, assigning the type assertion operation to it. This also allows the possiblity of a faster type assertion – getting and ignoring a failure code instead of a `panic`. When debugging, you might want to restore the `panic` possiblity until you're sure it's not an issue.
  * For example:
  ```go
  exports["foo"] = func(bar_ Any) Any {
    bar, _ := bar_.(int) // the ", _" gets and ignores the error flag
    ...
  }
  ```
  * And when debugging:
  ```go
  exports["foo"] = func(bar_ Any) Any {
    bar := bar_.(int) // remove ", _" to get friendly type assertion panics
    ...
  }
