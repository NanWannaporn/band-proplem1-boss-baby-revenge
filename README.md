## Problem 1: Boss Baby's Revenge 

This repository checks if Boss Baby's behavior has revenge on the neighbourhood kids who shoot water guns at his house.
Input is a string containing 'S' and 'R', which represents a shot from the neighborhood and Boss Baby's revenge.
Ouput is Boss Baby's behavior as either Good boy or Bad boy.

**Built with** 
- Go

### Get Started
check if Go is already installed on your computer by running the following command
```
go version
```
if Go is not installed, please install it first.

**Installation**
1. Clone this repository
```
git clone https://github.com/NanWannaporn/band-proplem1-boss-baby-revenge.git
cd band-proplem1-boss-baby-revenge
```
2. Download Dependencies
```
go mod tiny
```
3. Run project
```
go run .
```

**Output**
```
Input:  SRSSRRR -> Output: Good boy
Input:  RSSRR -> Output: Bad boy
Input:  SSSRRRRS -> Output: Bad boy
Input:  SRRSSR -> Output: Bad boy
Input:  SSRSRR -> Output: Good boy
Input:  C -> Output: Invalid
Input:  R -> Output: Bad boy
Input:  S -> Output: Bad boy
Input:  RSSRR -> Output: Bad boy
```
