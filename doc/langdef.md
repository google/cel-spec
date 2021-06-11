# Language Definition

This page constitutes the reference for CEL. For a gentle introduction, see
[Intro](intro.md).

## Contents

- [Overview](#overview)
- [Syntax](#syntax)
    - [Name Resolution](#name-resolution)
- [Values](#values)
    - [Numeric Values](#numeric-values)
    - [Enumerations](#enumerations)
    - [String and Bytes Values](#string-and-bytes-values)
    - [Aggregate Values](#aggregate-values)
    - [Booleans and Null](#booleans-and-null)
    - [Type Values](#type-values)
    - [Abstract Types](#abstract-types)
    - [Protocol Buffer Data Conversion](#protocol-buffer-data-conversion)
    - [Dynamic Values](#dynamic-values)
- [JSON Data Conversion](#json-data-conversion)
- [Gradual Type Checking](#gradual-type-checking)
- [Evaluation](#evaluation)
    - [Evaluation Environment](#evaluation-environment)
    - [Runtime Errors](#runtime-errors)
    - [Logical Operators](#logical-operators)
    - [Macros](#macros)
    - [Field Selection](#field-selection)
- [Performance](#performance)
    - [Abstract Sizes](#abstract-sizes)
    - [Time Complexity](#time-complexity)
    - [Space Complexity](#space-complexity)
    - [Macro Performance](#macro-performance)
    - [Performance Limits](#performance-limits)
- [Functions](#functions)
    - [Extension Functions](#extension-functions)
    - [Receiver Call Style](#receiver-call-style)
- [Standard Definitions](#standard-definitions)
    - [Equality and Ordering](#equality-and-ordering)
    - [Overflow](#overflow)
    - [Timezones](#timezones)
    - [Regular Expressions](#regular-expressions)
    - [List of Standard Definitions](#list-of-standard-definitions)
- [Appendix 1: Legacy Behavior](#appendix-1-legacy-behavior)
    - [Enums as Ints](#enums-as-ints)

## Overview

In the taxonomy of programming languages, CEL is:

*   **memory-safe:** programs cannot access unrelated memory, such as
    out-of-bounds array indexes or use-after-free pointer dereferences;
*   **side-effect-free:** a CEL program only computes an output from its inputs;
*   **terminating:** CEL programs cannot loop forever;
*   **strongly-typed:** values have a well-defined type, and operators and
    functions check that their arguments have the expected types;
*   **dynamically-typed:** types are associated with values, not with variables
    or expressions, and type safety is enforced at runtime;
*   **gradually-typed:** an optional type-checking phase before runtime can
    detect and reject some programs which would violate type constraints.

## Syntax

The grammar of CEL is defined below, using `|` for alternatives, `[]` for
optional, `{}` for repeated, and `()` for grouping.

```grammar
Expr           = ConditionalOr ["?" ConditionalOr ":" Expr] ;
ConditionalOr  = [ConditionalOr "||"] ConditionalAnd ;
ConditionalAnd = [ConditionalAnd "&&"] Relation ;
Relation       = [Relation Relop] Addition ;
Relop          = "<" | "<=" | ">=" | ">" | "==" | "!=" | "in" ;
Addition       = [Addition ("+" | "-")] Multiplication ;
Multiplication = [Multiplication ("*" | "/" | "%")] Unary ;
Unary          = Member
               | "!" {"!"} Member
               | "-" {"-"} Member
               ;
Member         = Primary
               | Member "." IDENT ["(" [ExprList] ")"]
               | Member "[" Expr "]"
               | Member "{" [FieldInits] "}"
               ;
Primary        = ["."] IDENT ["(" [ExprList] ")"]
               | "(" Expr ")"
               | "[" [ExprList] "]"
               | "{" [MapInits] "}"
               | LITERAL
               ;
ExprList       = Expr {"," Expr} ;
FieldInits     = IDENT ":" Expr {"," IDENT ":" Expr} ;
MapInits       = Expr ":" Expr {"," Expr ":" Expr} ;
```

Implementations are required to support at least:

*   32 levels of nested expressions;
*   32 repetitions of self-recursive or repetitive rules, i.e.:
    *   32 terms separated by `||` in a row;
    *   32 terms separated by `&&` in a row;
    *   32 relations in a row;
    *   32 binary arithmetic operators of the same precedence in a row;
    *   32 selection (`.`) operators in a row;
    *   32 indexing (`[_]`) operators in a row;
    *   32 function call arguments;
    *   list literals with 32 elements;
    *   map or message literals with 32 fields.

This grammar corresponds to the following operator precedence and associativity:

Precedence | Operator        | Description                    | Associativity
---------- | --------------- | ------------------------------ | -------------
1          | ()              | Function call                  | Left-to-right
&nbsp;     | .               | Qualified name or field access |
&nbsp;     | []              | Indexing                       |
&nbsp;     | {}              | Field initialization           |
2          | - (unary)       | Negation                       | Right-to-left
&nbsp;     | !               | Logical NOT                    |
3          | *               | Multiplication                 | Left-to-right
&nbsp;     | /               | Division                       |
&nbsp;     | %               | Remainder                      |
4          | +               | Addition                       |
&nbsp;     | - (binary)      | Subtraction                    |
5          | == != < > <= >= | Relations                      |
&nbsp;     | in              | Inclusion test                 |
6          | &&              | Logical AND                    |
7          | \|\|            | Logical OR                     |
8          | ?:              | Conditional                    | Right-to-left

The lexis is defined below. As is typical, the `WHITESPACE` and `COMMENT`
productions are only used to recognize separate lexical elements and are ignored
by the grammar.

```
IDENT          ::= [_a-zA-Z][_a-zA-Z0-9]* - RESERVED
LITERAL        ::= INT_LIT | UINT_LIT | FLOAT_LIT | STRING_LIT | BYTES_LIT
                 | BOOL_LIT | NULL_LIT
INT_LIT        ::= -? DIGIT+ | -? 0x HEXDIGIT+
UINT_LIT       ::= INT_LIT [uU]
FLOAT_LIT      ::= -? DIGIT* . DIGIT+ EXPONENT? | -? DIGIT+ EXPONENT
DIGIT          ::= [0-9]
HEXDIGIT       ::= [0-9abcdefABCDEF]
EXPONENT       ::= [eE] [+-]? DIGIT+
STRING_LIT     ::= [rR]? ( "    ~( " | NEWLINE )*  "
                         | '    ~( ' | NEWLINE )*  '
                         | """  ~"""*              """
                         | '''  ~'''*              '''
                         )
BYTES_LIT      ::= [bB] STRING_LIT
ESCAPE         ::= \ [bfnrt"'\]
                 | \ x HEXDIGIT HEXDIGIT
                 | \ u HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
                 | \ U HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
                 | \ [0-3] [0-7] [0-7]
NEWLINE        ::= \r\n | \r | \n
BOOL_LIT       ::= "true" | "false"
NULL_LIT       ::= "null"
RESERVED       ::= BOOL_LIT | NULL_LIT | "in"
                 | "as" | "break" | "const" | "continue" | "else"
                 | "for" | "function" | "if" | "import" | "let"
                 | "loop" | "package" | "namespace" | "return"
                 | "var" | "void" | "while"
WHITESPACE     ::= [\t\n\f\r ]+
COMMENT        ::= '//' ~NEWLINE* NEWLINE
```

For the sake of a readable representation, the escape
sequences (`ESCAPE`) are kept implicit in string tokens. This means that strings
without the `r` or `R` (raw) prefix process `ESCAPE` sequences, while in strings
with the raw prefix they stay uninterpreted. See documentation of string
literals below.

The following identifiers are reserved due to their use as literal values or in
the syntax:

    false in null true

The following identifiers are reserved to allow easier embedding of CEL into a
host language.

    as break const continue else for function if import let loop package
    namespace return var void while

In general it is a bad idea for those defining contexts or extensions to use
identifiers that are reserved words in programming languages which might embed
CEL.

### Name Resolution

A CEL expression is parsed in the scope of a specific protocol buffer package or
message, which controls the interpretation of names. The scope is set by the
application context of an expression. A CEL expression can contain simple names
as in `a` or qualified names as in `a.b`. The meaning of such expressions is a
combination of one or more of:

*   Variables and Functions: some simple names refer to variables in the
    execution context, standard functions, or other name bindings provided by
    the CEL application.
*   Field selection: appending a period and identifier to an expression could
    indicate that we're accessing a field within a protocol buffer or map.
*   Protocol buffer package names: a simple or qualified name could represent an
    absolute or relative name in the protocol buffer package namespace. Package
    names must be followed by a message type, enum type, or enum constant.
*   Protocol buffer message types, enum types, and enum constants: following an
    optional protocol buffer package name, a simple or qualified name could
    refer to a message type, and enum type, or an enum constant in the
    package's namespace.

Resolution works as follows. If `a.b` is a name to be resolved in the context of
a protobuf declaration with scope `A.B`, then resolution is attempted, in order,
as `A.B.a.b`, `A.a.b`, and finally `a.b`. To override this behavior, one can use
`.a.b`; this name will only be attempted to be resolved in the root scope, i.e.
as `a.b`.

If name qualification is mixed with field selection, the longest prefix of the
name which resolves in the current lexical scope is used. For example, if
`a.b.c` resolves to a message declaration, and `a.b` does so as well with `c` a
possible field selection, then `a.b.c` takes priority over the interpretation
`(a.b).c`. Explicit parentheses can be used to choose the field selection
interpretation.

## Values

Values in CEL represent any of the following:

Type          | Description
------------- | ---------------------------------------------------------------
`int`         | 64-bit signed integers
`uint`        | 64-bit unsigned integers
`double`      | 64-bit IEEE floating-point numbers
`bool`        | Booleans (`true` or `false`)
`string`      | Strings of Unicode code points
`bytes`       | Byte sequences
`list`        | Lists of values
`map`         | Associative arrays with `int`, `uint`, `bool`, or `string` keys
`null_type`   | The value `null`
message names | Protocol buffer messages
`type`        | Values representing the types in the first column

### Numeric Values

CEL supports only 64-bit integers and 64-bit IEEE double-precision
floating-point. We only support positive, decimal integer literals; negative
integers are produced by the unary negation operator. Note that the integer 7 as
an `int` is a different value than 7 as a `uint`, which would be written `7u`.
Double-precision floating-point is also supported, and the integer 7 would be
written `7.`, `7.0`, `7e0`, or any equivalent representation using a decimal
point or exponent.

Note that currently there are no automatic arithmetic conversions for the
numeric types (`int`, `uint`, and `double`). The arithmetic operators typically
contain overloads for arguments of the same numeric type, but not for mixed-type
arguments. Therefore an expression like `1 + 1u` is going to fail to dispatch.
To perform mixed-type arithmetic, use explicit conversion functions such as
`uint(1) + 1u`. Such explicit conversions will maintain their meaning even if
arithmetic conversions are added in the future.

CEL provides no way to control the finer points of floating-point arithmetic,
such as expression evaluation, rounding mode, or exception handling. However,
any two not-a-number values will compare equal even if their underlying
properties are different.

### Enumerations

Each protocol buffer enumeration is its own CEL type.  Enumerations have no
operations other than equality (and inequality), obtaining the type from a
value, and conversion to an `int`. (No conversion to string is supported, as
an enum value might have 0, 1, or many names.)  Values of type `int` can be
converted to enum values as long as they are in the range `-2^31` to `2^31 - 1`.
A value of type `string` can be converted to enum values as long as it is a
defined for that enum.

### String and Bytes Values

Strings are sequences of Unicode code points. Bytes are sequences of octets
(eight-bit data).

Quoted string literals are delimited by either single- or double-quote
characters, where the closing delimiter must match the opening one, and can
contain any unescaped character except the delimiter or newlines (either CR or
LF).

Triple-quoted string literals are delimited by three single-quotes or three
double-quotes, and may contain any unescaped characters except for the delimiter
sequence. Again, the closing delimiter must match the opening one. Triple-quoted
strings may contain newlines.

Both sorts of strings can include escape sequences, described below.

If preceded by an `r` or `R` character, the string is a _raw_ string and does
not interpret escape sequences. Raw strings are useful for expressing strings
which themselves must use escape sequences, such as regular expressions or
program text.

Bytes literals are represented by string literals preceded by a `b` or `B`
character. The bytes literal is the sequence of bytes given by the UTF-8
representation of the string literal. In addition, the octal escape sequence are
interpreted as octet values rather than as Unicode code points. Both raw and
multiline string literals can be used for byte literals.

Escape sequences are a backslash (`\ `) followed by one of the following:

*   A punctuation mark representing itself:
    *   `\ `: backslash
    *   `?`: question mark
    *   `"`: double quote
    *   `'`: single quote
    *   `` ` ``: backtick
*   A code for whitespace:
    *   `a`: bell
    *   `b`: backspace
    *   `f`: form feed
    *   `n`: line feed
    *   `r`: carriage return
    *   `t`: horizontal tab
    *   `v`: vertical tab
*   A `u` followed by four hexadecimal characters, encoding a Unicode code point
    in the
    [BMP](https://en.wikipedia.org/wiki/Plane_\(Unicode\)#Basic_Multilingual_Plane).
    Characters in other Unicode planes can be represented with surrogate pairs.
    Valid only for string literals.
*   A `U` followed by eight hexadecimal characters, encoding a Unicode code
    point. Valid only for string literals.
*   A `x` or `X` followed by two hexadecimal characters. For strings, it denotes
    the unicode code point. For bytes, it represents an octet value.
*   Three octal digits, in the range `000` to `377`. For strings, it denotes the
    unicode code point. For bytes, it represents an octet value.

Examples:

CEL Literal   | Meaning
------------- | ---------------------------------------------------
`""`          | Empty string
`'""'`        | String of two double-quote characters
`'''x''x'''`  | String of four characters "x''x"
`"\""`        | String of one double-quote character
`"\\"`        | String of one backslash character
`r"\\"`       | String of two backslash characters
`b"abc"`      | Byte sequence of 97, 98, 99
`b"ÿ"`        | Sequence of bytes 195 and 191 (UTF-8 of &yuml;)
`b"\303\277"` | Also sequence of bytes 195 and 191
`"\303\277"`  | String of "&Atilde;&iquest;" (code points 195, 191)
`"\377"`      | String of "&yuml;" (code point 255)
`b"\377"`     | Sequence of byte 255 (_not_ UTF-8 of &yuml;)
`"\xFF"`      | String of "&yuml;" (code point 255)
`b"\xFF"`     | Sequence of byte 255 (_not_ UTF-8 of &yuml;)

While strings must be sequences of valid Unicode code points, no Unicode
normalization is attempted on strings, as there are several normal forms, they
can be expensive to convert, and we don't know which is desired. If Unicode
normalization is desired, it should be performed outside of CEL, or done as a
custom extension function.

Likewise, no advanced collation is attempted on strings, as this depnds on the
normalization and can be locale-dependent. Strings are simply treated as
sequences of code points and are ordered with lexicographic ordering based on
the numeric value of the code points.

### Aggregate Values

Lists are ordered sequences of values.

Maps are a set of key values, and a mapping from these keys to arbitrary values.
Key values must be an allowed key type: `int`, `uint`, `bool`, or `string`. Thus
maps are the union of what's allowed in protocol buffer maps and JSON objects.

Note that the type checker uses a finer-grained notion of list and map types.
Lists are `list(A)` for the homogenous type `A` of list elements. Maps are
`map(K, V)` for maps with keys of type `K` and values of type `V`. The type
`dyn` is used for heterogeneous values See
[Gradual Type Checking](#gradual-type-checking). But these constraints are only
enforced within the type checker; at runtime, lists and maps can have
heterogeneous types.

Any protocol buffer message is a CEL value, and each message type is its own CEL
type, represented as its fully-qualified name.

A list can be denoted by the expression `[e1, e2, ..., eN]`, a map by `{ek1:
ev1, ek2: ev2, ..., ekN: evN}`, and a message by `M{f1: e1, f2: e2, ..., fN:
eN}`, where `M` must be a simple or qualified name which resolves to a message
type (see [Name Resolution](#name-resolution)). For a map, the entry keys are
sub-expressions that must evaluate to values of an allowed type (`int`, `uint`,
`bool`, or `string`). For a message, the field names are identifiers. It is an
error to have duplicate keys or field names. The empty list, map, and message
are `[]`, `{}`, and `M{}`, respectively.

See [Field Selection](#field-selection) for accessing elements of lists, maps,
and messages.

### Booleans and Null

CEL has `true` and `false` as the literals for the `bool` type, with the usual
meanings.

The null value is written `null`. It is used in conversion to and from protocol
buffer and JSON data, but otherwise has no built-in meaning in CEL. In
particular, null has its own type (`null_type`) and is not necessarily allowed
where a value of some other type is expected.

### Type Values

Every value in CEL has a runtime type which is itself a value. The standard
function `type(x)` returns the type of expression `x`.

As types are values, those values (`int`, `string`, etc.) also have a type: the
type `type`, which is an expression by itself which in turn also has type
`type`. So

*   `type(1)` evaluates to `int`
*   `type("a")` evaluates to `string`
*   `type(1) == string` evaluates to `false`
*   `type(type(1)) == type(string)` evaluates to `true`

### Abstract Types

A CEL implementation can add new types to the language. These types will be
given names in the same namespace as the other types, but will have no special
support in the language syntax. The only way to construct or use values of these
abstract types is through functions which the implementor must also provide.

Commonly, an abstract type will have a representation as a protocol buffer, so
that it can be stored or transmitted across a network. In this case, the
abstract type will be given the same name as the protocol buffer, which will
prevent CEL programs from being able to use that particular protocol buffer
message type; they will not be able to construct values of that type by message
expressions nor access the message fields. The abstract type remains abstract.

By default, CEL uses `google.protobuf.Timestamp` and `google.protobuf.Duration`
as abstract types. The standard functions provide ways to construct and
manipulate these values, but CEL programs cannot construct them with message
expressions or access their message fields.

### Protocol Buffer Data Conversion

Protocol buffers have a richer range of types than CEL, so Protocol buffer data
is converted to CEL data when read from a message field, and CEL data is
converted in the other direction when initializing a field. In general, protocol
buffer data can be converted to CEL without error, but range errors are possible
in the other direction.

Protocol Buffer Field Type                       | CEL Type
------------------------------------------------ | --------
int32, int64, sint32, sint64, sfixed32, sfixed64 | `int`
uint32, uint64, fixed32, fixed64                 | `uint`
float, double                                    | `double`
bool, string, bytes                              | same
enum E                                           | E
repeated                                         | `list`
map<K, V>                                        | `map`
oneof                                            | options expanded individually, at most one is set
message M                                        | M, except for conversions below

Signed integers, unsigned integers, and floating point numbers are converted to
the singular CEL type of the same sort. The CEL type is capable of expressing
the full range of protocol buffer values. When converting from CEL to protocol
buffers, an out-of-range CEL value results in an error.

Boolean, string, and bytes types have identical ranges and are converted without
error.

Repeated fields are converted to CEL lists of converted values, preserving the
order. In the other direction, the CEL list elements must be of the right type
and value to be converted to the corresponding protocol buffer type. Similarly,
protocol buffer maps are converted to CEL maps, and CEL map keys and values must
have the right type and value to be converted in the other direction.

Oneof fields are represented by the translation of each of their options as a
separate field, but at most one of these fields will be "set", as detected by
the `has()` macro. See [Macros](#macros).

Since protocol buffer messages are first-class CEL values, message-valued fields
are used without conversion.

Every protocol buffer field has a default value, and there is no semantic
difference between a field set to this default value, and an unset field. For
message fields, there default value is just the unset state, and an unset
message field is distinct from one set to an empty (i.e. all-unset) message.

The `has()` macro (see [Macros](#macros)) tells whether a message field is set
(i.e. not unset, hence not set to the default value). If an unset field is
nevertheless selected, it evaluates to its default value, or if it is a message
field, it evaluates to an empty (i.e. all-unset) message. This allows
expressions to use iterative field selection to examine the state of fields in
deeply nested messages without needing to test whether every intermediate field
is set. (See exception for wrapper types, below.)

### Dynamic Values

CEL automatically converts certain protocol buffer messages in the
`google.protobuf` package to other types.

google.protobuf message | CEL Conversion
----------------------- | --------------
`Any`                   | dynamically converted to the contained message type, or error
`ListValue`             | list of `Value` messages
`Struct`                | map (with string keys, `Value` values)
`Value`                 | dynamically converted to the contained type (null, double, string, bool, `Struct`, or `ListValue`)
wrapper types           | converted as eponymous field type

The wrapper types are `BoolValue`, `BytesValue`, `DoubleValue`, `FloatValue`,
`Int32Value`, `Int64Value`, `NullValue`, `StringValue`, `Uint32Value`, and
`Uint64Value`. Values of these wrapper types are converted to the obvious type.
Additionally, field selection of an unset message field of wrapper type will
evaluate to `null`, instead of the default message. This is an exception to the
usual evaluation of unset message fields.

Note that this implies some cascading conversions. An `Any` message might be
converted to a `Struct`, one of whose `Value`-typed values might be converted to
a `ListValue` of more values, and so on.

Also note that all of these conversions are dynamic at runtime, so CEL's static
type analysis cannot avoid the possibility of type-related errors in expressions
using these dynamic values.

## JSON Data Conversion

CEL can also work with JSON data. Since there is a natural correspondence of
most CEL data with protocol buffer data, and protocol buffers have a
[defined mapping](https://developers.google.com/protocol-buffers/docs/proto3#json)
to JSON, this creates a natural mapping of CEL to JSON. This creates an exact
bidirectional mapping between JSON types and a subset of CEL data:

JSON Type | CEL Type
--------- | -----------------------------------------------
`null`    | `null`
Boolean   | `bool`
Number    | `double` (except infinities or NaN)
String    | `string`
Array     | `list` of bi-convertible elements
Object    | `map` (with string keys, bi-convertible values)

We define JSON mappings for much of the remainder of CEL data, but note that
this data will not map back in to CEL as the same value:

CEL Data                                               | JSON Data
------------------------------------------------------ | ---------
`int`                                                  | Number if in interoperable range, otherwise decimal String.
`uint`                                                 | Number if in interoperable range, otherwise decimal String.
double infinity                                        | String `"Infinity"` or `"-Infinity"`
double NaN                                             | String "NaN"
`bytes`                                                | String of base64-encoded bytes
message                                                | JSON conversion of protobuf message.
`list` of convertible elements                         | JSON Array of converted values
`list` with a non-convertible element                  | none
`map` with string keys and convertible values          | JSON Object with converted values
`map` with a non-string key or a non-convertible value | none
`type`                                                 | none

The "interoperable" range of integer values is `-(2^53-1)` to `2^53 - 1`.

## Gradual Type Checking

CEL is a dynamically-typed language, meaning that the types of the values of the
variables and expressions might not be known until runtime. However, CEL has an
optional type-checking phase that takes annotation giving the types of all
variables and tries to deduce the type of the expression and of all its
sub-expressions. This is not always possible, due to the dynamic expansion of
certain messages like `Struct`, `Value`, and `Any` (see
[Dynamic Values](#dynamic-values)). However, if a CEL program does not use
dynamically-expanded messages, it can be statically type-checked.

The type checker uses a richer type system than the types of the dynamic values:
lists have a type parameter for the type of the elements, and maps have two
parameters for the types of keys and values, respectively. These richer types
preserve the stronger type guarantees that protocol buffer messages have. We can
infer stronger types from the standard functions, such as accessing list
elements or map fields. However, the `type()` function and dynamic dispatch to
particular function overloads only use the coarser types of the dynamic values.

The type checker also introduces the `dyn` type, which is the union of all other
types. Therefore the type checker could accept a list of heterogeneous values as
`dyn([1, 3.14, "foo"])`, which is given the type `list(dyn)`. The standard
function `dyn` has no effect at runtime, but signals to the type checker that
its argument should be considered of type `dyn`, `list(dyn)`, or a `dyn`-valued
map.

A CEL type checker attempts to identify possible runtime errors (see
[Runtime Errors](#runtime-errors)), particularly `no_matching_overload` and
`no_such_field`, ahead of runtime. It also serves to optimize execution speed
by narrowing down the number of possible matching overloads for a function
call, and by allowing for a more efficient (unboxed) runtime representation of
values.

By construction, a CEL expression that does not use the dynamic features coming
from `Struct`, `Value`, or `Any`, can be fully statically type checked and all
overloads can be resolved ahead of runtime.

If a CEL expression uses a mixture of dynamic and static features, a type
checker will still attempt to derive as much information as possible and
delegate undecidable type decisions to runtime.

The type checker is an optional phase of evaluation. Running the type checker
does not affect the result of evaluation, it can only reject expressions as
ill-typed in a given typing context.

## Evaluation

For a given evaluation environment, a CEL expression will deterministically
evaluate to either a value or an error. Here are how different expressions are
evaluated:

*   **Literals:** the various kinds of literals (numbers, booleans, strings,
    bytes, and `null`) evaluate to the values they represent.
*   **Variables:** variables are looked up in the binding environment. An
    unbound variable evaluates to an error.
*   **List, Map, and Message expressions:** each sub-expression is evaluated and
    if any sub-expression results in an error, this expression results in an
    error. Otherwise, it results in the list, map, or message of the
    sub-expression results, or an error if one of the values is of the wrong
    type.
*   **Field selection:** see [Field Selection](#field-selection).
*   **Macros:** see [Macros](#macros).
*   **Logical operators:** see [Logical Operators](#logical-operators).
*   **Other operators:** operators are translated into specially-named functions
    and the sub-expressions become their arguments, for instance `e1 + e2`
    becomes `_+_(e1, e2)`, which is then evaluated as a normal function.
*   **Normal functions:** all argument sub-expressions are evaluated and if any
    results in an error, then this expression results in an error. Otherwise,
    the function is identified by its name and dispatched to a particular
    overload based on the types of the sub-expression values. See
    [Functions](#functions).

Because CEL is free of side-effects, the order of evaluation among
sub-expressions is not guaranteed. If multiple subexpressions would evaluate to
errors causing the enclosing expression to evaluate to an error, it will
propagate one or more of the sub-expression erors, but it is not specified which
ones.

### Evaluation Environment

A CEL expression is parsed and evaluated in the scope of a particular protocol
buffer package, which controls name resolution as described above, and a binding
context, which binds identifiers to values, errors, and functions. A given
identifier has different meanings as a function name or as a variable, depending
on the use. For instance in the expression `size(requests) > size`, the first
`size` is a function, and the second is a variable.

The CEL implementation provides mechanisms for adding bindings of variable names
to either values or errors. The implementation will also provide function
bindings for at least all the standard functions listed below.

Some implementations might make use of a _context proto_, where a single
protocol buffer message represents all variable bindings: each field in the
message is a binding of the field name to the field value. This provides a
convenient encapsulation of the binding environment.

The evaluation environment can also specify the expected type of the result. If
the expected type is one of the protocol buffer wrapper messages, then CEL will
attempt to convert the result to the wrapper message, or will raise an error if
the conversion fails.

### Runtime Errors

In general, when a runtime error is produced, expression evaluation is
terminated; exceptions to this rule are discussed in
[Logical Operators](#logical-operators) and [Macros](#macros).

CEL provides the following built-in runtime errors:

*   `no_matching_overload`: this function has no overload for the types of the
    arguments.
*   `no_such_field`: a map or message does not contain the desired field.

There is no in-language representation of errors, no generic way to raise them,
and no way to catch or bypass errors, except for the short-circuiting behavior
of the logical operators, and macros.

### Logical Operators

In the conditional operator `e ? e1 : e2`, evaluates to `e1` if `e` evaluates to
`true`, and `e2` if `e` evaluates to `false`. The untaken branch is presumed to
not be executed, though that is an implementation detail.

In the boolean operators `&&` and `||`: if any of their operands uniquely
determines the result (`false` for `&&` and `true` for `||`) the other operand
may or may not be evaluated, and if that evaluation produces a runtime error, it
will be ignored. This makes those operators commutative (in contrast to
traditional boolean short-circuit operators). The rationale for this behavior is
to allow the boolean operators to be mapped to indexed queries, and align better
with SQL semantics.

To get traditional left-to-right short-circuiting evaluation of logical
operators, as in C or other languages (also called "McCarthy Evaluation"), the
expression `e1 && e2` can be rewritten `e1 ? e2 : false`. Similarly, `e1 || e2`
can be rewritten `e1 ? true : e2`.

### Macros

CEL supports a small set of predefined macros. Macro invocations have the same
syntax as function calls, but follow different type checking rules and runtime
semantics than regular functions. An application of CEL opts-in to which macros
to support, selecting from the predefined set of macros. The currently available
macros are:

*   `has(e.f)`: tests whether a field is available. See "Field Selection" below.
*   `e.all(x, p)`: tests whether a predicate holds for all elements of a list
    `e` or keys of a map `e`. Here `x` is a simple identifier to be used in `p`
    which binds to the element or key. The `all()` macro combines per-element
    predicate results with the "and" (`&&`) operator, so if any predicate
    evaluates to false, the macro evaluates to false, ignoring any errors from
    other predicates.
*   `e.exists(x, p)`: like the `all()` macro, but combines the predicate results
    with the "or" (`||`) operator.
*   `e.exists_one(x,p)`: like the `exists()` macro, but evaluates to `true` only
    if the predicate of exactly one element/key evaluates to `true`, and the
    rest to `false`. Any other combination of boolean results evaluates to
    `false`, and any predicate error causes the macro to raise an error.
*   `e.map(x, t)` transforms a list `e` by taking each element `x` to the
    element given by the expression `t`, which can use the variable `x`. For
    instance, `[1, 2, 3].map(n, n * n)` evaluates to `[1, 4, 9]`.  Any evaluation
    error for any element causes the macro to raise an error. The `map()` macro
    is not supported when `e` is a map.
*   `e.filter(x, p)` returns the sublist of all elements `x` of list `e` which
    evaluate to `true` in the predicate expression `p` (which can use variable
    `x`). For instance, `[1, 2, 3].filter(i, i % 2 > 0)` evaluates to `[1, 3]`.
    If no elements evaluate to `true`, the result is an empty list. Any
    evaluation error for any element causes the macro to raise an error. The
    `filter()` macro is not supported on maps.

### Field Selection

A field selection expression, `e.f`, can be applied both to messages and to
maps. For maps, selection is interpreted as the field being a string key.

The semantics depends on the type of the result of evaluating expression `e`:

1.  If `e` evaluates to a message and `f` is not declared in this message, the runtime
    error `no_such_field` is raised.
2.  If `e` evaluates to a message and `f` is declared, but the field is not set, the
    default value of the field's type will be produced. Note that this is `null`
    for messages or the according primitive default value as determined by
    proto2 or proto3 semantics.
3.  If `e` evaluates to a map, then `e.f` is equivalent to `e['f']` (where `f`
    is still being used as a meta-variable, e.g. the expression `x.foo` is
    equivalent to the expression `x['foo']` when `x` evaluates to a map).
4.  In all other cases, `e.f` evaluates to an error.

To test for the presence of a field, the boolean-valued macro `has(e.f)` can be
used.

1.  If `e` evaluates to a map, then `has(e.f)` indicates whether the string `f`
    is a key in the map (note that `f` must syntactically be an identifier).
2.  If `e` evaluates to a message and `f` is not a declared field for the
    message, `has(e.f)` raises a  `no_such_field` error.
3.  If `e` evaluates to a protocol buffers version 2 message and `f` is a
    defined field:
    -   If `f` is a repeated field or map field, `has(e.f)` indicates whether
        the field is non-empty.
    -   If `f` is a singular or  oneof field, `has(e.f)` indicates
        whether the field is set.
4.  If `e` evaluates to a protocol buffers version 3 message and `f` is a
    defined field:
    -   If `f` is a repeated field or map field, `has(e.f)` indicates whether
        the field is non-empty.
    -   If `f` is a oneof or singular message field, `has(e.f)` indicates
        whether the field is set.
    -   If `f` is some other singular field, `has(e.f)` indicates whether the
        field's value is its default value (zero for numeric fields, false for
        booleans, empty for strings and bytes).
5.  In all other cases, `has(e.f)` evaluates to an error.

## Performance

Since one of the main applications for CEL is for execution of untrusted
expressions with reliable containment, the time and space cost of evaluation
is an essential part of the specification of the language. But we also want to
give considerable freedom in how to implement the language. To balance these
concerns, we specify only the time and space computational complexity of
language constructs and standard functions (see [Functions](#functions)).

CEL applications are responsible for noting the computational complexity of
any extension functions they provide.

### Abstract Sizes

Space and time complexity will be measured in terms of an abstract size
measurment of CEL expressions and values. The size of a CEL value depends on
its type:

*   *string*: The size is its length, i.e. the number of code points, plus a
    constant.
*   *bytes*: The size is its length, i.e. the number of bytes, plus a constant.
*   *list*: The size is the sum of sizes of its entries, plus a constant.
*   *map*: The size is the sum of the key size plus the value size for all of
    its entries, plus a constant.
*   *message*: The size is the sum of the size of all fields, plus a constant.
*   All other values have constant size.

The size of a CEL program is:

*   *string literal*: The size of the resulting value.
*   *bytes literal*: The size of the resulting value.
*   Grammatical aggregates are the sum of the size of their components.
*   Gramatical primitives other than above have constant size.

Thus, the size of a CEL program is bounded by either the length of the source
text string or the bytes of the proto-encoded AST.

The inputs to a CEL expression are the _bindings_ given to the evaluator and
the _literals_ within the expression itself.

### Time Complexity

Unless otherwise noted, the time complexity of an expression is the sum of the
time complexity of its sub-expressions, plus the sum of the sizes of the
sub-expression values, plus a constant.

For instance, an expression `x` has constant time complexity since it has no
sub-expressions.  An expression `x != y` takes time proportional to the sum of
sizes of the bindings of `x` and `y`, plus a constant.

Some functions cost less than this:

*   The conditional expression `_?_:_`, only evaluates one of the alternative
    sub-expressions.
*   For the `size()` function on lists and maps, the time is proportional to
    the length of its input, not its total size (plus the time of the
    sub-expression).
*   The index operator on lists takes constant time (plus the time of the
    sub-expressions).
*   The select operator on messages takes constant time (plus the time of the
    sub-expression).

Some functions take more time than this.  The following functions take time
proportional to the _product_ of their input sizes (plus the time of the
sub-expressions):

*    The index operator on maps.
*    The select operator on maps.
*    The in operator.
*    The `contains`, `startsWith`, `endsWith`, and `matches` functions on
     strings.

See below for the time cost of macros.

Implementations are free to provide a more performant implementation. For
instance, a hashing implementation of maps would make indexing/selection
faster, but we do not require such sophistication from all implementations.

### Space Complexity

Unless otherwise noted, the space complexity of an expression is the sum of the
space complexity of its sub-expressions, plus a constant. The exceptions are:

*   *Literals*: Message, map, and list literals allocate new space for their
    output.
*   *Concatenation*: The `_+_` operator on lists and stings allocate new space
    for their output.

See below for the space cost of macros.

We'll assume that bytes-to-string and string-to-bytes conversions do not need
to allocate new space.

### Macro Performance

Macros can take considerably more time and space than other constructs, and
can lead to exponential behavior when nested or chained.  For instance,

```
[0,1].all(x,
  [0,1].all(x,
    ...
      [0,1].all(x, 1/0)...))
```

takes exponential (in the size of the expression) time to evaluate, while

```
["foo","bar"].map(x, [x+x,x+x]).map(x, [x+x,x+x])...map(x, [x+x,x+x])
```

is exponential in both time and space.

The time and space cost of macros is the cost of the range sub-expression `e`,
plus the follwing:

*   `has(e.f)`: Space is constant.
    *   If `e` is a map, time is linear in size of `e`.
    *   If `e` is a message, time is constant.
*   `e.all(x,p)`, `e.exists(x,p)`, and `e.exists_one(x,p)`
    *   Time is the sum of the time of `p` for each element of `e`.
    *   Space is constant.
*   `e.map(x,t)`
    *   Time is the sum of time of`t` for each element of `e`.
    *   Space is the sum of space of `t` for each element of `e`, plus a
        constant.
*   `e.filter(x,t)`
    *   Time is the sum of time of `t` for each element of `e`.
    *   Space is the space of `e`.

### Performance Limits

Putting this all together, we can make the following statements about the cost
of evaluation. Let `P` be the non-literal size of the expression, `L` be the
size of the literals, `B` be the size of the bindings, and `I=B+L` be the total
size of the inputs.

*   The macros other than `has()` are the only avenue for exponential
    behavior. This can be curtailed by the implementation allowing applications
    to set limits on the recursion or chaining of macros, or disable them
    entirely.
*   The concatenation operator `_+_` is the only operator that dramatically
    increases the space complexity, with the program `x + x + ... + x` taking
    time and space `O(B * P^2)`.
*   The string-detection functions (`contains()` and friends) yield a boolean
    result, thus cannot be nested to drive exponential or even higher
    polynomial cost.  We can bound the time cost by `O(B^2 * P)`, with a
    limiting case being `x.contains(y) || x.contains(y) || ...`.
*   The map indexing operators yield a smaller result than their input, and
    thus are also limited in their ability to increase the cost. A particularly
    bad case would be an expensive selection that returns a subcomponent that
    contains the majority of the size of the aggregate, resulting in a time
    cost of `O(P * I)`, and see below.
*   Eliminating all of the above and using only default-cost functions, plus
    aggregate literals, time and space are limited `O(P * I)`.
    A limiting time example is `size(x) + size(x) + ...`.
    A limiting time and space example is `[x, x, ..., x]`.

Note that custom function will alter this analysis if they are more expensive
than the default costs.

## Functions

CEL functions have no observable side-effects (there maybe side-effects like
logging or such which are not observable from CEL). The default argument
evaluation strategy for functions is strict, with exceptions from this rule
discussed in [Logical Operators](#logical-operators) and [Macros](#macros).

Functions are specified by a set of overloads. Each overload defines the number
and type of arguments and the type of the result, as well as an opaque
computation. Argument and result types can use type variables to express
overloads which work on lists and maps. At runtime, a matching overload is
selected and the according computation invoked. If no overload matches, the
runtime error `no_matching_overload` is raised (see also
[Runtime Errors](#errors)). For example, the standard function `size` is
specified by the following overloads:

<table border="1">
  <tr>
   <th rowspan="4">
      size
    </th>
    <td>
      (string) -> int
    </td>
    <td>
      string length
    </td>
  </tr>
  <tr>
    <td>
      (bytes) -> int
    </td>
    <td>
      bytes length
    </td>
  </tr>
  <tr>
    <td>
      (list(A)) -> int
    </td>
    <td>
      list size
    </td>
  </tr>
  <tr>
    <td>
      (map(A, B)) -> int
    </td>
    <td>
      map size
    </td>
  </tr>
</table>

Overloads must have non-overlapping argument types, after erasure of all type
variables (similar as type erasure in Java). Thus an implementation can
implement overload resolution by simply mapping all argument types to a strong
hash.

Operator subexpressions are treated as calls to specially-named built-in
functions. For instance, the expression `e1 + e2` is dispatched to the function
`_+_` with arguments `e1` and `e2`. Note that since`_+_` is not an identifier,
there would be no way to write this as a normal function call.

See [Standard Definitions](#standard-definitions) for the list of all predefined
functions and operators.

### Extension Functions

It is possible to add extension functions to CEL, which then behave in no way
different than standard functions. The mechanism how to do this is
implementation dependent and usually highly curated. For example, an application
domain of CEL can add a new overload to the `size` function above, provided this
overload's argument types do not overlap with any existing overload. For
methodological reasons, CEL disallows to add overloads to operators.

### Receiver Call Style

A function overload can be declared to use receiver call-style, so it must be
called as `e1.f(e2)` instead of `f(e1, e2)`. Overloads with different call
styles are non-overlapping per definition, regardless of their types.

## Standard Definitions

All predefined operators, functions and constants are listed in the table below.
For each symbol, the available overloads are listed. Operator symbols use a
notation like `_+_` where `_` is a placeholder for an argument.

### Equality and Ordering

Equality (`_==_`) and inequality (`_!=_`) are defined for all types. Inequality
is the logical negation of equality, i.e. `e1 != e2` is the same as `!(e1 ==
e2)` for all expressions `e1` and `e2`.

Equality and inequality are homogeneous; comparing values of different runtime
types results in a runtime error. Thus `2 == 3` is false, but `2 == 2.0` is an
error.

For `double`, all not-a-number (`NaN`) values compare equal. This is different
than the usual semantics of floating-point numbers, but it is more consistent
with the usual expectations of reflexivity, and is more compatible with the
usual notions of equality on protocol buffers.

Lists are unequal if their lengths are different. Otherwise, for lists `a` and
`b` with length `N`, `a == b` is equivalent to

```
a[0] == b[0] && a[1] == b[1] && ... && a[N-1] == b[N-1]
```

Maps are unequal if their key sets are different, otherwise for maps `a` and
`b` with keyset `k1, k2, ..., kN`, `a == b` is equivalent to

```
a[k1] == b[k1] && a[k2] == b[k2] && ... && a[kN] == b[kN]
```

So for equality of both lists and maps this means:

- if the list lengths / map key sets are different, the result is false;
- if one or more element comparisons is false, the result is false;
- if all element comparisons are true, the result is true;
- otherwise the result is an error.

Ordering operators are defined for `int`, `uint`, `double`, `string`, `bytes`,
`bool`, as well as `timestamp` and `duration`. Strings obey lexicographic
ordering of the code points, and bytes obey lexicographic ordering of the byte
values. The ordering operators obey the usual algebraic properties, i.e. `e1 <=
e2` gives the same result as `!(e1 > e2)` as well as `(e1 < e2) || (e1 == e2)`
when the expressions involved do not have side effects.

### Overflow

Arithmetic operations raise an error when the results exceed the range of the
integer type (int, uint) or the timestamp or duration type.  An error is also
raised for conversions which exceed the range of the target type.

### Timezones

Timezones are expressed in the following grammar:

```grammar
TimeZone = "UTC" | LongTZ | FixedTZ ;
LongTZ = ? list available at
           http://joda-time.sourceforge.net/timezones.html ? ;
FixedTZ = ( "+" | "-" ) Digit Digit ":" Digit Digit ;
Digit = "0" | "1" | ... | "9" ;
```

Fixed timezones are explicit hour and minute offsets from UTC. Long timezone
names are like `Europe/Paris`, `CET`, or `US/Central`.

### Regular Expressions

Regular expressions follow the
[RE2 syntax](https://github.com/google/re2/wiki/Syntax). Regular expression
matches succeed if they match a substring of the argument. Use explicit anchors
(`^` and `$`) in the pattern to force full-string matching, if desired.

### List of Standard Definitions

TODO: automatically generate these descriptions from the cel-go implementation.
See [cel-go/issues/9](https://github.com/google/cel-go/issues/9).

<table style="width=100%" border="1">
  <col width="15%">
  <col width="40%">
  <col width="45%">
  <tr>
    <th>Symbol</th>
    <th>Type</th>
    <th>Description</th>
  </tr>
  <tr>
    <th rowspan="1">
      !_
    </th>
    <td>
      (bool) -> bool
    </td>
    <td>
      logical not
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      -_
    </th>
    <td>
      (int) -> int
    </td>
    <td>
      negation
    </td>
  </tr>
  <tr>
    <td>
      (double) -> double
    </td>
    <td>
      negation
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      _!=_
    </th>
    <td>
      (A, A) -> bool
    </td>
    <td>
      inequality
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      _%_
    </th>
    <td>
      (int, int) -> int
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> uint
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      _&&_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      logical and
    </td>
  </tr>
  <tr>
    <td>
      (bool, ...) -> bool
    </td>
    <td>
      logical and (variadic)
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      _*_
    </th>
    <td>
      (int, int) -> int
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> uint
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> double
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <th rowspan="9">
      _+_
    </th>
    <td>
      (int, int) -> int
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> uint
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> double
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (string, string) -> string
    </td>
    <td>
      String concatenation. Space and time cost proportional to the sum of the
      input sizes.
    </td>
  </tr>
  <tr>
    <td>
      (bytes, bytes) -> bytes
    </td>
    <td>
      bytes concatenation
    </td>
  </tr>
  <tr>
    <td>
      (list(A), list(A)) -> list(A)
    </td>
    <td>
      List concatenation. Space and time cost proportional to the sum of the
      input sizes.
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Duration) -> google.protobuf.Timestamp
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Timestamp) -> google.protobuf.Timestamp
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> google.protobuf.Duration
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <th rowspan="6">
      _-_
    </th>
    <td>
      (int, int) -> int
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> uint
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> double
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Timestamp) -> google.protobuf.Duration
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Duration) -> google.protobuf.Timestamp
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> google.protobuf.Duration
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      _/_
    </th>
    <td>
      (int, int) -> int
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> uint
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> double
    </td>
    <td>
      arithmetic
    </td>
  </tr>
  <tr>
    <th rowspan="8">
      _<=_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (int, int) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (string, string) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (bytes, bytes) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Timestamp) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <th rowspan="8">
      _<_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (int, int) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (string, string) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (bytes, bytes) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Timestamp) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      _==_
    </th>
    <td>
      (A, A) -> bool
    </td>
    <td>
      equality
    </td>
  </tr>
  <tr>
    <th rowspan="8">
      _>=_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (int, int) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (string, string) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (bytes, bytes) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Timestamp) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <th rowspan="8">
      _>_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (int, int) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (uint, uint) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (double, double) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (string, string) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (bytes, bytes) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp, google.protobuf.Timestamp) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Duration, google.protobuf.Duration) -> bool
    </td>
    <td>
      ordering
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      _?_:_
    </th>
    <td>
      (bool, A, A) -> A
    </td>
    <td>
      The conditional operator. See above for evaluation semantics. Will
      evaluate the test and only one of the remaining sub-expressions.
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      _[_]
    </th>
    <td>
      (list(A), int) -> A
    </td>
    <td>
      list indexing. Constant time cost.
    </td>
  </tr>
  <tr>
    <td>
      (map(A, B), A) -> B
    </td>
    <td>
      map indexing.  For string keys, cost is proportional to the size of the
      map keys times the size of the index string.
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      in
    </th>
    <td>
      (A, list(A)) -> bool
    </td>
    <td>
      list membership. Time cost proportional to the product of the size of
      both arguments.
    </td>
  </tr>
  <tr>
    <td>
      (A, map(A, B)) -> bool
    </td>
    <td>
      map key membership. Time cost proportional to the product of the size of
      both arguments.
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      _||_
    </th>
    <td>
      (bool, bool) -> bool
    </td>
    <td>
      logical or
    </td>
  </tr>
  <tr>
    <td>
      (bool, ...) -> bool
    </td>
    <td>
      logical or (variadic)
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      bool
    </th>
    <td>
      type(bool)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      bytes
    </th>
    <td>
      type(bytes)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (string) -> bytes
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      contains
    </th>
    <td>
      string.(string) -> bool
    </td>
    <td>
      Tests whether the string operand contains the substring. Time cost
      proportional to the product of sizes of the arguments.
    </td>
  </tr>
  <tr>
    <th rowspan="4">
      double
    </th>
    <td>
      type(double)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (int) -> double
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (uint) -> double
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (string) -> double
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      duration
    </th>
    <td>
      (string) -> google.protobuf.Duration
    </td>
    <td>
      type conversion, duration should end with "s", which stands for seconds
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      dyn
    </th>
    <td>
      type(dyn)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (A) -> dyn
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      endsWith
    </th>
    <td>
      string.(string) -> bool
    </td>
    <td>
      Tests whether the string operand ends with the suffix argument. Time cost
      proportional to the product of the sizes of the arguments.
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getDate
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get day of month from the date in UTC, one-based indexing
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get day of month from the date with timezone, one-based indexing
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getDayOfMonth
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get day of month from the date in UTC, zero-based indexing
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get day of month from the date with timezone, zero-based indexing
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getDayOfWeek
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get day of week from the date in UTC, zero-based, zero for Sunday
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get day of week from the date with timezone, zero-based, zero for Sunday
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getDayOfYear
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get day of year from the date in UTC, zero-based indexing
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get day of year from the date with timezone, zero-based indexing
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getFullYear
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get year from the date in UTC
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get year from the date with timezone
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      getHours
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get hours from the date in UTC, 0-23
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get hours from the date with timezone, 0-23
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Duration.() -> int
    </td>
    <td>
      get hours from duration
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      getMilliseconds
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get milliseconds from the date in UTC, 0-999
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get milliseconds from the date with timezone, 0-999
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Duration.() -> int
    </td>
    <td>
      milliseconds from duration, 0-999
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      getMinutes
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get minutes from the date in UTC, 0-59
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get minutes from the date with timezone, 0-59
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Duration.() -> int
    </td>
    <td>
      get minutes from duration
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      getMonth
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get month from the date in UTC, 0-11
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get month from the date with timezone, 0-11
    </td>
  </tr>
  <tr>
    <th rowspan="3">
      getSeconds
    </th>
    <td>
      google.protobuf.Timestamp.() -> int
    </td>
    <td>
      get seconds from the date in UTC, 0-59
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Timestamp.(string) -> int
    </td>
    <td>
      get seconds from the date with timezone, 0-59
    </td>
  </tr>
  <tr>
    <td>
      google.protobuf.Duration.() -> int
    </td>
    <td>
      get seconds from duration
    </td>
  </tr>
  <tr>
    <th rowspan="6">
      int
    </th>
    <td>
      type(int)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (uint) -> int
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (double) -> int
    </td>
    <td>
      Type conversion. Rounds toward zero, then errors if result is out of
      range.
    </td>
  </tr>
  <tr>
    <td>
      (string) -> int
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (enum E) -> int
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (google.protobuf.Timestamp) -> int
    </td>
    <td>
      Convert timestamp to int64 in seconds since Unix epoch.
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      list
    </th>
    <td>
      type(list(dyn))
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      map
    </th>
    <td>
      type(map(dyn, dyn))
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      matches
    </th>
    <td>
      (string, string) -> bool
    </td>
    <td>
      Matches first argument against regular expression in second argument.
      Time cost proportional to the product of the sizes of the arguments.
    </td>
  </tr>
  <tr>
    <td>
      string.(string) -> bool
    </td>
    <td>
      Matches the self argument against regular expression in first argument.
      Time cost proportional to the product of the sizes of the arguments.
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      null_type
    </th>
    <td>
      type(null)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <th rowspan="4">
      size
    </th>
    <td>
      (string) -> int
    </td>
    <td>
      string length
    </td>
  </tr>
  <tr>
    <td>
      (bytes) -> int
    </td>
    <td>
      bytes length
    </td>
  </tr>
  <tr>
    <td>
      (list(A)) -> int
    </td>
    <td>
      list size. Time cost proportional to the length of the list.
    </td>
  </tr>
  <tr>
    <td>
      (map(A, B)) -> int
    </td>
    <td>
      map size. Time cost proportional to the number of entries.
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      startsWith
    </th>
    <td>
      string.(string) -> bool
    </td>
    <td>
      Tests whether the string operand starts with the prefix argument. Time
      cost proportional to the product of the sizes of the arguments.
    </td>
  </tr>
  <tr>
    <th rowspan="7">
      string
    </th>
    <td>
      type(string)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (int) -> string
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (uint) -> string
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (double) -> string
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (bytes) -> string
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (timestamp) -> string
    </td>
    <td>
      type conversion, using the same format as timestamp string parsing
    </td>
  </tr>
  <tr>
    <td>
      (duration) -> string
    </td>
    <td>
      type conversion, using the same format as duration string parsing
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      timestamp
    </th>
    <td>
      (string) -> google.protobuf.Timestamp
    </td>
    <td>
      Type conversion of strings to timestamps according to RFC3339. Example: "1972-01-01T10:00:20.021-05:00"
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      type
    </th>
    <td>
      type(dyn)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (A) -> type(dyn)
    </td>
    <td>
      returns type of value
    </td>
  </tr>
  <tr>
    <th rowspan="4">
      uint
    </th>
    <td>
      type(uint)
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (int) -> uint
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <td>
      (double) -> uint
    </td>
    <td>
      Type conversion. Rounds toward zero, then errors if result is out of
      range.
    </td>
  </tr>
  <tr>
    <td>
      (string) -> uint
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      E (for fully-qualified enumeration E)
    </th>
    <td>
      (int) -> enum E
    </td>
    <td>
      type conversion when in int32 range, otherwise error
    </td>
  </tr>
  <tr>
    <td>
      (string) -> enum E
    </td>
    <td>
      type conversion for unqualified symbolic name, otherwise error
    </td>
  </tr>
</table>

## Appendix 1: Legacy Behavior

### Enums as Ints

In many pre-1.0 implementations, protocol buffer enums are all treated as CEL
type `int`, and are legal arguments whenever an `int` is expected. Int values
in the range `-2**31` to `2**31 - 1` can be used whenever an enum is expected.
Values outside that range will raise an error. There are no standard functions
for conversion from strings, and no conversion to `int` is needed.

```
type(google.protobuf.Field{}.kind) # was int, now google.protobuf.Field.Kind
```
