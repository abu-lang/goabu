# Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
# SPDX-License-Identifier: Apache-2.0

# GNU Make >= 4.3

P_DIR = internal/antlr
SRC_COMMIT = $(shell sed -n -e '/^https/p' ../LICENSE | cut -d '/' -f 7)

ANTLR = antlr4
CKSUM = cksum
CURL = curl

ANTLR_VERSION = $(shell $(ANTLR) | head -n 1 | cut -d ' ' -f 6)
ANTLRFLAGS = -Dlanguage=Go -o '${P_DIR}' -package '$(notdir ${P_DIR})'
CURLFLAGS = -\# --proto =https

TOK = ${P_DIR}/EcaruleLexer.interp ${P_DIR}/EcaruleLexer.tokens
TOK_PART = $(addsuffix .part,${TOK})
PAR = ${P_DIR}/ecarule_parser.go ${P_DIR}/ecaruleparser_listener.go
PAR_PART = $(addsuffix .part,${PAR})

GRAMMARS = EcaruleLexer.g4 Grulev3Parser.g4 EcaruleParser.g4
PTOKS = ${P_DIR}/EcaruleParser.interp ${P_DIR}/EcaruleParser.tokens
TMP = EcaruleLexer.g4.part ${TOK_PART} ${PAR_PART} ${P_DIR}/ecaruleparser_base_listener.go grulev3.g4

notlast = $(filter-out $(lastword $(2)),$(1))

.PHONY: clean all
.INTERMEDIATE: ${TMP}
.DELETE_ON_ERROR:

all: ${PAR} ${GRAMMARS} ${TOK} ${PTOKS} ${P_DIR}/ecarule_lexer.go

define PRINT_ANTLR_HEADER
@echo "// Code generated from EcaruleParser.g4 by ANTLR ${ANTLR_VERSION} and MODIFIED by ../Makefile." > $@
@echo "" >> $@
endef

${P_DIR}/ecarule_parser.go: ${P_DIR}/ecarule_parser.go.part
	$(PRINT_ANTLR_HEADER)
	sed -n -e '/github\.com/{G;s/\(."github\.com.*\)\(\n\)/\1\2@"github.com\/hyperjumptech\/grule-rule-engine\/antlr\/parser\/grulev3"/;}'\
		-e '/type GrlContext struct/,$${/listenerT\.Enter/{G;s/listenerT\.Enter\([A-Za-z]*\)(s)\(\n\)/c := grulev3.New\1Context(s.parser, s.BaseParserRuleContext, -1)\2@@c.BaseParserRuleContext = s.BaseParserRuleContext\2@@listenerT.Enter\1(c)/;};}'\
		-e '/type GrlContext struct/,$${/listenerT\.Exit/{G;s/listenerT\.Exit\([A-Za-z]*\)(s)\(\n\)/c := grulev3.New\1Context(s.parser, s.BaseParserRuleContext, -1)\2@@c.BaseParserRuleContext = s.BaseParserRuleContext\2@@listenerT.Exit\1(c)/;};}'\
		-e '/^package/,$$p' $^ | tr '@' '\011' >> $@

${P_DIR}/ecaruleparser_listener.go: ${P_DIR}/ecaruleparser_listener.go.part
	$(PRINT_ANTLR_HEADER)
	sed -n -e '/^import/{G;s/import \(.*\)\(\n\)/import (\2@\1\2@"github.com\/hyperjumptech\/grule-rule-engine\/antlr\/parser\/grulev3"\2)/;}'\
		-e '/^.EnterGrl/,/^.EnterBooleanLiteral/{s/^.Enter\([a-zA-Z]*\).*/@Enter\1(c *grulev3.\1Context)/;}'\
		-e '/^.ExitGrl/,/^.ExitBooleanLiteral/{s/^.Exit\([a-zA-Z]*\).*/@Exit\1(c *grulev3.\1Context)/;}'\
		-e '/^package/,$$p' $^ | tr '@' '\011' >> $@

