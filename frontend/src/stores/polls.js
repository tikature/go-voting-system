import { defineStore } from 'pinia'
import api from '@/utils/api'

export const usePollsStore = defineStore('polls', {
    state: () => ({
        polls: [],
        currentPoll: null,
        loading: false,
        error: null
    }),

    actions: {
        async fetchPolls() {
            this.loading = true
            try {
                const response = await api.get('/polls')
                this.polls = response.data
                this.error = null
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to fetch polls'
            } finally {
                this.loading = false
            }
        },

        async fetchPollById(id, isPublic = false) {
            this.loading = true
            try {
                const endpoint = isPublic ? `/public/polls/${id}` : `/polls/${id}`
                const response = await api.get(endpoint)
                this.currentPoll = response.data
                this.error = null
                return response.data
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to fetch poll'
                throw error
            } finally {
                this.loading = false
            }
        },

        async createPoll(pollData) {
            this.loading = true
            try {
                const response = await api.post('/polls', pollData)
                this.polls.unshift(response.data)
                this.error = null
                return response.data
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to create poll'
                throw error
            } finally {
                this.loading = false
            }
        },

        async votePoll(voteData, isPublic = false) {
            try {
                const endpoint = isPublic ? '/public/vote' : '/vote'
                const response = await api.post(endpoint, voteData)
                this.error = null
                return response.data
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to vote'
                throw error
            }
        },

        async getPollResults(id, isPublic = false) {
            try {
                const endpoint = isPublic ? `/public/polls/${id}/results` : `/polls/${id}/results`
                const response = await api.get(endpoint)
                return response.data
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to get results'
                throw error
            }
        },

        async deletePoll(id) {
            try {
                await api.delete(`/polls/${id}`)
                this.polls = this.polls.filter(poll => poll.id !== id)
                this.error = null
            } catch (error) {
                this.error = error.response?.data?.error || 'Failed to delete poll'
                throw error
            }
        }
    }
})