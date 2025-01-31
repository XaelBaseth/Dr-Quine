section .data
    fmt db 'section .data',10,'    fmt db %c%s%c,10,%c%s%c,10,0',10
    newline db 10,0

section .text
    global _start

_start:
    mov rdi, fmt
    mov rsi, 39     ; ASCII for '
    mov rdx, fmt
    mov rcx, 34     ; ASCII for "
    mov r8, newline
    call print
    call exit

print:
    mov rax, 1       ; sys_write
    mov rdi, 1       ; stdout
    mov rdx, 100     ; length
    syscall
    ret

exit:
    mov rax, 60      ; sys_exit
    xor rdi, rdi
    syscall
