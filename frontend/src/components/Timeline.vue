<template>
    <div class="timeline">
    {{ new Date(this.min*1000).toString() }} - {{ new Date(this.max*1000).toString() }}
    </div>
</template>
<script>
export default {
    data() {
        return {
            constraints: {},
            max: 0,
            min: 0,
            current: 0,
        }
    },
    mounted() {
        this.fetchData()
        this.interval = window.setInterval(this.tick, 300);
    },
    dismounted() {
        clearInterval(this.interval)
    },
    methods: {
        tick() {
            this.current += 300
            this.$emit('timeChange', this.current)
        },
        fetchData() {
            var API_PATH = 'https://twotired.math.party/timestamps';

            var request = new XMLHttpRequest();
            request.open("GET", API_PATH, false);
            request.send(null)

            var constraints = JSON.parse(request.responseText);
            this.max = constraints.max;
            this.min = this.max-24*60*60;
            this.current = this.min*1000;
        }
    }
}

</script>