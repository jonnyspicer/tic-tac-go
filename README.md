# ❌⭕❌ Tic-Tac-Go ⭕❌⭕

### About 📑

I wrote tic-tac-go as part of my efforts to learn Go; it is in fact the
first program I've written from scratch in the language, and as such is
exceptionally basic. I mainly wanted to gain some familiarity with structs
and methods, hence it being very light on logic. I could've coded the
best move for the computer at each point without just picking randomly,
but it would've been a lot of repetition and I wouldn't have learned
anything new, plus this way you get the pleasure of winning virtually
every time 😊.

### Installation 🚀

If you don't already have Go installed on your machine, follow the
instructions found [here.](https://golang.org/doc/install) Clone the repo, navigate to the directory and run ```go run main.go``` then simply follow the instructions
in your CLI.

### Possible Improvements 🛠

Obviously there are a lot of things that could be improved here, I am a Go novice after all.
Some of the most obvious ones are:

- Add error handling;
- Add tests;
- Add functionality for two human players;
- Add logic to computer moves (even simple heuristics like blocks if two in a row);
- Refactor the board state to be a multidimensional array, which would make it less tedious to check if there's a winner;
- Refactor the way winner and draw checking works (I completely forgot about draws until I'd finished the rest of the code, so it was tacked on at the end);
- Abstract some of the functionality out of main(), at the moment it's doing more than one thing.