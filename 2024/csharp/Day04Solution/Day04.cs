using Utils;

namespace Day04Solution;

public class Day04
{
    private const int UP = -1;
    private const int DOWN = 1;
    private const int LEFT = -1;
    private const int RIGHT = 1;
    private const int CENTER = 0;
    private static string[] grid = Array.Empty<string>();

    static void Main(string[] args)
    {
        grid = Input.GetStrings("04", "");
        
        Console.WriteLine($"Part 1: {Part1(grid)}");
    }

    public static int Part1(string[] input)
    {
        if (input.Length != 0) grid = input;
        int count = 0;
        for (int row = 0; row < grid.Length; row++)
        {
            for (int col = 0; col < grid[row].Length; col++)
            {
                if (grid[row][col] != 'X') continue;
                count += CheckGrid(row, col, UP, LEFT);
                count += CheckGrid(row, col, UP, CENTER);
                count += CheckGrid(row, col, UP, RIGHT);
                count += CheckGrid(row, col, CENTER, LEFT);
                count += CheckGrid(row, col, CENTER, RIGHT);
                count += CheckGrid(row, col, DOWN, LEFT);
                count += CheckGrid(row, col, DOWN, CENTER);
                count += CheckGrid(row, col, DOWN, RIGHT);
            }
        }
        
        return count;
    }

    private static int CheckGrid(int row, int col, int rowDir, int colDir, string target = "XMAS")
    {
        foreach (var t in target)
        {
            if (row < 0 || row >= grid.Length) return 0;
            if (col < 0 || col >= grid[row].Length) return 0;
            if (grid[row][col] != t) return 0;
            row += rowDir;
            col += colDir;
        }

        return 1;
    }
}