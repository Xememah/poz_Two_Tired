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
  height: 650px;
  width: 100%;
}
</style>