ANTLR_JAR=lib/antlr-complete.jar
GRAMMAR=QueryLang.g4
OUT_DIR=antlr
PACKAGE=antlr

.PHONY: all generate clean

all: generate

generate:
	java -jar $(ANTLR_JAR) -visitor -no-listener -Dlanguage=Go -package $(PACKAGE) -o $(OUT_DIR) $(GRAMMAR)

clean:
	rm -rf $(OUT_DIR)/*.go
