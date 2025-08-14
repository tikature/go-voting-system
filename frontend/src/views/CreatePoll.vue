<template>
  <div class="create-poll">
    <div class="container">
      <div class="page-header">
        <h1>Create New Poll</h1>
        <p>Ask questions and get instant feedback from your audience</p>
      </div>
      
      <div class="create-poll-content">
        <PollForm @submit="handleCreatePoll" />
      </div>
    </div>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'
import { usePollsStore } from '@/stores/polls'
import PollForm from '@/components/PollForm.vue'

export default {
  name: 'CreatePoll',
  components: {
    PollForm
  },
  setup() {
    const router = useRouter()
    const pollsStore = usePollsStore()
    
    const handleCreatePoll = async (pollData) => {
      try {
        const newPoll = await pollsStore.createPoll(pollData)
        router.push(`/poll/${newPoll.id}`)
      } catch (error) {
        alert('Failed to create poll: ' + error)
      }
    }
    
    return {
      handleCreatePoll
    }
  }
}
</script>

<style scoped>
.create-poll {
  padding: 40px 20px;
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  text-align: center;
  margin-bottom: 50px;
}

.page-header h1 {
  font-size: 2.5rem;
  color: var(--text-dark);
  font-weight: 700;
  margin-bottom: 15px;
}

.page-header p {
  color: var(--text-light);
  font-size: 1.1rem;
}

@media (max-width: 768px) {
  .create-poll {
    padding: 20px 15px;
  }
  
  .page-header h1 {
    font-size: 2rem;
  }
}
</style>