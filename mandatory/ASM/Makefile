MAKEFLAGS	+= --silent
CC 			= gcc
CCFLAGS		= -no-pie
AS 			= nasm
ASFLAGS 	= -f elf64

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
	$(CC) $(CCFLAGS) $()$(COLLEEN_OBJ) -o $(COLLEEN)

test_Colleen: Colleen
	./Colleen > tmp_Colleen ; diff --color=auto tmp_Colleen Colleen.s

Grace: $(GRACE_OBJ)
	$(CC) $(GRACE_OBJ) -o $(GRACE)

test_Grace: Grace
	./Grace; diff Grace.s Grace_kid.s

Sully: $(SULLY_OBJ)
	$(CC) $(SULLY_OBJ) -o $(SULLY)

test_Sully: Sully
	./Sully

clean:
	$(RM) Colleen.o Grace.o Sully.o tmp_Colleen Grace_kid.c Sully_*

fclean:
	@make clean
	$(RM) Colleen Grace Sully

re: fclean all

.PHONY: all clean re
