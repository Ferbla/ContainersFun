<template>
    <div class="currentTime">
        <p>Current Time: {{  data.ct  }}</p>
        <p>Language: {{ data.language }}</p>
        <p>Endpoint Version: {{ data.version }}</p>
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

    let data = ref({} as CurrentTimeData);
    async function loadData() {
        // TODO: Move this URL to a configuration
        const response = await axios.get<CurrentTimeData>('http://localhost:3000/currentTime');
        data.value = response.data;
    }

    onMounted(async() => {
        await loadData();
    })

</script>

<style scoped>



</style>