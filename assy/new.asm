segment .data
	LF 			equ 0xA  ; Line Feed
	NULL 		equ 0xD  ; End of the String OR NULL char
	SYS_CALL 	equ 0x80 ; Send value to SO
	; EAX
	SYS_EXIT 	equ 0x1  ; End call code
	SYS_READ 	equ 0x3  ; Read Operation
	SYS_WRITE 	equ 0x4  ; Write operation
	; EBX
	RET_EXIT 	equ 0x0  ; Succesefull operation
	STD_IN 		equ 0x0  ; Normal entry
	STD_OUT 	equ 0x1  ; Exit value
section .data
	msg db "Enter your name: ", LF, NULL
	tam equ $- msg
section .bss
	name resb 1

section .text

gobal_start

_start:
	mov EAX, SYS_WRITE
	mov EBX, STD_OUT
	mov ECX, msg
	mov ECX, tam
	int SYS_CALL

	mov EAX, SYS_READ
	mov EBX, STD_IN
	mov ECX, name
	mov ECX, 0xA
	int SYS_CALL
	
	mov EAX, SYS_EXIT
	mov EBX, RET_EXIT
	int 0x0