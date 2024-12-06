using System.Text.RegularExpressions;
using Utils;

namespace Day03;

public class Day03
{
    static void Main(string[] args)
    {
        string input = Input.GetInput("03");
        int part1 = Part1(input);
        Console.WriteLine($"Part 1: {part1}");
        int part2 = Part2(input);
        Console.WriteLine($"Part 2: {part2}");
    }

    public static int Part1(string input)
    {
        int result = 0;
        const string mulPattern = @"mul\((\d{1,3}),(\d{1,3})\)";
        foreach (Match match in Regex.Matches(input, mulPattern, RegexOptions.NonBacktracking))
        {
            result += int.Parse(match.Groups[1].Value) * int.Parse(match.Groups[2].Value);
        }

        return result;
    }

    public static int Part2(string input)
    {
        int result = 0;
        const string mulPattern = @"mul\((\d{1,3}),(\d{1,3})\)";
        const string doPattern = @"do\(\)";
        const string dontPattern = @"don't\(\)";
        List<Match> allMatches = [];
        allMatches.AddRange(Regex.Matches(input, mulPattern, RegexOptions.NonBacktracking));
        allMatches.AddRange(Regex.Matches(input, doPattern, RegexOptions.NonBacktracking));
        allMatches.AddRange(Regex.Matches(input, dontPattern, RegexOptions.NonBacktracking));
        allMatches.Sort((a, b) => a.Index.CompareTo(b.Index));
        bool shouldDo = true;
        foreach (Match match in allMatches)
        {
            if (match.Groups.Count == 3 && shouldDo)
            {
                result += int.Parse(match.Groups[1].Value) * int.Parse(match.Groups[2].Value);
            }
            else if (match.Value == "do()")
            {
                shouldDo = true;
            }
            else
            {
                shouldDo = false;
            }
        }

        return result;
    }
}