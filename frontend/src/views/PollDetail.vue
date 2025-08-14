<template>
  <div class="poll-detail">
    <div class="container">
      <div v-if="loading" class="loading">
        Loading poll...
      </div>
      
      <div v-else-if="poll" class="poll-content">
        <div class="poll-header">
          <h1>{{ poll.title }}</h1>
          <p v-if="poll.description">{{ poll.description }}</p>
          <div class="poll-meta">
            <span class="poll-status" :class="{ expired: isExpired }">
              {{ isExpired ? 'Expired' : 'Active' }}
            </span>
            <span v-if="poll.expires_at" class="poll-expires">
              {{ isExpired ? 'Expired' : 'Expires' }}: {{ formatDate(poll.expires_at) }}
            </span>
          </div>
        </div>
        
        <div class="poll-voting">
          <div v-if="!hasVoted && !isExpired" class="voting-section">
            <h3>Cast Your Vote</h3>
            <div class="voting-options">
              <div 
                v-for="option in poll.options" 
                :key="option.id"
                class="vote-option"
                :class="{ selected: selectedOption === option.id }"
                @click="selectedOption = option.id"
              >
                <input 
                  type="radio" 
                  :id="`option-${option.id}`"
                  :value="option.id"
                  v-model="selectedOption"
                />
                <label :for="`option-${option.id}`">{{ option.option_text }}</label>
              </div>
            </div>
            <button 
              @click="submitVote" 
              class="btn-primary"
              :disabled="!selectedOption || submitting"
            >
              {{ submitting ? 'Voting...' : 'Submit Vote' }}
            </button>
          </div>
          
          <div v-else class="results-section">
            <h3>Poll Results</h3>
            <VoteChart :poll="poll" />
            <div class="results-summary">
              <div class="total-votes">
                Total Votes: {{ poll.total_votes || 0 }}
              </div>
              <div v-if="hasVoted" class="voted-message">
                ✅ You have already voted in this poll
              </div>
              <div v-if="isExpired" class="expired-message">
                ⏰ This poll has expired
              </div>
            </div>
          </div>
        </div>
        
        <div class="poll-actions">
          <button @click="copyPollLink" class="poll-btn cherry">
            📋 Copy Poll Link
          </button>
          <button @click="refreshResults" class="poll-btn cherry">
            🔄 Refresh Results
          </button>
        </div>
      </div>
      
      <div v-else class="error-state">
        <h2>Poll Not Found</h2>
        <p>The poll you're looking for doesn't exist or has been removed.</p>
        <router-link to="/" class="btn-primary">Go Home</router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePollsStore } from '@/stores/polls'
import VoteChart from '@/components/VoteChart.vue'

export default {
  name: 'PollDetail',
  components: {
    VoteChart
  },
  setup() {
    const route = useRoute()
    const authStore = useAuthStore()
    const pollsStore = usePollsStore()
    
    const poll = ref(null)
    const loading = ref(true)
    const selectedOption = ref(null)
    const submitting = ref(false)
    const hasVoted = ref(false)
    
    const isExpired = computed(() => {
      if (!poll.value?.expires_at) return false
      return new Date(poll.value.expires_at) < new Date()
    })
    
    const fetchPoll = async () => {
      try {
        const isPublic = !authStore.isLoggedIn
        const pollData = await pollsStore.fetchPollById(route.params.id, isPublic)
        poll.value = pollData
        
        // Check if user has already voted (you might want to add this to the API response)
        if (pollData.user_has_voted) {
          hasVoted.value = true
        }
      } catch (error) {
        console.error('Failed to fetch poll:', error)
      } finally {
        loading.value = false
      }
    }
    
    const submitVote = async () => {
      if (!selectedOption.value) return
      
      submitting.value = true
      try {
        const isPublic = !authStore.isLoggedIn
        await pollsStore.votePoll({
          poll_id: parseInt(route.params.id),
          option_id: selectedOption.value
        }, isPublic)
        
        hasVoted.value = true
        await refreshResults()
      } catch (error) {
        alert('Failed to submit vote: ' + error)
      } finally {
        submitting.value = false
      }
    }
    
    const refreshResults = async () => {
      try {
        const isPublic = !authStore.isLoggedIn
        const results = await pollsStore.getPollResults(route.params.id, isPublic)
        poll.value = { ...poll.value, ...results }
      } catch (error) {
        console.error('Failed to refresh results:', error)
      }
    }
    
    const copyPollLink = () => {
      const url = window.location.href
      navigator.clipboard.writeText(url).then(() => {
        alert('Poll link copied to clipboard!')
      }).catch(() => {
        alert('Failed to copy link')
      })
    }
    
    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleString()
    }
    
    onMounted(fetchPoll)
    
    return {
      poll,
      loading,
      selectedOption,
      submitting,
      hasVoted,
      isExpired,
      submitVote,
      refreshResults,
      copyPollLink,
      formatDate
    }
  }
}
</script>

