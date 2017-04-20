# splat!

[![Build Status](https://travis-ci.org/eljuanchosf/splat.svg?branch=master)](https://travis-ci.org/eljuanchosf/splat)
[![codecov](https://codecov.io/gh/eljuanchosf/splat/branch/master/graph/badge.svg)](https://codecov.io/gh/eljuanchosf/splat)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d4619fde24e74c9ebed73b64c2ab31d6)](https://www.codacy.com/app/eljuanchosf/splat?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=eljuanchosf/splat&amp;utm_campaign=Badge_Grade)

splat! is a simple yet powrful way to replace values in a YAML file.

## Define your YAML

```yaml
a: Easy!
b:
  c: 2
  d: [3, 4]
    f: ((< fileContent(./myfile) >))`
```

## Run splat!

```sh
splat myfile.yml
```

splat! will run the command you defined in your YML, automatically replacing the values, in the case of the example, with the contents of the file `myfile`.

## Implemented commands

* `lookup`: search for a key and retrieve the variable
  * in files
  * in ENV variables

## Planned commands

* `fileContent`: insert the content of the file
* `certificate`: generate a self-signed certificate
* `cmd`: runs an arbitrary command in the OS and inserts the result (STDOUT) of that command.

## Suggestions?

YES, please!!
User GitHub's issues to suggest features.

More to come soon...

