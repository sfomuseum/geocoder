window.addEventListener("load", function load(event){

    const null_island = [ 0.0, 0.0 ];
    
    const map = sfomuseum.geocoder.map();
    map.setView(null_island, 0);

    const geocoder_onselect = function(place){
	
	console.log("GOT PLACE", place);
        return new Promise((resolve, reject) => {
	    
            try {
	        const bounds = place.bbox;
		
                const leaflet_bounds = L.latLngBounds(
                    [ bounds[1], bounds[0] ],
		    [ bounds[3], bounds[2] ],
                );
		
                map.fitBounds(leaflet_bounds);
	        resolve();
		
            } catch(err) {
	        console.error("Failed to update map", err);
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
