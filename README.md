[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# honey-do
**Honey-Do** is a simple build / automation tool written entirely in Go. The primary goal is to be an easier and cross-platform alternative to GNU Make. The name is based on the concept of a **_"honey do" list_**.

## Getting started
This currently requires a [Go](https://golang.org/) environment to be available. Simply run:

```bash
$ go get -v -u github.com/elliottpolk/honey-do/cmd/honey
```

## Usage
The ```honey``` command expects a ```Honeyfile``` to be present in the path where the command is executed. **YAML** and **JSON** are currently supported formats and the ```Honeyfile``` naming should reflect appropriately (**i.e.** ```Honeyfile.yml```). The remainder of the ```README``` will use the **YAML** format.

### Sample ```Honeyfile.yml```

```yaml
version: "1"

vars:
  DERP: "duup"
  BIN: "honey"

targets:
  

  build:
    deps: 
      - test
    actions:
      - echo -n "hello {{.BIN}}"
      - echo -n "{{.DERP}}"
    platforms:
      "!windows":
        actions:
          - echo -n "hello for {{.DERP}}"
        vars:
          DERP: "not windows"

      windows:
        actions:
          - echo -n "hello for {{.DERP}}"
        vars:
          DERP: "windows"

      darwin:
        actions:
          - echo -n "hello for {{.DERP}}"
        vars:
          DERP: "macOS"

  test:
    actions:
      - echo -n "{{.DERP}}"
    vars:
      DERP: "narf"

```

### Running the sample (macOS)

```bash
INFO[0000] narf
INFO[0000] dependency target test complete
INFO[0000] hello for not windows
INFO[0000] hello for macOS
INFO[0000] hello honey
INFO[0000] duup
INFO[0000] target build complete
INFO[0000] target test complete
```

If no targets are specified, ```all``` is assumed. Not that ```all``` is also a valid (and **reserved**) target.

### Actions and Variables
The ```actions``` are passed through the [Go text/template engine](https://golang.org/pkg/text/template/) prior to execution. There should be a valid ```VAR``` set for every ```{{.VAR}}```. If there is no set var, the output will show ```<no value>``` in place of the ```{{.VAR}}```.

There are 2 levels of variables. The **global-level** is defined outside of the ```targets``` via ```vars```. The **target-level** is defined within a target. The **target-level** variables will override the **global-level** if they share the same name. This can be seen in the sample above with the variable ```DERP```. 

### Platforms
```platforms``` have the ability to run ```actions``` and set ```vars``` specific to an OS. This uses the result of the Go var ```runtime.GOOS``` (see [pkg/runtime](https://golang.org/pkg/runtime/#pkg-constants)). To run for any platform except for a specific one, the key needs to be wrapped in double quotes and prefixed with **!** (**e.g.** ```"!windows"``` will run on all platforms except Windows).

## Additional Notes
Similar to GNU Make, each action is run in its own ```shell```. If 4 actions are specified like:

```yaml
actions:
 - cd project_dir/
 - mkdir -p build/bin
 - go test -v ./...
 - go build -o build/bin/{{.BIN}}
vars:
 BIN: foo
```

This will not behave as expected. It will run the ```cd``` command to the appropriate directory. The subsequent commands will then proceed to be executed in the location where the ```honey``` command was originally executed from, producing a ```build/bin``` directory and (likely) errors. The current work-around for this is to chain the commands in the same ```action```:

```yaml
actions:
 - cd project_dir/; mkdir -p build/bin && go test -v ./... && go build -o build/bin/{{.BIN}}
vars:
 BIN: foo
```

## TODOs
* ~~OS specific actions~~
* Set a **working directory** for a given target




