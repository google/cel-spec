name: "go-0.1.0"
decl: <
  name: "int"
  ident: <
    type: <
      type: <
        primitive: INT64
      >
    >
  >
>
decl: <
  name: "uint"
  ident: <
    type: <
      type: <
        primitive: UINT64
      >
    >
  >
>
decl: <
  name: "bool"
  ident: <
    type: <
      type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "double"
  ident: <
    type: <
      type: <
        primitive: DOUBLE
      >
    >
  >
>
decl: <
  name: "bytes"
  ident: <
    type: <
      type: <
        primitive: BYTES
      >
    >
  >
>
decl: <
  name: "string"
  ident: <
    type: <
      type: <
        primitive: STRING
      >
    >
  >
>
decl: <
  name: "list"
  ident: <
    type: <
      type: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
    >
  >
>
decl: <
  name: "map"
  ident: <
    type: <
      type: <
        map_type: <
          key_type: <
            type_param: "A"
          >
          value_type: <
            type_param: "B"
          >
        >
      >
    >
  >
>
decl: <
  name: "null_type"
  ident: <
    type: <
      type: <
        null: NULL_VALUE
      >
    >
  >
>
decl: <
  name: "type"
  ident: <
    type: <
      type: <
	type: <>
      >
    >
  >
>
decl: <
  name: "_?_:_"
  function: <
    overloads: <
      overload_id: "conditional"
      params: <
        primitive: BOOL
      >
      params: <
        type_param: "A"
      >
      params: <
        type_param: "A"
      >
      type_params: "A"
      result_type: <
        type_param: "A"
      >
    >
  >
>
decl: <
  name: "_&&_"
  function: <
    overloads: <
      overload_id: "logical_and"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_||_"
  function: <
    overloads: <
      overload_id: "logical_or"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "!_"
  function: <
    overloads: <
      overload_id: "logical_not"
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_<_"
  function: <
    overloads: <
      overload_id: "less_bool"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_bytes"
      params: <
        primitive: BYTES
      >
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_<=_"
  function: <
    overloads: <
      overload_id: "less_equals_bool"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_bytes"
      params: <
        primitive: BYTES
      >
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "less_equals_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_>_"
  function: <
    overloads: <
      overload_id: "greater_bool"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_bytes"
      params: <
        primitive: BYTES
      >
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_>=_"
  function: <
    overloads: <
      overload_id: "greater_equals_bool"
      params: <
        primitive: BOOL
      >
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_bytes"
      params: <
        primitive: BYTES
      >
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "greater_equals_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_==_"
  function: <
    overloads: <
      overload_id: "equals"
      params: <
        type_param: "A"
      >
      params: <
        type_param: "A"
      >
      type_params: "A"
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_!=_"
  function: <
    overloads: <
      overload_id: "not_equals"
      params: <
        type_param: "A"
      >
      params: <
        type_param: "A"
      >
      type_params: "A"
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "_-_"
  function: <
    overloads: <
      overload_id: "subtract_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "subtract_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "subtract_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
    overloads: <
      overload_id: "subtract_timestamp_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        well_known: DURATION
      >
    >
    overloads: <
      overload_id: "subtract_timestamp_duration"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
    overloads: <
      overload_id: "subtract_duration_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        well_known: DURATION
      >
    >
  >
>
decl: <
  name: "_*_"
  function: <
    overloads: <
      overload_id: "multiply_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "multiply_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "multiply_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
  >
>
decl: <
  name: "_/_"
  function: <
    overloads: <
      overload_id: "divide_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "divide_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "divide_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
  >
>
decl: <
  name: "_%_"
  function: <
    overloads: <
      overload_id: "modulo_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "modulo_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
  >
