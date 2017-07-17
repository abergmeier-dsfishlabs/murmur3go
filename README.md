C port of Murmur3 hash
==============

This is a port of the [Murmur3](http://code.google.com/p/smhasher/wiki/MurmurHash3) hash function. Murmur3 is a non-cryptographic hash, designed to be fast and excellent-quality for making things like hash tables or bloom filters.

This is a port of the [C implementation by Peter Scott](https://github.com/PeterScott). We only use x86/128 so this is the only implementation we have.

How to use it
-----------

All code lives in `murmur` package. Signatures and code in general are kept as similar to C implementation as possible.

    func MurmurHash3_x86_128 (key []byte, seed uint32) ([2]uint64)

The interface is: You give a `key`, a byte slice of the data you wish to hash; `seed`, an arbitrary seed number which you can use to tweak the hash. You get an array of two `uint64` (Go <= 1.9 does not yet have a `uint128`).

The hash functions differ in both their internal mechanisms and in their outputs. They are specialized for different use cases:

**MurmurHash3_x86_32** has the lowest throughput, but also the lowest latency. If you're making a hash table that usually has small keys, this is probably the one you want to use on 32-bit machines. It has a 32-bit output. **(not yet ported)**


**MurmurHash3_x86_128** is also designed for 32-bit systems, but produces a 128-bit output, and has about 30% higher throughput than the previous hash. Be warned, though, that its latency for a single 16-byte key is about 86% longer! **(ported - for LittleEndian only)**

**MurmurHash3_x64_128** is the best of the lot, if you're using a 64-bit machine. Its throughput is 250% higher than MurmurHash3_x86_32, but it has roughly the same latency. It has a 128-bit output. **(not yet ported)**

License and contributing
--------------------

All this code is in the public domain. Murmur3 was created by Austin Appleby, and the C port and general tidying up was done by Peter Scott. The Go port was done by Andreas Bergmeier and licensed under Apache License 2.0.