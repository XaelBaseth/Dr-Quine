AS = nasm
ASFLAGS = -f elf64
MAKEFLAGS    += --silent

# Rules

all: Colleen Grace Sully

Colleen: Colleen.o
	$(LD) $< -o $@ $(LDFLAGS)

Grace: Grace.o
	$(LD) $< -o $@ $(LDFLAGS)

Sully: Sully.o
	$(LD) $< -o $@ $(LDFLAGS)

%.o: %.s
	$(AS) $(ASFLAGS) $< -o $@

clean:
	$(RM) Colleen.o Grace.o Sully.o 

fclean:
	@make clean
	$(RM) Colleen Grace Sully

re: fclean all

.PHONY: all clean fclean re
