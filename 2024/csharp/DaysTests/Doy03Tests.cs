namespace DaysTests;

public class Doy03Tests
{
    [TestCase("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161)]
    public void Part1_returns_expected_result(string input, int expected)
    {
        int actual = Day03.Day03.Part1(input);
        Assert.That(actual, Is.EqualTo(expected), $"Expected result to be {expected}, but was {actual}");
    }

    [TestCase("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48)]
    public void Part2_returns_expected_result(string input, int expected)
    {
        int actual = Day03.Day03.Part2(input);
        Assert.That(actual, Is.EqualTo(expected), $"Expected result to be {expected}, but was {actual}");
    }
}