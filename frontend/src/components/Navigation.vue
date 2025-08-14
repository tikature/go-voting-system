<template>
  <nav class="navbar">
    <div class="nav-container">
      <div class="nav-brand">
        <router-link to="/" class="brand-link">
          <span class="brand-icon"><img 
      src="https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExMGVhMW5ic2dxNno0NzM1NDJqcGZpMjFlemYzNDA1NXR2ajRqaXNxcyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/78eF7jSvqDxcZK19ps/giphy.gif" 
       alt="Fun GIF"
      width="50"    <!-- lebar dalam px -->
      height="20"   <!-- tinggi dalam px -->
    /></span>
          VotingApp
        </router-link>
      </div>
      
      <div class="nav-menu" :class="{ active: isMobileMenuOpen }">
        <router-link to="/" class="nav-link" @click="closeMobileMenu">Home</router-link>
        
        <template v-if="authStore.isLoggedIn">
          <router-link to="/dashboard" class="nav-link" @click="closeMobileMenu">Dashboard</router-link>
          <router-link to="/create-poll" class="nav-link" @click="closeMobileMenu">Create Poll</router-link>
          <div class="nav-user">
            <span class="user-name">{{ authStore.user?.username }}</span>
            <button @click="logout" class="btn-logout">Logout</button>
          </div>
        </template>
        
        <template v-else>
          <router-link to="/login" class="nav-link" @click="closeMobileMenu">Login</router-link>
          <router-link to="/register" class="nav-link btn-register" @click="closeMobileMenu">Register</router-link>
        </template>
      </div>
      
      <div class="mobile-menu-toggle" @click="toggleMobileMenu">
        <span></span>
        <span></span>
        <span></span>
      </div>
    </div>
  </nav>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

export default {
  name: 'Navigation',
  setup() {
    const router = useRouter()
    const authStore = useAuthStore()
    const isMobileMenuOpen = ref(false)
    
    const toggleMobileMenu = () => {
      isMobileMenuOpen.value = !isMobileMenuOpen.value
    }
    
    const closeMobileMenu = () => {
      isMobileMenuOpen.value = false
    }
    
    const logout = () => {
      authStore.logout()
      router.push('/')
      closeMobileMenu()
    }
    
    return {
      authStore,
      isMobileMenuOpen,
      toggleMobileMenu,
      closeMobileMenu,
      logout
    }
  }
}
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  width: 100%;
  background: var(--white);
  box-shadow: 0 2px 20px rgba(220, 20, 60, 0.1);
  z-index: 1000;
  border-bottom: 2px solid var(--cherry-pink);
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
}

.brand-link {
  display: flex;
  align-items: center;
  text-decoration: none;
  color: var(--cherry-red);
  font-size: 24px;
  font-weight: 700;
  gap: 8px;
}

.brand-icon {
  font-size: 28px;
}

.nav-menu {
  display: flex;
  align-items: center;
  gap: 30px;
}

.nav-link {
  text-decoration: none;
  color: var(--text-dark);
  font-weight: 500;
  padding: 8px 16px;
  border-radius: 6px;
  transition: all 0.3s ease;
}

.nav-link:hover {
  color: var(--cherry-red);
  background-color: var(--cherry-pink);
}

.nav-link.router-link-active {
  color: var(--cherry-red);
  background-color: var(--cherry-pink);
}

.btn-register {
  background-color: var(--cherry-red);
  color: white !important;
}

.btn-register:hover {
  background-color: var(--cherry-red-dark);
  color: white !important;
}

.nav-user {
  display: flex;
  align-items: center;
  gap: 15px;
}

.user-name {
  color: var(--text-dark);
  font-weight: 600;
}

.btn-logout {
  background: none;
  border: 2px solid var(--cherry-red);
  color: var(--cherry-red);
  padding: 6px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.3s ease;
}

.btn-logout:hover {
  background-color: var(--cherry-red);
  color: white;
}

.mobile-menu-toggle {
  display: none;
  flex-direction: column;
  cursor: pointer;
}

.mobile-menu-toggle span {
  width: 25px;
  height: 3px;
  background-color: var(--cherry-red);
  margin: 3px 0;
  transition: 0.3s;
}

@media (max-width: 768px) {
  .mobile-menu-toggle {
    display: flex;
  }
  
  .nav-menu {
    position: fixed;
    top: 70px;
    left: -100%;
    width: 100%;
    height: calc(100vh - 70px);
    background: var(--white);
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    padding-top: 50px;
    gap: 20px;
    transition: left 0.3s ease;
  }
  
  .nav-menu.active {
    left: 0;
  }
  
  .nav-user {
    flex-direction: column;
    gap: 15px;
  }
}
</style>