import { readFileSync } from "fs";

// Part 1
let Part1 = () => {
  const data = readFileSync("./numbers1.csv", { encoding: "utf8", flag: "r" });
  const str = data.split("\n");
  const numbers = str.map((n) => parseInt(n));
  const targetNumber = 2020;

  for (let i = 0; i < numbers.length; i++) {
    for (let j = numbers.length + 1; j > i; j--) {
      if (numbers[i] + numbers[j] == targetNumber) {
        console.log(`${numbers[i]} + ${numbers[j]} = ${numbers[i] + numbers[j]}`);
        console.log(`${numbers[i]} * ${numbers[j]} = ${numbers[i] * numbers[j]}`);
        process.exit();
      }
    }
  }
};

//Part 2
let Part2 = () => {
  const data = readFileSync("./numbers2.csv", { encoding: "utf8", flag: "r" });
  const str = data.split("\n");
  const numbers = str.map((n) => parseInt(n));
  const targetNumber = 2020;

  for (let i = 0; i < numbers.length; i++) {
    for (let j = numbers.length + 1; j > i; j--) {
      for (let k = i + 1; k < j; k++) {
        if (numbers[i] + numbers[j] + numbers[k] == targetNumber) {
          console.log(`${numbers[i]} + ${numbers[j]} + ${numbers[k]} = ${numbers[i] + numbers[j] + numbers[k]}`);
          console.log(`${numbers[i]} * ${numbers[j]} * ${numbers[k]} = ${numbers[i] * numbers[j] * numbers[k]}`);
          process.exit();
        }
      }
    }
  }
};

Part2();