# Shell
Simple shell based on the command pattern

https://refactoring.guru/ru/design-patterns/command

## Commands

* cd [dir]
* pwd
* echo [string]
* kill [pid]
* ps
* fork [command [argument ...]]
* exec [command [argument ...]]

## Pipes

It is possible to combine commands into a pipeline using pipe. Commands are executed 
concurrently, each in its own goroutine, data transfer between commands is carried 
out using channels. For example:

```shell
cd / | pwd
```

## Class diagram

![](../../../doc/other_tasks/shell/shell_class.png?raw=true)

## Sequence diagram

![](../../../doc/other_tasks/shell/shell_sequence.png?raw=true)
