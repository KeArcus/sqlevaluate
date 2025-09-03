grammar SQL;

// Parser rules
sql_statement
    : select_statement EOF
    ;

select_statement
    : SELECT select_list FROM table_name (WHERE where_clause)?
    ;

select_list
    : STAR
    | column_list
    ;

column_list
    : column_name (',' column_name)*
    ;

table_name
    : identifier
    ;

column_name
    : identifier
    ;

where_clause
    : condition
    ;

condition
    : condition AND condition                          # andCondition
    | condition OR condition                           # orCondition
    | '(' condition ')'                                # parenthesisCondition
    | comparison_expression                            # comparisonCondition
    | expression IN '(' value_list ')'               # inCondition
    | expression NOT IN '(' value_list ')'           # notInCondition
    | expression LIKE expression                     # likeCondition
    | expression NOT LIKE expression                 # notLikeCondition
    | expression IS NULL                             # isNullCondition
    | expression IS NOT NULL                         # isNotNullCondition
    ;

comparison_expression
    : expression comparison_operator expression
    ;

expression
    : function_call
    | column_name
    | value
    ;

comparison_operator
    : '='
    | '!='
    | '<>'
    | '<'
    | '<='
    | '>'
    | '>='
    ;

value
    : string_literal
    | number_literal
    | boolean_literal
    | NULL
    ;

value_list
    : expression (',' expression)*
    ;

string_literal
    : STRING
    ;

number_literal
    : NUMBER
    ;

boolean_literal
    : TRUE
    | FALSE
    ;

function_call
    : function_name '(' function_args? ')'
    ;

function_name
    : identifier
    ;

function_args
    : expression (',' expression)*
    ;

identifier
    : IDENTIFIER
    ;

// Lexer rules
SELECT      : S E L E C T ;
FROM        : F R O M ;
WHERE       : W H E R E ;
AND         : A N D ;
OR          : O R ;
NOT         : N O T ;
IN          : I N ;
LIKE        : L I K E ;
IS          : I S ;
NULL        : N U L L ;
TRUE        : T R U E ;
FALSE       : F A L S E ;

// Functions are handled as regular identifiers

STAR        : '*' ;

STRING      : '\'' (~'\'' | '\'\'')* '\'' ;
NUMBER      : [0-9]+ ('.' [0-9]+)? ;
IDENTIFIER  : [a-zA-Z_][a-zA-Z0-9_]* ;

WS          : [ \t\r\n]+ -> skip ;

// Case-insensitive fragments
fragment A : [aA] ;
fragment B : [bB] ;
fragment C : [cC] ;
fragment D : [dD] ;
fragment E : [eE] ;
fragment F : [fF] ;
fragment G : [gG] ;
fragment H : [hH] ;
fragment I : [iI] ;
fragment J : [jJ] ;
fragment K : [kK] ;
fragment L : [lL] ;
fragment M : [mM] ;
fragment N : [nN] ;
fragment O : [oO] ;
fragment P : [pP] ;
fragment Q : [qQ] ;
fragment R : [rR] ;
fragment S : [sS] ;
fragment T : [tT] ;
fragment U : [uU] ;
fragment V : [vV] ;
fragment W : [wW] ;
fragment X : [xX] ;
fragment Y : [yY] ;
fragment Z : [zZ] ;
