# lexer
A framework for building lexers. See
[examples/json](https://github.com/TimSatke/lexer/blob/master/examples/json) for
a working example on how to build a JSON lexer with this framework.

## How to use
Go get it with
```
go get -u github.com/TimSatke/lexer
```
Define your own states as you can see in
[examples/json/lexer.go](https://github.com/TimSatke/lexer/blob/master/examples/json/lexer.go).

## Why should you use this
* You want to build a parser and need a lexer, and don't want to build the base
  yourself, because it is effort to build and test it, and this framework
  already did that part

## Will there be a parser framework?
Eventually, yes. At the moment though, no. This is due to my lack of experience
with parsers and their architecture and design. Feel free to create one using
this framework, though.

## How does this lexer work?
First, you create a lexer. Second, you pass a starting state. This state will be
executed (a state is a function that can drain/accept runes from the lexer's
input), and must return the next state to go into.

The base lexer implementation provided with this package works on a byte slice.
It has two markers, `start` and `pos`. `pos` is the current position in the
input byte slice, `start` is the position, where the last token was emitted.
Upon emit, `start` will be set to `pos`, marking the start position of the next
token.

Whenever the lexer encounters unexpected runes, it is recommended to emit an
error token and `nil` as the next state. This will stop the lexer to process the
input. See
[this](https://github.com/TimSatke/lexer/blob/master/examples/json/lexer.go#L237-L248)
as an example on how to emit errors. The error tokens can then be processed by
your parser, which will consume the lexer's token stream (retrieve it with
`lexer.TokenStream().Tokens()`).