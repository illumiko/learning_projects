/* Just some code based on the aritcals preceeding challenge */
Random dice = new Random();
int roll = dice.Next();
int roll1 = dice.Next(1,7);
int roll2 = dice.Next(1,7);
/* Console.WriteLine($"{roll},{roll1},{roll2}"); */

/* Code challenge: Implement a method of the Math class that returns the larger of two numbers */

int firstValue = 500;
int secondValue = 600;
int largerValue = Math.Max(firstValue,secondValue);

Console.WriteLine(largerValue);
