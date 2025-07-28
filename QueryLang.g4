grammar QueryLang;

command: createBenchmark | insertAll | insert | lookup | delete | compareLookup;

createBenchmark: 'CREATE' 'BENCHMARK' BENCHMARK_NAME;
insertAll: 'INSERT_ALL' BENCHMARK_NAME ',' dataList;
insert: 'INSERT' BENCHMARK_NAME ',' BTREE_TYPE ',' dataList;
lookup: 'LOOKUP' BENCHMARK_NAME ',' BTREE_TYPE ',' keyList;
delete: 'DELETE' BENCHMARK_NAME ',' BTREE_TYPE ',' keyList;
compareLookup: 'COMPARE' 'LOOKUP' BENCHMARK_NAME ',' keyList;

dataList: dataEntry (',' dataEntry)*;
dataEntry: '(' INT_LIT ',' INT_LIT ')';

keyList: INT_LIT (',' INT_LIT)*;

BENCHMARK_NAME: [a-zA-Z_][a-zA-Z_0-9]*;
BTREE_TYPE: 'BTREE' | 'B+TREE';

INT_LIT: [0-9]+;

WS: [ \t\r\n]+ -> skip;
