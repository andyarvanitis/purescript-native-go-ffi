package ffi

// Load FFI packages. Comment out the ones you don't need for
// faster/smaller builds.

import "purescript-prelude"
import "purescript-arrays"
import "purescript-console"

type Loader = bool

type _ = prelude.Loader
type _ = console.Loader
type _ = arrays.Loader
