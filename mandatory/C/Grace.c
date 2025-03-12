#include <stdio.h>
/*
One comment
*/
#define FN "Grace_kid.c"
#define printout(...)fprintf(fopen(FN,"w"),__VA_ARGS__);
#define FT(x)int main(){printout(x, 10, 34, x)}
char *s = "#include <stdio.h>%1$c/*%1$cOne comment%1$c*/%1$c#define FN %2$cGrace_kid.c%2$c%1$c#define printout(...)fprintf(fopen(FN,%2$cw%2$c),__VA_ARGS__);%1$c#define FT(x)int main(){printout(x, 10, 34, x)}%1$cchar *s = %2$c%3$s%2$c;%1$cFT(s)%1$c";
FT(s)
