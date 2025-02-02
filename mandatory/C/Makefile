MAKEFLAGS    += --silent
CC = gcc
CFLAGS = -Wall -Wextra -Werror

# Rules

all: Colleen Grace Sully

Colleen: Colleen.c
	$(CC) $(CFLAGS) $< -o $@

test_Colleen: Colleen
	clang $(CFLAGS) -o Colleen Colleen.c; ./Colleen > tmp_Colleen ; diff tmp_Colleen Colleen.c

Grace: Grace.c
	$(CC) $(CFLAGS) $< -o $@

Sully: Sully.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	$(RM) Colleen Grace Sully tmp_Colleen

fclean:
	@make clean
	$(RM) Colleen Grace Sully

re: fclean all

.PHONY: all clean fclean re