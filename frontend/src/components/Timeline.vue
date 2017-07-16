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
            inter: {},
        }
    },
    props: ['interval', 'addition'],
    mounted() {
        this.fetchData()
        this.inter = window.setInterval(this.tick, Number(this.interval));
    },
    dismounted() {
        clearInterval(this.inter)
    },
    methods: {
        tick() {
            this.current += Number(this.addition)
            //console.log(this.current)
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
            this.current = 1483225486;
        }
    }
}

</script>