# Fndamentals of Operating Systems

Notes from [Fundamentals of Operating Systems](https://www.udemy.com/course/fundamentals-of-operating-systems) course by Hussein Nasser.

## Content

## Simple Process Execution
```
cd 01-simple-process
# Compile to assembly code
gcc -S main.c
less main.s

# debug execution
# see documentation at: https://sourceware.org/gdb/current/onlinedocs/gdb
gdb ./main

# set function breakpoint
break sum

# step to next line
n

# examine memory
# x/<number_of_bytes><display_format><unit_size>
# display the next 4 machine instructions (in memory) from the program counter address
x/4ib $pc

# examine registers
info registers

# print register values
# p/<display_format>
# display_format can be "x: hexadecimal", "d: decimal", etc.
p/x $pc
```