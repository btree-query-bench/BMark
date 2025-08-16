grammar QueryLang;

query: createBenchmark | search | modify | start;

createBenchmark: 'CREATE' 'BENCHMARK' BENCHMARK_NAME;

//search
search: 'SEARCH' BENCHMARK_NAME ',' (INDEX_TYPE ',')? keyList;

//modify
modify: insert | delete;
insert: 'INSERT' BENCHMARK_NAME ','(INDEX_TYPE ',')? dataList;
delete: 'DELETE' BENCHMARK_NAME ',' (INDEX_TYPE ',')? keyList;

// start
start: 'START' BENCHMARK_NAME;

// non-terminals
dataList: dataEntry (',' dataEntry)*;
dataEntry: '(' INT_LIT ',' INT_LIT ')';
keyList: INT_LIT (',' INT_LIT)*;

BENCHMARK_NAME: [a-zA-Z_][a-zA-Z_0-9]*;
INDEX_TYPE: 'BTREE' | 'B+TREE' | 'LIST';

INT_LIT: [0-9]+;

WS: [ \t\r\n]+ -> skip;
