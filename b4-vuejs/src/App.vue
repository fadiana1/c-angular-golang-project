<script setup>
import { ref, reactive, computed } from 'vue'

// State Utama untuk Navigasi Wizard (Step 1, 2, atau 3)
const currentStep = ref(1)

// 1. STATE FORM UTAMA (Semua field digabung agar state terjaga saat bolak-balik)
const formData = reactive({
  // Step 1: Data Pribadi
  fullName: '',
  nik: '',
  birthDate: '',
  gender: '',

  // Step 2: Data Pekerjaan
  department: '',
  position: '',
  startDate: '',
  baseSalary: null
})

const departments = ['Engineering', 'HR', 'Finance', 'Marketing']

// 2. VALIDASI REACTIVE (Menggunakan Computed agar tombol 'Next' otomatis mengunci/membuka)
const step1Errors = computed(() => {
  const errors = {}
  
  // Validasi Nama Lengkap
  if (!formData.fullName) {
    errors.fullName = 'Nama lengkap wajib diisi.'
  } else if (formData.fullName.length < 3) {
    errors.fullName = 'Nama lengkap minimal 3 karakter.'
  }

  // Validasi NIK (Harus 16 digit angka)
  const nikhRegex = /^\d{16}$/
  if (!formData.nik) {
    errors.nik = 'NIK wajib diisi.'
  } else if (!nikhRegex.test(formData.nik)) {
    errors.nik = 'NIK harus tepat 16 digit angka.'
  }

  // Validasi Tanggal Lahir (Minimal 18 Tahun)
  if (!formData.birthDate) {
    errors.birthDate = 'Tanggal lahir wajib diisi.'
  } else {
    const birth = new Date(formData.birthDate)
    const today = new Date()
    let age = today.getFullYear() - birth.getFullYear()
    const monthDiff = today.getMonth() - birth.getMonth()
    if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
      age--
    }
    if (age < 18) {
      errors.birthDate = 'Usia minimal harus 18 tahun.'
    }
  }

  // Validasi Jenis Kelamin
  if (!formData.gender) {
    errors.gender = 'Jenis kelamin wajib dipilih.'
  }

  return errors
})

const step2Errors = computed(() => {
  const errors = {}

  // Validasi Departemen
  if (!formData.department) {
    errors.department = 'Departemen wajib dipilih.'
  }

  // Validasi Posisi
  if (!formData.position) {
    errors.position = 'Posisi wajib diisi.'
  }

  // Validasi Tanggal Mulai Kerja (Tidak boleh tanggal yang sudah lewat)
  if (!formData.startDate) {
    errors.startDate = 'Tanggal mulai kerja wajib diisi.'
  } else {
    const start = new Date(formData.startDate).setHours(0,0,0,0)
    const today = new Date().setHours(0,0,0,0)
    if (start < today) {
      errors.startDate = 'Tanggal mulai kerja tidak boleh tanggal yang sudah lewat.'
    }
  }

  // Validasi Gaji Pokok (Minimal 3 Juta)
  if (formData.baseSalary === null || formData.baseSalary === '') {
    errors.baseSalary = 'Gaji pokok wajib diisi.'
  } else if (isNaN(formData.baseSalary)) {
    errors.baseSalary = 'Gaji pokok harus berupa angka.'
  } else if (formData.baseSalary < 3000000) {
    errors.baseSalary = 'Gaji pokok minimal Rp 3.000.000.'
  }

  return errors
})

// Rule Tambahan 1: Tombol "Next" disabled jika tidak valid
const isStep1Valid = computed(() => Object.keys(step1Errors.value).length === 0)
const isStep2Valid = computed(() => Object.keys(step2Errors.value).length === 0)

// Fungsi Navigasi Tombol
const nextStep = () => {
  if (currentStep.value === 1 && isStep1Valid.value) currentStep.value = 2
  else if (currentStep.value === 2 && isStep2Valid.value) currentStep.value = 3
}

const prevStep = () => {
  if (currentStep.value > 1) currentStep.value--
}

// Fungsi Format Rupiah untuk halaman Review
const formatRupiah = (value) => {
  if (!value) return 'Rp 0'
  return 'Rp ' + value.toLocaleString('id-ID')
}

