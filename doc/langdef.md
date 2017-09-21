# Language Definition

This page constitutes the reference for CEL. For a gentle introduction, see
[Intro](intro.md).

## Syntax

The grammar of CEL is defined below. We use `T?` for an optional token, but not
other meta-notation beyond BNF.

```grammar
Expr           ::= Conditional
Conditional    ::= LogicalOr
                 | LogicalOr '?' Conditional ':' LogicalOr
LogicalOr      ::= LogicalAnd
                 | LogicalOr '||' LogicalAnd
LogicalAnd     ::= Equality
                 | LogicalAnd '&&' Equality
Equality       ::= Relation
                 | Relation ( '==' | '!=') Relation
Relation       ::= Addition
                 | Addition ( '<' | '<=' | '>' | '>=' ) Addition
Addition       ::= Multiplication
                 | Addition ( '+' | '-' ) Multiplication
Multiplication ::= Primary
                 | Multiplication ( '*' | '/' | '%' ) Primary
Primary        ::= '!' Primary
                 | '-' Primary
                 | Primary '.' IDENT
                 | Primary '.' IDENT '(' ')' | Primary '.' IDENT '(' ExprList ')'
                 | Primary '[' Expr ']'
                 | QualifiedIdent '{' '}' | QualifiedIdent '{' FieldInits ','? '}'
                 | '{' '}' | '{' MapInits ','? '}'
                 | '[' ']' | '[' ExprList ','? ']'
                 | '(' Expr ')'
                 | '.'? IDENT
                 | LITERAL
ExprList       ::= Expr | Expr ',' ExprList
FieldInits     ::= IDENT ':' Expr | FieldInits ',' IDENT ':' Expr
MapInits       ::= Expr ':' Expr | MapInits ',' Expr ':' Expr
QualifiedIdent ::= '.'? IDENT | QualifiedIdent '.' IDENT
```

The lexis is defined below. A form of regular expressions which are white-space
insensitive is used:

```
IDENT          ::= [_a-zA-Z][_a-zA-Z0-9]*
LITERAL        ::= INT_LIT | UINT_LIT | FLOAT_LIT | STRING_LIT | BYTES_LIT
                 | DURATION_LIT | TIMESTAMP_LIT
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
                 | \ [0-7] [0-7] [0-7]
NEWLINE        ::= \r\n | \r | \n
DURATION_LIT   ::= FLOAT_LIT [sS]
TIMESTAMP_LIT  ::= [tT] STRING_LIT
```

For the sake of a readable representation, the escape sequences (`ESCAPE`) are
kept implicit in string tokens. The meaning is that in strings without the `r`
or `R` (raw) prefix, `ESCAPE` is processed, whereas in strings with they stay
uninterpreted.

## Values

Values in CEL represent any of the following:

1.  *Simple Values*, as they come from protocol buffers: booleans, integers,
    floating-point numbers, strings, byte strings, and messages. Numbers are
    normalized to the highest precision supported by protobuf, i.e. int32 as
    int64, uint32 as uint64, and float as double. Enum values are represented as
    integers. In addition, the well-known types `Duration` and `Timestamp` are
    not treated as messages, but as primitive values of the given type.
2.  *Null Value*, a special value to represent an absent value, for example, the
    default value of messages.
3.  *Aggregate Values*, as they are implied by `repeated` and `map` in protobuf
    messages: lists of the above values, and maps of the above values, with the
    restriction on allowed map keys implied by protobuf.
