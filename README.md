# Dr-Quine

Outside of being the series you watch during the lunch, a quine is also an auto-replication program that, when executed, outputs its own source code without any external input. The challenge is to achieve this without directly reading its own source file.
## The concept

The key to writing a quine is to find an alternative way to use a quote character other than the literal character itself. This is because the quote character would prematurely end the string and therefore requires escaping with a backslash. However, the escaping character also requires escaping, and so on.

In C the printf function can use a char placeholder (%c) in a format string and then, because a char is just an int, interpreted as an ASCII character, the integer value of an ASCII **quote (34)** can be passed as an argument. **Newlines** can be passed as decimal **10** to avoid the backslash escape hell there as well.