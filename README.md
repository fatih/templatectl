# templatectl
Simple templating CLI

## Install

```
GO111MODULE=on go get github.com/fatih/templatectl/cmd/templatectl@latest
```

## Usage


```sh
# By default templatectl prints to stdout
echo 'This is {{ env "ENV_FOO" }}' > input.tmpl
export ENV_FOO="foo"
$ templatectl --input input.tmpl
This is foo

# Or output to a file
$ templatectl --input input.tmpl --output templated.txt
$ cat templated.txt
This is foo
```

