using System;
using System.Collections.Generic;
using System.Linq;

public class AttendanceRecord
{
    public string EmployeeId { get; set; }
    public string Name       { get; set; }
    public string Date       { get; set; }
    public string Status     { get; set; } // "Hadir", "Izin", "Alpha"
}

class Program
{
    static void Main()
    {
        var records = new List<AttendanceRecord>
        {
            new AttendanceRecord { EmployeeId = "E001", Name = "Budi",  Date = "2024-01-02", Status = "Hadir" },
            new AttendanceRecord { EmployeeId = "E001", Name = "Budi",  Date = "2024-01-03", Status = "Izin"  },
            new AttendanceRecord { EmployeeId = "E001", Name = "Budi",  Date = "2024-01-04", Status = "Hadir" },
            new AttendanceRecord { EmployeeId = "E002", Name = "Sari",  Date = "2024-01-02", Status = "Hadir" },
            new AttendanceRecord { EmployeeId = "E002", Name = "Sari",  Date = "2024-01-03", Status = "Alpha" },
            new AttendanceRecord { EmployeeId = "E002", Name = "Sari",  Date = "2024-01-04", Status = "Alpha" },
            new AttendanceRecord { EmployeeId = "E003", Name = "Dani",  Date = "2024-01-02", Status = "Hadir" },
            new AttendanceRecord { EmployeeId = "E003", Name = "Dani",  Date = "2024-01-03", Status = "Hadir" },
            new AttendanceRecord { EmployeeId = "E003", Name = "Dani",  Date = "2024-01-04", Status = "Izin"  },
        };

        var groupedRecords = records.GroupBy(r => r.EmployeeId);
        var summaryList = new List<(string Name, int AlphaCount)>();

        Console.WriteLine("=== REKAP ABSENSI ===");

        foreach (var group in groupedRecords)
        {
            var empId = group.Key;
            var name = group.First().Name;
            
            int hadir = group.Count(r => r.Status == "Hadir");
            int izin = group.Count(r => r.Status == "Izin");
            int alpha = group.Count(r => r.Status == "Alpha");
            int totalHari = group.Count();

            double persentaseHadir = totalHari > 0 ? ((double)hadir / totalHari) * 100 : 0;

            // Logika deteksi Alpha berturut-turut
            bool perluPerhatian = false;
            var sortedRecords = group.OrderBy(r => r.Date).ToList();
            for (int i = 0; i < sortedRecords.Count - 1; i++)
            {
                if (sortedRecords[i].Status == "Alpha" && sortedRecords[i + 1].Status == "Alpha")
                {
                    perluPerhatian = true;
                    break;
                }
            }

            string labelPerhatian = perluPerhatian ? " ⚠ Perlu Perhatian" : "";
            Console.WriteLine($"{empId} - {name}   : Hadir={hadir}, Izin={izin}, Alpha={alpha} | Kehadiran: {persentaseHadir:F2}%{labelPerhatian}");

            summaryList.Add((name, alpha));
        }

        Console.WriteLine();
        var palingSeringAlpha = summaryList.OrderByDescending(s => s.AlphaCount).FirstOrDefault();
        Console.WriteLine($"Karyawan paling sering Alpha: {palingSeringAlpha.Name} ({palingSeringAlpha.AlphaCount} kali)");
    }
}