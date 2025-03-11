#include <stdio.h>
/*
This is an external comment
*/

void function(char*s) {
printf(s, 10, 34, s);
}

int main(void) {
/*
This is a comment inside the function
*/
char*s="#include <stdio.h>%1$c/*%1$cThis is an external comment%1$c*/%1$c%1$cvoid function(char*s) {%1$cprintf(s, 10, 34, s);%1$c}%1$c%1$cint main(void) {%1$c/*%1$cThis is a comment inside the function%1$c*/%1$cchar*s=%2$c%3$s%2$c;%1$cfunction(s);%1$creturn 0;%1$c}%1$c";
function(s);
return 0;
}
