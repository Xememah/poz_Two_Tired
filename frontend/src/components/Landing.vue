<template>
  <mapView></mapView>
</template>
<script>

import MapView from '@/components/MapView'

export default {
  components: {
    mapView: MapView
  },
  data() {
    return {}
  },
  mounted: function () {
  }
}
</script>
<style>
#map {
  height: 700px;
  width: 100%;
}
</style>