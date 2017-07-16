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
      interval: 400,
      addition: 4,
      future: 10,
      markers: {},
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
      if (this.markersData === undefined || (current - this.lastUpdate) > this.future-this.addition) {
        this.lastUpdate = current;
        this.fetchData(current, this.future);
      } else {
        var slider = document.getElementById('slider-connect');
        slider.noUiSlider.set(current);
        //console.log(current*1000)
        this.updatePaths(current);
        this.pathUpdate = current;
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
          let currPath = line.getPath();
          let prev;
          for (let step of markerData.steps) {
            if (Number(step.ts) > this.pathUpdate && Number(step.ts) <= time) {
              prev = step;
              currPath.push(new google.maps.LatLng(step.lat, step.lng));
              if(step.last) {
                line.setMap(null);
                break;
              }
            } else if(Number(step.ts) > this.pathUpdate && prev && Number(step.ts)>time) {
              continue;
              let dlat = Number(prev.lat)+(Number(step.lat)-Number(prev.lat))/(Number(step.ts)-time);
              let dlng = Number(prev.lng)+(Number(step.lng)-Number(prev.lng))/(Number(step.ts)-time);
              currPath.push(new google.maps.LatLng(dlat, dlng));
              break;
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

      if (this.colors[marker['activity_type']]) {
        lineSymbol.strokeColor = this.colors[marker['activity_type']];
        markerPath.strokeColor = this.colors[marker['activity_type']];
      } else {
        lineSymbol.strokeColor = this.colors["DEFAULT"];
        markerPath.strokeColor = this.colors["DEFAULT"];
      }

      var path = [];
      for (let step of marker.steps) {
        if (Number(step.ts) <= current) {
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