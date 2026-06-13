```markdown
# 📋 Hasil Uji Kompetensi — Bagian B: Modul Terpisah

Hasil pengerjaan projek pada soal test yang mencakup test kemampuan logika, API, dan develop tampilan telah dilakukan dalam struktur kode berikut:
1. b1-csharp
2. b2-golang
3. b3-angular
4. b4-vuejs

---

## 🛠️ Versi Runtime & Environment
Aplikasi ini dikembangkan dan diuji menggunakan versi lingkungan standar berikut untuk memastikan stabilitas eksekusi:
* **.NET SDK Ecosystem:** .NET v8.0 / v9.0 (C# 12)
* **Go Engine Environment:** go1.22 atau versi di atasnya
* **Node.js Environment:** v22.11.0 (LTS)
* **Package Manager:** npm v10.x / v8.19.1+

---

## 🚀 Panduan Menjalankan Masing-Masing Soal (Step-by-Step)

Pastikan posisi terminal Anda berada pada akar direktori utama (`bagian-b/`) sebelum menelusuri sub-modul di bawah ini.

### Soal B1 — C# (Rekap Absensi Karyawan)
Program berbasis konsol (*Console Application*) yang memproses kalkulasi log mentah presensi secara sekuensial, menghitung metrik persentase kehadiran, serta mendeteksi anomali tren ketidakhadiran beruntun.
1. Masuk ke direktori proyek C#:
   cd b1-csharp
2. untuk menjalankan program, ketik kode
 dotnet run
3. Output yang dihasilkan:
PS D:\bagian-b\b1-csharp> dotnet run
=== REKAP ABSENSI ===
E001 - Budi   : Hadir=2, Izin=1, Alpha=0 | Kehadiran: 66,67%
E002 - Sari   : Hadir=1, Izin=0, Alpha=2 | Kehadiran: 33,33% ⚠ Perlu Perhatian
E003 - Dani   : Hadir=2, Izin=1, Alpha=0 | Kehadiran: 66,67%

Karyawan paling sering Alpha: Sari (2 kali)

4. Logika Kode:
### Berisikan kumpulan data karyawan
records.GroupBy(r => r.EmployeeId);

### Variabel untuk mengambil ID dan nama karyawan
var empId = group.Key;
var name = group.First().Name;

### Logika untuk menghitung jumlah hadir, izin, alpha, total absensi dan persentase kehadiran
int hadir = group.Count(r => r.Status == "Hadir");
int izin = group.Count(r => r.Status == "Izin");
int alpha = group.Count(r => r.Status == "Alpha");
int totalHari = group.Count();
double persentaseHadir = totalHari > 0
    ? ((double)hadir / totalHari) * 100
    : 0;

### Logika mendeteksi alpha berturut
bool perluPerhatian = false;
var sortedRecords = group.OrderBy(r => r.Date).ToList(); -> urutkan berdasarkan tanggal
for (int i = 0; i < sortedRecords.Count - 1; i++) -> periksa masing-masing record

### Soal B2 — Golang (REST API Manajemen Cuti)
Layanan Backend RESTful API menggunakan HTTP Web Framework Gin Gonic dengan penyimpanan internal (Slice In-Memory).
1. Masuk ke direktori proyek Go:
cd b2-golang

2. Sinkronisasi dependensi modul:
go mod tidy

3. Jalankan server lokal:
go run main.go

4. Penjelasan Kode:
Logika Validasi Overlap:
Mencegah satu karyawan memiliki dua pengajuan cuti dengan tanggal yang tumpang tindih:
for _, req := range leaveRequests {
    if req.EmployeeID == newRequest.EmployeeID && req.Status != "Rejected" {
        if newRequest.StartDate <= req.EndDate && newRequest.EndDate >= req.StartDate {
            c.JSON(http.StatusBadRequest, gin.H{"error": "karyawan sudah memiliki pengajuan cuti yang overlap pada tanggal tersebut"})
            return
        }
    }
}

### Soal B3 — Angular (Halaman Direktori Karyawan)
Aplikasi direktori karyawan menggunakan Angular 18+ dengan fitur pencarian interaktif responsif berbasis Angular Signals.
1. Masuk ke direktori projek Angular
cd b3-angular

2. Memasang paket dependensi npm (jika belum pernah run project)
npm install

3. Jalankan server di lokal (http://localhost:4200)
npm run start

4. Penjelasan kode filtering
Untuk filtering data nama, departemen dan status aktif menerapkan fungsi computed()

this.filteredEmployees = computed(() => {
  const search = this.searchTerm().toLowerCase();
  const dept = this.selectedDepartment();
  const showInactive = this.includeInactive();

  return this.employees().filter(emp => {
    const matchName = emp.name.toLowerCase().includes(search);
    const matchDept = dept === 'All' || emp.department === dept;
    const matchActive = showInactive ? true : emp.active;
    return matchName && matchDept && matchActive;
  });
});

### Soal B4— Vue.js (Form Onboarding Karyawan Baru)
Menggunakan Vue 3 (Composition API) dan bundler Vite 5 yang mempertahankan state data antar tahapan navigasi.
1. Masuk direktori projek Vue.js
cd b4-vuejs

2. Pasang paket dependensi (jika belum pernah run)
npm install

3. Jalankan server di lokal (http://localhost:5173)
npm run dev

4. Output:
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2026/06/13 - 14:01:03 | 201 |   1.69ms |             ::1 | POST     "/leave"
[GIN] 2026/06/13 - 14:06:18 | 200 |       0s |             ::1 | GET      "/leave"
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2026/06/13 - 14:01:03 | 201 |   1.69ms |             ::1 | POST     "/leave"
[GIN] 2026/06/13 - 14:06:18 | 200 |       0s |             ::1 | GET      "/leave"
[GIN] 2026/06/13 - 14:07:19 | 201 |       0s |             ::1 | POST     "/leave"
[GIN] 2026/06/13 - 14:07:49 | 201 |       0s |             ::1 | POST     "/leave"
[GIN] 2026/06/13 - 14:07:57 | 200 |       0s |             ::1 | GET      "/leave"
[GIN] 2026/06/13 - 14:08:57 | 409 |       0s |             ::1 | POST     "/leave"
[GIN] 2026/06/13 - 14:15:35 | 200 |       0s |             ::1 | PUT      "/leave/1/approve"
[GIN] 2026/06/13 - 14:20:33 | 200 |       0s |             ::1 | GET      "/leave/employee/E003"

Link Github: https://github.com/fadiana1/c-angular-golang-project.git