<template>
  <div id="map"></div>
</template>

<script>
export default {
  data () {
    var API_PATH = 'http://twotired.math.party/data?timestamp=1483228952&duration=100';

    var request = new XMLHttpRequest();
    request.open("GET", API_PATH, false);
    request.send(null)

    var markers = JSON.parse(request.responseText);
    return {
      markers: markers
    }
  },
  mounted: function () {
    this.initMap();
    this.populateMarkers();
  },
  methods: {
    initMap: function () {
      var uluru = {lat: 52.385392, lng: 16.992963};
      this.map = new google.maps.Map(document.getElementById('map'), {
          scrollwheel: false,
          zoom: 13,
          center: uluru
      })
      // this.changeTarget();
    },
    populateMarkers: function () {
      var map = this.map
      for (let i in this.markers) {
        let data = this.markers[i].steps;
        // console.log(data)
        /*let latLng = new google.maps.LatLng(data[0].lat, data[0].lng);

        let marker = new google.maps.Marker({
          position: latLng,
          map: this.map,
          clickable: true,
          title: data.name,
          data: data
        });
        this.markers[i].marker = marker;*/

        var lineSymbol = {
          path: google.maps.SymbolPath.CIRCLE,
          scale: 10,
          strokeColor: '#cc0000'
        };

        var path = [];
        for (let j in data) {
          path.push({lat: data[j].lat, lng: data[j].lng});
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

        this.animateCircle(markerPath)

        markerPath.setMap(this.map);


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
      }
    },
    animateCircle: function (line) {
      var count = 0;
      window.setInterval(function() {
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