${PAR_PART} ${PTOKS} ${P_DIR}/ecaruleparser_base_listener.go &: ${GRAMMARS} ${TOK}
	$(ANTLR) ${ANTLRFLAGS} -listener EcaruleParser.g4
	$(foreach var,${PAR},mv ${var} ${var}.part $(if $(call notlast,${var},${PAR}),&&))

${TOK} ${P_DIR}/ecarule_lexer.go &: EcaruleLexer.g4 ${TOK_PART}
	$(foreach var,${TOK},cp ${var}.part ${var} $(if $(call notlast,${var},${TOK}),&&))
	$(ANTLR) ${ANTLRFLAGS} EcaruleLexer.g4

EcaruleLexer.g4: ${TOK_PART}
EcaruleLexer.g4: EcaruleLexer.g4.part Unshared.g4
	sed -e '/lexer grammar EcaruleLexer/{G;s/\(lexer grammar EcaruleLexer.\)\(\n\)/\1\2\2options \{tokenVocab=EcaruleLexer\;\}/;}'\
		-e '/^SIMPLENAME/,$$d' EcaruleLexer.g4.part > $@
	cat Unshared.g4 >> $@
	@echo "" >> $@
	sed -n -e '/^SIMPLENAME/,/START EcaruleParser/{/START EcaruleParser/!p;}' EcaruleLexer.g4.part >> $@

${TOK_PART} &: EcaruleLexer.g4.part
	cp $^ EcaruleLexer.g4
	$(ANTLR) ${ANTLRFLAGS} EcaruleLexer.g4
	rm EcaruleLexer.g4 ${P_DIR}/ecarule_lexer.go
	$(foreach var,${TOK},mv ${var} ${var}.part $(if $(call notlast,${var},${TOK}),&&))

define PRINT_GRULE_HEADER
@echo "// this $(if $(findstring exer,$@),lexer,parser) grammar was obtained by slightly MODIFYING the grulev3 grammar" > $@
@echo "// of the Grule Rule Engine which is released under the following license:" >> $@
@echo "" >> $@
sed -n -e '/Copyright hyperjumptech/,/limitations/{s/^ /\/\//;s/^$$/\/\//;p;}' ../LICENSE >> $@
@echo "" >> $@
@echo '// source: "https://github.com/hyperjumptech/grule-rule-engine/blob/${SRC_COMMIT}/antlr/grulev3.g4"'  >> $@
@echo "// see Makefile for details" >> $@
@echo "" >> $@
endef

EcaruleLexer.g4.part: grulev3.g4 ../LICENSE Unshared.g4
	$(PRINT_GRULE_HEADER)
	@echo "lexer grammar EcaruleLexer;" >> $@
	@echo "" >> $@
	sed -n -e '$$ ! s/.$$//'\
		-e "/^PLUS .*/{G;s/\(^PLUS .*\)\(\n\)/T__0                        : ',' \;\2\1/;}"\
		-e '1,/LEXER HERE/!p' grulev3.g4 >> $@
	@echo "" >> $@
	cat Unshared.g4 >> $@

Grulev3Parser.g4: grulev3.g4 ../LICENSE
	$(PRINT_GRULE_HEADER)
	@echo "parser grammar Grulev3Parser;" >> $@
	@echo "" >> $@
	sed -n -e "s/','/T__0/"\
		-e 's/.$$//'\
		-e '1,/PARSER HERE/!{/LEXER HERE/,/PARSER HERE/!p;}' $^ >> $@

grulev3.g4: ../LICENSE
	$(CURL) ${CURLFLAGS} 'https://raw.githubusercontent.com/hyperjumptech/grule-rule-engine/${SRC_COMMIT}/antlr/grulev3.g4' -o $@
	@CK=$$($(CKSUM) $@) && test "1672161330 7392 $@" = "$$CK"
	@echo "" >> $@

clean:
	rm -f ${TMP}
	rm -f ${PAR} ${TOK} ${PTOKS} ${P_DIR}/ecarule_lexer.go EcaruleLexer.g4 Grulev3Parser.g4