<style scoped>
.poll-detail {
  padding: 40px 20px;
  max-width: 1000px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  padding: 60px;
  font-size: 1.2rem;
  color: var(--text-light);
}

.poll-header {
  text-align: center;
  margin-bottom: 40px;
  padding: 40px;
  background: var(--white);
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(220, 20, 60, 0.1);
  border: 2px solid var(--cherry-pink);
}

.poll-header h1 {
  font-size: 2.5rem;
  color: var(--text-dark);
  font-weight: 700;
  margin-bottom: 15px;
}

.poll-header p {
  color: var(--text-light);
  font-size: 1.1rem;
  margin-bottom: 20px;
}

.poll-meta {
  display: flex;
  justify-content: center;
  gap: 20px;
  flex-wrap: wrap;
}

.poll-status {
  padding: 6px 16px;
  border-radius: 20px;
  font-weight: 600;
  font-size: 0.9rem;
  background-color: #10B981;
  color: white;
}

.poll-status.expired {
  background-color: #EF4444;
}

.poll-expires {
  color: var(--text-light);
  font-size: 0.9rem;
}

.poll-voting {
  margin-bottom: 40px;
}

.poll-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.poll-btn.blue {
  background-color: #007bff;
}

.poll-btn.cherry {
  background-color: #c23854ff;
}
.poll-btn.cherry:hover {
  background-color: #c23854ff;
}

.poll-btn.blue:hover {
  background-color: #0069d9;
}

.poll-btn.green {
  background-color: #28a745;
}

.poll-btn.green:hover {
  background-color: #218838;
}


.poll-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 10px 16px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  cursor: pointer;
  font-weight: 500;
  color: #fff;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: all 0.25s ease;
}


.voting-section, .results-section {
  background: var(--white);
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 4px 20px rgba(220, 20, 60, 0.1);
  border: 2px solid var(--cherry-pink);
}

.voting-section h3, .results-section h3 {
  color: var(--text-dark);
  font-size: 1.5rem;
  margin-bottom: 25px;
  text-align: center;
}

.voting-options {
  margin-bottom: 30px;
}

.vote-option {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 20px;
  border: 2px solid var(--gray-200);
  border-radius: 12px;
  margin-bottom: 15px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.vote-option:hover {
  border-color: var(--cherry-red-light);
  background-color: var(--cherry-bg);
}

.vote-option.selected {
  border-color: var(--cherry-red);
  background-color: var(--cherry-pink);
}

.vote-option input[type="radio"] {
  width: 20px;
  height: 20px;
  accent-color: var(--cherry-red);
}

.vote-option label {
  font-size: 1.1rem;
  font-weight: 500;
  color: var(--text-dark);
  cursor: pointer;
  flex: 1;
}

.results-summary {
  margin-top: 30px;
  text-align: center;
}

.total-votes {
  font-size: 1.2rem;
  font-weight: 600;
  color: var(--text-dark);
  margin-bottom: 15px;
}

.voted-message {
  color: #10B981;
  font-weight: 500;
  margin-bottom: 10px;
}

.expired-message {
  color: #EF4444;
  font-weight: 500;
}

.poll-actions {
  display: flex;
  justify-content: center;
  gap: 20px;
  flex-wrap: wrap;
}

.error-state {
  text-align: center;
  padding: 80px 20px;
  background: var(--white);
  border-radius: 16px;
  border: 2px solid var(--cherry-pink);
}

.error-state h2 {
  color: var(--text-dark);
  font-size: 2rem;
  margin-bottom: 15px;
}

.error-state p {
  color: var(--text-light);
  font-size: 1.1rem;
  margin-bottom: 30px;
}

@media (max-width: 768px) {
  .poll-detail {
    padding: 20px 15px;
  }
  
  .poll-header h1 {
    font-size: 2rem;
  }
  
  .poll-header, .voting-section, .results-section {
    padding: 30px 20px;
  }
  
  .poll-actions {
    flex-direction: column;
    align-items: center;
  }
}
</style>