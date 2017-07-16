<template>
  <div>
    <mapView></mapView>
    <div id="timeView">Map timestamp: {{ timestamp }}</div>
    <div>{{ log }}</div>
  </div>
</template>
<script>

import MapView from '@/components/MapView'

export default {
  components: {
    mapView: MapView
  },
  data() {
    var timestamp = 1483228952
    return {
      timestamp: timestamp,
      log: "enter"
    }
  },
  mounted: function () {
    this.updateTimestamp()
    console.log(this)
  },
  methods: { 
    updateTimestamp: function() {
      setInterval(function() {
         this.timestamp += 1000;
         console.log("done")
         this.log += " done "
    }, 1000);
    }
  }
}
</script>
<style>
#map {
  height: 650px;
  width: 100%;
}
#timeView {

}
</style>