# Benchmarks

A repository to measure performance across different languages in respect to operations commonly used in backends. 
Currently features Golang (with the Fiber Framework), NodeJS (with the Express Framework) and Python (with the Django Framework). 

Obviously some tests will tell you more about the strong and weak points of languages, whereas some others might be showing off 
the difference between library implementations across different programming languages. 

As an exmaple: 
Even though Golang is considered way faster than NodeJS and Python, the image decoding and encoding is slower than in NodeJS f.e. because 
the underlying library in NodeJS commonly used for this (sharp) is probably written in C or C++ making it perform way better than standard NodeJS code.

The first table (measuring performance for algorithms) might be the best representation of how fast a certain language is, but I think it is also 
fair to compare standard operations involving files, images and SQL queries to showcase that real-life performance including libraries written in different languages 
and optimized differently can impact the speed of applications heavily and therefore making statements about application performance solely based on programming language used 
is not a very good thing to do for complex applications.

## More languages

In the future, Rust/Java/C# will be added as well.
