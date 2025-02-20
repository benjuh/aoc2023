# Advent of Code 2023

This is my solutions for [Advent of Code 2023](https://adventofcode.com/2023)

If you pay attention to the commit dates, you'll notice that I am doing this in 2025.
I did the 2024 AoC back in Jan 2025 -> Feb 2025 and I loved it so I decided to do the previous year in Go.

# Running

```bash
# run a single day
go run src/day01/main.go

# I will make a script that runs all the days at once later
```
# Solutions

## Day 1

[Problem](https://adventofcode.com/2023/day/1)

For this I utilized some built-in string functions like `strconv.Atoi`, `strings.Index`, and `strings.LastIndex`.
Used a map to map string like "two" to the number 2. Used an array of size `len(string)` where the value at each element is the number found at that index.

For example:

```
string = "twoonethree5"
array = [2,0,0,1,0,0,3,0,0,0,0,5]

output => 25
```

then I just took the far left non-zero element of the array as my "left" and the far right non-zero element as my "right".

the problem then needed the answer in the form of left + right as a 2 digit number. i.e. => (left * 10) + right

Time Complexity: O(n)
Space Complexity: O(n)

## Day 2

[Problem](https://adventofcode.com/2023/day/2)

Parsing was the hardest part of this one. My approach was to make a struct `Game` that has fields of an `id` and a `game_states` array that keeps track of used [R, G, B] dice in the game state. Then I can iterate over each game state in the game and make sure they meet the limits given in the problem. This approach amde it really easy to complete part 2.

Time Complexity: O(n*m) where n is the number of games and m is the number of times each game is played

Space Complexity: O(n*m) for each game `n` there is an array of `m` game states in which each game state is an array of size 3 to represent number of R,G,B dice used in each game.

## Day 3

[Problem](https://adventofcode.com/2023/day/3)

Iterated over the input string and made a map of Ranges (Range is a struct that has an `i` index *specifically the row of the number*, a `start` index, an `end` index, and a `final_val` value) that keeps track of the range of numbers that are used in the string. 

Then I iterated over the map of ranges looked for if at any point there was a symbol adjacent to the range. If there was, I added the final value of the range to the sum.

For part 2 this was the same but instead of there was a symbol adjacent to the range AND the symbol was a `*` I added it to a 
`potential_gears` map that mapped a Point to a list of adjacent numbers. Then I iterated over the map of potential gears and for each potential_gear that had EXACTLY 2 adjacent numbers, multiplied those two numbers together and added them to the return sum.

Time Complexity: O(n) where n is the length of the input string
Space Complexity: O(n*m) where n is the length of the input string and m is the amount of numbers in the string
