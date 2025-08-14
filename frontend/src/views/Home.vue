<template>
  <div class="home">
    <section class="hero">
      <div class="hero-content">
        <h1 class="hero-title">Create & Share Polls Instantly</h1>
              <div class="hero-visual">
        <div class="poll-preview">
          <div class="poll-card-demo">
            <h3>What's your favorite fruit? 🍓</h3>
            <div class="option-demo">Cherry <div class="progress-bar"><div class="progress" style="width: 65%"></div></div> 65%</div>
            <div class="option-demo">Apple <div class="progress-bar"><div class="progress" style="width: 25%"></div></div> 25%</div>
            <div class="option-demo">Banana <div class="progress-bar"><div class="progress" style="width: 10%"></div></div> 10%</div>
          </div>
        </div>
      </div>
        <p class="hero-subtitle">
          Build beautiful polls, collect votes, and see real-time results with our modern voting platform.
        </p>
        <div class="hero-actions">
          <router-link to="/create-poll" class="btn-primary" v-if="authStore.isLoggedIn">
            Create Your First Poll
          </router-link>
          <router-link to="/register" class="btn-primary" v-else>
            Get Started Free
          </router-link>
          <button @click="scrollToFeatures" class="btn-secondary">
            Learn More
          </button>
        </div>
      </div>

    </section>

    <section class="features" ref="featuresRef">
      <div class="container">
        <h2 class="section-title">Why Choose Our Platform?</h2>
        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon">⚡</div>
            <h3>Instant Results</h3>
            <p>See real-time voting results with beautiful charts and animations.</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🔒</div>
            <h3>Secure Voting</h3>
            <p>One vote per person with IP and user authentication protection.</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">📱</div>
            <h3>Mobile Friendly</h3>
            <p>Works perfectly on all devices - desktop, tablet, and mobile.</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">⏰</div>
            <h3>Time Limits</h3>
            <p>Set expiration dates for your polls to control voting periods.</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">📊</div>
            <h3>Beautiful Charts</h3>
            <p>Visualize results with stunning charts and export data easily.</p>
          </div>
          <div class="feature-card">
            <div class="feature-icon">🚀</div>
            <h3>Easy Sharing</h3>
            <p>Share polls with direct links, no account required for voters.</p>
          </div>
        </div>
      </div>
    </section>

    <section class="cta">
      <div class="container">
        <h2>Ready to Create Your First Poll?</h2>
        <p>Join thousands of users creating engaging polls every day.</p>
        <router-link to="/register" class="btn-primary btn-large" v-if="!authStore.isLoggedIn">
          Start Polling Now
        </router-link>
        <router-link to="/create-poll" class="btn-primary btn-large" v-else>
          Create New Poll
        </router-link>
      </div>
    </section>
    <footer class="footer-small">
    <p>&copy; 2025 VotingApp by <a href="github.com/tikature">tikature</a></p>
  </footer>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'

export default {
  name: "FooterSmall",
  methods: {
    scrollTop() {
      window.scrollTo({ top: 0, behavior: 'smooth' });
    },
  },
  name: 'Home',
  setup() {
    const authStore = useAuthStore()
    const featuresRef = ref(null)
    
    const scrollToFeatures = () => {
      featuresRef.value?.scrollIntoView({ behavior: 'smooth' })
    }
    
    return {
      authStore,
      featuresRef,
      scrollToFeatures
    }
  }
}
</script>

<style scoped>
.home {
  overflow-x: hidden;
}

.hero {
  min-height: 90vh;
  display: flex;
  align-items: center;
  padding: 40px 20px;
  background: linear-gradient(135deg, var(--cherry-bg) 0%, var(--white) 50%, var(--cherry-pink) 100%);
}

.hero-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 800;
  color: var(--text-dark);
  line-height: 1.1;
  margin-bottom: 20px;
  background: linear-gradient(135deg, var(--cherry-red), var(--cherry-red-light));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.hero-subtitle {
  font-size: 1.2rem;
  color: var(--text-light);
  line-height: 1.6;
  margin-bottom: 40px;
}

.hero-actions {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.btn-secondary {
  background: transparent;
  border: 2px solid var(--cherry-red);
  color: var(--cherry-red);
  padding: 12px 24px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-block;
}

.btn-secondary:hover {
  background-color: var(--cherry-red);
  color: white;
  transform: translateY(-2px);
}

.poll-preview {
  display: flex;
  justify-content: center;
  align-items: center;
}

.poll-card-demo {
  background: var(--white);
  border-radius: 16px;
  padding: 30px;
  box-shadow: 0 10px 40px rgba(220, 20, 60, 0.2);
  border: 2px solid var(--cherry-pink);
  animation: float 6s ease-in-out infinite;
}

.poll-card-demo h3 {
  color: var(--text-dark);
  margin-bottom: 20px;
  font-size: 1.3rem;
}

.option-demo {
  display: flex;
  align-items: center;
  gap: 15px;
  margin: 15px 0;
  font-weight: 500;
  color: var(--text-dark);
}

.progress-bar {
  flex: 1;
  height: 8px;
  background: var(--gray-200);
  border-radius: 4px;
  overflow: hidden;
}

.progress {
  height: 100%;
  background: linear-gradient(90deg, var(--cherry-red), var(--cherry-red-light));
  border-radius: 4px;
  transition: width 2s ease;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.features {
  padding: 100px 20px;
  background: var(--white);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

.section-title {
  text-align: center;
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--text-dark);
  margin-bottom: 60px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 40px;
}

.feature-card {
  text-align: center;
  padding: 40px 20px;
  border-radius: 16px;
  background: var(--cherry-bg);
  border: 2px solid var(--cherry-pink);
  transition: all 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 20px 40px rgba(220, 20, 60, 0.15);
}

.feature-icon {
  font-size: 3rem;
  margin-bottom: 20px;
}

.feature-card h3 {
  font-size: 1.3rem;
  color: var(--text-dark);
  margin-bottom: 15px;
  font-weight: 600;
}

.feature-card p {
  color: var(--text-light);
  line-height: 1.6;
}

.cta {
  padding: 80px 20px;
  background: linear-gradient(135deg, var(--cherry-red), var(--cherry-red-dark));
  text-align: center;
  color: white;
}

.cta h2 {
  font-size: 2.5rem;
  margin-bottom: 20px;
}

.cta p {
  font-size: 1.2rem;
  margin-bottom: 40px;
  opacity: 0.9;
}

.btn-large {
  padding: 16px 40px;
  font-size: 1.1rem;
  background: var(--white);
  color: var(--cherry-red);
}

.btn-large:hover {
  background: var(--gray-100);
  color: var(--cherry-red-dark);
}

@media (max-width: 768px) {
  .hero-content {
    grid-template-columns: 1fr;
    text-align: center;
    gap: 40px;
  }
  
  .hero-title {
    font-size: 2.5rem;
  }
  
  .hero-actions {
    justify-content: center;
  }
  
  .section-title {
    font-size: 2rem;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
  }
}

.footer-small {
  background-color: #ea5c79ff; /* Red Cherry */
  color: #fff;
  text-align: center;
  padding: 10px 0;
  font-family: sans-serif;
  position: relative;
}

.footer-btn {
  background-color: #fff;
  color: #FF3B30;
  border: none;
  padding: 6px 12px;
  border-radius: 5px;
  cursor: pointer;
  font-weight: bold;
  margin-bottom: 5px;
  transition: background-color 0.3s, color 0.3s;
}

.footer-btn:hover {
  background-color: #ffe5e0;
  color: #FF3B30;
}
</style>