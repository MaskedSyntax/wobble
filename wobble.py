#!/bin/python3

from ascii_art import ascii_art


# Function to print letters side by side
def print_ascii_side_by_side(letters):
    # Prepare each letter's ASCII art
    letter_lines = {letter: art.splitlines() for letter, art in ascii_art.items()}
    
    # Print each row of ASCII art side by side
    for i in range(len(next(iter(letter_lines.values())))):  # Loop over rows
        row = "  ".join(letter_lines[letter][i] for letter in letters)
        print(row)

letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
text = "Hello I am Aftaab"
print_ascii_side_by_side(text.upper())
