<template>
    <div class="timeline">
        <div id="slider-connect"></div>
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
        this.interval = window.setInterval(this.tick, 300);

        var timeline = document.getElementById('slider-connect');
        
        console.log(this.min + ' ; ' + this.max)

        var base = this
        function toDate (v) {
            return base.formatDate(new Date(v*1000), true);
        }

        function toDatePips (v) {
            return base.formatDate(new Date(v*1000), false);
        }

        function reverseDate(str) {
            var date = str.slice(-33)
            date = date.replace(/<[^>]*>/g, '');
            var dateParts = date.match(/(\d+).(\d+).(\d+) (\d+):(\d+)/);
            console.log(dateParts)
            new Date()
            date = new Date(dateParts[3], parseInt(dateParts[2], 10) - 1, dateParts[1], dateParts[4], dateParts[5]);
            console.log(date.getTime());
            return date
        }

        var slider = noUiSlider.create(timeline, {
            start: this.min,
            connect: [true, false],
            tooltips: [true],
            format: { to: toDate, from: Number },
            range: {
                'min': this.min-5000,
                'max': this.max+5000
            },
            step: 2000,
            pips: {
                format: { to: toDatePips, from: Number },
                mode: 'values',
                values: [this.min, this.max]
            }
        });

        var conf = this
        slider.target.noUiSlider.on('change', function (values, handle) {
            if (reverseDate(values[handle]) < 1000*conf.min) {
                console.log("za mały")
                slider.target.noUiSlider.set(conf.min);
            } else if (reverseDate(values[handle]) > 1000*conf.max) {
                console.log("that's what she said")
                slider.target.noUiSlider.set(conf.max);
            }
        });
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
        formatDate(date, pip) {
            var
                weekdays = [
                    "Niedziela", "Poniedziałek", "Wtorek",
                    "Środa", "Czwartek", "Piątek",
                    "Sobota"
                ],
                months = [
                    "Styczeń", "Luty", "Marzec",
                    "Kwiecień", "Maj", "Czerwiec", "Lipiec",
                    "Sierpień", "Wrzesień", "Październik",
                    "Listopad", "Grudzień"
                ];

            function nth(d, e) {
                if(e) d+=1
                if(d<10) return '0' + d
                else return d
            }

            if(pip)
                return weekdays[date.getDay()] + ", " +
                    nth(date.getDate(), false) + "." +
                    nth(date.getMonth(), true) + "." +
                    date.getFullYear() + " <strong>" + nth(date.getHours(), false) +
                    ":" + nth(date.getMinutes(), false) + "</strong>";
            else
                return nth(date.getDate(), false) + "." +
                    nth(date.getMonth(), true) + "." +
                    date.getFullYear();
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
<style>
    #slider-connect {
        width: 80%;
        margin: 0 auto;
        margin-top: 3em;
        margin-bottom: 3em;
    }
</style>