<template>
  <div class="poll-form-container">
    <form @submit.prevent="handleSubmit" class="poll-form">
      <div class="form-group">
        <label for="title">Poll Title *</label>
        <input
          v-model="form.title"
          type="text"
          id="title"
          class="input-field"
          placeholder="What would you like to ask?"
          required
        />
      </div>
      
      <div class="form-group">
        <label for="description">Description (Optional)</label>
        <textarea
          v-model="form.description"
          id="description"
          class="input-field textarea"
          placeholder="Add more context to your poll..."
          rows="3"
        ></textarea>
      </div>
      
      <div class="form-group">
        <label>Poll Options *</label>
        <div class="options-list">
          <div 
            v-for="(option, index) in form.options" 
            :key="index"
            class="option-input"
          >
            <input
              v-model="form.options[index]"
              type="text"
              class="input-field"
              :placeholder="`Option ${index + 1}`"
              required
            />
            <button 
              v-if="form.options.length > 2"
              type="button"
              @click="removeOption(index)"
              class="btn-remove"
            >
              ✕
            </button>
          </div>
        </div>
        <button 
          type="button"
          @click="addOption"
          class="btn-add-option"
          :disabled="form.options.length >= 10"
        >
          + Add Option
        </button>
      </div>
      
      <div class="form-group">
        <label for="expiresAt">Expiration Date (Optional)</label>
        <input
          v-model="form.expiresAt"
          type="datetime-local"
          id="expiresAt"
          class="input-field"
          :min="minDateTime"
        />
        <small class="form-hint">Leave empty for no expiration</small>
      </div>
      
      <div class="form-actions">
        <router-link to="/dashboard" class="btn-secondary">
          Cancel
        </router-link>
        <button type="submit" class="btn-primary" :disabled="loading">
          {{ loading ? 'Creating...' : 'Create Poll' }}
        </button>
      </div>
      
      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </form>
  </div>
</template>

<script>
import { ref, computed } from 'vue'

export default {
  name: 'PollForm',
  emits: ['submit'],
  setup(_, { emit }) {
    const form = ref({
      title: '',
      description: '',
      options: ['', ''],
      expiresAt: ''
    })
    
    const loading = ref(false)
    const error = ref('')
    
    const minDateTime = computed(() => {
      const now = new Date()
      now.setMinutes(now.getMinutes() + 30) // Minimum 30 minutes from now
      return now.toISOString().slice(0, 16)
    })
    
    const addOption = () => {
      if (form.value.options.length < 10) {
        form.value.options.push('')
      }
    }
    
    const removeOption = (index) => {
      if (form.value.options.length > 2) {
        form.value.options.splice(index, 1)
      }
    }
    
    const handleSubmit = async () => {
      error.value = ''
      
      // Validation
      if (!form.value.title.trim()) {
        error.value = 'Poll title is required'
        return
      }
      
      const validOptions = form.value.options.filter(opt => opt.trim())
      if (validOptions.length < 2) {
        error.value = 'At least 2 options are required'
        return
      }
      
      if (form.value.expiresAt) {
        const expiryDate = new Date(form.value.expiresAt)
        if (expiryDate <= new Date()) {
          error.value = 'Expiration date must be in the future'
          return
        }
      }
      
      loading.value = true
      
      try {
        const pollData = {
          title: form.value.title.trim(),
          description: form.value.description.trim(),
          options: validOptions,
          expires_at: form.value.expiresAt ? new Date(form.value.expiresAt).toISOString() : null
        }
        
        emit('submit', pollData)
      } catch (err) {
        error.value = err.message || 'Failed to create poll'
      } finally {
        loading.value = false
      }
    }
    
    return {
      form,
      loading,
      error,
      minDateTime,
      addOption,
      removeOption,
      handleSubmit
    }
  }
}
</script>

<style scoped>
.poll-form-container {
  max-width: 600px;
  margin: 0 auto;
}

.poll-form {
  background: var(--white);
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 4px 20px rgba(220, 20, 60, 0.1);
  border: 2px solid var(--cherry-pink);
}

.form-group {
  margin-bottom: 25px;
}

.form-group label {
  display: block;
  color: var(--text-dark);
  font-weight: 600;
  margin-bottom: 8px;
}

.textarea {
  resize: vertical;
  min-height: 80px;
}

.options-list {
  margin-bottom: 15px;
}

.option-input {
  display: flex;
  gap: 10px;
  margin-bottom: 15px;
  align-items: center;
}

.option-input .input-field {
  flex: 1;
}

.btn-remove {
  width: 36px;
  height: 36px;
  border: none;
  background: #FEE2E2;
  color: #B91C1C;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 600;
}

.btn-remove:hover {
  background: #FECACA;
}

.btn-add-option {
  background: var(--cherry-pink);
  color: var(--cherry-red);
  border: 2px dashed var(--cherry-red);
  padding: 12px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
  width: 100%;
}

.btn-add-option:hover:not(:disabled) {
  background: var(--cherry-red);
  color: white;
  border-style: solid;
}

.btn-add-option:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.form-hint {
  color: var(--text-light);
  font-size: 0.85rem;
  margin-top: 5px;
  display: block;
}

.form-actions {
  display: flex;
  gap: 20px;
  justify-content: flex-end;
  margin-top: 40px;
}

.error-message {
  background-color: #FEE2E2;
  color: #B91C1C;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #FCA5A5;
  text-align: center;
  font-size: 0.9rem;
  margin-top: 20px;
}

@media (max-width: 768px) {
  .poll-form {
    padding: 30px 20px;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .btn-secondary, .btn-primary {
    width: 100%;
  }
}
</style>