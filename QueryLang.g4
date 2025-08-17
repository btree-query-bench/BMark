grammar QueryLang;

query: createBenchmark | search | modify | start;

// create
createBenchmark: 'CREATE' 'BENCHMARK' NAME;

//search
search: SEARCH NAME COMMA (INDEX_TYPE COMMA)? keyList;

//modify
modify: insert | delete;
insert: INSERT NAME COMMA (INDEX_TYPE COMMA)? dataList;
delete: DELETE NAME COMMA (INDEX_TYPE COMMA)? keyList;

// start
start: START NAME;

dataList: dataEntry (',' dataEntry)*;
dataEntry: '(' INT_LIT ',' INT_LIT ')';
keyList: INT_LIT (',' INT_LIT)*;

NAME: [a-zA-Z_][a-zA-Z_0-9]*;
INDEX_TYPE: 'BTREE' | 'B+TREE' | 'LIST';
CREATE: 'CREATE';
BENCHMARK: 'BENCHMARK';
SEARCH: 'SEARCH';
INSERT: 'INSERT';
DELETE: 'DELETE';
START: 'START';
COMMA: ',';
LPAREN: '(';
RPAREN: ')';

INT_LIT: [0-9]+;

WS: [ \t\r\n]+ -> skip;
