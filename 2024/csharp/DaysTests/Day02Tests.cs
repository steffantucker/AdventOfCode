using Day02;

namespace DaysTests;

public class Day02Tests
{
    [TestCase(new[] { 1, 2, 7, 8, 9 })]
    [TestCase(new[] { 9, 7, 6, 2, 1 })]
    [TestCase(new[] { 1, 3, 2, 4, 5 })]
    [TestCase(new[] { 8, 6, 4, 4, 1 })]
    [TestCase(new[] { 1, 1, 2 })]
    [TestCase(new[] { 2, 6, 1 })]
    [TestCase(new[] { 1, 2, 5, 9 })]
    public void IsSafe_ReturnsFalse(int[] input)
    {
        bool actual = Day02.Day02.IsSafe(input);
        Assert.That(actual, Is.False, $"{input} should not have been safe");
    }

    [TestCase(new[] { 7, 6, 4, 2, 1 })]
    [TestCase(new[] { 1, 3, 6, 7, 9 })]
    public void IsSafe_ReturnsTrue(int[] input)
    {
        bool actual = Day02.Day02.IsSafe(input);
        Assert.That(actual, Is.True, $"{input} should have been safe");
    }

    [TestCase(new[] { 1, 2, 7, 8, 9 })]
    [TestCase(new[] { 9, 7, 6, 2, 1 })]
    public void IsPermissivelySafe_ReturnsFalse(int[] input)
    {
        bool actual = Day02.Day02.IsPermissivelySafe(input);
        Assert.That(actual, Is.False, $"{input} should not have been safe");
    }

    [TestCase(new[] { 7, 6, 4, 2, 1 })]
    [TestCase(new[] { 1, 3, 2, 4, 5 })]
    [TestCase(new[] { 8, 6, 4, 4, 1 })]
    [TestCase(new[] { 1, 3, 6, 7, 9 })]
    [TestCase(new[] { 1, 1, 2 })]
    [TestCase(new[] { 2, 6, 1 })]
    [TestCase(new[] { 1, 2, 5, 9 })]
    public void IsPermissivelySafe_ReturnsTrue(int[] input)
    {
        bool actual = Day02.Day02.IsPermissivelySafe(input);
        Assert.That(actual, Is.True, $"{input} should have been safe");
    }
}