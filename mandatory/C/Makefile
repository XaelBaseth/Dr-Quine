MAKEFLAGS    += --silent
CC = gcc
CFLAGS = -Wall -Wextra -Werror

# Rules

all: Colleen Grace Sully

Colleen: Colleen.c
	$(CC) $(CFLAGS) $< -o $@

Grace: Grace.c
	$(CC) $(CFLAGS) $< -o $@

Sully: Sully.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	$(RM) Colleen Grace Sully

fclean:
	@make clean
	$(RM) Colleen Grace Sully

re: fclean all

.PHONY: all clean fclean re