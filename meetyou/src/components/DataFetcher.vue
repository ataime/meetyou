<template>
    <div>
        <h1>数据从 Go 应用获取：</h1>
        <p>{{ data }}</p>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            data: ''
        };
    },
    mounted() {
        this.fetchData();
    },
    methods: {
        fetchData() {
            axios.get('http://localhost:8081/api/user')
                .then(response => {
                    this.data = response.data;
                })
                .catch(error => {
                    console.error('请求出错:', error);
                });
        }
    }
};
</script>
