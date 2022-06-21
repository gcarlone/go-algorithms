# Dynamic Connectivity

Given a set of N objects.
- **Union command**: connect two objects.
- **Find/connected query**¢: is there a path connecting the two objects?

<p align="center">

![Connectivity](.img/connectivity.png)

</p>

We assume "is connected to" is an **equivalence relation**:
- Reflexive: _p_ is connected to _p_.
- Symmetric: if _p_ is connected to _q_, then _q_ is connected to _p_.
- Transitive: if _p_ is connected to _q_ and _q_ is connected to _r_,
then _p_ is connected to _r_.

> **Connected components**: Maximal set of objects that are mutually
connected.

<p align="center">{ 0 } { 1 4 5 } { 2 3 6 7 }</p>
<br>

<p align="center">

![Union](.img/union.png)

</p>

<br>

# Quick-Find [eager approach]
Quick-find is too slow

> **Cost model** Number of array accesses (for read or write)

<p align="center">

|algorithm|initialize|union|find|
|---------|:--------:|:---:|:--:|
|`quick-find`|N|N|1|

</p>
<p align="center">order of growth of number of array accesses</p>

> **Union** is too expensive It takes N² array accesses to process a sequence of N union commands on N objects.

<br>

# Quick-Union [lazy approach]
Quick-union is also too slow
<br>

> Data structure
- Integer array id[] of length N
- Interpretation: id[i] is parent of i
- Root of i is id[id[id[...id[i]...]]]

> **Find** check if p and q have the same root.

> **Union** to merge components containing p and q, set the id of p's root to the id of q's root.

<p align="center">

![Quick Union](.img/quick-union.png)

</p>

> **Cost model** Number of array accesses (for read or write)

<p align="center">

|algorithm|initialize|union|find|
|---------|:--------:|:---:|:--:|
|`quick-find`|N|N|1|
|`quick-union`*|N|N**|N|

_*  worst case_<br>
_** includes cost of finding roots_<br>
</p>
<br>

***

> **Quick-find defect**
- Union too expensive (N array accesses)
- Trees are flat, but too expensive to keep them flat
> **Quick-union defect**
- Trees can get tall
- Find too expensive (could be N array accesses)

<br>

# Weighted quick-union.
- Modify quick-union to avoid tall trees
- Keep track of size (number of objects) of each tree
- Balance by linking root of smaller tree to root of larger tree

<br>

> **Data structure** Same as quick-union, but maintain extra array size[i] to count number of objects in the tree rooted at i.

> **Find** Identical to quick-union but takes time proportional to depth of p and q (depth of any node x is at most lg N)

> **Union** Modify quick-union to link root of smaller tree to root of larger tree. Takes constant time, given roots

<br>

<p align="center">

![QU-WQU](.img/qu-wqu.png)

</p>

> **Cost model** Number of array accesses (for read or write)
<p align="center">

|algorithm|initialize|union|find|
|---------|:--------:|:---:|:--:|
|`quick-find`|N|N|1|
|`quick-union`*|N|N*|N|
|`weighted quick-union`*|N|ln N*|ln N|

_* includes cost of finding roots_<br>
</p>

<br>

# Quick union with path compression
Just after computing the root of p, set the id of each examined node to point to that root.

<br>

# Worst-Case Time
 
<p align="center">

|algorithm|worst-case time|
|---------|:-------------:|
|quick-find|M * N|
|quick-union|M * N|
|weighted QU|N + M * log N|
|QU + path compression|N + M * log N|
|weighted QU + path compression| N + M * lg N|

_M union-find operations on a set of N objects_

</p>

