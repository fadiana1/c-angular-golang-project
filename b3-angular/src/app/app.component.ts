import { Component, computed, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

// Mock data sesuai dengan spesifikasi soal B3
export const EMPLOYEES = [
  { id: 1, name: 'Budi Santoso',   department: 'Engineering', position: 'Backend Dev',   phone: '081234567890', active: true  },
  { id: 2, name: 'Sari Dewi',      department: 'HR',          position: 'HR Specialist', phone: '082345678901', active: true  },
  { id: 3, name: 'Dani Pratama',   department: 'Engineering', position: 'Frontend Dev',  phone: '083456789012', active: false },
  { id: 4, name: 'Rina Wati',      department: 'Finance',     position: 'Accountant',    phone: '084567890123', active: true  },
  { id: 5, name: 'Joko Susilo',    department: 'Engineering', position: 'DevOps',        phone: '085678901234', active: true  },
];

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  // Menggunakan Angular Signals untuk Reactive Programming sesuai penilaian
  employees = signal(EMPLOYEES);
  searchTerm = signal('');
  selectedDept = signal('');
  showInactive = signal(true); // Default true agar semua tampil, bisa di-toggle
  selectedEmployee = signal<any | null>(null);

  departments = ['Engineering', 'HR', 'Finance'];

  // Fungsi Reactive Filtering menggunakan 'computed' signal
  filteredEmployees = computed(() => {
    return this.employees().filter(emp => {
      const matchName = emp.name.toLowerCase().includes(this.searchTerm().toLowerCase());
      const matchDept = this.selectedDept() ? emp.department === this.selectedDept() : true;
      const matchActive = this.showInactive() ? true : emp.active;
      
      return matchName && matchDept && matchActive;
    });
  });

  selectEmployee(emp: any) {
    this.selectedEmployee.set(emp);
  }
}