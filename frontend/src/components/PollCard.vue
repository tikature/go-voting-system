<template>
  <div class="poll-card">
    <div class="poll-card-header">
      <h3 class="poll-title">{{ poll.title }}</h3>
      <div class="poll-meta">
        <span class="poll-date">{{ formatDate(poll.created_at) }}</span>
        <span class="poll-status" :class="{ expired: isExpired, active: !isExpired }">
          {{ isExpired ? 'Expired' : 'Active' }}
        </span>
      </div>
    </div>
    
    <div class="poll-description" v-if="poll.description">
      {{ poll.description }}
    </div>
    
    <div class="poll-stats">
      <div class="stat">
        <span class="stat-number">{{ poll.options?.length || 0 }}</span>
        <span class="stat-label">Options</span>
      </div>
      <div class="stat">
        <span class="stat-number">{{ poll.total_votes || 0 }}</span>
        <span class="stat-label">Votes</span>
      </div>
      <div v-if="poll.expires_at" class="stat">
        <span class="stat-number">{{ timeRemaining }}</span>
        <span class="stat-label">{{ isExpired ? 'Expired' : 'Remaining' }}</span>
      </div>
    </div>
    
    <div class="poll-actions">
      <router-link :to="`/poll/${poll.id}`" class="btn-primary btn-small">
        {{ showActions ? 'View Details' : 'Vote Now' }}
      </router-link>
      
      <div v-if="showActions" class="poll-admin-actions">
        <button @click="copyPollLink" class="btn-icon" title="Copy Link">
          📋
        </button>
        <button @click="$emit('delete', poll.id)" class="btn-icon btn-danger" title="Delete">
          🗑️
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'

export default {
  name: 'PollCard',
  props: {
    poll: {
      type: Object,
      required: true
    },
    showActions: {
      type: Boolean,
      default: false
    }
  },
  emits: ['delete'],
  setup(props) {
    const isExpired = computed(() => {
      if (!props.poll.expires_at) return false
      return new Date(props.poll.expires_at) < new Date()
    })
    
    const timeRemaining = computed(() => {
      if (!props.poll.expires_at) return 'No limit'
      
      const now = new Date()
      const expiry = new Date(props.poll.expires_at)
      const diff = expiry - now
      
      if (diff <= 0) return 'Expired'
      
      const days = Math.floor(diff / (1000 * 60 * 60 * 24))
      const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
      
      if (days > 0) return `${days}d ${hours}h`
      if (hours > 0) return `${hours}h`
      return '<1h'
    })
    
    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleDateString()
    }
    
    const copyPollLink = () => {
      const url = `${window.location.origin}/poll/${props.poll.id}`
      navigator.clipboard.writeText(url).then(() => {
        alert('Poll link copied to clipboard!')
      }).catch(() => {
        alert('Failed to copy link')
      })
    }
    
    return {
      isExpired,
      timeRemaining,
      formatDate,
      copyPollLink
    }
  }
}
</script>

<style scoped>
.poll-card {
  background: var(--white);
  border-radius: 16px;
  padding: 25px;
  box-shadow: 0 4px 20px rgba(220, 20, 60, 0.1);
  border: 2px solid var(--cherry-pink);
  transition: all 0.3s ease;
  height: fit-content;
}

.poll-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 30px rgba(220, 20, 60, 0.15);
}

.poll-card-header {
  margin-bottom: 15px;
}

.poll-title {
  color: var(--text-dark);
  font-size: 1.3rem;
  font-weight: 600;
  margin-bottom: 10px;
  line-height: 1.3;
}

.poll-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 15px;
}

.poll-date {
  color: var(--text-light);
  font-size: 0.85rem;
}

.poll-status {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 0.8rem;
  font-weight: 600;
}

.poll-status.active {
  background-color: #D1FAE5;
  color: #065F46;
}

.poll-status.expired {
  background-color: #FEE2E2;
  color: #991B1B;
}

.poll-description {
  color: var(--text-light);
  font-size: 0.95rem;
  line-height: 1.5;
  margin-bottom: 20px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.poll-stats {
  display: flex;
  justify-content: space-around;
  margin: 20px 0;
  padding: 15px 0;
  border-top: 1px solid var(--cherry-pink);
  border-bottom: 1px solid var(--cherry-pink);
}

.stat {
  text-align: center;
}

.stat-number {
  display: block;
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--cherry-red);
}

.stat-label {
  display: block;
  font-size: 0.8rem;
  color: var(--text-light);
  font-weight: 500;
}

.poll-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
}

.btn-small {
  padding: 10px 20px;
  font-size: 0.9rem;
}

.poll-admin-actions {
  display: flex;
  gap: 10px;
}

.btn-icon {
  width: 36px;
  height: 36px;
  border: none;
  border-radius: 8px;
  background: var(--gray-100);
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
}

.btn-icon:hover {
  background: var(--cherry-pink);
}

.btn-danger:hover {
  background: #FEE2E2;
}

@media (max-width: 768px) {
  .poll-card {
    padding: 20px;
  }
  
  .poll-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  
  .poll-actions {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .poll-admin-actions {
    justify-content: center;
  }
}
</style>