>
decl: <
  name: "_+_"
  function: <
    overloads: <
      overload_id: "add_int64"
      params: <
        primitive: INT64
      >
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "add_uint64"
      params: <
        primitive: UINT64
      >
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "add_double"
      params: <
        primitive: DOUBLE
      >
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
    overloads: <
      overload_id: "add_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "add_bytes"
      params: <
        primitive: BYTES
      >
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BYTES
      >
    >
    overloads: <
      overload_id: "add_list"
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      type_params: "A"
      result_type: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
    >
    overloads: <
      overload_id: "add_timestamp_duration"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
    overloads: <
      overload_id: "add_duration_timestamp"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
    overloads: <
      overload_id: "add_duration_duration"
      params: <
        well_known: DURATION
      >
      params: <
        well_known: DURATION
      >
      result_type: <
        well_known: DURATION
      >
    >
  >
>
decl: <
  name: "-_"
  function: <
    overloads: <
      overload_id: "negate_int64"
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "negate_double"
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
  >
>
decl: <
  name: "_[_]"
  function: <
    overloads: <
      overload_id: "index_list"
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      params: <
        primitive: INT64
      >
      type_params: "A"
      result_type: <
        type_param: "A"
      >
    >
    overloads: <
      overload_id: "index_map"
      params: <
        map_type: <
          key_type: <
            type_param: "A"
          >
          value_type: <
            type_param: "B"
          >
        >
      >
      params: <
        type_param: "A"
      >
      type_params: "A"
      type_params: "B"
      result_type: <
        type_param: "B"
      >
    >
  >
>
decl: <
  name: "size"
  function: <
    overloads: <
      overload_id: "string_size"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "bytes_size"
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "list_size"
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      type_params: "A"
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "map_size"
      params: <
        map_type: <
          key_type: <
            type_param: "A"
          >
          value_type: <
            type_param: "B"
          >
        >
      >
      type_params: "A"
      type_params: "B"
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "size_string"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "size_bytes"
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "size_list"
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      type_params: "A"
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "size_map"
      params: <
        map_type: <
          key_type: <
            type_param: "A"
          >
          value_type: <
            type_param: "B"
          >
        >
      >
      type_params: "A"
      type_params: "B"
      result_type: <
        primitive: INT64
      >
    >
  >
>
decl: <
  name: "@in"
  function: <
    overloads: <
      overload_id: "in_list"
      params: <
        type_param: "A"
      >
      params: <
        list_type: <
          elem_type: <
            type_param: "A"
          >
        >
      >
      type_params: "A"
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "in_map"
      params: <
        type_param: "A"
      >
      params: <
        map_type: <
          key_type: <
            type_param: "A"
          >
          value_type: <
            type_param: "B"
          >
        >
      >
      type_params: "A"
      type_params: "B"
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "type"
  function: <
    overloads: <
      overload_id: "type"
      params: <
        type_param: "A"
      >
      type_params: "A"
      result_type: <
        type: <
          type_param: "A"
        >
      >
    >
  >
>
decl: <
  name: "int"
  function: <
    overloads: <
      overload_id: "int64_to_int64"
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "uint64_to_int64"
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "double_to_int64"
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "string_to_int64"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "timestamp_to_int64"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
    >
    overloads: <
      overload_id: "duration_to_int64"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: INT64
      >
    >
  >
>
decl: <
  name: "uint"
  function: <
    overloads: <
      overload_id: "uint64_to_uint64"
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "int64_to_uint64"
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "double_to_uint64"
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: UINT64
      >
    >
    overloads: <
      overload_id: "string_to_uint64"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: UINT64
      >
    >
  >
>
decl: <
  name: "double"
  function: <
    overloads: <
      overload_id: "double_to_double"
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: DOUBLE
      >
    >
    overloads: <
      overload_id: "int64_to_double"
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: DOUBLE
      >
    >
    overloads: <
      overload_id: "uint64_to_double"
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: DOUBLE
      >
    >
    overloads: <
      overload_id: "string_to_double"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: DOUBLE
      >
    >
  >
>
decl: <
  name: "bool"
  function: <
    overloads: <
      overload_id: "bool_to_bool"
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: BOOL
      >
    >
    overloads: <
      overload_id: "string_to_bool"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
    >
  >
