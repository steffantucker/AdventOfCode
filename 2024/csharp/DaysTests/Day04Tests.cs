using Utils;

namespace DaysTests;

public class Day04Tests
{
    [TestCase(@"MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX", 18)]
    public void Part1_returns_expected(string input, int expected)
    {
        string[] grid = input.Split("\n");
        int actual = Day04Solution.Day04.Part1(grid);
        Assert.That(actual, Is.EqualTo(expected), "$\"Expected result to be {expected}, but was {result}\"");
    }
}