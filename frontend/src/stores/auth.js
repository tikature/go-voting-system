import { defineStore } from 'pinia'
import api from '@/utils/api'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
        token: localStorage.getItem('token'),
        isAuthenticated: false
    }),

    getters: {
        isLoggedIn: (state) => !!state.token && !!state.user
    },

    actions: {
        async login(credentials) {
            try {
                const response = await api.post('/login', credentials)
                const { token, user } = response.data

                this.token = token
                this.user = user
                this.isAuthenticated = true

                localStorage.setItem('token', token)
                localStorage.setItem('user', JSON.stringify(user))

                return response.data
            } catch (error) {
                throw error.response?.data?.error || 'Login failed'
            }
        },

        async register(userData) {
            try {
                const response = await api.post('/register', userData)
                const { token, user } = response.data

                this.token = token
                this.user = user
                this.isAuthenticated = true

                localStorage.setItem('token', token)
                localStorage.setItem('user', JSON.stringify(user))

                return response.data
            } catch (error) {
                throw error.response?.data?.error || 'Registration failed'
            }
        },

        logout() {
            this.token = null
            this.user = null
            this.isAuthenticated = false

            localStorage.removeItem('token')
            localStorage.removeItem('user')
        },

        initAuth() {
            const token = localStorage.getItem('token')
            const user = localStorage.getItem('user')

            if (token && user) {
                this.token = token
                this.user = JSON.parse(user)
                this.isAuthenticated = true
            }
        }
    }
})