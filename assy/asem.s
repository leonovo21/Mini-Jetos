TITLE PROGRAM TO COMPUTE TWO NUMBERS
.model small
.stack 100H
.data 
        string1 db 10,13, 'Enter first number $'
        string2 db 10,13, 'Enter second number $'
        string3 db 10,13, 'Sum is $'
.code 
.global _start
.intel_syntax noprefix


//craet a stack

ret;\