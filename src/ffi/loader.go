package ffi

// Load FFI packages. Comment out the ones you don't need for
// faster/smaller builds.

import "purescript-prelude"
import "purescript-arrays"
import "purescript-console"

const Loader = true

const _ = prelude.Loader
const _ = console.Loader
const _ = arrays.Loader
