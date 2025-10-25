# 🚀 ROCKET'S REACTION / SUDOKU SPEED TEST 🚀

This little program is designed to test how fast your meat-brain can process information and enter a specific number. Think of it as a crucial emergency procedure where you need to input the **missing Sudoku digit** before the bomb goes off.

---

## 💥 The Mission

The goal is simple: **Enter the Secret Code** as fast as you can.

The test runs for a set number of rounds (see [launch instructions](## 🛠️ How to Launch This Thing)). The moment the "Enter Code:" prompt appears, the clock starts. The clock only stops when you correctly type the missing number and hit Enter. If you mess up, the clock keeps ticking while the program yells at you to try again.

---

## 🛠️ How to Launch This Thing

You need the Go language installed, obviously.

You'll be passing in 3 values:

1. The number of tests to run - this is how many times a test grid will be spat out you have to answer.
2. The width of the grid
3. The height of the grid

```bash
go run . 5 3 3
```
