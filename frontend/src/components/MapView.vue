<template>
  <div>
    <div id="map"></div>
    <timeline @timeChange="timeChange" :interval="interval" :addition="addition"></timeline>
  </div>
</template>

<script>
import Timeline from './Timeline.vue'
import HttpClient from '../httpclient.js'

export default {
  components: {
    'timeline': Timeline
  },
  data() {
    return {
      interval: 500,
      addition: 6,
      future: 60,
      markers:{},
      markersData: undefined,
      lastUpdate: 0,
    }
  },
  mounted: function () {
    this.initMap();
  },
  methods: {
    fetchData: function (time, howManySecondsToSeeInToFuture) {
      let apiPath = 'https://twotired.math.party/data?timestamp=' + time + '&duration=' + howManySecondsToSeeInToFuture;
      let client = new HttpClient();
      client.get(apiPath, (response) => {
        this.markersData = JSON.parse(response).slice(0, 80);
        this.updatePaths(time);
      });
    },
    timeChange: function (current) {
      if ((current-this.lastUpdate) > (this.future-this.addition)*(this.interval/1000)) {
        this.lastUpdate = current;
        this.fetchData(current, this.future);
      } else {
        console.log(new Date(current*1000).toString());
        this.updatePaths(current);
      }
    },
    updatePaths: function (time) {
      if (!this.markersData) {
        return;
      }
      for (let markerData of this.markersData) {
        if (!this.markers || !this.markers[markerData.hash]) {
          let markerPath = this.initMarker(markerData, time);
          this.markers[markerData.hash] = markerPath;
        } else {
          let line = this.markers[markerData.hash];
          let newp = line.getPath();
          outer:
          for(let i=0;i<markerData.steps.length;i++) {
            let step = markerData.steps[i];
            for(let s of line.getPath().getArray()) {
              if(step.lat == s.lat && step.lng == s.lng) {
                continue outer;
              }
            }
            if(step.ts<=time) {
              newp.push(new google.maps.LatLng(step.lat, step.lng))
              markerData.steps.splice(i,1);
              i--;
            }
          }
       }
      }
    },
    initMarker: function (marker, current) {
      var lineSymbol = {
        path: google.maps.SymbolPath.FORWARD_CLOSED_ARROW,
        scale: 3,
        strokeColor: '#cc0000'
      };

      var path = [];
      for (let step of marker.steps) {
        if (step.ts <= current) {
          path.push({ lat: step.lat, lng: step.lng });
        }
      }

      var markerPath = new google.maps.Polyline({
        path: path,
        icons: [{
          icon: lineSymbol,
          offset: '100%'
        }],
        geodesic: true,
        strokeColor: '#cc0000',
        strokeOpacity: 0.7,
        strokeWeight: 5
      });
      markerPath.setMap(this.map);
      return markerPath;
    },
    initMap: function () {
      var uluru = { lat: 52.385392, lng: 16.992963 };
      this.map = new google.maps.Map(document.getElementById('map'), {
        scrollwheel: false,
        zoom: 13,
        center: uluru
      })
      // this.changeTarget();
    },
    populateMarkers: function () {
      this.animateCircle(markerPath)
      /*google.maps.event.addListener(marker, 'click', function () {
        map.panTo(this.position);
        // if a tooltip is present, dispose it
        if (tooltip) {
          tooltip.close();
        }
        tooltip = new google.maps.InfoWindow({
          content: createTooltip(this.data, pict)
        });
        tooltip.open(map, this);
      })*/
      //}
    },
    animateCircle: function (line) {
      var count = 0;
      window.setInterval(function () {
        count = (count + 1) % 200;

        var icons = line.get('icons');
        icons[0].offset = (count / 2) + '%';
        line.set('icons', icons);
      }, 300);
    },
    changeTarget: function () {
      var t = this.city + ', ' + this.target + ', Polska';
      var geocoder = new google.maps.Geocoder();
      var b = this.map;
      geocoder.geocode({ 'address': t }, (results, status) => {
        if (status === 'OK') {
          b.panTo(results[0].geometry.location);
          b.setZoom(15);
          if (this.marker) {
            this.marker.setMap(null);
          }
          if (this.marker === null) {
            var marker = new google.maps.Marker({
              position: results[0].geometry.location,
              title: "Tu jesteś!"
            });
            this.marker = marker;
            marker.setMap(this.map);
          }
          this.marker = null

        }
      })
    }
  }
}
</script>