>
decl: <
  name: "string"
  function: <
    overloads: <
      overload_id: "string_to_string"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "bool_to_string"
      params: <
        primitive: BOOL
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "int64_to_string"
      params: <
        primitive: INT64
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "uint64_to_string"
      params: <
        primitive: UINT64
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "double_to_string"
      params: <
        primitive: DOUBLE
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "bytes_to_string"
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "timestamp_to_string"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: STRING
      >
    >
    overloads: <
      overload_id: "duration_to_string"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: STRING
      >
    >
  >
>
decl: <
  name: "bytes"
  function: <
    overloads: <
      overload_id: "bytes_to_bytes"
      params: <
        primitive: BYTES
      >
      result_type: <
        primitive: BYTES
      >
    >
    overloads: <
      overload_id: "string_to_bytes"
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BYTES
      >
    >
  >
>
decl: <
  name: "timestamp"
  function: <
    overloads: <
      overload_id: "timestamp_to_timestamp"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
    overloads: <
      overload_id: "string_to_timestamp"
      params: <
        primitive: STRING
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
    overloads: <
      overload_id: "int64_to_timestamp"
      params: <
        primitive: INT64
      >
      result_type: <
        well_known: TIMESTAMP
      >
    >
  >
>
decl: <
  name: "duration"
  function: <
    overloads: <
      overload_id: "duration_to_duration"
      params: <
        well_known: DURATION
      >
      result_type: <
        well_known: DURATION
      >
    >
    overloads: <
      overload_id: "string_to_duration"
      params: <
        primitive: STRING
      >
      result_type: <
        well_known: DURATION
      >
    >
    overloads: <
      overload_id: "int64_to_duration"
      params: <
        primitive: INT64
      >
      result_type: <
        well_known: DURATION
      >
    >
  >
>
decl: <
  name: "dyn"
  function: <
    overloads: <
      overload_id: "to_dyn"
      params: <
        type_param: "A"
      >
      type_params: "A"
      result_type: <
        dyn: <
        >
      >
    >
  >
>
decl: <
  name: "contains"
  function: <
    overloads: <
      overload_id: "contains_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "endsWith"
  function: <
    overloads: <
      overload_id: "ends_with_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "matches"
  function: <
    overloads: <
      overload_id: "matches_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "startsWith"
  function: <
    overloads: <
      overload_id: "starts_with_string"
      params: <
        primitive: STRING
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: BOOL
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getFullYear"
  function: <
    overloads: <
      overload_id: "timestamp_to_year"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_year_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getMonth"
  function: <
    overloads: <
      overload_id: "timestamp_to_month"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_month_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getDayOfYear"
  function: <
    overloads: <
      overload_id: "timestamp_to_day_of_year"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_day_of_year_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getDayOfMonth"
  function: <
    overloads: <
      overload_id: "timestamp_to_day_of_month"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_day_of_month_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getDate"
  function: <
    overloads: <
      overload_id: "timestamp_to_day_of_month_1_based"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_day_of_month_1_based_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getDayOfWeek"
  function: <
    overloads: <
      overload_id: "timestamp_to_day_of_week"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_day_of_week_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getHours"
  function: <
    overloads: <
      overload_id: "timestamp_to_hours"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_hours_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "duration_to_hours"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getMinutes"
  function: <
    overloads: <
      overload_id: "timestamp_to_minutes"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_minutes_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "duration_to_minutes"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getSeconds"
  function: <
    overloads: <
      overload_id: "timestamp_to_seconds"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_seconds_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "duration_to_seconds"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
decl: <
  name: "getMilliseconds"
  function: <
    overloads: <
      overload_id: "timestamp_to_milliseconds"
      params: <
        well_known: TIMESTAMP
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "timestamp_to_milliseconds_with_tz"
      params: <
        well_known: TIMESTAMP
      >
      params: <
        primitive: STRING
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
    overloads: <
      overload_id: "duration_to_milliseconds"
      params: <
        well_known: DURATION
      >
      result_type: <
        primitive: INT64
      >
      is_instance_function: true
    >
  >
>
