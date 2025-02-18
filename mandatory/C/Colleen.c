#include <stdio.h>
/* This is an external comment */
void function(char*s){/* This is a comment inside the function */printf(s, 10, 10, 10, 34, s, 34, 10, 10);}
int main(void){char*s="#include <stdio.h>%c/* This is an external comment */%cvoid function(char*s){/* This is a comment inside the function */printf(s, 10, 10, 10, 34, s, 34, 10, 10);}%cint main(void){char*s=%c%s%c;function(s);return 0;}%c";function(s);return 0;}
