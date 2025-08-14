<template>
  <div class="register-page">
    <div class="register-container">
      <div class="register-card">
        <div class="card-header">
          <h1>Create Account</h1>
          <p>Join our community and start creating polls</p>
        </div>
        
        <form @submit.prevent="handleRegister" class="register-form">
          <div class="form-group">
            <label for="username">Username</label>
            <input
              v-model="form.username"
              type="text"
              id="username"
              class="input-field"
              placeholder="Choose a username"
              required
            />
          </div>
          
          <div class="form-group">
            <label for="email">Email</label>
            <input
              v-model="form.email"
              type="email"
              id="email"
              class="input-field"
              placeholder="Enter your email"
              required
            />
          </div>
          
          <div class="form-group">
            <label for="password">Password</label>
            <input
              v-model="form.password"
              type="password"
              id="password"
              class="input-field"
              placeholder="Create a password (min 6 characters)"
              required
              minlength="6"
            />
          </div>
          
          <div class="form-group">
            <label for="confirmPassword">Confirm Password</label>
            <input
              v-model="form.confirmPassword"
              type="password"
              id="confirmPassword"
              class="input-field"
              placeholder="Confirm your password"
              required
            />
          </div>
          
          <button type="submit" class="btn-primary btn-full" :disabled="loading">
            {{ loading ? 'Creating Account...' : 'Create Account' }}
          </button>
          
          <div class="form-footer">
            <p>Already have an account? 
              <router-link to="/login" class="link-primary">Sign in</router-link>
            </p>
          </div>
        </form>
        
        <div v-if="error" class="error-message">
          {{ error }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

export default {
  name: 'Register',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    
    const form = ref({
      username: '',
      email: '',
      password: '',
      confirmPassword: ''
    })
    
    const loading = ref(false)
    const error = ref('')
    
    const handleRegister = async () => {
      error.value = ''
      
      if (form.value.password !== form.value.confirmPassword) {
        error.value = 'Passwords do not match'
        return
      }
      
      loading.value = true
      
      try {
        await authStore.register({
          username: form.value.username,
          email: form.value.email,
          password: form.value.password
        })
        router.push('/dashboard')
      } catch (err) {
        error.value = err
      } finally {
        loading.value = false
      }
    }
    
    return {
      form,
      loading,
      error,
      handleRegister
    }
  }
}
</script>

<style scoped>
.register-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  background: linear-gradient(135deg, var(--cherry-bg) 0%, var(--white) 100%);
}

.register-container {
  width: 100%;
  max-width: 450px;
}

.register-card {
  background: var(--white);
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 10px 40px rgba(220, 20, 60, 0.15);
  border: 2px solid var(--cherry-pink);
}

.card-header {
  text-align: center;
  margin-bottom: 30px;
}

.card-header h1 {
  color: var(--text-dark);
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 8px;
}

.card-header p {
  color: var(--text-light);
  font-size: 0.95rem;
}

.register-form {
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  color: var(--text-dark);
  font-weight: 500;
  margin-bottom: 8px;
}

.btn-full {
  width: 100%;
  margin: 30px 0 20px 0;
}

.form-footer {
  text-align: center;
}

.form-footer p {
  color: var(--text-light);
  font-size: 0.9rem;
}

.link-primary {
  color: var(--cherry-red);
  text-decoration: none;
  font-weight: 600;
}

.link-primary:hover {
  color: var(--cherry-red-dark);
}

.error-message {
  background-color: #FEE2E2;
  color: #B91C1C;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #FCA5A5;
  text-align: center;
  font-size: 0.9rem;
}
</style>