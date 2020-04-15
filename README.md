Fx Hash implementation for Go
=============================

The Fx Hash algorithm from the Rust compiler, ported to Go.
https://github.com/rust-lang/rustc-hash/blob/master/src/lib.rs

I used this in a project where I needed a basic hash function of adequate
quality to go as fast as possible on a Raspberry Pi, fnv1-a was too slow and I
didn't fancy getting even more distracted by a Go Assembly yak-shaving exercise
than I already was while making this lib (though that'll probably happen at
some point).

There are fancier hashes than this, but this totally solved my performance
problem so here it is.


## Divergence

Hashes produced by this library are not compatible with hashes produced by the
Rust library. This is partly because Rust doesn't guarantee the stability of
these values anyway.

For this reason, this library diverges slightly from the Rust version:

1. Rust uses host endianness, this uses little endian encoding in order to
   guarantee that values hashed by this library will remain stable.

2. Hashes using `fxhash.Sum64()` and `fxhash.New()` use a seed value, whereas
   in the Rust implementations, the hash is initialised to 0. Rust hashes the
   size of the slice as well as the slice itself, which creates a de-facto seed
   for slices of length 1 or more. Because of divergence #1, I don't bother
   doing that, but you can emulate it easily.

   If you want to use your own seed value, use `fxhash.Seed(seed)` instead of
   `fxhash.New()`, or `fxhash.Append64(seed, ...)` instead of `fxhash.Sum64(...)`.


## Expectation Management

This API should be stable, but this hash function might be no good. I can't
vouch for its quality, and I have only cursorily tested its compatibility with
the Rust implementation it's derived from. I can only state that it solved my
immediate problem, which was one of performance, without an unacceptable level
of degredation due to a bad hash.

If you want to use this, I recommend vendoring it into your project.

I will try to make sure the hashes produced by this library remain stable, but
won't absolutely guarantee it.

Here is the disclaimer from the original repo:
    
    It is not a cryptographically secure hash, so it is strongly recommended
    that you do not use this hash for cryptographic purproses. Furthermore,
    this hashing algorithm was not designed to prevent any attacks for
    determining collisions which could be used to potentially cause quadratic
    behavior in HashMaps. So it is not recommended to expose this hash in
    places where collissions or DDOS attacks may be a concern.


## Silly Benchmarke Game

After just 2 bytes on my i7-8550U, fxhash starts to beat fnv1 pretty comfortably:

    BenchmarkFNV1a64/0-1-8          320184895                3.69 ns/op
    BenchmarkFNV1a64/1-2-8          266694536                4.48 ns/op
    BenchmarkFNV1a64/2-4-8          197219958                6.06 ns/op
    BenchmarkFNV1a64/3-8-8          129225718                9.27 ns/op
    BenchmarkFNV1a64/4-16-8         76044976                16.0 ns/op
    BenchmarkFNV1a64/5-32-8         30411601                33.4 ns/op
    BenchmarkFNV1a64/6-128-8         8789748               130 ns/op
    BenchmarkFxHash/0-1-8           206682140                5.80 ns/op
    BenchmarkFxHash/1-2-8           198642433                5.90 ns/op
    BenchmarkFxHash/2-4-8           205205641                5.84 ns/op
    BenchmarkFxHash/3-8-8           149115448                7.92 ns/op
    BenchmarkFxHash/4-16-8          126367027                9.25 ns/op
    BenchmarkFxHash/5-32-8          101090066               11.8 ns/op
    BenchmarkFxHash/6-128-8         37033078                28.8 ns/op
