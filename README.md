# aocgo - Advent of Code solutions in Go

## What is Advent of Code?

[This](https://adventofcode.com). Every year since 2015, starting December 1 to 25, 25 small programming problems are presented on http://adventofcode.com. The challenges are fun to solve in programming language of your choice. I mostly solved those with [Elixir in the past](https://github.com/code-shoily/advent_of_code), but am looking forward to play with Golang in future challenges. And as I am trying really learn Go, I felt like I should backfill the old challenges throughout 2023. 

More about Advent of Code here: [https://adventofcode.com/2022/about](https://adventofcode.com/2022/about).

## How do I run these?
* To create stubs for solving for instance, 2022/1, run: `go run main.go gen 2022 1` (also `g` instead of `gen`)
* To run solution for the same, run: `go run main.go solve 2022 1` (also `run`, `r` or `s` instead of `solve`)

Note: The additional commands (i.e. `run`, `g` etc) are there because when I am rapidly testing solutions, I have often typed those and failed. So totally for my convenience that one.

Note, empty `input.txt` is create after `gen` (or `g`) command for now. You will need to copy/paste the input data to the file.

## FAQ

**What version of Go do I need to use for these?**

I am using `embed` and `generics` to solve problems. So at least go `1.18` is needed to run these. 

## Progress - 157 of 450

| Day | [2015](year15) | [2016](year16) | [2017](year17) | [2018](year18) | [2019](year19) | [2020](year20) | [2021](year21) | [2022](year22) | [2023](year23) |
|:---:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:| :-: |
| :star: | 37 | 26 | 15 | 11 | 12 | 18 | 19 | 15 | 2 |
|1| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal:| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: |
|2| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | |
|3| :1st_place_medal: | :1st_place_medal: | | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | |
|4| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | |
|5| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | |
|6| :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | :1st_place_medal: | |
|7| :1st_place_medal: | :1st_place_medal: | | :2nd_place_medal: | | :1st_place_medal: |  :1st_place_medal: | :1st_place_medal: | |
|8| :1st_place_medal: | | :1st_place_medal: | | | :1st_place_medal: | | | |
|9| :1st_place_medal: | | :1st_place_medal: | | | :1st_place_medal: | | | |
|10| :1st_place_medal: | | | | | | :1st_place_medal: | | |
|11| :1st_place_medal: | | | | | | :1st_place_medal: | | |
|12| :1st_place_medal: | :1st_place_medal: | | | | | | :1st_place_medal: | |
|13| :1st_place_medal: | :1st_place_medal: | | | | | | | |
|14| :1st_place_medal: | | | | | | | | |
|15| :1st_place_medal: | | | | | | | | |
|16| | :1st_place_medal: | | | | | | | |
|17| :1st_place_medal: | | | | | | | | |
|18| | | | | | | | | |
|19| | | | | | | | | |
|20| :1st_place_medal: | :1st_place_medal: | | | | | | | |
|21| | :1st_place_medal: | | | | | | | |
|22| | | | | | | | | |
|23| | :2nd_place_medal: | | | | | | | |
|24| :1st_place_medal: | | | | | | | | |
|25| :2nd_place_medal: | :2nd_place_medal: | :2nd_place_medal: | | | | :2nd_place_medal: | :2nd_place_medal: | |
