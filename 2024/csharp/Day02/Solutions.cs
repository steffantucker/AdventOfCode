namespace Day02;

public class Solutions
{
    static void Main(string[] args)
    {
        int[][] input = Utils.Input.GetMatrixInput("02");
        int part1 = input.Select(IsExcessivelySafe).Count(b => b);
        Console.WriteLine($"Part 1: {part1}");
    }

    public static bool IsExcessivelySafe(int[] report)
    {
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
        if (IsExcessivelySafe(report)) return true;
        return true;
    }
}