#include <stdio.h>
/*
One comment
*/
#define FN "Grace_kid.c"
#define printout(...)fprintf(fopen(FN,"w"),__VA_ARGS__);
#define FT(x)int main(){printout(x,10, 10, 10, 10, 34, 34, 10, 34, 34, 10, 10, 34, x, 34, 10, 10)}
char*s="#include <stdio.h>%c/*%cOne comment%c*/%c#define FN %cGrace_kid.c%c%c#define printout(...)fprintf(fopen(FN,%cw%c),__VA_ARGS__);%c#define FT(x)int main(){printout(x,10, 10, 10, 10, 34, 34, 10, 34, 34, 10, 10, 34, x, 34, 10, 10)}%cchar*s=%c%s%c;%cFT(s)%c";
FT(s)
