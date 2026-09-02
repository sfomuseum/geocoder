var sfomuseum = sfomuseum || {};

sfomuseum.geocoder = (function(){

    const null_island = [ 0.0, 0.0 ];    
    var endpoint = "http://localhost:8080";
    
    var self = {

	/**
         * Perform a geocoding query against the API.
	 * @param {FormData} params - Query parameters.
         * @returns {Promise<Object>} Resolves with the parsed JSON response.
         */
	query: function(params) {

	    return new Promise ((resolve, reject) => {
		
		const uri = endpoint + "/api/query/";
		
		const args = {
		    method: 'POST',
		    body: params,
		};
		
		fetch(uri, args).then((rsp) => {

		    // Check status code here...
			
		    resolve(rsp.json());
		}).catch((err) => {
		    reject(err);
		});
	    });
	},

	/**
         * Create a Leaflet map in a specified container.
         * @param {string} map_id - ID of the container element.
         * @returns {L.Map} The Leaflet map instance.
         */
	newMap: function(map_id) {

	    const map_args = {};
	    
	    map = L.map(map_id, map_args);
	    
	    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
		maxZoom: 19,
		attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
	    }).addTo(map);

	    map.setView(null_island, 1);
	    return map;
	},

	/**
         * Override the default API endpoint.
         * @param {string} url - Base URL for the geocoding API.
         */
	setEndpoint: function(url) {
	    endpoint = url;
	},
    };

    return self;

})();
