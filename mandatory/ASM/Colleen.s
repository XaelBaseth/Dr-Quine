global main
extern printf
;One comment
colleen:
mov rdi, format
mov rsi, 10
mov rcx, 34
mov r8, format
push rbx
call printf wrt ..plt
pop rbx
ret

main:
push rbx
call colleen
pop rbx
;Another comment
mov rax, 0
ret

section .data
format:
db "global main%1$cextern printf%1$c;One comment%1$ccolleen:%1$c%2$cmov rdi, format%1$cmov rsi, 10%1$cmov rcx, 34%1$cmov r8, format%1$cpush rbx%1$ccall printf wrt ..plt%1$cpop rbx%1$cret%1$c%1$cmain:%1$cpush rbx%1$ccall colleen%1$cpop rbx%1$c;Another comment%1$c; Return 0%1$c;%1$cmov rax, 0%1$c%ret%1$c%1$csection .data%1$cformat:%1$cdb %2$c%3$s%2$c%1$c%1$c"