4.  *Type Values*. See [Types](#types) below.
5.  *Function Values*. See [Functions](#functions) below.

### Dynamic Values

Newer versions of protobuf support the well-known types `Struct`, `Value`,
`Any`, and various wrapper messages for primitives like `Int32Value`. Those
types have no equivalent in CEL, but instead are automatically converted to one
of the basic representations introduced in [Values](#values) above as follows:

1.  A `Struct` value is converted to an equivalent of `map<String, Value>`.
2.  A `ListValue` is converted to an equivalent of `repeated Value`.
3.  A `Value` union is converted to the according value case: double, string,
    boolean, map, list, or null.
4.  An `Any` value is on-the-fly deserialized and converted to the message it
    represents.
5.  A wrapper is automatically converted into its underlying value, or `null` if
    the wrapper message is not set.

NOTE(go/api-expr-open): heterogenous vs homogeneous lists and maps, as well as
treatment of Any.

## Types

Every value in CEL has a runtime type which is a value by itself. The standard
function `type(x)` denotes this type.

As types are values, denotations of aggregate types are regular function
applications. For example, `map(int, string)` denotes the type of maps from
`int` to `string`, where `map` is simply a function applied to the values for
types `int` and `string`.

As types are values, those values also have a type: the expression `int` has
type `type(int)`, which is an expression by itself which in turn has type
`type(type(int))`. Types of types have no actual use in CEL, but are included
for consistency as excluding them would be more complicated than otherwise.

## Functions

A function value represents a computation which delivers a new value based on
zero or more arguments. Function applications have no observable side-effects
(there maybe side-effects like logging or such which are not observable from
CEL). The default argument evaluation strategy for functions is strict, with
exceptions from this rule discussed in [Argument Evaluation](#arg-eval).

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

Operators like `x + y` are considered syntactic sugar for function calls, using
symbols like `_+_` for the function name.

See [Standard Operators and Functions](#standard) for the list of all predefined
functions and operators.

### Extension Functions

It is possible to add extension functions to CEL, which then behave in no way
different than standard functions. The mechanism how to do this is
implementation dependent and usually highly curated. For example, an application
domain of CEL can add a new overload to the `size` function above, provided this
overload's argument types do not overlap with any existing overload. For
methodological reasons, CEL disallows to add overloads to operators.

### Macros

CEL supports a small set of predefined macros. Macro invocations have the same
syntax as function calls, but follow different type checking rules and runtime
semantics than regular functions. An application of CEL opts-in to which macros
to support, selecting from the predefined set of macros. The currently available
macros are:

-   `has(e.f)`: tests whether a field is available. See [discussion of the
    select](#select) expression.
-   `e.all(x, p), e.exists(x, p), e.exists_one(x,p)`: tests where a predicate
    holds for all/at least one/exactly one element of a list `e` or keys of a
    map `e`. Here `x` is a simple identifier to be used in `p` which binds to
    the element or key.

### Argument Evaluation

By default, all arguments to function calls are strictly evaluated
(call-by-value). Any errors produced during argument evaluation are propagated.
Because CEL is free of side-effects, evaluation order is not guaranteed.

One exception is the conditional operator `e ? e1 : e2`: as common, `e1` is only
executed if `e` evaluates to true, and `e2` if `e` evaluates to false.

Another exception are the boolean operators `&&` and `||`: if any of their
operands uniquely determines the result (`false` for `&&` and `true` for `||`)
the other operand may or may not be evaluated, and if that evaluation produces a
runtime error, it will be ignored. This makes those operators commutative (in
contrast to traditional boolean short-circuit operators). The rationale for this
behavior is to allow the boolean operators to be mapped to indexed queries, and
align better with SQL semantics.

### Receiver Call Style

A function overload can be declared to use receiver call-style, so it must be
called as `e1.f(e2)` instead of `f(e1, e2)`. Overloads with different call
styles are non-overlapping per definition, regardless of their types.

## Field Selection

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

## Aggregates

A list can be denoted by the expression `[e1,e2]`, a map by `{e11:e12,e21:e22}`.
Message aggregates are also supported, and take the form `M{f1:e1, f2:e2}`,
where `M` must be a simple or qualified name which resolves to a messsage type
(see [Name Resolution](#resolution)).

NOTE(go/api-expr-open): heterogenous vs homogeneous lists and maps.

## Runtime Errors

In general, when a runtime error is produced, expression evaluation is
terminated; exceptions to this rule are discussed in [Argument
Evaluation](#arg-eval).

CEL provides two built-in runtime errors; implementations may add more. The
built-in errors are `no_matching_overload` and `no_such_field`; those errors
correspond to type errors in static typing (see [Gradual Type
Checking](#type-checking)).

## Name Resolution

A CEL expression can contain simple names as in `a` or qualified names as in
`a.b`. Such names are resolved in the lexical scope of a protobuf package or
message declaration, following the same rules as protobuf itself (which in turn
follows the C++ rules). The scope is set by the application context of an
expression.

All message and enum constant declarations from the protobuf scope are available
to CEL expressions. In addition, all [Standard Definitions](#standard) are
injected into the root scope. Applications of CEL may inject additional name
bindings in arbitrary scopes.

Resolution works as follows. If `a.b` is a name to be resolved in the context of
a protobuf declaration with scope `A.B`, then resolution is attempted, in order,
as `A.B.a.b`, `A.a.b`, and finally `a.b`. To override this behavior, one can use
`.a.b`; this name will only be attempted to be resolved in the root scope, i.e.
as `a.b`.

If name qualification is mixed with field selection, the longest prefix of the
name which resolves in the current lexical scope is used. For example, if
`a.b.c` resolves to a message declaration, and `a.b` does so as well with `c` a
possible field selection, then `a.b.c` takes priority over the interpretation
`(a.b).c`.

## Standard Definitions

All predefined operators, functions and constants are listed in the table below.
For each symbol, the available overloads are listed. Operator symbols use a
notation like `_+_` where `_` is a placeholder for an argument.

TODO https://issuetracker.google.com/67014381 : have better descriptions. The table is auto-generated so the
descriptions need to be updated in the code.

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

## Gradual Type Checking

A CEL type checker attempts to identify occurrences of `no_matching_overload`
and `no_such_field` runtime errors ahead of runtime. It also serves to optimize
execution speed, by narrowing down the number of possible matching overloads for
a function call, and by allowing for a more efficient (unboxed) runtime
representation of values.

By construction, a CEL expression that does not use the dynamic features coming
from `Struct`, `Value`, or `Any`, can be fully statically type checked and all
overloads can be resolved ahead of runtime.

If a CEL expression uses a mixture of dynamic and static features, a type
checker will still attempt to derive as much information as possible and
delegate undecidable type decisions to runtime.
