#include <stdio.h>
/*
This is an external comment
*/
void function(char*s){/*
This is a comment inside the function
*/printf(s, 10, 10, 10, 10, 10, 10, 10, 34, s, 34, 10, 10);}
int main(void){char*s="#include <stdio.h>%c/*%cThis is an external comment%c*/%cvoid function(char*s){/*%cThis is a comment inside the function%c*/printf(s, 10, 10, 10, 10, 10, 10, 10, 34, s, 34, 10, 10);}%cint main(void){char*s=%c%s%c;function(s);return 0;}%c";function(s);return 0;}