// Submit Akhir
const submitForm = () => {
  alert('Pendaftaran Onboarding Karyawan Baru Berhasil Disubmit!')
  // Reset Form kembali ke Step 1
  currentStep.value = 1
  formData.fullName = ''
  formData.nik = ''
  formData.birthDate = ''
  formData.gender = ''
  formData.department = ''
  formData.position = ''
  formData.startDate = ''
  formData.baseSalary = null
}
</script>

<template>
  <div class="wizard-container">
    <h2>Form Onboarding Karyawan Baru</h2>

    <div class="step-indicator">
      <div :class="['step-box', { active: currentStep >= 1 }]">1. Data Pribadi</div>
      <div class="step-line" :class="{ active: currentStep >= 2 }"></div>
      <div :class="['step-box', { active: currentStep >= 2 }]">2. Data Pekerjaan</div>
      <div class="step-line" :class="{ active: currentStep === 3 }"></div>
      <div :class="['step-box', { active: currentStep === 3 }]">3. Review & Submit</div>
    </div>

    <div class="form-card">
      
      <div v-if="currentStep === 1">
        <h3 class="step-title">Step 1 — Data Pribadi</h3>
        
        <div class="form-group">
          <label>Nama Lengkap <span class="required">*</span></label>
          <input type="text" v-model="formData.fullName" placeholder="Masukkan nama lengkap" />
          <span class="error-msg" v-if="step1Errors.fullName">{{ step1Errors.fullName }}</span>
        </div>

        <div class="form-group">
          <label>NIK (Nomor Induk Kependudukan) <span class="required">*</span></label>
          <input type="text" v-model="formData.nik" placeholder="Contoh: 3201234567890123" maxlength="16" />
          <span class="error-msg" v-if="step1Errors.nik">{{ step1Errors.nik }}</span>
        </div>

        <div class="form-group">
          <label>Tanggal Lahir <span class="required">*</span></label>
          <input type="date" v-model="formData.birthDate" />
          <span class="error-msg" v-if="step1Errors.birthDate">{{ step1Errors.birthDate }}</span>
        </div>

        <div class="form-group">
          <label>Jenis Kelamin <span class="required">*</span></label>
          <div class="radio-group">
            <label class="radio-label">
              <input type="radio" value="Laki-laki" v-model="formData.gender" /> Laki-laki
            </label>
            <label class="radio-label">
              <input type="radio" value="Perempuan" v-model="formData.gender" /> Perempuan
            </label>
          </div>
          <span class="error-msg" v-if="step1Errors.gender">{{ step1Errors.gender }}</span>
        </div>

        <div class="button-group">
          <button type="button" @click="nextStep" :disabled="!isStep1Valid" class="btn btn-next">Next</button>
        </div>
      </div>

      <div v-if="currentStep === 2">
        <h3 class="step-title">Step 2 — Data Pekerjaan</h3>

        <div class="form-group">
          <label>Departemen <span class="required">*</span></label>
          <select v-model="formData.department">
            <option value="">-- Pilih Departemen --</option>
            <option v-for="dept in departments" :key="dept" :value="dept">{{ dept }}</option>
          </select>
          <span class="error-msg" v-if="step2Errors.department">{{ step2Errors.department }}</span>
        </div>

        <div class="form-group">
          <label>Posisi / Jabatan <span class="required">*</span></label>
          <input type="text" v-model="formData.position" placeholder="Contoh: Database Engineer" />
          <span class="error-msg" v-if="step2Errors.position">{{ step2Errors.position }}</span>
        </div>

        <div class="form-group">
          <label>Tanggal Mulai Kerja <span class="required">*</span></label>
          <input type="date" v-model="formData.startDate" />
          <span class="error-msg" v-if="step2Errors.startDate">{{ step2Errors.startDate }}</span>
        </div>

        <div class="form-group">
          <label>Gaji Pokok (Rp) <span class="required">*</span></label>
          <input type="number" v-model.number="formData.baseSalary" placeholder="Minimal 3.000.000" />
          <span class="error-msg" v-if="step2Errors.baseSalary">{{ step2Errors.baseSalary }}</span>
        </div>

        <div class="button-group">
          <button type="button" @click="prevStep" class="btn btn-back">Back</button>
          <button type="button" @click="nextStep" :disabled="!isStep2Valid" class="btn btn-next">Next</button>
        </div>
      </div>

      <div v-if="currentStep === 3">
        <h3 class="step-title">Step 3 — Review Ringkasan Data</h3>
        
        <div class="review-section">
          <h4>Data Pribadi</h4>
          <table class="review-table">
            <tr><th>Nama Lengkap</th><td>{{ formData.fullName }}</td></tr>
            <tr><th>NIK</th><td>{{ formData.nik }}</td></tr>
            <tr><th>Tanggal Lahir</th><td>{{ formData.birthDate }}</td></tr>
            <tr><th>Jenis Kelamin</th><td>{{ formData.gender }}</td></tr>
          </table>

          <h4>Data Pekerjaan</h4>
          <table class="review-table">
            <tr><th>Departemen</th><td>{{ formData.department }}</td></tr>
            <tr><th>Posisi</th><td>{{ formData.position }}</td></tr>
            <tr><th>Mulai Kerja</th><td>{{ formData.startDate }}</td></tr>
            <tr><th>Gaji Pokok</th><td class="salary-text">{{ formatRupiah(formData.baseSalary) }}</td></tr>
          </table>
        </div>

        <div class="button-group">
          <button type="button" @click="prevStep" class="btn btn-back">⚙️ Edit Data</button>
          <button type="button" @click="submitForm" class="btn btn-submit">Submit Data</button>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
