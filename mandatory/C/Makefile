MAKEFLAGS   += --silent
CC 			= clang
CFLAGS 		= -Wall -Wextra -Werror -g

COLLEEN_SRC	= Colleen.c
GRACE_SRC	= Grace.c
SULLY_SRC	= Sully.c

COLLEEN		= Colleen
GRACE		= Grace
SULLY		= Sully

COLLEEN_OBJ	= $(COLLEEN_SRC:.c=.o)
GRACE_OBJ	= $(GRACE_SRC:.c=.o)
SULLY_OBJ	= $(SULLY_SRC:.c=.o)

# Rules
%.o: %.s
	$(CC) $(CFLAGS) $<

all: Colleen Grace Sully

Colleen: $(COLLEEN_OBJ)
	$(CC) $(COLLEEN_OBJ) -o $(COLLEEN)

test_Colleen: Colleen
	./Colleen > tmp_Colleen ; diff tmp_Colleen $(COLLEEN_SRC)

Grace: $(GRACE_OBJ)
	$(CC) $(GRACE_OBJ) -o $(GRACE)

test_Grace: Grace
	./Grace; diff $(GRACE_SRC) Grace_kid.c

Sully: $(SULLY_OBJ)
	$(CC) $(SULLY_OBJ) -o $(SULLY)

test_Sully: Sully
	./Sully

clean:
	$(RM) Colleen.o Grace.o Sully.o tmp_Colleen Grace_kid.c Sully_*

fclean:
	@make clean
	$(RM) $(COLLEEN) $(GRACE) $(SULLY)

re: fclean all

.PHONY: all clean fclean re test_Colleen test_Grace test_Sully