<template>
  <div class="vote-chart">
    <div class="chart-container">
      <canvas ref="chartCanvas"></canvas>
    </div>
    
    <div class="results-breakdown">
      <div 
        v-for="(option, index) in chartData" 
        :key="option.id"
        class="result-item"
      >
        <div class="result-header">
          <span class="result-option">{{ option.option_text }}</span>
          <span class="result-percentage">{{ option.percentage }}%</span>
        </div>
        <div class="progress-bar">
          <div 
            class="progress-fill" 
            :style="{ 
              width: option.percentage + '%',
              backgroundColor: colors[index % colors.length]
            }"
          ></div>
        </div>
        <div class="result-votes">{{ option.vote_count }} votes</div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, watch } from 'vue'
import { Chart, registerables } from 'chart.js'

Chart.register(...registerables)

export default {
  name: 'VoteChart',
  props: {
    poll: {
      type: Object,
      required: true
    }
  },
  setup(props) {
    const chartCanvas = ref(null)
    const chartInstance = ref(null)
    
    const colors = [
      '#DC143C', // Cherry Red
      '#FF6B6B', // Cherry Red Light
      '#B91C3C', // Cherry Red Dark
      '#FF8A80', // Light Red
      '#F48FB1', // Pink
      '#CE93D8', // Light Purple
      '#90CAF9', // Light Blue
      '#81C784'  // Light Green
    ]
    
    const chartData = computed(() => {
      if (!props.poll.options) return []
      
      const totalVotes = props.poll.options.reduce((sum, opt) => sum + opt.vote_count, 0)
      
      return props.poll.options.map(option => ({
        ...option,
        percentage: totalVotes > 0 ? Math.round((option.vote_count / totalVotes) * 100) : 0
      })).sort((a, b) => b.vote_count - a.vote_count)
    })
    
    const createChart = () => {
      if (!chartCanvas.value) return
      
      // Destroy existing chart
      if (chartInstance.value) {
        chartInstance.value.destroy()
      }
      
      const ctx = chartCanvas.value.getContext('2d')
      
      chartInstance.value = new Chart(ctx, {
        type: 'doughnut',
        data: {
          labels: chartData.value.map(item => item.option_text),
          datasets: [{
            data: chartData.value.map(item => item.vote_count),
            backgroundColor: colors,
            borderColor: '#ffffff',
            borderWidth: 3,
            hoverOffset: 10
          }]
        },
        options: {
          responsive: true,
          maintainAspectRatio: false,
          plugins: {
            legend: {
              position: 'bottom',
              labels: {
                padding: 20,
                usePointStyle: true,
                font: {
                  size: 14,
                  family: 'Inter'
                }
              }
            },
            tooltip: {
              callbacks: {
                label: function(context) {
                  const label = context.label || ''
                  const value = context.parsed
                  const total = context.dataset.data.reduce((a, b) => a + b, 0)
                  const percentage = total > 0 ? Math.round((value / total) * 100) : 0
                  return `${label}: ${value} votes (${percentage}%)`
                }
              }
            }
          },
          cutout: '60%',
          animation: {
            animateScale: true,
            animateRotate: true
          }
        }
      })
    }
    
    onMounted(() => {
      createChart()
    })
    
    watch(() => props.poll.options, () => {
      createChart()
    }, { deep: true })
    
    return {
      chartCanvas,
      chartData,
      colors
    }
  }
}
</script>

<style scoped>
.vote-chart {
  margin: 20px 0;
}

.chart-container {
  position: relative;
  height: 300px;
  margin-bottom: 30px;
}

.results-breakdown {
  space-y: 20px;
}

.result-item {
  margin-bottom: 20px;
}

.result-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.result-option {
  font-weight: 600;
  color: var(--text-dark);
  font-size: 1rem;
}

.result-percentage {
  font-weight: 700;
  color: var(--cherry-red);
  font-size: 1.1rem;
}

.progress-bar {
  height: 12px;
  background-color: var(--gray-200);
  border-radius: 6px;
  overflow: hidden;
  margin-bottom: 5px;
}

.progress-fill {
  height: 100%;
  border-radius: 6px;
  transition: width 1s ease-in-out;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.result-votes {
  font-size: 0.9rem;
  color: var(--text-light);
  text-align: right;
}

@media (max-width: 768px) {
  .chart-container {
    height: 250px;
  }
  
  .result-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
  
  .result-votes {
    text-align: left;
  }
}
</style>