.wizard-container {
  max-width: 650px;
  margin: 30px auto;
  padding: 10 45px;
  font-family: 'Segoe UI', system-ui, sans-serif;
  color: #15634d;
}

h2 {
  text-align: center;
  color: #094c3f;
  margin-bottom: 25px;
}

.step-title {
  margin-top: 0;
  border-bottom: 2px solid #e2e8f0;
  padding-bottom: 8px;
  color: #2c3e50;
  font-size: 20px;
  font-weight: bold;
}

.step-indicator {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 25px;
}

.step-box {
  padding: 8px 14px;
  background-color: #edf2f7;
  border-radius: 20px;
  font-weight: 600;
  font-size: 12px;
  color: #718096;
  transition: all 0.3s ease;
}

.step-box.active {
  background-color: #41b883;
  color: white;
}

.step-line {
  flex: 1;
  height: 3px;
  background-color: #edf2f7;
  margin: 0 8px;
}

.step-line.active {
  background-color: #41b883;
}

.form-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 10px;
  padding: 24px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
}

.form-group {
  margin-bottom: 18px;
}

label {
  display: block;
  font-weight: 600;
  margin-bottom: 6px;
  font-size: 14px;
}

.required {
  color: #e53e3e;
}

input[type="text"], input[type="number"], select, input[type="date"] {
  width: 100%;
  padding: 10px;
  border: 1px solid #cbd5e0;
  border-radius: 6px;
  box-sizing: border-box;
  font-size: 14px;
  transition: border-color 0.2s;
}

input:focus, select:focus {
  outline: none;
  border-color: #41b883;
}

.radio-group {
  display: flex;
  gap: 20px;
  padding: 8px 0;
}

.radio-label {
  font-weight: normal;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 6px;
}

.error-msg {
  color: #e53e3e;
  font-size: 12px;
  margin-top: 5px;
  display: block;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin-top: 25px;
}

.btn {
  padding: 10px 22px;
  font-size: 14px;
  font-weight: bold;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn:disabled {
  background-color: #cbd5e0 !important;
  color: #718096 !important;
  cursor: not-allowed;
}

.btn-next {
  background-color: #41b883;
  color: white;
  margin-left: auto;
}

.btn-back {
  background-color: #718096;
  color: white;
}

.btn-submit {
  background-color: #41b883;
  color: white;
}

/* Styling Halaman Review */
.review-section h4 {
  margin: 15px 0 8px 0;
  color: #026448;
  border-left: 3px solid #3182ce;
  padding-left: 6px;
}

.review-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 15px;
}

.review-table th, .review-table td {
  border: 1px solid #e2e8f0;
  padding: 10px;
  font-size: 14px;
  text-align: left;
}

.review-table th {
  background-color: #f7fafc;
  width: 35%;
  color: #4a5568;
}

.salary-text {
  font-weight: bold;
  color: #2f855a;
}
</style>