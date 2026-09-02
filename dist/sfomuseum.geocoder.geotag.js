// Display a modal dialog to query the sfomuseum/geocoder API, display the results
// in a select menu and updating the map as the (select) menu is updated.
// This does not update anything except the map in the modal dialog but returns
// the currently selected place (in an "on_select") callback and closes the dialog
// on confirmation.
var sfomuseum = sfomuseum || {};
sfomuseum.geocoder = sfomuseum.geocoder || {};

// It may be tempting to try and reconcile/smush-up this code with the code
// sfomuseum.geocoder.georeference.js but they are sufficiently different,
// desipte sharing quite a lof of code, that it's not really worth it.

sfomuseum.geocoder.geotag = (function(){

    var _lookup = {};
    
    var self = {

	// TBD: Options to customize the following:
	//
	// Tile layer(s)/function - currently defaults to OSM but from a UX
	// perspective it might be desireable to mirror other map layers.
	//
	// API endpoint/function - currently this is calling the SFO Museum API
	// (as expected) but it would be good to allow/default to the
	// default Geocoder API.
	//
	// Which is to say yes, but maybe not "top of the list" right now.
	
	init: function(target, id, on_select){

	    var dialog = self.renderDialog(id);
	    target.prepend(dialog);
	    dialog.show();
	    
	    self.initForm(id, on_select);	    	    
	},

	renderDialog: function(id) {
	
	    var dialog_id = "geocoder-new-dialog-" + id;
	    var close_id = "geocoder-new-dialog-close" + id;	    

	    var dialog = document.createElement("dialog");
	    dialog.setAttribute("id", dialog_id);
	    dialog.setAttribute("class", "geocoder-dialog");
	    
	    var close = document.createElement("button");
	    close.setAttribute("id", close_id);
	    close.setAttribute("class", "btn btn-primary geocoder-dialog-close");
	    close.appendChild(document.createTextNode("close"));

	    close.onclick = function(){
		self.closeDialog(id);
		return false;
	    };
	    
	    var form = self.renderForm(id);

	    dialog.appendChild(close);
	    dialog.appendChild(form);
	    
	    return dialog;
	},

	closeDialog: function(id){

	    var dialog_id = "geocoder-new-dialog-" + id;
	    var dialog_el = document.getElementById(dialog_id);

	    if (dialog_el){
		dialog_el.close();
		dialog_el.parentNode.removeChild(dialog_el);
	    }
	},
	
	renderForm: function(id) {
	    
	    var form_id = "geocoder-new-form-" + id;
	    var map_id = "geocoder-new-map-" + id;
	    var search_id = "geocoder-new-search-" + id;
	    var candidates_id = "geocoder-new-candidates-" + id;
	    var status_id = "geocoder-new-candidates-status" + id;
	    var submit_id = "geocoder-new-submit" + id;	    	    	    	    
	    
	    var form = document.createElement("form");
	    form.setAttribute("id", form_id);
	    form.setAttribute("class", "form");

	    var map_div = document.createElement("div");
	    map_div.setAttribute("id", map_id);
	    map_div.setAttribute("class", "map");
	    map_div.setAttribute("style", "width: 100%; height: 200px; border:solid thin;");
	    
	    var search_label = document.createElement("label");
	    search_label.setAttribute("for", search_id);
	    search_label.setAttribute("class", "form-label");
	    search_label.appendChild(document.createTextNode("Search for a place"));

	    var search_input = document.createElement("input");
	    search_input.setAttribute("id", search_id);
	    search_input.setAttribute("type", "text");
	    search_input.setAttribute("class", "form-control");	    
	    search_input.setAttribute("value", "");
	    search_input.setAttribute("geocoder", "Search for a place to associate with the label");

	    var candidates_status = document.createElement("span");
	    candidates_status.setAttribute("id", status_id);
	    candidates_status.setAttribute("style", "font-style:italic;margin-left:.5rem");
	    
	    var candidates_label = document.createElement("label");
	    candidates_label.setAttribute("for", candidates_id);
	    candidates_label.setAttribute("class", "form-label");
	    candidates_label.appendChild(document.createTextNode("Select a place"));
	    candidates_label.appendChild(candidates_status);

	    var candidates_opt = document.createElement("option");
	    candidates_opt.setAttribute("value", "-1");
	    candidates_opt.appendChild(document.createTextNode(""));
	    
	    var candidates_select = document.createElement("select");
	    candidates_select.setAttribute("id", candidates_id);
	    candidates_select.setAttribute("class", "form-select");
	    candidates_select.appendChild(candidates_opt);
	    
	    var submit_button = document.createElement("button");
	    submit_button.setAttribute("id", submit_id);
	    submit_button.setAttribute("class", "btn btn-primary geocoder-dialog-add");
	    // To do: Add code to toggle this on/off depending on label, place values
	    // submit_button.setAttribute("disabled", "disabled");
	    submit_button.appendChild(document.createTextNode("Select"));
	    
	    form.appendChild(map_div);
	    form.appendChild(search_label);
	    form.appendChild(search_input);
	    form.appendChild(candidates_label);
	    form.appendChild(candidates_select);	    
	    form.appendChild(submit_button);

	    return form;
	},
	
	initForm: function(id, on_select){

	    var map_id = "geocoder-new-map-" + id;
	    var search_id = "geocoder-new-search-" + id;
	    var label_id = "geocoder-new-label-" + id;	    
	    var candidates_id = "geocoder-new-candidates-" + id;
	    var status_id = "geocoder-new-candidates-status" + id;
	    var submit_id = "geocoder-new-submit" + id;
	    
	    var map_el = document.getElementById(map_id);
	    var search_el = document.getElementById(search_id);
	    var label_el = document.getElementById(label_id);	    
	    var candidates_el = document.getElementById(candidates_id);
	    var status_el = document.getElementById(status_id);
	    var submit_el = document.getElementById(submit_id);	    	    
	    
	    var map = sfomuseum.maps.leaflet.createSFOMuseumMap(map_el);
	    var popup;
	    
	    // Please fix me...
	    // const tile_url = sfomuseum.maps.protomaps.tileURL({"api_key": "foo"});
	    // const tile_layer = sfomuseum.maps.protomaps.tileLayer(tile_url);
	    // tile_layer.addTo(map);
	    
	    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
		maxZoom: 19,
		attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
	    }).addTo(map);
	    
	    search_el.onchange = function(e){
		
		var el = e.target;
		var q = el.value;
		
		if (q.length < 3){
		    return;
		}
		
		if (popup){
		    popup.removeFrom(map);
		}
		
		var geocode_args = {
		    query: q,
		};
		
		status_el.innerText = "searching...";	

		var geocoder_uri;	// FIX ME
		
		const geocoder_params = new FormData();
		geocoder_params.set("query": q);
		
		const geocode_args = {
		    method: "POST",
		    body: geocoder_params,
		};	
		
		fetch(geocoder_uri, geocoder_args).then((rsp) => {

		    // Check status code...
			
		    return rsp.json();
		    
		}).then((rsp) => {
		    
		    // API returns GeoJSON FeatureCollection		    
		    var places = rsp.results.features;
		    var count = places.length;
		    
		    status_el.innerHTML = "";
		    _lookup = {};
		    
		    var opt = document.createElement("option");
		    opt.setAttribute("value", "-1");
		    opt.appendChild(document.createTextNode(""));
		    candidates_el.appendChild(opt);
		    
		    switch (count){
		    case 0:
			status_el.innerText = "no results found";		
			console.log("No matching places for query");
			return;
		    case 1:
			status_el.innerText = "one result found";
			break;
		    default:
			status_el.innerText = count + " results found";
			break;		
		    }
		    
		    for (var i=0; i < count; i++){
			
			const pl = places[i];
			const props = pl.properties;
			_lookup[props["geocoder:id"]] = pl;
			
			const label = props["geocoder:label"];
			
			var opt = document.createElement("option");
			opt.setAttribute("value", props["geocoder:id"]);
			opt.setAttribute("data-bbox", pl.bbox);
			opt.setAttribute("data-latitude", pl.geometry.coordinates[1]);
			opt.setAttribute("data-longitude", pl.geometry.coordinates[0]);				
			
			opt.appendChild(document.createTextNode(label + " (" + props["wof:placetype"] + ")"));
			
			candidates_el.appendChild(opt);
		    }
		    
		    candidates_el.onchange = function(candidates_e){
			
			const el = candidates_e.target;
			const selected = el.children[el.selectedIndex];
			
			const lat = selected.getAttribute("data-latitude");
			const lon = selected.getAttribute("data-longitude");
			const bbox = selected.getAttribute("data-bbox");				
			const bounds = bbox.split(",");
			
			const leaflet_bounds = [
			    [ bounds[1], bounds[0] ],
			    [ bounds[3], bounds[2] ],
			];
			
			map.fitBounds(leaflet_bounds);
			
			if (popup){
			    popup.removeFrom(map);
			}
			
			popup = L.popup()
			    .setLatLng([lat, lon])
			    .setContent(selected.textContent)
			    .openOn(map);
			
			return false;
		    };
		    
		}).catch((err) => {

		    console.error("Failed to geocode text", err);
		    
		    status_el.innerHTML = "";
		    status_el.innerText = "There was a problem executing the geocoding request, " + err;				    
		});
	    };

	    var _self = self;
	    
	    submit_el.onclick = function(e){

		try {
		    
		    if (! on_select){
			_self.closeDialog();
			return false;
		    }

		    const place_id = candidates_el.value;
		    const place = _lookup[place_id];
		    
		    // Check values here...
		    
		    on_select(place).then((rsp) => {
			_self.closeDialog(id);
		    }).catch((err) => {
			console.error("Failed to process geocoding, on_select callback error", err);
		    });
		    
		} catch (err) {
		    console.log("Failed to complete", err);
		}
		
		return false;
	    };
	},
	
    };

    return self;

})();
