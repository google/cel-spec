# Language Definition

This page constitutes the reference for CEL. For a gentle introduction, see
[Intro](intro.md).

## Overview

In the taxonomy of programming languages, CEL is:

*   **memory-safe:** programs cannot express access to unrelated memory, such as
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
INT_LIT        ::= DIGIT+ | 0x HEXDIGIT+
UINT_LIT       ::= INT_LIT [uU]
FLOAT_LIT      ::= DIGIT* . DIGIT+ EXPONENT? | DIGIT+ EXPONENT
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
                 | \ u HEXDIGIT HEXDIGIT HEXDIGIT HEXDIGIT
                 | \ [0-3] [0-7] [0-7]
NEWLINE        ::= \r\n | \r | \n
BOOL_LIT       ::= "true" | "false"
NULL_LIT       ::= "null"
RESERVED       ::= BOOL_LIT | NULL_LIT | "in"
                 | "for" | "if" | "function" | "return" | "void"
                 | "import" | "package" | "as" | "let" | "const"
WHITESPACE     ::= [\t\n\f\r ]+
COMMENT        ::= '//' ~NEWLINE* NEWLINE
```

Note that negative numbers are recognized as positives with the unary `-`
operator from the grammar. For the sake of a readable representation, the escape
sequences (`ESCAPE`) are kept implicit in string tokens. The meaning is that in
strings without the `r` or `R` (raw) prefix, `ESCAPE` is processed, whereas in
strings with they stay uninterpreted. See documentation of string literals
below.

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

A CEL expression is parsed in the scope of a specific protocol buffer package
or message, which controls the interpretation of names. The scope is set by
the application context of an expression. A CEL expression can contain simple
names as in `a` or qualified names as in `a.b`. The meaning of such expressions
is a combination of one or more of:

*   Variables and Functions: some simple names refer to variables in the
    execution context, standard functions, or other name bindings provided by
    the CEL application.
*   Field selection: appending a period and identifier to an expression could
    indicate that we're accessing a field within a protocol buffer or map.
*   Protocol buffer package names: a simple or qualified name could represent an
    absolute or relative name in the protocol buffer package namespace. Package
    names must be followed by a message type or enum constant.
*   Protocol buffer message types and enum constants: following an optional
    protocol buffer package name, a simple or qualified name could refer to a
    message type or enum constant in the package's namespace.

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
`string`      | Strings of UTF-8 code points
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
any two not-a-number values will compare unequal even if their underlying
properties are different.

### String and Bytes Values

Strings are sequences of Unicode code points. Bytes are sequences of octets
(eight-bit data).

Quoted string literals are delimited by either single- or double-quote
characters, where the closing delimiter must match the opening one, and can
contain any unescaped character except the delimiter or newlines (either CR or
LF).

Triple-quoted string literals are delimited by three single-quotes or three
double-quotes, and may contain any unescaped characters except for the delimiter
sequence.  Again, the closing delimiter must match the opening one. Triple-quoted
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

Escape sequences are a backslash (`\`) followed by one of the following:

*   A single-quote(`'`), double-quote(`"`), or backslash(`\`), representing
    itself.
*   A code for whitespace:
    *   `b`: backspace
    *   `f`: form feed
    *   `n`: line feed
    *   `r`: carriage return
    *   `t`: horizontal tab
*   A `u` followed by four hexadecimal characters, encoding a Unicode code point
    in the BMP. Characters in other Unicode planes can be represented with
    surrogate pairs.
*   Three octal digits, in the range `000` to `377`. For strings, it denotes the
    unicode code point. For bytes, it represents an octet value.

Examples:

