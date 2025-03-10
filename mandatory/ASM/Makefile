MAKEFLAGS	+= --silent
CC 			= ld
CCFLAGS		= -dynamic-linker /lib64/ld-linux-x86-64.so.2 -lc
AS 			= nasm
ASFLAGS 	= -felf64

COLLEEN_SRC	= Colleen.s
GRACE_SRC	= Grace.s
SULLY_SRC	= Sully.s

COLLEEN		= Colleen
GRACE		= Grace
SULLY		= Sully

COLLEEN_OBJ	= $(COLLEEN_SRC:.s=.o)
GRACE_OBJ	= $(GRACE_SRC:.s=.o)
SULLY_OBJ	= $(SULLY_SRC:.s=.o)

%.o: %.s
	$(AS) $(ASFLAGS) $<

all: Colleen Grace Sully

Colleen: $(COLLEEN_OBJ)
	$(CC) $(CCFLAGS) $(COLLEEN_OBJ) -o $(COLLEEN)

test_Colleen: Colleen
	@./Colleen > tmp_Colleen ; diff tmp_Colleen $(COLLEEN_SRC)

Grace: $(GRACE_OBJ)
	$(CC) $(CCFLAGS) $(GRACE_OBJ) -o $(GRACE)

test_Grace: Grace
	./Grace; diff Grace.s Grace_kid.s

Sully: $(SULLY_OBJ)
	$(CC) $(CCFLAGS) $(SULLY_OBJ) -o $(SULLY)

test_Sully: Sully
	./Sully

clean:
	$(RM) Colleen.o Grace.o Sully.o tmp_Colleen Grace_kid.s Sully_*

fclean:
	@make clean
	$(RM) $(COLLEEN) $(GRACE) $(SULLY)

re: fclean all

.PHONY: all clean re Colleen test_Colleen Grace test_Grace test_Sully Sully
