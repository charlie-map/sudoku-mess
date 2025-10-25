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

### Examples:

```bash
charlie-map@me speed-test % go run . 1 3 2
Building 1 tests... 3x2
Ready?
+---+---+---+
| 6 | 1 | 2 |
+---+---+---+
| 5 |   | 3 |
+---+---+---+
Answer: 4
Well done! Your average response time was 867.139333ms ms
```

You can do any shape as well, so if you want to practice your columns or rows:

```bash
charlie-map@me speed-test % go run . 1 1 6
Building 1 tests... 1x6
Ready?
+---+
| 1 |
+---+
| 6 |
+---+
| 3 |
+---+
| 5 |
+---+
| 2 |
+---+
|   |
+---+
Answer: 4
Well done! Your average response time was 1.91164975s ms
```
