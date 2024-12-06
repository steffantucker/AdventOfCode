namespace DaysTests;

public class Day04Tests
{
    public void Part1_returns_expected(string[][] grid, int expected)
    {
        int actual = Day04Solution.Day04.Part1(grid);
        Assert.That(actual, Is.EqualTo(expected), "$\"Expected result to be {expected}, but was {result}\"");
    }
}