window.addEventListener("load", function load(event){

    const null_island = [ 0.0, 0.0 ];

    const place_el = document.querySelector("#place");
    
    const map_id = "map";
    const map_args = {};

    const map = L.map(map_id, map_args);
    
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
	maxZoom: 19,
	attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    map.setView(null_island, 0);

    var place_layer;
    var place_popup;
    
    const place_layer_style = {
	radius: 8,
	fillColor: "#ff7800",
	color: "#000",
	weight: 1,
	opacity: 1,
	fillOpacity: 0.8
    };
        
    const geocoder_onselect = function(place){

	return new Promise((resolve, reject) => {

	    if (! place){
		resolve();
		return;
	    }
	    
	    try {
		
		if (place_layer){
		    map.removeLayer(place_layer);
		}
		
		const bounds = place.bbox;

		const leaflet_bounds = L.latLngBounds(
                    L.latLng(bounds[1], bounds[0]),
		    L.latLng(bounds[3], bounds[2]),
		);

		map.fitBounds(leaflet_bounds);
		
		const place_layer_args = {
		    pointToLayer: function (feature, latlng) {
			return L.circleMarker(latlng, place_layer_style);
		    }
		};
		
		place_layer = L.geoJSON(place, place_layer_args);
		place_layer.addTo(map);
		
		const str_place = JSON.stringify(place, null, 2);	
		
		const pre = document.createElement("pre");
		pre.appendChild(document.createTextNode(str_place));

		place_el.innerHTML = "";		
		place_el.appendChild(pre);

		if (place_popup){
		    place_popup.removeFrom(map);
		}

		const geom = place.geometry;
		const props = place.properties;
		
		place_popup = L.popup()
			 .setLatLng([geom.coordinates[1], geom.coordinates[0]])
			 .setContent(props["geocoder:label"])
			 .openOn(map);
		
		resolve();
		
	    } catch(err) {
		reject(err);
	    }
	    
	});
	
    };
    
    const geocoder_opts = {
        on_select: geocoder_onselect,
    };
    
    const geocoder_control = L.control.geocoder(geocoder_opts);
    geocoder_control.addTo(map);
});
