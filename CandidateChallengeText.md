## General Description

Welcome, Candidate! Please create a program executable from the command line that when given text(s) will return 100 of most common three word sequences in descending order of frequency.

For example, if I ran `ruby ./solution.rb texts/moby_dick.txt` or `java main.java texts/moby_dick.txt` the first lines of the result would (probably) be:

```
the sperm whale - 85
the white whale - 71
of the whale - 67
```

You can write this in any language, but if you choose something other than java or ruby, please give us some hints about how to set it up locally.

## How long to spend on the challenge

* We suggest spending 2 hours meeting these basic requirements:
  - The code has a README with examples of how to run the code AND the tests.
  - The "Basic requirements" below should be satisfied.
  - Automated tests should be written.
  - Code should be well structured.
* Then, if you so choose, spend 2 hours on the extra credit section at the bottom of these instructions.

## Documentation Requirements

* Your README should include:
  - How to run your program.
  - How to run the tests.
  - What you would do next, given more time (if anything)?
  - Are there bugs that you are aware of?

## Basic Requirements

* The program is written in any language (ruby and java are common, but all are welcome)
* The program has a README with instructions for running the program and the tests
  - If you choose a language other than java or ruby, please give us some hints about how you would recommend we get it running locally.
* The program accepts a list of one *or more* file paths (e.g. `ruby solution.rb texts/moby-dick.txt brothers-karamazov.txt ...).`
* The program *also* accepts input via stdin (e.g. `cat texts/*.txt | java solution.java`).
* The program outputs the first 100 most common three word sequences.
* The program ignores punctuation, line endings, and is case insensitive
  - `“I love\nsandwiches.”` should be treated the same as `"(I LOVE SANDWICHES!!)"`).
  - Contractions shouldn't be changed into two words (eg. `can't` should not become `can t`).
  - Hyphens on line-endings can be treated as continuations OR punctuation.
  - Unicode may also be handled in any way you like (including replacing it with a space)
* Bonus points for doing something extra and documenting it. Here are some examples that might be relevant to the team you're applying to:
  - The program is capable of processing large files and runs efficiently (>4s is normal in Ruby, 2.5s is quite quick). *Please do not* write incomprehensible code in the quest for speed.
  - Handling unicode or line-endings gracefully.
  - Running via Docker as well as via the command-line.
  - Something else that you think your team would find valuable (just make sure to document it so we know to look for it!)

## Testing Requirements

* The program should have automated tests that use assertions.
* The program should have documentation about how to run the program AND the tests.
* Tests should be runnable on the evaluator's machine with small effort.
* Tests should cover the important, public parts of the program.
* Bonus points for doing something extra and documenting it. Here are some examples that might be relevant to the team you're applying to:
  - The tests clearly demonstrate both functional *and* unit style testing.
  - The tests cover boundary conditions and positive/negative tests.
  - In addition to the main body of tests, the tests *also* cover other exploratory styles of testing such as fuzzing.
  - Something else that you think your team would find valuable (just make sure to document it so we know to look for it!)

## Code Structure/Style Requirements

* The program should be well structured and understandable. This is purposefully a bit vague, but generally we're looking for:
  - Encapsulation
  - Separation of concerns
  - Understandability
  - Appropriate use of data structures for the language of choice
* Please show us what _you_ think makes a program beautiful and understandable. In exchange, we'll try to open our minds and learn from you.
* There are also many bonus points you can earn in this category, but as with the previous categories, if you do attempt bonus points in this category, please document what you did that's extra special so we can properly appreciate it.
