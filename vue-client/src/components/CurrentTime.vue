<template>
    <div>
        <div v-if="loading">Loading...</div>

        <div v-else class="currentTime">
            <div v-if="error === ''">
                <p>Current Time: {{ data.ct }}</p>
                <p>Language: {{ data.language }}</p>
                <p>Endpoint Version: {{ data.version }}</p>
            </div>
            <div v-else>
                {{ error }}
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import { ref, onMounted } from 'vue';
    import axios from 'axios';

    interface CurrentTimeData {
        language: string,
        ct: string,
        version: string
    }

    const loading = ref(true);
    const error = ref("");
    const data = ref<CurrentTimeData>({ language: '', ct: '', version: '' });

    async function loadData() {
        try {
            const response = await axios.get<CurrentTimeData>('http://localhost:3000/currentTime');
            data.value = response.data;
            loading.value = false;
        } catch (err) {
            loading.value = false;
            error.value = "We ran into an error. Please try request again later";
            console.error(err);
        }
    }

    onMounted(async () => {
        await loadData();
    });
</script>

<style scoped>
/* Add your styles here */
</style>
