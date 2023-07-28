# Zo's solution to the Jotto Problem

## The problem

### Overview

The "Jotto Problem" is a challenge to find five 5-letter words where no letter is repeated in either a single word or accross all the words.

### Background

Taking inspiration from Autralia's very own Matt Parker, I have decided to attempt finding all the solutions to the "Jotto Problem" in Go. This is meant to serve as a coding exercis, continuing my learning of Go.

For more information on the problem itself and the inspiration, see [Matt's video](https://www.youtube.com/watch?v=c33AZBnRHks).

## Results

### Searching for only one result

| Word list                      | Type           | Time taken     |                                                       |
|--------------------------------|----------------|----------------|-------------------------------------------------------|
| **words_alpha**                | First solution | ~ 100 ms       | fconv, klutz, bejig, ampyx, hdqrs                     |
|                                | All solutions  | ~ 3 min 30 sec | [See full list](./SOLUTIONS.md#All-words)             |
| **wordle_data/accepted_words** | First solution | ~ 4 sec        | kempt, brung, vozhd, cylix/xylic, waqfs               |
|                                | All solutions  | ~ 46 sec       | [See full list](./SOLUTIONS.md#Wordle-accepted-words) |
| **wordle_data/answers**        | First solution | ~ 600 ms       | None found                                            |
|                                | All solutions  | ~ 600 ms       | None exist                                            |

## Sources

`words_alpha.txt` is sourced from [here](https://github.com/dwyl/english-words).

Both word lists from `wordle_data` are sourced from [this github user](https://gist.github.com/cfreshman).
