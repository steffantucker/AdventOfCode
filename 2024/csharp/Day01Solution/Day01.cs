class Day01
{
    static void Main(string[] args)
    {
        string appDataFolder = Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData);
        string file = File.ReadAllText(Path.Combine(appDataFolder, "AoC2024", "day01"));
        int part1 = Part1(file);
        int part2 = Part2(file);
        Console.WriteLine($"Part 1: {part1}");
        Console.WriteLine($"Part 2: {part2}");
    }

    private static int Part1(string input)
    {
        List<int> left = [];
        List<int> right = [];
        foreach (string line in input.Split(["\r\n", "\n"],
                     StringSplitOptions.TrimEntries | StringSplitOptions.RemoveEmptyEntries))
        {
            int[] nums = line.Split(" ", StringSplitOptions.RemoveEmptyEntries)
                .Select(int.Parse).ToArray();
            left.Add(nums[0]);
            right.Add(nums[1]);
        }
        left.Sort();
        right.Sort();
        return left.Zip(right, (l,r) => Math.Abs(l - r)).Sum();
    }

    private static int Part2(string input)
    {
        List<int> left = [];
        Dictionary<int, int> right = [];
        foreach (string line in input.Split(["\r\n","\n"], StringSplitOptions.TrimEntries|StringSplitOptions.RemoveEmptyEntries))
        {
            int[] nums = line.Split(" ", StringSplitOptions.RemoveEmptyEntries)
                .Select(int.Parse).ToArray();
            left.Add(nums[0]);
            if (right.TryGetValue(nums[1], out int rightValue)) right[nums[1]] = rightValue+1;
            else right.Add(nums[1], 1);
        }

        return left.Select(val => val * (right.GetValueOrDefault(val, 0))).Sum();
    }
}