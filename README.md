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

## Usage 

To use this code locally, you must ensure to install a few things beforehand: 
1) For Go, you need to install the language itself as well as air (when using VSC tasks, see below), which is a tool used 
to automatically recompile the source code since Go is not an interpreted language.
2) For NodeJS, you will obviously need node as well as run "npm i" to install the necessary packages.
3) For Python, please install python as well as django, Pillow and django-cors-headers. 
4) For the frontend, install the necessary packages using npm i.

I recommend using VSC tasks to start up all 4 servers (I pushed the task to the repository as well, you can just start it up like that).

## Suggestions

As for suggestions and possible unfair tests made as well as people wanting to help out by adding more tests/languages, you can contact me at d.steblin.dev@gmail.com.
