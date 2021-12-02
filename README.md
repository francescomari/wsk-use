# wsk-use

`wsk-use` allows you to update your OpenWhisk configuration depending on a list
of previously defined contexts.

## Install

```shell
go install github.com/francescomari/wsk-use@latest
```

## Configure

Create a configuration file named `.wsk-use` in your home directory. The file
contains JSON data with the following format:

```json
{
    "contexts": {
        "one": {
            "host": "example.com",
            "auth": "2a2407ff-5970-47ba-a06d-1f668a3dcffb"
        },
        "two": {
            "host": "example.com",
            "auth": "da615349-919f-4b96-a057-bb8e42c162b9"
        },
        "three": {
            "host": "example.com",
            "auth": "d7d7ac43-9ec2-48b9-a2c9-a2e9c0724e3d"
        }
    }
}
```

## List contexts

Invoke `wsk-use` without arguments to see the list of available contexts.

```shell
$ wsk-use
one
three
two
```

## Switch context

Invoke `wsk-use` with the name of the context you want to switch to. `wsk-use`
will read the data from the context and write a new `.wskprops` file in your
home directory.

```shell
$ wsk-use two
```

If you already have a `.wskprops` in your home directory, `wsk-use`
will overwrite it and won't create a backup.

