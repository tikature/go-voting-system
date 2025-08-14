<template>
  <div class="dashboard">
    <div class="container">
      <div class="dashboard-header">
        <h1>My Dashboard</h1>
        <p>Manage your polls and view statistics</p>
      </div>
      
      <div class="dashboard-stats">
        <div class="stat-card">
          <div class="stat-icon">📊</div>
          <div class="stat-info">
            <h3>{{ polls.length }}</h3>
            <p>Total Polls</p>
          </div>
        </div>
        <div class="stat-card">
          <div class="stat-icon">🗳️</div>
          <div class="stat-info">
            <h3>{{ totalVotes }}</h3>
            <p>Total Votes</p>
          </div>
        </div>
        <div class="stat-card">
          <div class="stat-icon">👥</div>
          <div class="stat-info">
            <h3>{{ activePolls }}</h3>
            <p>Active Polls</p>
          </div>
        </div>
      </div>
      
      <div class="dashboard-actions">
        <router-link to="/create-poll" class="btn-primary">
          Create New Poll
        </router-link>
      </div>
      
      <div class="polls-section">
        <h2>Your Polls</h2>
        <div v-if="loading" class="loading">
          Loading polls...
        </div>
        <div v-else-if="polls.length === 0" class="empty-state">
          <div class="empty-icon">📝</div>
          <h3>No polls yet</h3>
          <p>Create your first poll to get started</p>
          <router-link to="/create-poll" class="btn-primary">
            Create Poll
          </router-link>
        </div>
        <div v-else class="polls-grid">
          <PollCard
            v-for="poll in polls"
            :key="poll.id"
            :poll="poll"
            :show-actions="true"
            @delete="handleDeletePoll"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { usePollsStore } from '@/stores/polls'
import PollCard from '@/components/PollCard.vue'

export default {
  name: 'Dashboard',
  components: {
    PollCard
  },
  setup() {
    const pollsStore = usePollsStore()
    const loading = ref(false)
    
    const polls = computed(() => pollsStore.polls)
    const totalVotes = computed(() => {
      return polls.value.reduce((total, poll) => total + (poll.total_votes || 0), 0)
    })
    const activePolls = computed(() => {
      return polls.value.filter(poll => poll.is_active).length
    })
    
    const handleDeletePoll = async (pollId) => {
      if (confirm('Are you sure you want to delete this poll?')) {
        try {
          await pollsStore.deletePoll(pollId)
        } catch (error) {
          alert('Failed to delete poll')
        }
      }
    }
    
    onMounted(async () => {
      loading.value = true
      await pollsStore.fetchPolls()
      loading.value = false
    })
    
    return {
      polls,
      totalVotes,
      activePolls,
      loading,
      handleDeletePoll
    }
  }
}
</script>

<style scoped>
.dashboard {
  padding: 40px 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.dashboard-header {
  text-align: center;
  margin-bottom: 40px;
}

.dashboard-header h1 {
  font-size: 2.5rem;
  color: var(--text-dark);
  font-weight: 700;
  margin-bottom: 10px;
}

.dashboard-header p {
  color: var(--text-light);
  font-size: 1.1rem;
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.stat-card {
  background: var(--white);
  border-radius: 12px;
  padding: 30px 20px;
  text-align: center;
  box-shadow: 0 4px 20px rgba(220, 20, 60, 0.1);
  border: 2px solid var(--cherry-pink);
  transition: transform 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-5px);
}

.stat-icon {
  font-size: 2.5rem;
  margin-bottom: 15px;
}

.stat-info h3 {
  font-size: 2rem;
  color: var(--cherry-red);
  font-weight: 700;
  margin-bottom: 5px;
}

.stat-info p {
  color: var(--text-light);
  font-weight: 500;
}

.dashboard-actions {
  text-align: center;
  margin-bottom: 50px;
}

.polls-section h2 {
  font-size: 1.8rem;
  color: var(--text-dark);
  font-weight: 600;
  margin-bottom: 30px;
}

.loading {
  text-align: center;
  padding: 40px;
  color: var(--text-light);
  font-size: 1.1rem;
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
  background: var(--white);
  border-radius: 16px;
  border: 2px solid var(--cherry-pink);
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 20px;
}

.empty-state h3 {
  color: var(--text-dark);
  font-size: 1.5rem;
  margin-bottom: 10px;
}

.empty-state p {
  color: var(--text-light);
  margin-bottom: 30px;
}

.polls-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 30px;
}

@media (max-width: 768px) {
  .dashboard {
    padding: 20px 15px;
  }
  
  .dashboard-header h1 {
    font-size: 2rem;
  }
  
  .polls-grid {
    grid-template-columns: 1fr;
  }
}
</style>