namespace AoC_2015;

public class Day01
{
    public static int Part1(string input)
    {
        int floor = 0;
        foreach (char c in input)
        {
            switch (c)
            {
                case '(': floor++; break;
                case ')': floor--; break;
            }
        }
        return floor;
    }
}