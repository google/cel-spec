<!--
Copyright 2025 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Strings

## string.format(list) -> string

### Format

`%[.precision]conversion`

### Precision

Optional. In the form of a period `.` followed by a required positive decimal digit sequence. The default precision is `6`. Not all conversions support precision.

### Conversion

|  Character | Precision | Description |
| --- | --- | --- |
| `s` | N | <table><tbody><tr><td><code>bool</code></td><td>The value is foramtted as <code>true</code> or <code>false</code>.</td></tr><tr><td><code>int</code></td><td>The value is formatted in base 10 with a preceding <code>-</code> if the value is negative. No insignificant <code>0</code>s must be included.</td></tr><tr><td><code>uint</code></td><td>The value is formatted in base 10. No insignificant <code>0</code>s must be included.</td></tr><tr><td><code>double</code></td><td>The value is formatted in base 10. No insignificant <code>0</code>s must be included. If there are no significant digits after the <code>.</code> then it must be excluded.</td></tr><tr><td><code>bytes</code></td><td>The value is formatted as if `string(value)` was performed and any invalid UTF-8 sequences are replaced with <code>\ufffd</code>. Multiple adjacent invalid UTF-8 sequences must be replaced with a single <code>\ufffd</code>.</td></tr><tr><td><code>string</code></td><td>The value is included as is.</td></tr><tr><td><code>duration</code></td><td>The value is formatted as decimal seconds as if the value was converted to <code>double</code> and then formatted as <code>%ds</code>.</td></tr><tr><td><code>timestamp</code></td><td>The value is formatted according to RFC 3339 and is always in UTC.</td></tr><tr><td><code>null_type</code></td><td>The value is formatted as <code>null</code>.</td></tr><tr><td><code>type</code></td><td>The value is formatted as a string.</td></tr><tr><td><code>list</code></td><td>The value is formatted as if each element was formatted as <code>"%s".format([element])</code>, joined together with <code>, </code> and enclosed with <code>[</code> and <code>]</code>.</td></tr><tr><td><code>map</code></td><td>The value is formatted as if each entry was formatted as <code>"%s: %s".format([key, value])</code>, sorted by the formatted keys in ascending order, joined together with <code>, </code>, and enclosed with <code>{</code> and <code>}</code>.</td></tr></tbody></table> |
| `d` | N | <table><tbody><tr><td><code>int</code></td><td>The value is formatted in base 10 with a preceding <code>-</code> if the value is negative. No insignificant <code>0</code>s must be included.</td></tr><tr><td><code>uint</code></td><td>The value is formatted in base 10. No insignificant <code>0</code>s must be included.</td></tr><tr><td><code>double</code></td><td>The value is formatted in base 10. No insignificant <code>0</code>s must be included. If there are no significant digits after the <code>.</code> then it must be excluded.</td></tr></tbody></table> |
| `f` | Y | `int` `uint` `double`: The value is converted to the style `[-]dddddd.dddddd` where there is at least one digit before the decimal and exactly `precision` digits after the decimal. If `precision` is 0, then the decimal is excluded. |
| `e` | Y | `int` `uint` `double`: The value is converted to the style `[-]d.dddddde±dd` where there is one digit before the decimal and `precision` digits after the decimal followed by `e`, then the plus or minus, and then two digits. |
| `x` `X` | N | Values are formatted in base 16. For `x` lowercase letters are used. For `X` uppercase letters are used.<table><tbody><tr><td><code>int</code> <code>uint</code></td><td>The value is formatted in base 16 with no insignificant digits. If the value was negative <code>-</code> is prepended.</td></tr><tr><td><code>string</code></td><td>The value is formatted as if `bytes(value)` was used to convert the <code>string</code> to <code>bytes</code> and then each byte is formatted in base 16 with exactly 2 digits.</td></tr><tr><td><code>bytes</code></td><td>The value is formatted as if each byte is formatted in base 16 with exactly 2 digits.</td></tr></tbody></table> |
| `o` | N | `int` `uint`: The value is converted to base 8 with no insignificant digits. If the value was negative `-` is prepended. |
| `b` | N | `int` `uint` `bool`: The value is converted to base 2 with no insignificant digits.  If the value was negative `-` is prepended. |

> In all cases where `double` is accepted: if the value is NaN the result is `NaN`, if the value is infinity the result is `[-]Infinity`.

### Examples

```
"%s".format(["foo"])  // foo
"%s".format([b"foo"])  // foo
"%d".format([1])      // 1
"%d".format([1u])     // 1
"%d".format([3.14])   // 3.14
"%f".format([1])      // 1.000000
"%f".format([1u])     // 1.000000
"%f".format([3.14])   // 3.140000
"%.1f".format([3.14])   // 3.1
"%e".format([1])      // 1.000000e+00
"%e".format([1u])     // 1.000000e+00
"%e".format([3.14])   // 3.140000e+00
"%.1e".format([3.14])   // 3.1e+00
"%.1e".format([-3.14])   // -3.1e+00
```
