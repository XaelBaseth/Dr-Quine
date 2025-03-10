MAKEFLAGS	+= --silent
GOBUILD		= go build

COLLEEN_SRC	= Colleen.go
GRACE_SRC	= Grace.go
SULLY_SRC	= Sully.go

COLLEEN		= Colleen
GRACE		= Grace
SULLY		= Sully

all: Colleen Grace Sully

Colleen:
	$(GOBUILD) $(COLLEEN_SRC)

test_Colleen: Colleen
	./Colleen > tmp_Colleen ; diff tmp_Colleen $(COLLEEN_SRC)

Grace:
	$(GOBUILD) $(GRACE_SRC)

test_Grace: Grace
	./Grace; diff $(GRACE_SRC) Grace_kid.c

Sully:
	$(GOBUILD) $(SULLY_SRC)

test_Sully: Sully
	./Sully

clean:
	$(RM) tmp_Colleen Grace_kid.c Sully_*

fclean:
	@make clean
	$(RM) $(COLLEEN) $(GRACE) $(SULLY)

re: fclean all

.PHONY: all clean fclean re test_Colleen test_Grace test_Sully