CEL Literal  | Meaning
------------ | --------------------------------------------------------------
`""`         | Empty string
`'""'`       | String of two double-quote characters
`'''x''x'''` | String of four characters "x''x"
`"\""`       | String of one double-quote character
`"\\"`       | String of one backslash character
`r"\\"`      | String of two backslash characters
`b"abc"`     | Byte sequence of 61, 62, 63
`b"&yuml;"`  | Sequence of bytes 195 and 191 (UTF-8 of &yuml;)
`b"\303\277"`| Also sequence of bytes 195 and 191
`"\303\277"` | String of "&Atilde;&iquest;" (code points 195, 191)
`"\377"`     | String of "&yuml;" (code point 255)
`b"\377"`    | Sequence of byte 255 (_not_ UTF-8 of &yuml;)

### Aggregate Values

Lists are ordered sequences of values.

Maps are a set of key values, and a mapping from these keys to arbitrary values.
Key values must be an allowed key type: `int`, `uint`, `bool`, or `string`.
Thus maps are the union of what's allowed in protocol buffer maps and JSON
objects.

Note that the type checker uses a finer-grained notion of list and map types.
Lists are `list(A)` for the homogenous type `A` of list elements. Maps are
`map(K, V)` for maps with keys of type `K` and values of type `V`.  The type
`dyn` is used for heterogeneous values  See
[Gradual Type Checking](#gradual-type-checking). But these constraints are
only enforced within the type checker; at runtime, lists and maps can have
heterogeneous types.

Any protocol buffer message is a CEL value, and each message type is its own CEL
type, represented as its fully-qualified name.

A list can be denoted by the expression `[e1, e2, ..., eN]`, a map by `{ek1:
ev1, ek2: ev2, ..., ekN: evN}`, and a message by `M{f1: e1, f2: e2, ..., fN: eN}`,
where `M` must be a simple or qualified name which resolves to a message type
(see [Name Resolution](#name-resolution)). For a map, the entry keys are
sub-expressions that must evaluate to values of an allowed type (`int`, `uint`,
`bool`, or `string`).  For a message, the field names are identifiers.  It is
an error to have duplicate keys or field names. The empty list, map, and message
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

Every value in CEL has a runtime type which is a value by itself. The standard
function `type(x)` returns the type of expression `x`.

As types are values, those values (`int`, `string`, etc.) also have a type: the
type `type`, which is an expression by itself which in turn also has type
`type`. So

*   `type(1)` evaluates to `int`
*   `type("a")` evaluates to `string`
*   `type(1) == string` evaluates to `false`
*   `type(type(1)) == type(string)` evaluates to `true`

### Abstract Types

A CEL implementation can add new types to the language.  These types will be
given names in the same namespace as the other types, but will have no special 
upport in the language syntax.  The only way to construct or use values of these
abstract types is through functions which the implementor will also provide.

Commonly, an abstract type will have a representation as a protocol buffer, so
that it can be stored or transmitted across a network.  In this case, the abstract
type will be given the same name as the protocol buffer, which will prevent CEL
programs from being able to use that particular protocol buffer message type;
they will not be able to construct values of that type by message expressions
nor access the message fields.  The abstract type remains abstract.

By default, CEL uses `google.protobuf.Timestamp` and `google.protobuf.Duration`
as abstract types.  The standard functions provide ways to construct and manipulate
these values, but CEL programs cannot construct them with message expressions
or access their message fields.

### Protocol Buffer Data Conversion

Protocol buffers have a richer range of types than CEL, so Protocol buffer data
is converted to CEL data when read from a message field, and CEL data is
converted in the other direction when initializing a field. In general, protocol
buffer data can be converted to CEL without error, but range errors are possible
in the other direction.

| Protocol Buffer Type             | CEL Type                               |
| -------------------------------- | -------------------------------------- |
| int32, int64, sint32, sint64, sfixed32, sfixed64    | `int`               |
| uint32, uint64, fixed32, fixed64 | `uint`                                 |
| float, double                    | `double`                               |
| bool, string, bytes              | same                                   |
| enum                             | `int`                                  |
| repeated                         | `list`                                 |
| map<K, V>                        | `map`                                  |
| oneof                 | options expanded individually, at most one is set |
| message                          | same                                   |

Signed integers, unsigned integers, and floating point numbers are converted to
the singular CEL type of the same sort. The CEL type is capable of expressing
the full range of protocol buffer values. When converting from CEL to protocol
buffers, an out-of-range CEL value results in an error.

Boolean, string, and bytes types have identical ranges and are converted without
error.

Protocol buffer enum values are converted to the corresponding `int` value.
Protocol buffer enum fields can accept any signed 32-bit number, values outside
that range will raise an error.

Repeated fields are converted to CEL lists of converted values, preserving the
order. In the other direction, the CEL list elements must be of the right type
and value to be converted to the corresponding protocol buffer type. Similarly,
protocol buffer maps are converted to CEL maps, and CEL map keys and values must
have the right type and value to be converted in the other direction.

Oneof fields are represented by the translation of each of their options as a
separate field, but at most one of these fields will be "set", as detected by
the `has()` macro.  See [Macros](#macros).

Since protocol buffer messages are first-class CEL values, message-valued fields
are used without conversion.

Every protocol buffer field has a default value, and there is no semantic
difference between a field set to this default value, and an unset field. For
message fields, there default value is just the unset state, and an unset
message field is distinct from one set to an empty (i.e. all-unset) message.

The `has()` macro (see [Macros](#macros)) tells whether a message field is set
(i.e. not unset, hence not set to the default value). If an unset field is
nevertheless selected, it evaluates to its default value, or if it is a message
field, it evaluates to an empty (i.e. all-unset) message. This allows expressions
to use iterative field selection to examine the state of fields in deeply nested
messages without needing to test whether every intermediate field is set. (See
exception for wrapper types, below.)

### Dynamic Values

CEL automatically converts certain protocol buffer messages in the
google.protobuf package to other types.

| google.protobuf message | CEL Conversion                                     |
| ----------------------- | -------------------------------------------------- |
| `Any`        | dynamically converted to the contained message type, or error |
| `ListValue`             | list of `Value` messages                           |
| `Struct`                | map (with string keys, `Value` values)             |
| `Value` | dynamically converted to the contained type (null, double, string, bool, `Struct`, or `ListValue`) |
| wrapper types           | converted to eponymous type                        |

The wrapper types are `BoolValue`, `BytesValue`, `DoubleValue`, `EnumValue`, `FloatValue`,
`Int32Value`, `Int64Value`, `NullValue`, `StringValue`, `Uint32Value`, and `Uint64Value`.
Values of these wrapper types are converted to the obvious type. Additionally,
field selection of an unset message field of wrapper type will evaluate to
`null`, instead of the default message. This is an exception to the usual
evaluation of unset message fields.

Note that this implies some cascading conversions. An `Any` message might be
converted to a `Struct`, one of whose `Value`-typed values might be converted to a
`ListValue` of more values, and so on.

Also note that all of these conversions are dynamic at runtime, so CEL's static
type analysis cannot avoid the possibility of type-related errors in expressions
using these dynamic values.

## Gradual Type Checking

CEL is a dynamically-typed language, meaning that the types of the values of the
variables and expressions might not be known until runtime. However, CEL has an
optional type-checking phase that takes annotation giving the types of all
variables and tries to deduce the type of the expression and of all its
sub-expressions. This is not always possible, due to the dynamic expansion of
certain messages like `Struct`, `Value`, and `Any` (see [Dynamic Values](#dynamic-values)).
However, if a CEL program does not use dynamically-expanded messages, it can be
statically type-checked.

The type checker uses a richer type system than the types of the dynamic values:
lists have a type parameter for the type of the elements, and maps have two
parameters for the types of keys and values, respectively. These richer types
preserve the stronger type guarantees that protocol buffer messages have. We can
infer stronger types from the standard functions, such as accessing list
elements or map fields. However, the `type()` function and dynamic dispatch to
particular function overloads only use the coarser types of the dynamic values.

The type checker also introduces the `dyn` type, which is the union of all other
types. Therefore the type checker could accept a list of heterogeneous values as
`dyn([1, 3.14, "foo"])`, which is tiven the type `list(dyn)`. The standard
function `dyn` has no effect at runtime, but signals to the type checker that its
argument should be considered of type `dyn`, `list(dyn)`, or a `dyn`-valued map.

A CEL type checker attempts to identify occurrences of `no_matching_overload`
and `no_such_field` runtime errors (see [Runtime Errors](#runtime-errors))
ahead of runtime. It also serves to optimize execution speed, by narrowing
down the number of possible matching overloads for a function call, and by
allowing for a more efficient (unboxed) runtime representation of values.

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

*   **Literals:** the various kinds of literals (numbers, booleans, strings, bytes,
    and `null`) evaluate to the values they represent.
*   **Variables:** variables are looked up in the binding environment. An unbound
    variable evaluates to an error.
*   **List, Map, and Message expressions:** each sub-expression is evaluated and if
    any sub-expression results in an error, this expression results in an error.
    Otherwise, it results in the list, map, or message of the sub-expression
    results, or an error if one of the values is of the wrong type.
*   **Field selection:** see [Field Selection](#field-selection).
*   **Macros:** see [Macros](#macros).
*   **Logical operators:** see [Logical Operators](#logical-operators).
*   **Other operators:** operators are translated into specially-named functions and
    the sub-expressions become their arguments, for instance `e1 + e2` becomes
    `_+_(e1, e2)`, which is then evaluated as a normal function.
*   **Normal functions:** all argument sub-expressions are evaluated and if any
    results in an error, then this expression results in an error. Otherwise,
    the function is identified by its name and dispatched to a particular
    overload based on the types of the sub-expression values. See [Functions](#functions).

Because CEL is free of side-effects, the order of evaluation among
sub-expressions is not guaranteed.  If multiple subexpressions would evaluate to errors
causing the enclosing expression to evaluate to an error, it will propagate one or more
of the sub-expression erors, but it is not specified which ones.

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
terminated; exceptions to this rule are discussed in [Logical Operators](#logical-operators)
and [Macros](#macros).

CEL provides the following built-in runtime errors:

*   `no_matching_overload`: this function has no overload for the types of the
    arguments.
*   `no_such_field`: a map or message does not contain the desired field.

There is no in-language representation of errors, no generic way to raise them,
and no way to catch or bypass errors, except for the short-circuiting behavior
of the logical operators, and macros.

### Logical Operators

In the conditional operator `e ? e1 : e2`, evaluates to `e1` if `e`
evaluates to `true`, and `e2` if `e` evaluates to `false`.  The untaken branch
is presumed to not be executed, though that is an implementation detail.

In the boolean operators `&&` and `||`: if any of their operands uniquely
determines the result (`false` for `&&` and `true` for `||`) the other operand
may or may not be evaluated, and if that evaluation produces a runtime error, it
will be ignored. This makes those operators commutative (in contrast to
traditional boolean short-circuit operators). The rationale for this behavior is
to allow the boolean operators to be mapped to indexed queries, and align better
with SQL semantics.

### Macros

CEL supports a small set of predefined macros. Macro invocations have the same
syntax as function calls, but follow different type checking rules and runtime
semantics than regular functions. An application of CEL opts-in to which macros
to support, selecting from the predefined set of macros. The currently available
macros are:

*   `has(e.f)`: tests whether a field is available. See "Field Selection" below.
*   `e.all(x, p),`: tests where a predicate holds for all elements of a list `e`
    or keys of a map `e`. Here `x` is a simple identifier to be used in `p`
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

### Field Selection

A field selection expression, `e.f`, can be applied both to messages and to
maps. For maps, selection is interpreted as the field being a string key.

The semantics depends on the type of `e`:

1.  If `e` is a message and `f` is not declared in this message, the runtime
    error `no_such_field` is raised.
2.  If `e` is a message and `f` is declared, but the field is not set, the
    default value of the field's type will be produced. Note that this is `null`
    for messages or the according primitive default value as determined by
    proto2 or proto3 semantics.
3.  If `e` is a map and `f` is not present in the map, a runtime error will be
    produced. (Note the runtime error is not a well-known one like
    `no_such_field` but implementation dependent.) It holds that `e.f ==
    e['f']`.

To test for the presence of a field, the macro `has(e.f)` can be used.
`has(e.f)` behaves similar as `e.f`, except as where the former would produce
`null` or an error different than `no_such_field`, it will return false, and
true otherwise. This means means that `has(e.f)` applied to a message which does
not declare field `f` produces a `no_such_field` error, where it produces false
if `f` is declared but not set (or, in proto3, has its default value). Moreover,
`has(e.f)` where `e` is a map returns false if `f` is not defined in the map.

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
runtime error `no_matching_overload` is raised (see also [Runtime
Errors](#errors)). For example, the standard function `size` is specified by the
following overloads:

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

TODO https://issuetracker.google.com/67014381 : have better descriptions. The
table is auto-generated so the descriptions need to be updated in the code.

<!-- BEGIN GENERATED DECL TABLE; DO NOT EDIT BELOW -->

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
    <th rowspan="1">
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
    <th rowspan="6">
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
      string concatenation
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
      list concatenation
    </td>
  </tr>
  <tr>
    <th rowspan="3">
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
    <th rowspan="6">
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
    <th rowspan="6">
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
    <th rowspan="6">
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
    <th rowspan="6">
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
    <th rowspan="1">
      _?_:_
    </th>
    <td>
      (bool, A, A) -> A
    </td>
    <td>
      conditional
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
      list indexing
    </td>
  </tr>
  <tr>
    <td>
      (map(A, B), A) -> B
    </td>
    <td>
      map indexing
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      _in_
    </th>
    <td>
      (A, list(A)) -> bool
    </td>
    <td>
      list membership
    </td>
  </tr>
  <tr>
    <td>
      (A, map(A, B)) -> bool
    </td>
    <td>
      map key membership
    </td>
  </tr>
  <tr>
    <th rowspan="1">
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
      dyn
    </th>
    <td>
      (A) -> dyn
    </td>
    <td>
      type conversion
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
      get day of month from the date with timezone, zero-based indexing
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
      get milliseconds from the date in UTC, 0-999
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
    <th rowspan="2">
      in
    </th>
    <td>
      (A, list(A)) -> bool
    </td>
    <td>
      list membership
    </td>
  </tr>
  <tr>
    <td>
      (A, map(A, B)) -> bool
    </td>
    <td>
      map key membership
    </td>
  </tr>
  <tr>
    <th rowspan="4">
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
      type conversion
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
    <th rowspan="2">
      list
    </th>
    <td>
      (type(A)) -> type(list(A))
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (type(A), list(A)) -> list(A)
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="2">
      map
    </th>
    <td>
      (type(A), type(B)) -> type(map(A, B))
    </td>
    <td>
      type denotation
    </td>
  </tr>
  <tr>
    <td>
      (type(A), type(B), map(A, B)) -> map(A, B)
    </td>
    <td>
      type conversion
    </td>
  </tr>
  <tr>
    <th rowspan="1">
      matches
    </th>
    <td>
      (string, string) -> bool
    </td>
    <td>
      matches second argument against regular expression in first argument
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
  <tr>
    <th rowspan="5">
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
    <th rowspan="1">
      type
    </th>
    <td>
      (A) -> type(A)
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
      type conversion
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
</table>

<!-- END GENERATED DECL TABLE; DO NOT EDIT ABOVE -->
