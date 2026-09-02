var sfomuseum = sfomuseum || {};

sfomuseum.geocoder = (function(){

    var endpoint = "http://localhost:8080";
    
    var self = {

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

	queryFunc: function(custom_func) {

	    if (custom_func){
		return custom_func;
	    }

	    return self.query;
	},
	
	map: function(custom_map) {

	    if (custom_map){
		return custom_map;
	    }

	    const map_id = "map";
	    const map_args = {};
	    
	    map = L.map(map_id, map_args);
	    
	    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
		maxZoom: 19,
		attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
	    }).addTo(map);

	    return map;
	},
	
	setEndpoint: function(url) {
	    endpoint = url;
	},
    };

    return self;

})();
