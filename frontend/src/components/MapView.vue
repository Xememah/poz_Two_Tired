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
      pathUpdate: 0,
      colors: {
        "BIKE_CITY": "#EA2E49",
        "BIKE_MOUNTAIN": "#333745",
        "BIKE_ROAD": "#5C832F",
        "BIKE_TREKKING": "#2980B9",
        "DEFAULT": "#cc0000"
      }
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
        this.pathUpdate = time;
      });
    },
    timeChange: function (current) {
      if ((current-this.lastUpdate) > (this.future-this.addition)*(this.interval/1000)) {
        this.lastUpdate = current;
        this.fetchData(current, this.future);
      } else {
        this.updatePaths(current);
        this.pathUpdate = current;
      }
    },
    updatePaths: function (time, oldData) {
      if (!this.markersData) {
        return;
      }
      for (let markerData of this.markersData) {
        if (!this.markers || !this.markers[markerData.hash]) {
          let markerPath = this.initMarker(markerData, time);
          this.markers[markerData.hash] = markerPath;
        } else {
          let line = this.markers[markerData.hash];
          let currPath = line.getPath();
          for(let step of markerData.steps) {
            if(step.ts > this.pathUpdate && step.ts <= time) {
              currPath.push(new google.maps.LatLng(step.lat, step.lng));
            }
          }
          continue;
          outer:
          for(let i=0;i<markerData.steps.length;i++) {
            let step = markerData.steps[i];
            for(let s of line.getPath().getArray()) {
              if(step.lat == s.lat && step.lng == s.lng) {
                continue outer;
              }
            }
            if(step.ts<=time) {
              if(step.last) {
                line.setMap(null);
                line = null;
                break;
              }
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
        strokeWeight: 3
      };
      var markerPath = new google.maps.Polyline({
        icons: [{
          icon: lineSymbol,
          offset: '100%'
        }],
        geodesic: false,
        strokeColor: '#cc0000',
        strokeOpacity: 1,
        strokeWeight: 3
      });

      if(this.colors[marker['activity_type']]) {
        lineSymbol.strokeColor = this.colors[marker['activity_type']];
        markerPath.strokeColor = this.colors[marker['activity_type']];
      } else {
        lineSymbol.strokeColor = this.colors["DEFAULT"];
        markerPath.strokeColor = this.colors["DEFAULT"];
      }

      var path = [];
      for (let step of marker.steps) {
        if (step.ts <= current) {
          path.push({ lat: step.lat, lng: step.lng });
        }
      }
      markerPath.path = path;
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