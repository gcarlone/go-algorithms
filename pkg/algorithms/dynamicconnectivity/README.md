# Quick-find [eager approach] is too slow

´Cost model´ Number of array accesses (for read or write)

|algorithm|initialize|union|find|
|---------|----------|-----|----|
|quick-find|N|N|1|

order of growth of number of array accesses

´Union is too expensive´ It takes N² array accesses to process a sequence of N union commands on N objects.