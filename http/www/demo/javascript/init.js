window.addEventListener("load", function load(event){

    const null_island = [ 0.0, 0.0 ];
    var map;

    var markers_layer;    
    var results_layer;
    var feature_cache = {};

    var current_form;
    var current_params;	// API query parameters
    
    const map_el = document.querySelector("#map");    
    const query_el = document.querySelector("#query");
    const submit_el = document.querySelector("#query-submit");
    const results_el = document.querySelector("#results");
    const feedback_el = document.querySelector("#feedback");    

    const adv_details_el = document.querySelector("#advanced-details");    
    const adv_query_el = document.querySelector("#advanced-query");
    const adv_query_embeddings_el = document.querySelector("#advanced-query-embeddings");    
    const adv_lang_el = document.querySelector("#advanced-lang");
    const adv_tag_el = document.querySelector("#advanced-tag");
    const adv_placetype_el = document.querySelector("#advanced-placetype");
    const adv_country_el = document.querySelector("#advanced-country");
    const adv_belongsto_el = document.querySelector("#advanced-belongsto");
    const adv_submit_el = document.querySelector("#advanced-submit");                    

    var geocode_spinner = document.getElementById("geocode-spinner-svg");    
    var adv_geocode_spinner = document.getElementById("advanced-geocode-spinner-svg");

    const prepare_id = function(id){
	id = id.replace(":", "-");
	id = id.replace("=", "-");
	return id;
    };
    
    const select_row = function(id){

	unselect_row();
	
	const row = document.querySelector("#results-" + id);
	 
	if (row){
	    row.setAttribute("class", "selected");
	    row.scrollIntoView();
	}
    };

    const unselect_row = function(){

	var current = document.querySelector(".selected");
	
	if (current){
	    current.classList.remove("selected");
	}
    };
    
    const draw_results_map = function(data){

	if (results_layer){
	    map.removeLayer(results_layer);
	}

	if (markers_layer){
	    map.removeLayer(markers_layer);
	}

	feature_cache = {};
	
	const marker_opts = {
	    radius: 8,
	    fillColor: "#ff7800",
	    color: "#000",
	    weight: 1,
	    opacity: 1,
	    fillOpacity: 0.8
	};
	
	const args = {
	    pointToLayer: function (feature, latlng){
		return L.circleMarker(latlng, marker_opts);
	    },
	    onEachFeature: function(feature, layer){
		const id = feature.properties["geocoder:id"];
		const id_prep = prepare_id(id);
		
		layer.bindPopup(feature.properties["geocoder:label"]);
		feature_cache[id_prep] = layer;

		layer.on("click", function(){
		    console.log("CLICK", id_prep);
		    select_row(id_prep);
		});
		
	    }
	};
	
	results_layer =  L.geoJSON(data, args);

	markers_layer = L.markerClusterGroup();
	markers_layer.addLayer(results_layer);
	markers_layer.addTo(map);
	
	// results_layer.addTo(map);
    };
    
    const draw_results = function(rsp){

	feedback_el.innerHTML = "";
	
	const pg = rsp.pagination;
	const data = rsp.results;
	
	draw_results_map(data);

	const count = data.features.length;	

	const pagination = document.createElement("div");
	pagination.setAttribute("id", "pagination");

	const total = pg.total;
	const page = pg.page;
	const pages = pg.pages;

	const query_el = document.createElement("span");
	
	const quote_el = document.createElement("q");
	quote_el.appendChild(document.createTextNode(current_form.get("query")));

	query_el.appendChild(quote_el);
	
	if (current_form.get("query-embeddings") != ""){
	    query_el.appendChild(document.createTextNode(" (with vector embeddings)"));
	}
	
	switch (total) {
	    case 0:
		feedback_el.innerText = "There are no results for that query.";
		return;
	    case 1:
		pagination.appendChild(document.createTextNode("There is one result for"));
		pagination.appendChild(query_el);
		pagination.appendChild(document.createTextNode("."));
		break;
	    default:
		pagination.appendChild(document.createTextNode("There are " + total + " results for"));
		pagination.appendChild(query_el);
		pagination.appendChild(document.createTextNode("."));

		if (pages > 1) {
		    pagination.appendChild(document.createTextNode("This is page " + page + " of " + pages + "."));
		}
		
		const geocode_with_page = function(e){
		    const el = e.target;
		    const page = parseInt(el.getAttribute("data-page"));
		    current_form.set("page", page);
		    geocode(current_form);
		    return false;
		};
		
		if (pg.previous_page){
		    
		    const prev_wrapper = document.createElement("span");
		    prev_wrapper.setAttribute("class", "pagination-link pagination-prev");
		    prev_wrapper.setAttribute("data-page", pg.previous_page);
		    prev_wrapper.appendChild(document.createTextNode(" See previous page of results"));

		    if (pg.next_page){
			prev_wrapper.appendChild(document.createTextNode(" / "));
		    } else {
			prev_wrapper.appendChild(document.createTextNode("."));
		    }
		    
		    prev_wrapper.onclick = geocode_with_page;
		    pagination.appendChild(prev_wrapper);		    
		}
		
		if (pg.next_page){
		    
		    const next_wrapper = document.createElement("span");
		    next_wrapper.setAttribute("class", "pagination-link pagination-next");
		    next_wrapper.setAttribute("data-page", pg.next_page);
		    next_wrapper.appendChild(document.createTextNode(" See next page of results."));
		    
		    next_wrapper.onclick = geocode_with_page;		    
		    pagination.appendChild(next_wrapper);
		}
		
		break;
	}

	feedback_el.appendChild(pagination);
	
	const table = document.createElement("table");
	table.setAttribute("class", "table table-striped");

	const header_row = document.createElement("tr");

	const rank_header = document.createElement("th");
	rank_header.appendChild(document.createTextNode("Rank"));
	header_row.appendChild(rank_header);
	
	const id_header = document.createElement("th");
	id_header.setAttribute("scope", "col");
	id_header.appendChild(document.createTextNode("Id"));
	header_row.appendChild(id_header);

	/*
	const name_header = document.createElement("th");
	name_header.setAttribute("scope", "col");	
	name_header.appendChild(document.createTextNode("Name"));
	header_row.appendChild(name_header);
	 */
	
	const label_header = document.createElement("th");
	label_header.setAttribute("scope", "col");		
	label_header.appendChild(document.createTextNode("Label"));
	header_row.appendChild(label_header);

	const latitude_header = document.createElement("th");
	latitude_header.setAttribute("scope", "col");		
	latitude_header.appendChild(document.createTextNode("Latitude"));
	header_row.appendChild(latitude_header);

	const longitude_header = document.createElement("th");
	longitude_header.setAttribute("scope", "col");		
	longitude_header.appendChild(document.createTextNode("Longitude"));
	header_row.appendChild(longitude_header);
	
	const placetype_header = document.createElement("th");
	placetype_header.setAttribute("scope", "col");		
	placetype_header.appendChild(document.createTextNode("Placetype"));
	header_row.appendChild(placetype_header);

	const country_header = document.createElement("th");
	country_header.setAttribute("scope", "col");		
	country_header.appendChild(document.createTextNode("Country"));
	header_row.appendChild(country_header);
	
	const is_current_header = document.createElement("th");
	is_current_header.setAttribute("scope", "col");		
	is_current_header.appendChild(document.createTextNode("Current"));
	header_row.appendChild(is_current_header);

	const inception_header = document.createElement("th");
	inception_header.setAttribute("scope", "col");		
	inception_header.appendChild(document.createTextNode("Inception"));
	header_row.appendChild(inception_header);

	const cessation_header = document.createElement("th");
	cessation_header.setAttribute("scope", "col");			
	cessation_header.appendChild(document.createTextNode("Cessation"));
	header_row.appendChild(cessation_header);

	thead = document.createElement("thead");
	thead.appendChild(header_row);	
	table.appendChild(thead);

	tbody = document.createElement("tbody");
	
	for (var i=0; i < count; i++){
    
	    const f = data.features[i];
	    const props = f.properties;
	    const id = props["geocoder:id"];
	    const id_prep = prepare_id(id);
	    
	    const feature_row = document.createElement("tr");
	    feature_row.setAttribute("id", "results-" + id_prep);
	    
	    const rank_feature = document.createElement("td");
	    rank_feature.appendChild(document.createTextNode(props["geocoder:rank"]));
	    feature_row.appendChild(rank_feature);

	    const id_link = document.createElement("span");
	    id_link.setAttribute("class", "id-link");
	    id_link.setAttribute("data-id", id_prep);
	    id_link.appendChild(document.createTextNode(id));

	    id_link.onclick = function(e){

		const el = e.target;
		const id = el.getAttribute("data-id");

		console.log("DATA", id);
		
		const layer = feature_cache[id];

		console.log("LAYER", layer);
		
		if (layer) {
		    const pos = layer.getLatLng();

		    const popup = layer.getPopup();
		    popup.options.autoPan = false;
		    
		    map.setView(pos, 12); 		    
		    map.openPopup(popup, pos);
		    select_row(id);
		}
		
		return false;
	    };
	    
	    const id_feature = document.createElement("td");
	    id_feature.appendChild(id_link);
	    feature_row.appendChild(id_feature);

	    /*
	    const name_feature = document.createElement("td");
	    name_feature.appendChild(document.createTextNode(props["geocoder:name"]));
	    feature_row.appendChild(name_feature);
	    */
	    
	    const label_feature = document.createElement("td");
	    label_feature.appendChild(document.createTextNode(props["geocoder:label"]));
	    feature_row.appendChild(label_feature);

	    const latitude_feature = document.createElement("td");
	    latitude_feature.appendChild(document.createTextNode(f.geometry.coordinates[1]));
	    feature_row.appendChild(latitude_feature);

	    const longitude_feature = document.createElement("td");
	    longitude_feature.appendChild(document.createTextNode(f.geometry.coordinates[0]));
	    feature_row.appendChild(longitude_feature);
	    
	    var all_placetypes = [
		props["wof:placetype"],
	    ];

	    if (props["wof:placetype_alt"]){

		for (const i in props["wof:placetype_alt"]){
		    all_placetypes.push(props["wof:placetype_alt"][i]);
		}
	    }
	    
	    const placetype_feature = document.createElement("td");
	    placetype_feature.appendChild(document.createTextNode(all_placetypes.join("; ")));
	    feature_row.appendChild(placetype_feature);
	    
	    const country_feature = document.createElement("td");
	    country_feature.appendChild(document.createTextNode(props["wof:country"]));
	    feature_row.appendChild(country_feature);
	    
	    const is_current_feature = document.createElement("td");
	    is_current_feature.appendChild(document.createTextNode(props["mz:is_current"]));
	    feature_row.appendChild(is_current_feature);
	    
	    const inception_feature = document.createElement("td");
	    inception_feature.appendChild(document.createTextNode(props["edtf:inception"]));
	    feature_row.appendChild(inception_feature);
	    
	    const cessation_feature = document.createElement("td");
	    cessation_feature.appendChild(document.createTextNode(props["edtf:cessation"]));
	    feature_row.appendChild(cessation_feature);
	    	    
	    tbody.appendChild(feature_row);
	}

	table.appendChild(tbody);

	const tdiv = document.createElement("div");
	tdiv.setAttribute("class", "table-responsive");
	tdiv.appendChild(table);
	    
	results_el.appendChild(tdiv);
	results_el.style.display = "block";
    };
    
    const geocode = function(form_data){

	current_form = form_data;
	
	const u = new URL(location);
	u.pathname = u.pathname + "api/query/";

	const uri = u.toString();
	console.debug("Query API", uri);
	
	submit_el.setAttribute("disabled", "disabled");
	adv_submit_el.setAttribute("disabled", "disabled");	

	if (results_layer){
	    map.removeLayer(results_layer);
	}

	feedback_el.innerHTML = "";	
	results_el.innerHTML = "";

	results_el.style.display = "none";

	const fetch_args = {
	    method: 'POST',
	    body: form_data,
	};

	geocode_spinner.style.display = "inline-block";
	adv_geocode_spinner.style.display = "inline-block";	
	
	fetch(uri, fetch_args).then(rsp => {

	    geocode_spinner.style.display = "none";
	    adv_geocode_spinner.style.display = "none";	    
	    
	    if (! rsp.ok){
		throw new Error(`HTTP error ${rsp.status}`);
	    }
	    
	    return rsp.json();
	    
	}).then((data) => {
	    submit_el.removeAttribute("disabled");
	    adv_submit_el.removeAttribute("disabled");	    
	    draw_results(data);
	}).catch((err) => {
	    submit_el.removeAttribute("disabled");
	    adv_submit_el.removeAttribute("disabled");

	    feedback_el.innerText = "Failed to geocode query:" + err;
	    console.error("Failed to geocode query", uri, err);
	});
	
    };
    
    submit_el.onclick = function(){

	const q = query_el.value;

	if (q == ""){
	    alert("Missing query");
	    return false;
	}

	const form_data = new FormData();
	form_data.set("query", q);

	geocode(form_data);

	results_el.innerHTML = "";
	feedback_el.innerHTML = "";
	
	return false;
    };

    adv_submit_el.onclick = async function(e){

	e.preventDefault();
	
	const form_data = new FormData();
	
	const q = adv_query_el.value;

	if (q == ""){
	    alert("Missing query");
	    return false;
	}

	form_data.set("query", q);
	
	const lang = adv_lang_el.value;

	if (lang != ""){
	    form_data.set("lang", lang);
	}
	
	const tag = adv_tag_el.value;

	if (tag != ""){
	    form_data.set("tag", tag);
	}
	
	const pt = adv_placetype_el.value;

	if (pt != ""){
	    form_data.set("placetype", pt.split(","));
	}
	
	const co = adv_country_el.value;

	if (co != ""){
	    form_data.set("country", co.split(","));
	}
	
	const bt = adv_belongsto_el.value;

	if (bt != ""){
	    form_data.set("belongsto", bt.split(","));
	}

	results_el.style.display = "none";	
	results_el.innerHTML = "";
	feedback_el.innerHTML = "";
	
	const with_embeddings = adv_query_embeddings_el.checked;

	if (with_embeddings){

	    feedback_el.innerText = "Deriving vector embeddings for query text";

	    adv_geocode_spinner.style.display = "inline-block";
	    
	    try {
		
		const rsp = await window.getEmbedding(q);
		const data = rsp.data[0].embedding;

		if (! data){
		    console.debug("Embedding response missing data", rsp);
		    throw new Error("No embedding data");
		}

		form_data.set("query-embeddings", JSON.stringify(data));

		feedback_el.innerText = "";
		adv_geocode_spinner.style.display = "none";		
		
	    } catch(err) {		
		console.error("Failed to generate embeddings", err);
		adv_geocode_spinner.style.display = "none";
		feedback_el.innerText = "Failed to derive vector embeddings, " + err;
		return false;
	    }

	}
	
	// const lang = adv_lang_el.value;	

	console.debug("Geocode", form_data);
	geocode(form_data);
	
	adv_details_el.open = false;
	return false;
    };
    
    //
    
    const map_u = new URL(location);
    map_u.pathname = map_u.pathname + "map.json";
	
    fetch(map_u.toString()).then(
	rsp => rsp.json()
    ).then((cfg) => {

	map_el.style.display = "block";
	
	map = L.map('map');

	map.on("click", function(){
	    unselect_row();
	});
	
	map.setView(null_island, 1);
	
        var tile_url = cfg.tile_url;

        var tile_layer = L.tileLayer(tile_url, {
            maxZoom: 19,
	    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
        });
	
        tile_layer.addTo(map);

	submit_el.removeAttribute("disabled");
	adv_submit_el.removeAttribute("disabled");
    	
    }).catch((err) => {

	feedback_el.innerText = "Failed to retrieve map config:" + err;	
	console.error("Failed to retrieve map config", map_u.toString(), err);
    });

    // Load embeddings stuff (but only if wllama.js is present)

    fetch_args = {
	method: "HEAD",
    };

    const js_u = new URL(location);
    js_u.pathname = js_u.pathname + 'javascript/wllama.js';

    fetch(js_u.toString(), fetch_args).then(rsp => {

	if (! rsp.ok){
	    throw new Error("wwlama.js not found");
	}
	
    }).then(() => {
	
	(async () => {
            try {
		
		const module_u = new URL(location);
		module_u.pathname = module_u.pathname + 'javascript/embeddings.js';
		
		const { initModel, getEmbedding } = await import(module_u.toString());
		
		console.debug("Starting model download/initialization...");
		await initModel();
		
		console.debug("Model successfully loaded into memory!");
		
		window.getEmbedding = getEmbedding;
		
		const wrapper = document.querySelector("#advanced-query-embeddings-wrapper");
		wrapper.style.display = "block";
		
            } catch (error) {
		throw error
            }
	})();
	
    }).catch((err) => {
	console.error("Failed to load query embedding scaffolding", err);	
    });
    
});
