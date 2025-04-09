<template>
  <div v-if="chartData" style="height: 150px;">
  <Bar :chart-data="chartData" :chart-options="chartOptions" />
</div>


  <div v-else>
    Memuat data grafik...
  </div>
</template>

<script>
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
} from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
  name: 'GrafikStatistik',
  components: { Bar },
  props: {
    dataStatistik: {
      type: Array,
      required: true,
      default: () => []
    }
  },
  computed: {
    chartData() {
      if (!this.dataStatistik.length || this.dataStatistik.length === 0) return null;
      return {
        labels: this.dataStatistik.map(item => `${item.jurusan} ${item.kelas}`),
        datasets: [
          {
            label: 'Hadir',
            backgroundColor: '#28a745',
            data: this.dataStatistik.map(item => item.hadir)
          },
          {
            label: 'Sakit',
            backgroundColor: '#ffc107',
            data: this.dataStatistik.map(item => item.sakit)
          },
          {
            label: 'Izin',
            backgroundColor: '#17a2b8',
            data: this.dataStatistik.map(item => item.izin)
          },
          {
            label: 'Alfa',
            backgroundColor: '#dc3545',
            data: this.dataStatistik.map(item => item.alfa)
          }
        ]
      }
    },
    chartOptions() {
  return {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'top',
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
}

  }
}
</script>
