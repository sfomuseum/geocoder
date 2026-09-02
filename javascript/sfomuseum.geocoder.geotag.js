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

	init: function(target, id, on_select, custom_map){

	    const dialog = self.renderDialog(id);
	    target.prepend(dialog);
	    dialog.show();
	    
	    self.initForm(id, on_select, custom_map);	    	    
	},

	renderDialog: function(id) {
	
	    const dialog_id = "geocoder-new-dialog-" + id;
	    const close_id = "geocoder-new-dialog-close" + id;	    

	    const dialog = document.createElement("dialog");
	    dialog.setAttribute("id", dialog_id);
	    dialog.setAttribute("class", "geocoder-dialog");
	    
	    const close = document.createElement("button");
	    close.setAttribute("id", close_id);
	    close.setAttribute("class", "btn btn-primary geocoder-dialog-close");
	    close.appendChild(document.createTextNode("close"));

	    close.onclick = function(){
		self.closeDialog(id);
		return false;
	    };
	    
	    const form = self.renderForm(id);

	    dialog.appendChild(close);
	    dialog.appendChild(form);
	    
	    return dialog;
	},

	closeDialog: function(id){

	    const dialog_id = "geocoder-new-dialog-" + id;
	    const dialog_el = document.getElementById(dialog_id);

	    if (dialog_el){
		dialog_el.close();
		dialog_el.parentNode.removeChild(dialog_el);
	    }
	},
	
	renderForm: function(id) {
	    
	    const form_id = "geocoder-new-form-" + id;
	    const map_id = "geocoder-new-map-" + id;
	    const search_id = "geocoder-new-search-" + id;
	    const candidates_id = "geocoder-new-candidates-" + id;
	    const status_id = "geocoder-new-candidates-status" + id;
	    const submit_id = "geocoder-new-submit" + id;	    	    	    	    
	    
	    const form = document.createElement("form");
	    form.setAttribute("id", form_id);
	    form.setAttribute("class", "form");

	    const map_div = document.createElement("div");
	    map_div.setAttribute("id", map_id);
	    map_div.setAttribute("class", "map");
	    map_div.setAttribute("style", "width: 100%; height: 200px; border:solid thin;");
	    
	    const search_label = document.createElement("label");
	    search_label.setAttribute("for", search_id);
	    search_label.setAttribute("class", "form-label");
	    search_label.appendChild(document.createTextNode("Search for a place"));

	    const search_input = document.createElement("input");
	    search_input.setAttribute("id", search_id);
	    search_input.setAttribute("type", "text");
	    search_input.setAttribute("class", "form-control");	    
	    search_input.setAttribute("value", "");
	    search_input.setAttribute("geocoder", "Search for a place to associate with the label");

	    const candidates_status = document.createElement("span");
	    candidates_status.setAttribute("id", status_id);
	    candidates_status.setAttribute("style", "font-style:italic;margin-left:.5rem");
	    
	    const candidates_label = document.createElement("label");
	    candidates_label.setAttribute("for", candidates_id);
	    candidates_label.setAttribute("class", "form-label");
	    candidates_label.appendChild(document.createTextNode("Select a place"));
	    candidates_label.appendChild(candidates_status);

	    const candidates_opt = document.createElement("option");
	    candidates_opt.setAttribute("value", "-1");
	    candidates_opt.appendChild(document.createTextNode(""));
	    
	    const candidates_select = document.createElement("select");
	    candidates_select.setAttribute("id", candidates_id);
	    candidates_select.setAttribute("class", "form-select");
	    candidates_select.appendChild(candidates_opt);
	    
	    const submit_button = document.createElement("button");
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
	
	initForm: function(id, on_select, custom_map){

	    const map_id = "geocoder-new-map-" + id;
	    const search_id = "geocoder-new-search-" + id;
	    const label_id = "geocoder-new-label-" + id;	    
	    const candidates_id = "geocoder-new-candidates-" + id;
	    const status_id = "geocoder-new-candidates-status" + id;
	    const submit_id = "geocoder-new-submit" + id;
	    
	    const map_el = document.getElementById(map_id);
	    const search_el = document.getElementById(search_id);
	    const label_el = document.getElementById(label_id);	    
	    const candidates_el = document.getElementById(candidates_id);
	    const status_el = document.getElementById(status_id);
	    const submit_el = document.getElementById(submit_id);	    	    

	    const map = sfomuseum.geocoder.map(custom_map);
	    
	    const popup;
	    	    	    
	    search_el.onchange = function(e){
		
		const el = e.target;
		const q = el.value;
		
		if (q.length < 3){
		    return;
		}
		
		if (popup){
		    popup.removeFrom(map);
		}
		
		const geocode_args = {
		    query: q,
		};
		
		status_el.innerText = "searching...";	
		
		const geocoder_params = new FormData();
		geocoder_params.set("query": q);
		
		sfomuseum.geocoder.query(geocode_params).then((rsp) => {
		    
		    // API returns GeoJSON FeatureCollection		    
		    const places = rsp.results.features;
		    const count = places.length;
		    
		    status_el.innerHTML = "";
		    _lookup = {};
		    
		    const opt = document.createElement("option");
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
		    
		    for (const i=0; i < count; i++){
			
			const pl = places[i];
			const props = pl.properties;
			_lookup[props["geocoder:id"]] = pl;
			
			const label = props["geocoder:label"];
			
			const opt = document.createElement("option");
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

	    const _self = self;
	    
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
