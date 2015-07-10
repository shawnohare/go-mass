# go-rel

Package rel provides interfaces and functions that act on lists
of objects that have some embedding into the real numbers.
A common use case is to employ a single data structure modeling
a collection of (object, value) pairs that handles sorting, partitioning,
and statistical calculations over the corresponding set of real values.

## What's in a name?

The package name is a bit of a play on the word.  On the one hand, it
connotes a mathematical **rel**ation between between some set and the reals.
This relation is more appropriately viewed as a function, which gives rise
to an embedding into the reals, or a Real Embeding.  The interface modeling 
this situation is a list, hence we have a Real Embedded List.

## Version History

- 1.0.0 Fri, 10 Jul 2015 10:17:29 -0700
