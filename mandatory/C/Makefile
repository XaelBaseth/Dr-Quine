MAKEFLAGS   += --silent
CC 			= clang
CFLAGS 		= -Wall -Wextra -Werror -g

# Rules
all: Colleen Grace Sully

Colleen: Colleen.c
	$(CC) $(CFLAGS) $< -o $@

test_Colleen: Colleen
	./Colleen > tmp_Colleen ; diff tmp_Colleen Colleen.c

Grace: Grace.c
	$(CC) $(CFLAGS) $< -o $@

test_Grace: Grace
	./Grace; diff Grace.c Grace_kid.c

Sully: Sully.c
	$(CC) $(CFLAGS) $< -o $@

test_Sully: Sully
	./Sully

clean:
	$(RM) tmp_Colleen Grace_kid.c Sully_*

fclean:
	@make clean
	$(RM) Colleen Grace Sully

re: fclean all

.PHONY: all clean fclean re test_Colleen test_Grace test_Sully