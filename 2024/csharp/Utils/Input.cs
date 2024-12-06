namespace Utils;

public class Input
{
    // GetInput takes a zero-padded double-digit number for the day as a string
    public static string GetInput(string day)
    {
        string appDataFolder = Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData);
        return File.ReadAllText(Path.Combine(appDataFolder, "AoC", "2024", "day"+day));
    }

    public static int[][] GetNumberMatrix(string day, string sep)
    {
        string input = GetInput(day);
        return input.Split(["\r\n", "\n"], StringSplitOptions.TrimEntries | StringSplitOptions.RemoveEmptyEntries)
            .Select(x => x.Trim()
                .Split(sep, StringSplitOptions.RemoveEmptyEntries)
                .Select(int.Parse)
                .ToArray())
            .ToArray();
    }

    public static string[] GetStrings(string day, string sep)
    {
        string input = GetInput(day);
        return input.Split(["\r\n", "\n"], StringSplitOptions.TrimEntries | StringSplitOptions.RemoveEmptyEntries)
            .Select(x => x.Trim())
            .ToArray();
    }
}