namespace Day02;

public class Solutions
{
    static void Main(string[] args)
    {
        int[][] input = Utils.Input.GetMatrixInput("02");
        int part1 = input.Where(IsSafe).Count();
        Console.WriteLine($"Part 1: {part1}");
        int part2 = input.Where(IsPermissivelySafe).Count();
        Console.WriteLine($"Part 2: {part2}");
    }

    public static bool IsSafe(int[] report)
    {
        // report.Zip(report.Skip(1), (a, b) => a - b);
        // then all(x => x >= -3 && x <= -1) || all(x => x >= 1 && x <= 3)
        // ^ entirety of safe check, if run on diffs
        bool isDecreasing = (report[0] > report[1]);
        for (int i = 0; i < report.Length - 1; i++)
        {
            int diff = report[i] - report[i + 1];
            switch (diff)
            {
                case < 0 when isDecreasing:
                case > 0 when !isDecreasing:
                    return false;
            }

            if (Math.Abs(diff) > 3 || Math.Abs(diff) < 1) return false;
        }
        return true;
    }

    public static bool IsPermissivelySafe(int[] report)
    {
        return IsSafe(report) || Enumerable.Range(0, report.Count())
            .Any(i => IsSafe(report.Take(i).Concat(report.Skip(i+1)).ToArray()));
    }
}