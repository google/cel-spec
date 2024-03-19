#!/bin/sh
bazelisk build //proto/cel/expr:all
files=($(bazelisk aquery 'kind(proto, //proto/cel/expr:all)' | grep Outputs | grep "[.]pb[.]go" | sed 's/Outputs: \[//' | sed 's/\]//' | tr "," "\n"))
for src in ${files[@]};
do
  dst=$(echo $src | sed 's/\(.*\%\/github.com\/google\/cel-spec\/\(.*\)\)/\2/')
  $(cp $src ./go)
done
