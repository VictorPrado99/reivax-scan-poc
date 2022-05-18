# Reivax Scan Poc
### What is it?
Reivax is a cli tool, writen in Golang using Cobra, who mimic a static security code scanner. The objective wans't create something who should be used in production. Actually the logic in the scanner is really simple, and will just mimic a real scanner. But the design architeture could be used in real cli tool. Actually, the code is really decoupled and create a real logic for static scanner, would be really simple.

### How Works?
Basically is a cli created with cobra, so the framework take cares of the command creation. New commands can easily be created using the command
```
cobra add <new_command>
```
And that is it. Cobra will generate a file you can use as start point to create your logic.
Reivax works basically starting off from scan command, who receive a directory, and have some flags options to set the extension of files you wanna analyse, and the output format you wanna (json, plain text, etc). The cli itself can tell you more about it.

About the Logic itself, works in two packages.

- **analysis_output**
This package is responsible for generate the output, he will write you json or txt file. The way that works, is basically you have a manager who defines the abstraction, and the object who have the api other package use it. To create new output form, you basically need to create a new file and implement the interface with the function to generate your output type, and have a init function, who will subscribe to the manager.
- **code_scanner**
This package work pretty much the same as the package above, the difference the have a singleton to basically be used as a pseudo Observable Pattern. They implement an interface who works as a pseudo Abstract Factory Pattern, to create new code scanner (e.g. SQL Injection Scanner), you need to do preety much the same as to create a new output. You create a new object implementing the interface and subscribe to the singleton in the init function. The main difference, was thinking in implementing for different languages, we have one more abstraction layer a interface to define the method who should run.  
This was made thinking to give support to different languages with the same scanner, you can have for example, a object to analyse sql injection in golang, and another to analyse sql injection in java. The process to create new is the same, a file who implement the interface, and subscribe in init function. To define this logic, is needed you create a id for the scanner in both files. This is used after to generate a key pairing ScannerID | extension, in a dictionary (map) for the mechanisn know which analyse object to use.

### How I run it?
Right now, you need to use
```
go install && reivax-scan-poc <flags> <directory>
```
Eventually the cli will be dockerized with the real build.

### Objective
This cli was a great project to learn more about Golang. I can say for sure somethings could be improved, and the logic for the scan itself, is not well made. The regex will only match if the pattern is in the same line. This **isn't** a technical limitation, this happen because I did't wanted to spend more time in this logic, thinking this cli was just to mimic a real tool.
